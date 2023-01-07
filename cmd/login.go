/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/geetchoubey/gimme-vault/shared/http"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"golang.org/x/crypto/ssh/terminal"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to your account",
	Run: func(cmd *cobra.Command, args []string) {
		if f := config.IsConfigInitialized(); !f {
			fmt.Println("no configuration has been set. run 'gimme-vault configure' first")
			return
		}

		isLoginValid := checkLoginValid()
		if !isLoginValid {
			doLogin()
		}
		awsCreds := doWriteCredentials()
		doExportCredentials(awsCreds)
	},
}

func checkLoginValid() bool {
	token := config.GetString("token")
	if len(token) == 0 {
		return false
	}
	updatedAt := config.GetInt64("updatedAt")
	if updatedAt == 0 {
		return false
	}
	updatedAtTime := time.Unix(updatedAt, 0)
	return int(time.Since(updatedAtTime).Seconds()) < config.GetInt("leaseDuration")
}

func doLogin() {
	fmt.Printf("Logging in using [%s] profile\n", profile)
	fmt.Printf("Password: ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		panic(fmt.Errorf("error reading password %v", err))
	}
	fmt.Println("Logging in...")
	authResponse, err := http.Login(config.GetLoginUrl(), string(password))
	if err != nil {
		jww.DEBUG.Fatalln(err)
		jww.FEEDBACK.Println("Error occurred. Please retry")
		return
	}
	fmt.Println("Successfully logged in")
	config.Set("token", authResponse.ClientToken)
	config.Set("leaseDuration", authResponse.LeaseDuration)
	config.Set("updatedAt", time.Now().Unix())
	config.Save()
}

func doWriteCredentials() http.AWSCredentials {
	fmt.Println("Writing credentials")
	awsResponse, err := http.WriteCredentials(config.GetAWSWriteUrl(), config.GetString("token"))
	if err != nil {
		panic(fmt.Errorf("error writing credentials %v", err))
	}
	return awsResponse
}

func run(value string, field string) {

	if _, err := exec.Command("aws", "configure", "set", field, value).Output(); err != nil {
		panic(fmt.Errorf("error occurred while setting %s %v", field, err))
	} else {
		fmt.Printf("Set [%s] done\n", field)
	}
}

func doExportCredentials(awsCreds http.AWSCredentials) {
	run(awsCreds.AccessKey, "aws_access_key_id")
	run(awsCreds.SecretKey, "aws_secret_access_key")
	run(awsCreds.SecurityToken, "aws_session_token")
	run(config.GetString("region"), "default.region")

}

func init() {
	rootCmd.AddCommand(loginCmd)
}

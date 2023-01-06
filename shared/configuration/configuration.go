/*
Copyright © 2023 Geet Choubey <geetchoubey@gmail.com>
*/
package configuration

import (
	"os"

	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

type Config struct {
	Namespace *string
}

func (c *Config) GetQualifiedKey(key string) string {
	return *c.Namespace + "." + key
}

func (c *Config) GetString(key string) string {
	return viper.GetString(c.GetQualifiedKey(key))
}

func (c *Config) GetBool(key string) bool {
	return viper.GetBool(c.GetQualifiedKey(key))
}

func (c *Config) GetInt(key string) int {
	return viper.GetInt(c.GetQualifiedKey(key))
}

func (c *Config) Get(key string) interface{} {
	return viper.Get(c.GetQualifiedKey(key))
}

func (c *Config) Set(key string, value interface{}) {
	viper.Set(c.GetQualifiedKey(key), value)
}

func (c *Config) Save() error {
	if err := viper.SafeWriteConfig(); err != nil {
		if err := viper.WriteConfig(); err != nil {
			return err
		}
	}
	return nil
}

func Init(fileName string) {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".gimme-vault" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName("." + fileName)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is not found, display error.
	if err := viper.ReadInConfig(); err != nil {
		jww.INFO.Println("Configuration not initialized")
		// panic(fmt.Errorf("unable to use config file: %v", viper.ConfigFileUsed()))
	}
}

package configuration

var Keys = []string{"username", "awsAccount", "region", "url", "version", "resource", "action"}

func (c *Config) GetLoginUrl() string {
	return c.GetString("url") + "/" +
		c.GetString("version") + "/" +
		"auth/ldap/login/" +
		c.GetString("username")
}

func (c *Config) GetAWSWriteUrl() string {
	return c.GetString("url") + "/" +
		c.GetString("version") + "/" +
		"account/" + c.GetString("awsAccount") +
		"/sts/Owner"
}

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
		c.GetString("resource") + "/" +
		c.GetString("awsAccount") + "/" +
		c.GetString("action")
}

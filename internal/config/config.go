package config

import "github.com/spf13/viper"

type Config struct {
	v *viper.Viper
}

// ========== log ========== //

func (c *Config) LogLevel() int {

	return c.v.GetInt("log.level")
}

// ========== api ========== //

func (c *Config) ApiHost() string {

	return c.v.GetString("api.host")
}

func (c *Config) ApiPort() string {

	return c.v.GetString("api.port")
}

func (c *Config) ApiSslCrtLocation() string {

	return c.v.GetString("api.ssl_crt")
}

func (c *Config) ApiSslKeyLocation() string {

	return c.v.GetString("api.ssl_key")
}

func (c *Config) ApiDebug() bool {

	return c.v.GetBool("api.debug")
}

func (c *Config) WhitelistIp() []string {

	return c.v.GetStringSlice("api.whitelist_ip")
}

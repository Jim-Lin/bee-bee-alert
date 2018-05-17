package util

import (
	"path/filepath"

	"github.com/magiconair/properties"
)

type Config struct {
	EsUrl         string `properties:"es_url"`
	RedisUrl      string `properties:"redis_url"`
	SmtpHostname  string `properties:"smtp_hostname"`
	SmtpAddr      string `properties:"smtp_addr"`
	SmtpFrom      string `properties:"smtp_from"`
	SmtpPassword  string `properties:"smtp_password"`
	MarketingMail string `properties:"marketing_mail"`
}

var cfg Config

// singleton pattern
func GetConfig() Config {
	if cfg == (Config{}) {
		abspath, _ := filepath.Abs("./")
		p := properties.MustLoadFile(abspath+"/backend.properties", properties.UTF8)

		err := p.Decode(&cfg)
		CheckError(err)
	}

	return cfg
}

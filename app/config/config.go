package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	Service struct {
		LogLevel        string `yaml:"log_level"`
		ApiAddrPort     string `yaml:"api_addr_port"`
		AppUrl          string `yaml:"app_url"`
		GitInfo         string `yaml:"git_info"`
		TemplateWelcome string `yaml:"template_welcome"`
		TemplateStatic  string `yaml:"template_static"`
		DefaultLocale   string `yaml:"default_locale"`
	} `yaml:"service"`
	Realm struct {
		ID   int    `yaml:"id"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"realm"`
	DB struct {
		Auth       string `yaml:"auth"`
		Characters string `yaml:"characters"`
		World      string `yaml:"world"`
		Sait       string `yaml:"sait"`
		Verbose    bool   `yaml:"verbose"`
	} `yaml:"db"`
	Soap struct {
		Enable   bool   `yaml:"enable"`
		Host     string `yaml:"host"`
		Login    string `yaml:"login"`
		Password string `yaml:"password"`
	} `yaml:"soap"`
	Cors struct {
		AllowOrigins     string `yaml:"allow_origins"`
		AllowMethods     string `yaml:"allow_methods"`
		AllowHeaders     string `yaml:"allow_headers"`
		AllowCredentials bool   `yaml:"allow_credentials"`
	} `yaml:"cors"`
	Cookie struct {
		Name           string        `yaml:"name"`
		Domain         string        `yaml:"domain"`
		HTTPOnly       bool          `yaml:"httpOnly"`
		Secure         bool          `yaml:"secure"`
		SameSite       string        `yaml:"sameSite"`
		AccessDuration time.Duration `yaml:"access_duration"`
		JwtKey         string        `yaml:"jwt_key"`
	} `yaml:"cookie"`
	Mail struct {
		Server                   string `yaml:"server"`
		Port                     string `yaml:"port"`
		User                     string `yaml:"user"`
		Password                 string `yaml:"password"`
		SanderName               string `yaml:"sander_name"`
		FolderTemplates          string `yaml:"folder_templates"`
		TemplateVerify           string `yaml:"template_verify"`
		TemplatePassword         string `yaml:"template_password"`
		TemplatePasswordComplete string `yaml:"template_password_complete"`
		TemplateLogin            string `yaml:"template_login"`
		TemplateChangeEmail      string `yaml:"template_change_email"`
	} `yaml:"mail"`
}

var (
	config  *Config
	once    sync.Once
	cfgPath = "config/config.yml"
)

func init() {
	flag.StringVar(&cfgPath, "cfg", cfgPath, "")
	flag.Parse()
}

func Get() *Config {
	once.Do(func() {
		data, err := os.ReadFile(cfgPath)
		if err != nil {
			log.Fatal(err)
		}

		var cfg Config
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			log.Fatal(err)
		}
		config = &cfg
	})
	return config
}

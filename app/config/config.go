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
		LogLevel        string `env:"log_level"`
		LogFolder       string `yaml:"log_folder"`
		LogPrefix       string `yaml:"log_prefix"`
		ApiAddrPort     string `yaml:"api_addr_port"`
		AppUrl          string `yaml:"app_url"`
		GitInfo         string `yaml:"git_info"`
		TemplateWelcome string `yaml:"template_welcome"`
		TemplateStatic  string `yaml:"template_static"`
		DefaultLocale   string `yaml:"default_locale"`
	} `yaml:"service"`
	DB struct {
		Auth       string `yaml:"auth"`
		Characters string `yaml:"characters"`
		World      string `yaml:"world"`
		Sait       string `yaml:"sait"`
		Verbose    bool   `yaml:"verbose"`
	} `yaml:"db"`
	Cors struct {
		AllowOrigins     string `yaml:"allow_origins"`
		AllowMethods     string `yaml:"allow_methods"`
		AllowHeaders     string `yaml:"allow_headers"`
		AllowCredentials bool   `yaml:"allow_credentials"`
	} `yaml:"cors"`
	Cookie struct {
		Name           string        `env:"name"`
		Domain         string        `env:"domain"`
		HTTPOnly       bool          `env:"httpOnly"`
		Secure         bool          `env:"secure"`
		SameSite       string        `env:"sameSite"`
		AccessDuration time.Duration `env:"access_duration"`
		JwtKey         string        `env:"jwt_key"`
	} `yaml:"cookie"`
	Mail struct {
		Server                   string `env:"server"`
		Port                     string `env:"port"`
		User                     string `env:"user"`
		Password                 string `env:"password"`
		SanderName               string `env:"sander_name"`
		FolderTemplates          string `env:"folder_templates"`
		TemplateVerify           string `env:"template_verify"`
		TemplatePassword         string `env:"template_password"`
		TemplatePasswordComplete string `env:"template_password_complete"`
		TemplateLogin            string `env:"template_login"`
		TemplateChangeEmail      string `env:"template_change_email"`
	} `env:"mail"`
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

package config

import (
	"encoding/json"
	"github.com/caarlos0/env"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	LogLevel                     string        `env:"SERVICE_LOG_LEVEL"  envDefault:"debug"`
	ApiAddrPort                  string        `env:"SERVICE_API_ADDR_PORT" envDefault:"0.0.0.0:3011"`
	AppUrl                       string        `env:"SERVICE_APP_URL" envDefault:"https://example.com"`
	GitInfo                      string        `env:"SERVICE_GIT_INFO" envDefault:"static/gitinfo.json"`
	TemplateWelcome              string        `env:"SERVICE_TEMPLATE_WELCOME" envDefault:"templates/welcome.html"`
	TemplateStatic               string        `env:"SERVICE_TEMPLATE_STATIC" envDefault:"static/welcome.html"`
	DefaultLocale                string        `env:"SERVICE_DEFAULT_LOCALE" envDefault:"en"`
	UptraceEnable                bool          `env:"UPTRACE_ENABLE" envDefault:"false"`
	UptraceDSN                   string        `env:"UPTRACE_DSN"`
	UptraceName                  string        `env:"UPTRACE_NAME"`
	UptraceDeployment            string        `env:"UPTRACE_DEPLOYMENT"`
	UptraceVersion               string        `env:"UPTRACE_VERSION"`
	RealmID                      int           `env:"REALM_ID" envDefault:"1"`
	RealmRate                    string        `env:"REALM_RATE" envDefault:"x1"`
	RealmFlag                    string        `env:"REALM_FLAG" envDefault:"PvP"`
	RealmRealmlist               string        `env:"REALM_REALMLIST" envDefault:"logon.warhoop.su"`
	DBAuth                       string        `env:"DB_AUTH" envDefault:"root:root@tcp(localhost:3306)/auth?parseTime=true"`
	DBCharacters                 string        `env:"DB_CHARACTERS" envDefault:"root:root@tcp(localhost:3306)/characters?parseTime=true"`
	DBWorld                      string        `env:"DB_WORLD" envDefault:"root:root@tcp(localhost:3306)/world?parseTime=true"`
	DBNexus                      string        `env:"DB_NEXUS" envDefault:"root:root@tcp(localhost:3306)/nexus?parseTime=true"`
	DBVerbose                    bool          `env:"DB_VERBOSE"  envDefault:"false"`
	RedisEnable                  bool          `env:"REDIS_ENABLE" envDefault:"false"`
	RedisDSN                     string        `env:"REDIS_DSN" envDefault:"127.0.0.1:6379"`
	RedisPassword                string        `env:"REDIS_PASSWORD" envDefault:""`
	SoapEnable                   bool          `env:"SOAP_ENABLE" envDefault:"false"`
	SoapHost                     string        `env:"SOAP_HOST"`
	SoapLogin                    string        `env:"SOAP_LOGIN"`
	SoapPassword                 string        `env:"SOAP_PASSWORD"`
	CorsAllowOrigins             string        `env:"CORS_ALLOW_ORIGINS" envDefault:"http://localhost:5173,https://test.warhoop.su"`
	CorsAllowMethods             string        `env:"CORS_ALLOW_METHODS" envDefault:"GET,POST,PUT,PATCH,DELETE,OPTIONS"`
	CorsAllowHeaders             string        `env:"CORS_ALLOW_HEADERS" envDefault:"Authorization,Content-Type"`
	CorsAllowCredentials         bool          `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	CookieName                   string        `env:"COOKIE_NAME" envDefault:"warhoop"`
	CookieDomain                 string        `env:"COOKIE_DOMAIN" envDefault:""`
	CookieHTTPOnly               bool          `env:"COOKIE_HTTP_ONLY" envDefault:"true"`
	CookieSecure                 bool          `env:"COOKIE_SECURE" envDefault:"true"`
	CookieSameSite               string        `env:"COOKIE_SAME_SITE" envDefault:"None"`
	CookieAccessDuration         time.Duration `env:"COOKIE_ACCESS_DURATION" envDefault:"12h"`
	CookieJwtKey                 string        `env:"COOKIE_JWT_KEY" envDefault:"supersecretkey"`
	MailServer                   string        `env:"MAIL_SERVER"`
	MailPort                     string        `env:"MAIL_PORT"`
	MailUser                     string        `env:"MAIL_USER"`
	MailPassword                 string        `env:"MAIL_PASSWORD"`
	MailSanderName               string        `env:"MAIL_SANDER_NAME"`
	MailFolderTemplates          string        `env:"MAIL_FOLDER_TEMPLATES" envDefault:"templates"`
	MailTemplateVerify           string        `env:"MAIL_TEMPLATE_VERIFY" envDefault:"verify.html"`
	MailTemplatePassword         string        `env:"MAIL_TEMPLATE_PASSWORD" envDefault:"password.html"`
	MailTemplatePasswordComplete string        `env:"MAIL_TEMPLATE_PASSWORD_COMPLETE" envDefault:"change_password_complete.html"`
	MailTemplateLogin            string        `env:"MAIL_TEMPLATE_LOGIN" envDefault:"login.html"`
	MailTemplateChangeEmail      string        `env:"MAIL_TEMPLATE_CHANGE_EMAIL" envDefault:"change_email.html"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := env.Parse(&config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		infoLog := log.New(os.Stdout, "INIT\t", log.Ldate|log.Ltime|log.LUTC)
		infoLog.Println("Configuration:", string(configBytes))
	})
	return &config
}

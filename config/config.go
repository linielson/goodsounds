package config

import (
	"fmt"
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

var (
	config *conf
	db     *gorm.DB
)

// Its executed before the main func
func Init() error {
	var err error
	config, err = loadConfig()
	if err != nil {
		return fmt.Errorf("error loading config: ", err)
	}
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: ", err)
	}
	return nil
}

func loadConfig() (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal config", err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, err
}

func GetTokenAuth() *jwtauth.JWTAuth {
	return config.TokenAuth
}

func GetJwtExpiresIn() int {
	return config.JwtExpiresIn
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}

func GetSQLite() *gorm.DB {
	return db
}

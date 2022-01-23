package config

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
)

var ErrConfigFileIsNotSet = errors.New("empty config file")

type Config struct {
	Port        int        `yaml:"port" json:"port,omitempty" envconfig:"PORT"`
	Loglevel    string     `yaml:"loglevel" json:"loglevel,omitempty" envconfig:"LOG_LEVEL"`
	StoragePath string     `yaml:"storage_path" json:"storage_path,omitempty" envconfig:"STORAGE"`
	AuthConfig  AuthConfig `yaml:"auth_config" json:"auth_config"`
	DBConfig    DBConfig   `yaml:"db_config" json:"db_config"`
}

type AuthConfig struct {
	JWTSecret string        `yaml:"jwt_secret" json:"-" envconfig:"JWT_SECRET"`
	JWTTTL    time.Duration `yaml:"jwt_ttl" json:"jwt_ttl,omitempty" envconfig:"JWT_TTL"`
}

type DBConfig struct {
	DBUrl    string `yaml:"db_url" json:"db_url,omitempty" envconfig:"DATABASE_URL"`
	Host     string `yaml:"host" json:"host,omitempty" envconfig:"DB_HOST"`
	User     string `yaml:"user" json:"user,omitempty" envconfig:"DB_USER"`
	Password string `yaml:"password" json:"-" envconfig:"DB_PASSWORD"`
	DBName   string `yaml:"db_name" json:"db_name,omitempty" envconfig:"DB_NAME"`
	Port     int    `yaml:"port" json:"port" envconfig:"DB_PORT"`
}

func (c *Config) ReadFromFile(logger echo.Logger) error {
	configPath := flag.String("config", "", "path yo yaml config")
	flag.Parse()

	if *configPath == "" {
		return ErrConfigFileIsNotSet
	}

	data, err := os.ReadFile(*configPath)
	if err != nil {
		logger.Fatalf("can't read config: %v", err)
	}

	if err = yaml.Unmarshal(data, c); err != nil {
		logger.Fatalf("can't unmarshal config: %v", err)
	}

	c.printInLog(logger)

	return nil
}

func (c *Config) ReadFromEnv(logger echo.Logger) {
	err := envconfig.Process("", c)
	if err != nil {
		logger.Fatalf("failed to read config from env: %v", err)
	}

	c.printInLog(logger)
}

func (c *Config) printInLog(logger echo.Logger) {
	//nolint:errcheck
	jsn, _ := json.Marshal(c)
	logger.Infof("have read config %s", string(jsn))
}

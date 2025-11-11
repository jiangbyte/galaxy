package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	RabbitMQ RabbitMQConfig `yaml:"rabbitmq"`
	Etcd     EtcdConfig     `yaml:"etcd"`
	Monitor  MonitorConfig  `yaml:"monitor"`
	Queue    QueueConfig    `yaml:"queue"`
	JWT      JWTConfig      `yaml:"jwt"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Env     string `yaml:"env"`
	Port    int    `yaml:"port"`
	Mode    string `yaml:"mode"`
}

type ServerConfig struct {
	Web   ServerPortConfig `yaml:"web"`
	Admin ServerPortConfig `yaml:"admin"`
	Judge ServerPortConfig `yaml:"judge"`
	Clone ServerPortConfig `yaml:"clone"`
}

type ServerPortConfig struct {
	Port int `yaml:"port"`
}

type DatabaseConfig struct {
	Postgres PostgresConfig `yaml:"postgres"`
}

type PostgresConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	User            string `yaml:"user"`
	Password        string `yaml:"password"`
	DBName          string `yaml:"dbname"`
	SSLMode         string `yaml:"sslmode"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	PoolSize int    `yaml:"pool_size"`
}

type RabbitMQConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	VHost    string `yaml:"vhost"`
}

type EtcdConfig struct {
	Endpoints   []string `yaml:"endpoints"`
	DialTimeout int      `yaml:"dial_timeout"`
}

type MonitorConfig struct {
	Prometheus PrometheusConfig `yaml:"prometheus"`
	Grafana    GrafanaConfig    `yaml:"grafana"`
}

type PrometheusConfig struct {
	Port int `yaml:"port"`
}

type GrafanaConfig struct {
	Port int `yaml:"port"`
}

type QueueConfig struct {
	JudgeQueue string `yaml:"judge_queue"`
	CloneQueue string `yaml:"clone_queue"`
}

type JWTConfig struct {
	Secret      string `yaml:"secret"`
	ExpireHours int    `yaml:"expire_hours"`
}

var (
	instance *Config
	once     sync.Once
)

func Load(configPath string) *Config {
	once.Do(func() {
		if configPath == "" {
			configPath = getDefaultConfigPath()
		}

		data, err := os.ReadFile(configPath)
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}

		instance = &Config{}
		if err := yaml.Unmarshal(data, instance); err != nil {
			log.Fatalf("Failed to unmarshal config: %v", err)
		}
	})
	return instance
}

func Get() *Config {
	if instance == nil {
		log.Fatal("Config not loaded. Call Load() first.")
	}
	return instance
}

func getDefaultConfigPath() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	configName := "config_" + env + ".yaml"
	return filepath.Join("configs", configName)
}

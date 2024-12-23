package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Username     string
	Password     string
	Host         string
	Port         uint16
	DBName       string
	SSLMode      string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

type Config struct {
	Port    string
	Env     string
	DB      DatabaseConfig
	Limiter struct {
		RPS     float64
		Burst   int
		Enabled bool
	}
}

func LoadConfig() (*DatabaseConfig, error) {
	username, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_USER env variable set")
	}

	password, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_PASSWORD env variable set")
	}

	host, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_HOST env variable set")
	}

	portStr, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_PORT env variable set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("failed to convert port to int: %w", err)
	}

	dbname, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return nil, fmt.Errorf("no POSTGRES_DB env variable set")
	}

	sslmode, ok := os.LookupEnv("POSTGRES_SSLMODE")
	if !ok {
		return nil, fmt.Errorf("no SSLMode env variable set")
	}

  config := &DatabaseConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     uint16(port),
		DBName:   dbname,
		SSLMode:  sslmode,
	}

	flag.IntVar(&config.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&config.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&config.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	return config, nil
}

func (cfg *Config) URL() string {
  // "host=localhost user=admin password=admin dbname=swim-results port=5432 sslmode=disable TimeZone=Europe/Vienna"
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Europe/Vienna",
    cfg.DB.Host,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.DBName,
    cfg.DB.Port,
		cfg.DB.SSLMode,
	)
}

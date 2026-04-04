package config

import "github.com/Akkato47/go-boilerplate/internal/core/common/env"

type Config struct {
	App      AppConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	JWT      JWTConfig
	HTTP     HTTPConfig
	GRPC     GRPCConfig
	Kafka    KafkaConfig
}

type AppConfig struct {
	Port string
	Env  string
}

type PostgresConfig struct {
	URL      string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SslMode  string
}

type RedisConfig struct {
	URL string
}

type JWTConfig struct {
	Secret string
}

type HTTPConfig struct {
	AllowedOrigins []string
	CsrfSecret     string
}

type GRPCConfig struct {
	Port string
}

type KafkaConfig struct {
	Brokers            []string
	TopicUserEvents    string
	TopicNotifications string
	GroupID            string
}

func newAppConfig() *AppConfig {
	return &AppConfig{
		Port: env.GetString("APP_PORT", "8000"),
		Env:  env.GetString("APP_ENV", "development"),
	}
}

func newPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		URL:      env.GetString("POSTGRES_URL", ""),
		Host:     env.GetString("POSTGRES_HOST", "localhost"),
		Port:     env.GetString("POSTGRES_PORT", "5432"),
		User:     env.GetString("POSTGRES_USER", "postgres"),
		Password: env.GetString("POSTGRES_PASSWORD", "postgres"),
		Name:     env.GetString("POSTGRES_NAME", ""),
		SslMode:  env.GetString("POSTGRES_SSL_MODE", "disable"),
	}
}

func newRedisConfig() *RedisConfig {
	return &RedisConfig{
		URL: env.GetString("REDIS_URL", "redis://localhost:6379"),
	}
}

func newJWTConfig() *JWTConfig {
	return &JWTConfig{
		Secret: env.GetString("JWT_SECRET", "secret"),
	}
}

func newHTTPConfig() *HTTPConfig {
	return &HTTPConfig{
		AllowedOrigins: env.GetStrings("ALLOWED_ORIGINS", ",", []string{}),
		CsrfSecret:     env.GetString("CSRF_SECRET", "csrf_secret"),
	}
}

func newGRPCConfig() *GRPCConfig {
	return &GRPCConfig{
		Port: env.GetString("GRPC_PORT", "50051"),
	}
}

func newKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers:            env.GetStrings("KAFKA_BROKERS", ",", []string{"localhost:9092"}),
		TopicUserEvents:    env.GetString("KAFKA_TOPIC_USER_EVENTS", "user.events"),
		TopicNotifications: env.GetString("KAFKA_TOPIC_NOTIFICATIONS", "notification.events"),
		GroupID:            env.GetString("KAKFA_GROUP_ID", "user-service-group"),
	}
}

func NewConfig() *Config {
	return &Config{
		App:      *newAppConfig(),
		Postgres: *newPostgresConfig(),
		Redis:    *newRedisConfig(),
		JWT:      *newJWTConfig(),
		HTTP:     *newHTTPConfig(),
		GRPC:     *newGRPCConfig(),
		Kafka:    *newKafkaConfig(),
	}
}

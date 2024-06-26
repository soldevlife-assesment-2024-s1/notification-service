package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Redis         RedisConfig
	HttpServer    HttpServerConfig
	HttpClient    HttpClientConfig
	Logger        LoggerConfig
	UserService   UserService
	MessageStream MessageStreamConfig
	Email         EmailConfig
	ServiceName   string `envconfig:"service_name"`
	OpenTelemetry OpenTelemetryConfig
}

type EmailConfig struct {
	Server       string `envconfig:"email_server"`
	SmtpPort     string `envconfig:"email_smtp_port"`
	Username     string `envconfig:"email_username"`
	Password     string `envconfig:"email_password"`
	SkipSSL      bool   `envconfig:"email_skip_ssl"`
	EmailAddress string `envconfig:"email_address"`
}

type MessageStreamConfig struct {
	Host           string `envconfig:"message_stream_host"`
	Port           string `envconfig:"message_stream_port"`
	Username       string `envconfig:"message_stream_username"`
	Password       string `envconfig:"message_stream_password"`
	ExchangeName   string `envconfig:"message_stream_exchange_name"`
	PublishTopic   string `envconfig:"message_stream_publish_topic"`
	SubscribeTopic string `envconfig:"message_stream_subscribe_topic"`
	SSL            bool   `envconfig:"message_stream_ssl"`
}

type RedisConfig struct {
	Host            string        `envconfig:"redis_host"`
	Port            string        `envconfig:"redis_port"`
	Username        string        `envconfig:"redis_username"`
	Password        string        `envconfig:"redis_password"`
	DB              int           `envconfig:"redis_db"`
	MaxRetries      int           `envconfig:"redis_max_retries"`
	PoolFIFO        bool          `envconfig:"redis_pool_fifo"`
	PoolSize        int           `envconfig:"redis_pool_size"`
	PoolTimeout     time.Duration `envconfig:"redis_pool_timeout"`
	MinIdleConns    int           `envconfig:"redis_min_idle_conns"`
	MaxIdleConns    int           `envconfig:"redis_max_idle_conns"`
	ConnMaxIdleTime time.Duration `envconfig:"redis_conn_max_idle_time"`
	ConnMaxLifetime time.Duration `envconfig:"redis_conn_max_lifetime"`
}

type HttpClientConfig struct {
	Host                string  `envconfig:"http_client_host"`
	Port                string  `envconfig:"http_client_port"`
	Timeout             int     `envconfig:"http_client_timeout"`
	ConsecutiveFailures int     `envconfig:"http_client_consecutive_failures"`
	ErrorRate           float64 `envconfig:"http_client_error_rate"` // 0.001 - 0.999
	Threshold           int     `envconfig:"http_client_threshold"`
	Type                string  `envconfig:"http_client_type"` // consecutive, error_rate
}

type UserService struct {
	Host string `envconfig:"user_service_host"`
	Port string `envconfig:"user_service_port"`
}

type HttpServerConfig struct {
	Host string `envconfig:"http_server_host"`
	Port string `envconfig:"http_server_port"`
}

type LoggerConfig struct {
	IsVerbose       bool   `envconfig:"logger_is_verbose"`
	LoggerCollector string `envconfig:"logger_logger_collector"`
}
type OpenTelemetryConfig struct {
	Endpoint     string `envconfig:"otel_endpoint"`
	HttpEndpoint string `envconfig:"otel_http_endpoint"`
}

func InitConfig() *Config {
	var Cfg Config

	err := envconfig.Process("user_service", &Cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &Cfg
}

package config

type Config struct {
	HttpPort string `envconfig:"HTTP_API_PORT" default:"8080"`
}

package config

type Config struct {
	HttpPort     string `envconfig:"HTTP_API_PORT" default:"8080"`
	MysqlDsn     string `envconfig:"MYSQL_DSN" default:""`
	MysqlTestDsn string `envconfig:"MYSQL_TEST_DSN" default:""`
}

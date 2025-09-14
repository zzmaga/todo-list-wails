package backend

import "fmt"

// Config содержит конфигурацию приложения
type Config struct {
	AppName string `json:"appName"`

	// Database конфигурация
	Database DatabaseConfig `json:"database"`
}

// DatabaseConfig содержит настройки подключения к базе данных
type DatabaseConfig struct {
	// Type тип базы данных: "file" или "postgres"
	Type string `json:"type"`

	// Postgres настройки для PostgreSQL
	Postgres PostgresConfig `json:"postgres"`
}

// PostgresConfig содержит настройки подключения к PostgreSQL
type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

// GetConnectionString возвращает строку подключения к PostgreSQL
func (c *PostgresConfig) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

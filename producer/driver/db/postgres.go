package db

import (
	"fmt"
	"github/Shitomo/producer/ent"
	"os"

	_ "github.com/lib/pq"
)

type config struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

func (c config) dsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Host, c.Port, c.User, c.DbName, c.Password)
}

func newConfig() config {
	return config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		DbName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func NewClient() (*ent.Client, error) {
	config := newConfig()
	client, err := ent.Open("postgres", config.dsn())
	if err != nil {
		return nil, err
	}
	return client, nil
}

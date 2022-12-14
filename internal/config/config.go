package config

import "yla_parser/pkg/repository"

var dbConfig = repository.Config{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "1234",
	DBName:   "avito_test",
	SSLMode:  "disable",
}

func GetDbConfig() (repository.Config, error) {
	return dbConfig, nil
}

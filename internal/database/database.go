package database

import (
	"account/internal/env"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Db       string `yaml:"db"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Init() *gorm.DB {
	cfg := Config{}
	env.Load("database", &cfg)

	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s sslmode=disable user=%s password=%s",
		cfg.Host,
		cfg.Port,
		cfg.Db,
		cfg.Username,
		cfg.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(
		fmt.Sprintf("Initialize database connection at %s:%d/%s",
			cfg.Host,
			cfg.Port,
			cfg.Db),
	)
	DB = db
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

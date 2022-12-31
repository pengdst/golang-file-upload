package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(conf *EnvConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		conf.DbHost, conf.DbUser, conf.DbPassword, conf.DbName, conf.DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}

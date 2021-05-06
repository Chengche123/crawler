package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysqlGorm(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("cannot init gorm object: %v", err)
	}

	return db, nil
}

func NewMysqlGormWithTable(dsn string, model interface{}) (*gorm.DB, error) {
	db, err := NewMysqlGorm(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create mysql gorm: %v", err)
	}

	if !db.Migrator().HasTable(model) {
		err := db.Migrator().CreateTable(model)
		if err != nil {
			return nil, fmt.Errorf("cannot create new table: %v", err)
		}
	}

	return db, nil
}

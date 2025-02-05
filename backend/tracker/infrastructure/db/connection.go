package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"net/url"
	"os"
	"tracker/infrastructure/entity"
	"tracker/infrastructure/entity/data"
)

type DbConnection struct {
	Db *gorm.DB
}

func NewDBConnection(uri string, maxOpenConns int) *DbConnection {
	u, err := url.Parse(uri)
	if err != nil {
		fmt.Printf("failed parsing input: %+v", err)
		os.Exit(4)
	}

	dsn := fmt.Sprintf("%s@tcp(%s)%s?%sparseTime=true", u.User.String(), u.Host, u.Path, u.RawQuery)
	return NewDBConnectionFromDSN(dsn, maxOpenConns)
}

func NewDBConnectionFromDSN(dsn string, maxOpenConns int) *DbConnection {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("failed to connect database")
		os.Exit(4)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxOpenConns(maxOpenConns)

	fmt.Printf("DB Connection successful!\n")

	if autoMigrateError := db.AutoMigrate(
		&entity.Settings{},
	); autoMigrateError != nil {
		fmt.Printf("Error migrate table: %s", autoMigrateError.Error())
		os.Exit(4)
	}
	fmt.Printf("DB Migration successful!\n")

	fillDbWithData(db)

	return &DbConnection{Db: db}
}

func fillDbWithData(db *gorm.DB) {
	db.Create(data.Settings)
}

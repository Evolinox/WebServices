package api

import (
	"os"
	"tracker/controller/repositories"
	"tracker/infrastructure/db"
)

type DIC struct {
	dbConnection   *db.DbConnection
	settingsRepo   *repositories.SettingsRepository
	productAPIRepo *repositories.ProductAPIRepository
}

func (d *DIC) GetDbConnection() *db.DbConnection {
	if d.dbConnection == nil {
		maxOpenConns := 100
		if "" == os.Getenv("ENV_ENVIRONMENT") {
			maxOpenConns = 25
		}
		d.dbConnection = db.NewDBConnection(os.Getenv("DATABASE_URI"), maxOpenConns)
	}
	return d.dbConnection
}

func (d *DIC) GetProductAPIRepository() *repositories.ProductAPIRepository {
	if d.productAPIRepo == nil {
		d.productAPIRepo = repositories.NewProductAPIRepository("http://localhost:8081")
	}
	return d.productAPIRepo
}

func (d *DIC) GetSettingsRepository() *repositories.SettingsRepository {
	if d.settingsRepo == nil {
		d.settingsRepo = repositories.NewSettingsRepository(d.GetDbConnection().Db)
	}
	return d.settingsRepo
}

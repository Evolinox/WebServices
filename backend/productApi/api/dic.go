package api

import (
	"os"
	"productApi/controller/repositories"
	"productApi/infrastructure/db"
)

type DIC struct {
	dbConnection *db.DbConnection
	productRepo  *repositories.ProductRepository
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

func (d *DIC) GetProductRepository() *repositories.ProductRepository {
	if d.productRepo == nil {
		d.productRepo = repositories.NewProductRepository(d.GetDbConnection().Db)
	}
	return d.productRepo
}

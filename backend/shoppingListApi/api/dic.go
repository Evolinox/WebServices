package api

import (
	"os"
	"shoppingListApi/controller/repositories"
	"shoppingListApi/infrastructure/db"
)

type DIC struct {
	dbConnection     *db.DbConnection
	shoppingListRepo *repositories.ShoppingListRepository
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

func (d *DIC) GetShoppingListRepository() *repositories.ShoppingListRepository {
	if d.shoppingListRepo == nil {
		d.shoppingListRepo = repositories.NewShoppingListRepository(d.GetDbConnection().Db)
	}
	return d.shoppingListRepo
}

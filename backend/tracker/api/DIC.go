package api

import (
	"os"
	"tracker/controller/repositories"
	"tracker/infrastructure/db"
)

type DIC struct {
	dbConnection                 *db.DbConnection
	consumeProductRepo           *repositories.ConsumeProductRepository
	dailyNutritionStatisticsRepo *repositories.NutritionStatisticsRepository
	diaryRepo                    *repositories.DiaryRepository
	settingsRepo                 *repositories.SettingsRepository
	productAPIRepo               *repositories.ProductAPIRepository
	shopListAPIRepo              *repositories.ShopListAPIRepository
	calendarAPIRepo              *repositories.CalendarAPIRepository
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

func (d *DIC) GetConsumeProductRepository() *repositories.ConsumeProductRepository {
	if d.consumeProductRepo == nil {
		d.consumeProductRepo = repositories.NewConsumeProductRepository(d.GetDbConnection().Db)
	}
	return d.consumeProductRepo
}

func (d *DIC) GetNutritionStatisticsRepository() *repositories.NutritionStatisticsRepository {
	if d.dailyNutritionStatisticsRepo == nil {
		d.dailyNutritionStatisticsRepo = repositories.NewNutritionStatisticsRepository(d.GetDbConnection().Db)
	}
	return d.dailyNutritionStatisticsRepo
}

func (d *DIC) GetDiaryRepository() *repositories.DiaryRepository {
	if d.diaryRepo == nil {
		d.diaryRepo = repositories.NewDiaryRepository(d.GetDbConnection().Db)
	}
	return d.diaryRepo
}

func (d *DIC) GetSettingsRepository() *repositories.SettingsRepository {
	if d.settingsRepo == nil {
		d.settingsRepo = repositories.NewSettingsRepository(d.GetDbConnection().Db)
	}
	return d.settingsRepo
}

func (d *DIC) GetProductAPIRepository() *repositories.ProductAPIRepository {
	if d.productAPIRepo == nil {
		d.productAPIRepo = repositories.NewProductAPIRepository("http://product-api:8081")
	}
	return d.productAPIRepo
}

func (d *DIC) GetCalendarAPIRepository() *repositories.CalendarAPIRepository {
	if d.calendarAPIRepo == nil {
		d.calendarAPIRepo = repositories.NewCalendarAPIRepository("http://calendar-api:8083")
	}
	return d.calendarAPIRepo
}

func (d *DIC) GetShopListAPIRepository() *repositories.ShopListAPIRepository {
	if d.shopListAPIRepo == nil {
		d.shopListAPIRepo = repositories.NewShopListAPIRepository("http://shop-list-api:8084")
	}
	return d.shopListAPIRepo
}

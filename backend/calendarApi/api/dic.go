package api

import (
	"calendarApi/controller/repositories"
	"calendarApi/infrastructure/db"
	"os"
)

type DIC struct {
	dbConnection *db.DbConnection
	calendarRepo *repositories.CalendarRepository
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

func (d *DIC) GetCalendarRepository() *repositories.CalendarRepository {
	if d.calendarRepo == nil {
		d.calendarRepo = repositories.NewCalendarRepository(d.GetDbConnection().Db)
	}
	return d.calendarRepo
}

package graph

import (
	"database/sql"
	"time"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/dataloader"
	localDb "github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/db"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DbClient    influxdb2.Client
	CsvDbClient *sql.DB
	Dataloader  dataloader.Dataloader
	QueryApi    api.QueryAPI
}

func NewResolver() *Resolver {
	db, queryApi := localDb.ConnectToDatabase()
	csvDbClient := localDb.ConnectToCsvq()

	return &Resolver{
		DbClient:    db,
		Dataloader:  dataloader.NewDataloader(queryApi),
		QueryApi:    queryApi,
		CsvDbClient: csvDbClient,
	}
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

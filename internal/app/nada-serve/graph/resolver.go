package graph

import (
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/dataloader"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/db"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DbClient   influxdb2.Client
	Dataloader dataloader.Dataloader
	QueryApi   api.QueryAPI
}

func NewResolver() *Resolver {
	db, queryApi := db.ConnectToDatabase()
	return &Resolver{
		DbClient:   db,
		Dataloader: dataloader.NewDataloader(queryApi),
		QueryApi:   queryApi,
	}
}

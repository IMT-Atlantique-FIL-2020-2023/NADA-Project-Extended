package dataloader

import (
	dataloader "github.com/graph-gophers/dataloader/v6"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

func NewDataloader(queryApi api.QueryAPI) Dataloader {
	dataloaderLocal := Dataloader{
		queryApi: queryApi,
	}
	dataloaderLocal.SensorDataloader = dataloader.NewBatchedLoader(dataloaderLocal.getSensors, dataloader.WithCache(&dataloader.NoCache{}))
	dataloaderLocal.MeanValuesDataloader = dataloader.NewBatchedLoader(dataloaderLocal.getMeanData, dataloader.WithCache(&dataloader.NoCache{}))
	return dataloaderLocal
}

type Dataloader struct {
	queryApi             api.QueryAPI
	SensorDataloader     *dataloader.Loader
	MeanValuesDataloader *dataloader.Loader
}

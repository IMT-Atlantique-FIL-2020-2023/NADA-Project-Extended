package dataloader

import (
	"context"
	"fmt"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/db"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/model"
	dataloader "github.com/graph-gophers/dataloader/v6"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type Dataloader struct {
	queryApi         api.QueryAPI
	SensorDataloader *dataloader.Loader
}

func NewDataloader(queryApi api.QueryAPI) Dataloader {
	dataloaderLocal := Dataloader{
		queryApi: queryApi,
	}
	dataloader.NewBatchedLoader(dataloaderLocal.getSensors, dataloader.WithCache(&dataloader.NoCache{}))

	return dataloaderLocal
}
func (r Dataloader) getSensors(c context.Context, airportIds dataloader.Keys) []*dataloader.Result {
	result, err := r.queryApi.Query(c, fmt.Sprintf(`
keys = [%v]

from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> keep(columns: ["airportId", "sensorId"])
    |> filter(fn: (r) => exists r.sensorId and exists r.airportId and contains(value: r.airportId, set: keys))
    |> distinct(column: "sensorId")
	`, db.SanitizeStringArray(airportIds.Keys())))
	var sensors []*dataloader.Result
	if err == nil {
		for result.Next() {
			sensors = append(sensors, &dataloader.Result{Data: &model.Sensor{
				ID: result.Record().Value().(string),
				Measurement: &model.Measurement{
					ID:   "TEST",
					Name: "TEST",
					Unit: "TEST",
				},
			}})
		}
	}
	return sensors
}

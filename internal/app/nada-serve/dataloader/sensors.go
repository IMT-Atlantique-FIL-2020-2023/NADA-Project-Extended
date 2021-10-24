package dataloader

import (
	"context"
	"fmt"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/config"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/db"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/model"
	dataloader "github.com/graph-gophers/dataloader/v6"
)

type SensorDlKey struct {
	Airport *model.Airport
	Sensors []string
}

func (k SensorDlKey) Raw() interface{} {
	return k
}

func (k SensorDlKey) String() string {
	return fmt.Sprintf("%v-%v", k.Airport.ID, k.Sensors)
}

func NewSensorDlKey(airport *model.Airport, sensors ...string) SensorDlKey {
	return SensorDlKey{airport, sensors}
}

func SensorDlAirportKeys(l dataloader.Keys) ([]string, map[string][]string, map[string]*model.Airport) {
	list := make([]string, len(l))
	m := make(map[string][]string, len(l))
	mAiportId := make(map[string]*model.Airport)
	for i := range l {
		v := l[i].Raw().(SensorDlKey)
		list[i] = v.Airport.ID
		m[v.Airport.ID] = v.Sensors
		mAiportId[v.Airport.ID] = v.Airport
	}
	return list, m, mAiportId
}

func (r *Dataloader) getSensors(c context.Context, airportIds dataloader.Keys) []*dataloader.Result {
	airportIdsKeys, sensorsMapId, mAiportId := SensorDlAirportKeys(airportIds)
	query := fmt.Sprintf(`
// Load in batch unique sensors for each airport specified by keys, filter by airport
import "dict"

keys = [%v]
sensorsIds = ["":[],%v]

filterFunc = (r) => {
    item = dict.get(dict: sensorsIds, key: r.airportId, default: [])
    len = length(arr: item) 
    return (len == 0 or contains(value: r.sensorId, set: item))
}

from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> filter(fn: (r) => exists r.sensorId and exists r.airportId)
    |> filter(fn: filterFunc)
    |> distinct(column: "sensorId")
    |> keep(columns: ["airportId", "_measurement", "_value"])
    |> rename(columns: {airportId: "_field"})`,
		db.SanitizeStringArray(airportIdsKeys),
		db.SanitizeMapArrayString(sensorsMapId))
	result, err := r.queryApi.Query(c, query)
	sensors := make(map[string][]*model.Sensor, 0)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		measurement := result.Record().Measurement()
		airportId := result.Record().Field()
		sensorConfig := config.GetSensorConfig(measurement)
		sensors[airportId] = append(sensors[airportId], &model.Sensor{
			ID:      result.Record().Value().(string),
			Airport: mAiportId[airportId],
			Measurement: &model.Measurement{
				ID:   measurement,
				Name: sensorConfig.MeasurementName,
				Unit: sensorConfig.MeasurementUnit,
			},
		})
	}

	var dataloaderResults []*dataloader.Result
	for _, v := range airportIdsKeys {
		dataloaderResults = append(dataloaderResults, &dataloader.Result{
			Data: sensors[v],
		})
	}
	return dataloaderResults
}

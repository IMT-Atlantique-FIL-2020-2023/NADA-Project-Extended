package dataloader

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/db"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/model"
	"github.com/google/uuid"
	dataloader "github.com/graph-gophers/dataloader/v6"
)

var space uuid.UUID = uuid.MustParse("dde99ccc-5765-467b-b20f-8bc3f9f14d16")

type MeanDataDlKey struct {
	AirportId      string
	Start          time.Time
	End            time.Time
	EveryValue     string
	DiscretizeMode *model.MeanMeasureMode
	Sensor         *model.Sensor
}

func (k MeanDataDlKey) Raw() interface{} {
	return k
}

func (k MeanDataDlKey) String() string {
	return fmt.Sprintf("%v-%v", k.AirportId, k.Sensor.ID)
}

func MeanDataDlKeys(l dataloader.Keys) ([]string, map[string][]string, map[string]map[string]*model.Sensor, time.Time, time.Time, string, *model.MeanMeasureMode) {
	list := make([]string, len(l))
	m := make(map[string][]string, len(l))
	m1 := make(map[string]map[string]*model.Sensor, len(l))

	var start time.Time
	var end time.Time
	var everyValue string
	var discretizeMode *model.MeanMeasureMode
	for i := range l {
		v := l[i].Raw().(MeanDataDlKey)
		list[i] = v.AirportId
		m[v.AirportId] = append(m[v.AirportId], v.Sensor.ID)
		if _, ok := m1[v.AirportId]; !ok {
			m1[v.AirportId] = make(map[string]*model.Sensor)
		}
		m1[v.AirportId][v.Sensor.ID] = v.Sensor
		start = v.Start
		end = v.End
		everyValue = v.EveryValue
		discretizeMode = v.DiscretizeMode
	}
	return list, m, m1, start, end, everyValue, discretizeMode
}
func getDurationEveryValue(everyValue string, mode model.MeanMeasureMode) string {
	if mode == model.MeanMeasureModeFluxDuration {
		return db.SanitizeDuration(everyValue)
	}
	return "1h"
}
func getDiscritzeIntoValue(everyValue string, mode model.MeanMeasureMode) int32 {
	i, err := strconv.ParseInt(everyValue, 10, 32)
	if err != nil {
		return 1
	}
	return int32(i)
}

func (r *Dataloader) getMeanData(c context.Context, airportIds dataloader.Keys) []*dataloader.Result {
	airportIdsKeys, sensorsMapId, sensorMap, start, end, everyValue, mode := MeanDataDlKeys(airportIds)
	query := fmt.Sprintf(`
import "dict"

keys = [%v]
sensorsIds = ["":[],%v]
timeRangeStart = %v
timeRangeEnd = %v
everyValue = %v
discritzeInto = %v
discretizeMode = %t
duration = if discretizeMode then duration(v: (int(v: timeRangeEnd) - int(v:timeRangeStart)) / discritzeInto) else everyValue

filterFunc = (r) => {
    item = dict.get(dict: sensorsIds, key: r.airportId, default: [])
    return (contains(value: r.sensorId, set: item))
}

from(bucket: "nada-bucket")
    |> range(start: timeRangeStart, stop: timeRangeEnd)
    |> filter(fn: (r) => exists r.sensorId and exists r.airportId)
    |> filter(fn: filterFunc)
    |> window(every: duration,  createEmpty: false)
    |> mean()
    |> keep(columns: ["airportId", "_measurement", "_value", "_start", "_stop", "_time", "sensorId"])
    |> rename(columns: {airportId: "_field"})`,
		db.SanitizeStringArray(airportIdsKeys),
		db.SanitizeMapArrayString(sensorsMapId),
		db.SanitizeTime(start),
		db.SanitizeTime(end),
		getDurationEveryValue(everyValue, *mode),
		getDiscritzeIntoValue(everyValue, *mode),
		*mode == model.MeanMeasureModeForInterval,
	)
	result, err := r.queryApi.Query(c, query)
	measureData := make(map[string]map[string]map[string][]*model.MeasureMeanData, 0)
	if err != nil {
		panic(err)
	}
	for _, v := range airportIdsKeys {
		measureData[v] = make(map[string]map[string][]*model.MeasureMeanData)
	}

	for result.Next() {
		airportId := result.Record().Field()
		sensorId := result.Record().ValueByKey("sensorId").(string)
		measurement := result.Record().Measurement()

		start := result.Record().Start()
		end := result.Record().Stop()
		sensor := sensorMap[airportId][sensorId]
		if _, ok := measureData[airportId][sensorId]; !ok {
			measureData[airportId][sensorId] = make(map[string][]*model.MeasureMeanData)
		}
		measureData[airportId][sensorId][measurement] = append(measureData[airportId][sensorId][measurement], &model.MeasureMeanData{
			ID:        uuid.NewMD5(space, []byte(start.String()+end.String()+airportId+sensorId+measurement)).String(),
			Sensor:    sensor,
			Airport:   sensor.Airport,
			StartDate: start,
			EndDate:   end,
			Value:     result.Record().Value().(float64),
		})
	}

	var dataloaderResults []*dataloader.Result
	for _, v := range airportIds {
		k := v.Raw().(MeanDataDlKey)
		dataloaderResults = append(dataloaderResults, &dataloader.Result{
			Data: measureData[k.AirportId][k.Sensor.ID][k.Sensor.Measurement.ID],
		})
	}
	return dataloaderResults
}

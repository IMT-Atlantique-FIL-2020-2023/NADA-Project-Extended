package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/db"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/generated"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/model"
	"github.com/graph-gophers/dataloader/v6"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *airportResolver) Sensors(ctx context.Context, obj *model.Airport) ([]*model.Sensor, error) {
	thunk := r.Dataloader.SensorDataloader.Load(ctx, dataloader.StringKey(obj.ID))
	result, err := thunk()
	return result.([]*model.Sensor), err
}

func (r *airportResolver) GetMeanMeasures(ctx context.Context, obj *model.Airport, day *time.Time) ([]*model.MeasureMeanData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *airportResolver) GetSubsetOfSensors(ctx context.Context, obj *model.Airport, sensorIds []string) ([]*model.Sensor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Airports(ctx context.Context) ([]*model.Airport, error) {
	result, err := r.Resolver.QueryApi.Query(ctx,
		`
from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> keep(columns: ["airportId"])
    |> filter(fn: (r) => exists r.airportId)
    |> distinct(column: "airportId")`)
	airports := make([]*model.Airport, 0, 10)
	if err == nil {
		for result.Next() {
			airports = append(airports, &model.Airport{
				ID:   result.Record().Value().(string),
				Name: "",
			})
		}
	}
	return airports, err
}

func (r *queryResolver) GetAirportByID(ctx context.Context, id string) (*model.Airport, error) {
	result, err := r.Resolver.QueryApi.Query(ctx,
		fmt.Sprintf(`
from(bucket: "nada-bucket")
    |> range(start: -1y, stop: -0m)
    |> keep(columns: ["airportId"])
    |> filter(fn: (r) => r.airportId == "%v")
    |> rename(columns: {airportId: "_value"})
    |> first()`, db.SanitizeString(id)))
	if !result.Next() {
		return nil, gqlerror.Errorf("Airport not found %v", id)
	}
	airport := &model.Airport{
		ID:   result.Record().Value().(string),
		Name: "",
	}

	return airport, err
}

func (r *sensorResolver) GetMeanMeasureInterval(ctx context.Context, obj *model.Sensor, start time.Time, end time.Time, discretize *int, discretizeMode *model.MeanMeasureMode) ([]*model.MeasureMeanData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Airport returns generated.AirportResolver implementation.
func (r *Resolver) Airport() generated.AirportResolver { return &airportResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Sensor returns generated.SensorResolver implementation.
func (r *Resolver) Sensor() generated.SensorResolver { return &sensorResolver{r} }

type airportResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type sensorResolver struct{ *Resolver }

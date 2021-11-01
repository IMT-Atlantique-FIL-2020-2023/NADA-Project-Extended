package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/config"
	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/graph/model"
	_ "github.com/mithrandie/csvq-driver"
	"github.com/rs/zerolog/log"
)

func ConnectToCsvq() *sql.DB {
	db, err := sql.Open("csvq", "") // csvq path default to cwd
	if err != nil {
		log.Fatal().Err(err).Msg("Error loading csvq")
	}
	return db
}

func ExecQuery(ctx context.Context, db *sql.DB, iataCode string) *model.Airport {
	var (
		ident        sql.NullString
		typeAirport  sql.NullString
		name         sql.NullString
		elevation_ft sql.NullInt32
		continent    sql.NullString
		iso_country  sql.NullString
		iso_region   sql.NullString
		municipality sql.NullString
		gps_code     sql.NullString
		iata_code    sql.NullString
		local_code   sql.NullString
		coordinates  sql.NullString
	)
	query := fmt.Sprintf("SELECT ident, type, name, elevation_ft, continent, iso_country, iso_region, municipality, gps_code, iata_code, local_code, coordinates FROM `%v` WHERE iata_code = \"%v\" LIMIT 1", config.CurrentConfig.CsvAirportFile, iataCode)
	err := db.QueryRowContext(ctx, query).Scan(&ident, &typeAirport, &name, &elevation_ft, &continent, &iso_country, &iso_region, &municipality, &gps_code, &iata_code, &local_code, &coordinates)

	if err != nil {
		log.Error().Err(err).Msgf("Error loading from CSV file airport %v", iataCode)
		return &model.Airport{
			ID: iataCode,
		}
	}
	coords := strings.Split(coordinates.String, ",")
	lon, _ := strconv.ParseFloat(strings.TrimSpace(coords[0]), 64)
	lat, _ := strconv.ParseFloat(strings.TrimSpace(coords[len(coords)-1]), 64)
	elevation := &elevation_ft.Int32
	converted := int(*elevation)
	return &model.Airport{
		Name:         &name.String,
		ElevationFt:  &converted,
		Continent:    &continent.String,
		IsoCountry:   &iso_country.String,
		Municipality: &municipality.String,
		IsoRegion:    &iso_region.String,
		GpsCode:      &gps_code.String,
		LocalCode:    &local_code.String,
		Lon:          &lon,
		Lat:          &lat,
		ID:           iataCode,
	}

}

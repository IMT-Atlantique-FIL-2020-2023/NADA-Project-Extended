package db

import (
	"context"
	"fmt"
	"time"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/config"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/rs/zerolog/log"
)

func ConnectToDatabase() (influxdb2.Client, api.QueryAPI) {
	client := influxdb2.NewClientWithOptions(
		config.CurrentConfig.InfluxDb.Host,
		config.CurrentConfig.InfluxDb.Token,
		influxdb2.DefaultOptions(),
	)

	result, err := client.Health(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot check database health")
	}
	if result.Status == domain.HealthCheckStatusPass {
		log.Info().Msgf("Database HealthCheckStatusPass - %v %v %v", result.Name, *result.Version, *result.Message)
	}
	ticker := time.NewTicker(60 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				_, err := client.Health(context.Background())
				if err != nil {
					log.Error().Err(err).Msg("Cannot check database health")
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	queryApi := client.QueryAPI(config.CurrentConfig.InfluxDb.Org)
	return client, queryApi
}

func SanitizeString(value string) string {
	if value == "" {
		return ""
	}
	i := 0
	retVal := ""
	prepareRetVal := func() {
		if retVal == "" {
			retVal = value[:i]
		}
	}
	for ; i < len(value); i++ {
		c := value[i]
		switch c {
		case '\r':
			prepareRetVal()
			retVal += "\\r"

		case '\n':
			prepareRetVal()
			retVal += "\\n"
		case '\t':
			prepareRetVal()
			retVal += "\\t"
		case '"', '\\':
			prepareRetVal()
			retVal = retVal + string('\\') + string(c)
		case '$':
			if i+1 < len(value) && value[i+1] == '{' {
				prepareRetVal()
				i++
				retVal += "'\\${'"
			}
			if retVal != "" {
				retVal += string(c)
			}
		default:
			if retVal != "" {
				retVal += string(c)
			}

		}
	}
	if retVal != "" {
		return retVal
	}
	return value
}

func SanitizeStringArray(strs []string) string {
	str := ""
	for _, v := range strs {
		str += "\"" + SanitizeString(v) + "\","
	}
	return str
}

func SanitizeMapArrayString(m map[string][]string) string {
	str := ""
	for k, v := range m {
		str += fmt.Sprintf("\"%v\":[%v],", k, SanitizeStringArray(v))
	}
	return str
}

func SanitizeTime(t time.Time) string {
	return fmt.Sprintf("time(v: \"%v\")", t.Format(time.RFC3339))
}

func SanitizeDuration(d string) string {
	return fmt.Sprintf(`duration(v: "%v")`, SanitizeString(d))
}

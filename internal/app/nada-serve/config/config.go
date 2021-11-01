package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/IMT-Atlantique-FIL-2020-2023/NADA-extended/internal/app/nada-serve/config/dotenv"
	"github.com/go-playground/validator/v10"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var k = koanf.New(".")
var mutexConfig sync.Mutex

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type NadaServeConfig struct {
	Port           int                     `koanf:"port" validate:"required"`
	LogLevel       zerolog.Level           `koanf:"loglevel" validate:"omitempty,required"`
	InfluxDb       InfluxbConfig           `koanf:"influxdb" validate:"required"`
	CsvAirportFile string                  `koanf:"csvairportfile" validate:"required"`
	AllowedOrigins []string                `koanf:"allowedorigins" validate:"required"`
	Sensors        map[string]SensorConfig `koanf:"sensors" validate:"required,min=1"` // we need at least one sensor available in config
}

type InfluxbConfig struct {
	Host   string `koanf:"host" validate:"required"`
	Token  string `koanf:"token" validate:"required"`
	Org    string `koanf:"org" validate:"required"`
	Bucket string `koanf:"bucket" validate:"required"`
}

type SensorConfig struct {
	MeasurementName string `koanf:"name" validate:"required"`
	MeasurementUnit string `koanf:"unit" validate:"required"`
}

var CurrentConfig NadaServeConfig
var watcherAdded bool = false

func watch(f *file.File) func(event interface{}, err error) {
	return func(event interface{}, err error) {
		if err != nil {
			log.Error().Msgf("watch error: %v", err)
			return
		}

		log.Info().Msg("config changed. Reloading ...")
		LoadConfig()

	}

}

func init() {
	validate = validator.New()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func LoadConfig() {
	mutexConfig.Lock()
	defer mutexConfig.Unlock()
	// Load default values using the structs provider.
	// We provide a struct along with the struct tag `koanf` to the
	// provider.
	k.Load(structs.Provider(NadaServeConfig{
		Port:     9090,
		LogLevel: zerolog.DebugLevel,
	}, "koanf"), nil)
	// Load YAML config and merge into the previously loaded config (because we can).
	f := file.Provider("nada-serve.yaml")
	f2 := file.Provider("nada-serve.yml")

	k.Load(f, yaml.Parser())
	k.Load(f2, yaml.Parser())

	d := file.Provider("nada-serve.env")
	k.Load(d, dotenv.Parser())
	if !watcherAdded {
		d.Watch(watch(d))
		f.Watch(watch(f))
		f2.Watch(watch(f2))
	}
	// Load environment variables and merge into the loaded config.
	// "MYVAR" is the prefix to filter the env vars by.
	// "." is the delimiter used to represent the key hierarchy in env vars.
	// The (optional, or can be nil) function can be used to transform
	// the env var names, for instance, to lowercase them.
	//
	// For example, env vars: MYVAR_TYPE and MYVAR_PARENT1_CHILD1_NAME
	// will be merged into the "type" and the nested "parent1.child1.name"
	// keys in the config file here as we lowercase the key,
	// replace `_` with `.` and strip the MYVAR_ prefix so that
	// only "parent1.child1.name" remains.
	k.Load(env.ProviderWithValue("NADA_SERVE_", ".", dotenv.TransformKeyWithValue), nil)
	if err := k.UnmarshalWithConf("", &CurrentConfig, koanf.UnmarshalConf{
		DecoderConfig: &mapstructure.DecoderConfig{
			DecodeHook: mapstructure.ComposeDecodeHookFunc(
				mapstructure.StringToTimeDurationHookFunc()),
			Metadata:         nil,
			Result:           &CurrentConfig,
			WeaklyTypedInput: true,
		},
	}); err != nil {
		log.Fatal().Err(err)
	}
	if err := validate.Struct(&CurrentConfig); err != nil {
		log.Fatal().Err(err).Msg("Error incorrect config")
	}
	zerolog.SetGlobalLevel(CurrentConfig.LogLevel)
	log.Debug().Fields(k.All()).Msg("Koanf loaded:")
	log.Debug().Msgf("Config loaded: %v", CurrentConfig)

	watcherAdded = true
}

func GetSensorConfig(sensorType string) SensorConfig {
	if v, ok := CurrentConfig.Sensors[sensorType]; ok {
		return v
	}
	return SensorConfig{
		MeasurementName: fmt.Sprintf("Unknown sensor type %v", sensorType),
		MeasurementUnit: "Unknown unit",
	}
}

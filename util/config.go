package util

import (
	"errors"
	"os"
	"time"

	"gitlab.com/karte/mongo-lib/mserver"
)

// Config is used to hold information about the configuration of the Karte server.
type Config struct {
	// ServiceName
	ServiceName string

	// Version
	Version string

	// ServerURL is the full URL for the root of the server. This may be used
	// by other middleware to compute redirect URLs
	ServerURL string

	// IndexConfigPath is the path to an indexes.conf configuration file, specifying
	// what mongo indexes the server should create (or verify) on startup
	IndexConfigPath string

	// DatabaseHost is the url of the running mongo instance to use  database.
	DatabaseUsername string

	// DatabaseHost is the url of the running mongo instance to use  database.
	DatabasePassword string

	// DatabaseHost is the url of the running mongo instance to use  database.
	DatabaseHost string

	// DatabasePort is the url of the running mongo instance to use  database.
	DatabasePort string

	// DatabaseName is the name of the mongo database used  database.
	// Typically this will be the default DatabaseName "karte-health-records".
	DatabaseName string

	// DatabaseSocketTimeout is the amount of time the mgo driver will wait for a response
	// from mongo before timing out.
	DatabaseSocketTimeout time.Duration

	// DatabaseOpTimeout is the amount of time GoKarte will wait before killing a long-running
	// database process. This defaults to a reasonable upper bound for slow, pipelined queries: 30s.
	DatabaseOpTimeout time.Duration

	// DatabaseKillOpPeriod is the length of time between scans of the database to kill long-running ops.
	DatabaseKillOpPeriod time.Duration

	// CountTotalResults toggles whether the searcher should also get a total
	// count of the total results of a search. In practice this is a performance hit
	// for large datasets.
	CountTotalResults bool

	// EnableCISearches toggles whether the mongo searches uses regexes to maintain
	// case-insesitivity when performing searches on string fields, codes, etc.
	EnableCISearches bool

	// ReadOnly toggles whether the server is in read-only mode. In read-only
	// mode any HTTP verb other than GET, HEAD or OPTIONS is rejected.
	ReadOnly bool

	// Debug toggles debug-level logging.
	Debug bool

	// Local where we serve http for graphql
	Local string

	// LogFormat
	LogFormat string

	// AppPort where we serve http for graphql
	AppPort string

	//AuthAud URL
	AuthAud string

	//AuthISS URL
	AuthISS string
}

var DefaultConfig = Config{
	ServerURL:             "",
	IndexConfigPath:       "db/indexes.hr.conf",
	DatabaseHost:          "localhost",
	DatabasePort:          "27017",
	DatabaseName:          "karte_health_records",
	DatabaseSocketTimeout: 2 * time.Minute,
	DatabaseOpTimeout:     90 * time.Second,
	DatabaseKillOpPeriod:  10 * time.Second,
	EnableCISearches:      true,
	CountTotalResults:     true,
	ReadOnly:              false,
	Debug:                 true,
	Local:                 "DEV",
	ServiceName:           "HealthRecord_Repository",
	Version:               "0.0.1",
	LogFormat:             "%{color}%{time:2006/01/02 15:04:05 -07:00 MST} [%{level:.6s}] %{shortfile} : %{color:reset}%{message}",
	AppPort:               "5000",
	AuthAud:               "https://healthrecord-repository.karte.io",
	AuthISS:               "https://karte-dev.auth0.com/",
}

func SetupConfig() Config {
	config := DefaultConfig

	val, bool := os.LookupEnv("ServiceName")
	if bool {
		config.ServiceName = val
	}
	val, bool = os.LookupEnv("Version")
	if bool {
		config.Version = val
	}
	val, bool = os.LookupEnv("DATABASE_USERNAME")
	if bool {
		config.DatabaseUsername = val
	}
	val, bool = os.LookupEnv("DATABASE_PASSWORD")
	if bool {
		config.DatabasePassword = val
	}
	val, bool = os.LookupEnv("DATABASE_HOST")
	if bool {
		config.DatabaseHost = val
	}
	val, bool = os.LookupEnv("DATABASE_PORT")
	if bool {
		config.DatabasePort = val
	}
	val, bool = os.LookupEnv("DATABASE_NAME")
	if bool {
		config.DatabaseName = val
	}
	val, bool = os.LookupEnv("EnableCISearches")
	if bool {
		config.EnableCISearches = strToBool(val)
	}
	val, bool = os.LookupEnv("ReadOnly")
	if bool {
		config.ReadOnly = strToBool(val)
	}
	val, bool = os.LookupEnv("Debug")
	if bool {
		config.Debug = strToBool(val)
	}
	val, bool = os.LookupEnv("LOCAL")
	if bool {
		config.Local = val
	}
	val, bool = os.LookupEnv("LogFormat")
	if bool {
		config.LogFormat = val
	}
	val, bool = os.LookupEnv("AppPort")
	if bool {
		config.AppPort = val
	}
	val, bool = os.LookupEnv("AUTH_AUD")
	if bool {
		config.AuthAud = val
	}
	val, bool = os.LookupEnv("AUTH_ISS")
	if bool {
		config.AuthISS = val
	}

	return config
}

func strToBool(str string) bool {
	if str == "true" || str == "TRUE" {
		return true
	} else if str == "FALSE" || str == "false" {
		return false
	} else {
		panic(errors.New("STRING IS NOT OF TYPE BOOLEAN"))
		return false
	}
}

func (c *Config) MongoConfig() mserver.Config {
	return mserver.Config{
		ServerURL:             c.ServerURL,
		IndexConfigPath:       c.IndexConfigPath,
		DatabaseHost:          c.DatabaseHost + ":" + c.DatabasePort,
		DatabaseName:          c.DatabaseName,
		DatabaseSocketTimeout: c.DatabaseSocketTimeout,
		DatabaseOpTimeout:     c.DatabaseOpTimeout,
		DatabaseKillOpPeriod:  c.DatabaseKillOpPeriod,
		EnableCISearches:      c.EnableCISearches,
		CountTotalResults:     c.CountTotalResults,
		ReadOnly:              c.ReadOnly,
		Debug:                 c.Debug,
	}
}

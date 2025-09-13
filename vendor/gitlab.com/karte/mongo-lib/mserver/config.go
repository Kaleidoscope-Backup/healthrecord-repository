package mserver

import (
	"time"
)

// Config is used to hold information about the configuration of the Karte server.
type Config struct {
	// ServerURL is the full URL for the root of the server. This may be used
	// by other middleware to compute redirect URLs
	ServerURL string

	// IndexConfigPath is the path to an indexes.conf configuration file, specifying
	// what mongo indexes the server should create (or verify) on startup
	IndexConfigPath string

	// DatabaseHost is the url of the running mongo instance to use for the fhir database.
	DatabaseHost string

	// DatabaseName is the name of the mongo database used for the fhir database.
	// Typically this will be the default DatabaseName "fhir".
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
}

// DefaultConfig is the default server configuration
var DefaultConfig = Config{
	ServerURL:             "",
	IndexConfigPath:       "db/indexes.conf",
	DatabaseHost:          "localhost:27017",
	DatabaseName:          "karte",
	DatabaseSocketTimeout: 2 * time.Minute,
	DatabaseOpTimeout:     90 * time.Second,
	DatabaseKillOpPeriod:  10 * time.Second,
	EnableCISearches:      true,
	CountTotalResults:     true,
	ReadOnly:              false,
	Debug:                 false,
}

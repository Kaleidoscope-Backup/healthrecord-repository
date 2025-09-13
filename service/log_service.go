package service

import (
	"os"

	"github.com/Kaleidoscope-Backup/healthrecord-repository/util"
	"github.com/op/go-logging"
)

// NewLogger creates a new logger for Karte
func NewLogger(config *util.Config) *logging.Logger {
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter(config.LogFormat)
	backendFormatter := logging.NewBackendFormatter(backend, format)

	backendLeveled := logging.AddModuleLevel(backendFormatter)
	backendLeveled.SetLevel(logging.INFO, "")
	if config.Debug {
		backendLeveled.SetLevel(logging.DEBUG, "")
	}

	logging.SetBackend(backendLeveled)
	logger := logging.MustGetLogger(config.ServiceName)
	return logger
}

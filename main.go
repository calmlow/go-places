package main

import (
	"log"
	"os"
	"path"

	"github.com/calmlow/go-places/internal/ui"
)

const (
	PAGE1_MAIN         = "main"
	PAGE2_REPO_HOME    = "repoHome"
	PAGE3_CONTEXT_HOME = "contextHome"
	COLS_MAX           = 435
)

func main() {
	logFile, err := setupLogging()
	if err != nil {
		log.Panicf("Problem initializing the logging system: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	ui.RunGui()
}

func setupLogging() (*os.File, error) {
	stateEnv := os.Getenv("XDG_STATE_HOME")
	if stateEnv == "" {
		stateEnv = path.Join(os.Getenv("HOME"), ".local", "state")
	}
	logFileName := path.Join(stateEnv, "go-places.log")
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	return os.OpenFile(logFileName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
}

package main

import (
	"time"

	"github.com/devbird007/logging"
)

func main() {
	logger := logging.New(time.RFC3339, true)

	logger.Log("info", "starting up service")
	logger.Log("warning", "no tasks found")
	logger.Log("error", "exiting: no work performed")

}

package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.TraceLevel)

	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"file":       "image.jpg",
		"size_bytes": 132932,
	}).Info("upload successful!")

	log.Trace("Hello from Logrus!")
	log.Debug("Hello from Logrus!")
	log.Info("Hello from Logrus!")
	log.Warn("Hello from Logrus!")
	log.Error("Hello from Logrus!")
	log.Fatal("Hello from Logrus!")
	log.Panic("Hello from Logrus!")
}

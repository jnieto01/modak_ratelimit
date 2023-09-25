package logger

import (

	"github.com/sirupsen/logrus"
	"os"
)

var (
	log = logrus.New()
)

func init() {
	setupOut()
}


func setupOut() {
	log.Out = os.Stdout
	level := logrus.DebugLevel
	if os.Getenv("GO_ENV") == "PROD" {

		level = logrus.InfoLevel
	}

	log.SetLevel(level)
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
}


func Info(message string) {
	log.Info(message)
}

func Error(message string, err error) {
	log.WithError(err).Error(message)
}

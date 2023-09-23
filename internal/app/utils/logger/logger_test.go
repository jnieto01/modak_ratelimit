package logger

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"errors"
	"bytes"
	"os"
	"github.com/sirupsen/logrus"
)



func TestInitialization(t *testing.T) {

    setupOut()

    if os.Getenv("GO_ENV") == "PROD" {
        assert.Equal(t, logrus.InfoLevel, log.GetLevel())
    } else {
        assert.Equal(t, logrus.DebugLevel, log.GetLevel())
    }

    formatter, ok := log.Formatter.(*logrus.TextFormatter)
    assert.True(t, ok)
    assert.True(t, formatter.FullTimestamp)
    assert.Equal(t, "2006-01-02 15:04:05", formatter.TimestampFormat)
}


func TestSetupOut(t *testing.T) {

    setupOut()

    if os.Getenv("GO_ENV") == "PROD" {
        assert.Equal(t, logrus.InfoLevel, log.GetLevel())
    } else {
        assert.Equal(t, logrus.DebugLevel, log.GetLevel())
    }

    formatter, ok := log.Formatter.(*logrus.TextFormatter)
    assert.True(t, ok)
    assert.True(t, formatter.FullTimestamp)
    assert.Equal(t, "2006-01-02 15:04:05", formatter.TimestampFormat)
}

func TestInfo(t *testing.T) {

    oldOutput := log.Out
    defer func() { log.Out = oldOutput }()

	var buf bytes.Buffer
    log.Out = &buf

    Info("This is an info message")

	logOutput := buf.String()
	assert.Contains(t, logOutput, "This is an info message")

}


func TestError(t *testing.T) {

    oldOutput := log.Out
    defer func() { log.Out = oldOutput }()

	var buf bytes.Buffer
    log.Out = &buf

    err := errors.New("Test error")
    Error("This is an error message", err)


	logOutput := buf.String()
	assert.Contains(t, logOutput, "This is an error message")
}


func TestIntegration(t *testing.T) {

	var buf bytes.Buffer
    log.Out = &buf

	Info("Testing started")

	logOutput := buf.String()
	assert.Contains(t, logOutput, "Testing started")


	File := "archivo.txt"
    file, err := os.Open(File)
    if err != nil {
		Error("Fail to open file", err)
		logOutput = buf.String()
		assert.Contains(t, logOutput, "Fail to open file")
    }else{
		assert.Contains(t, logOutput, "Testing started")
	}
    defer file.Close() 

}


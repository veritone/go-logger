package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogLevel(t *testing.T) {
	assert.Equal(t, InfoLevel, GetLogLevel("info"))
	assert.Equal(t, ErrorLevel, GetLogLevel("ERROR"))
}

func TestStdout(t *testing.T) {
	log, err := NewDefault()
	assert.Equal(t, nil, err)

	log.Info("Testing logger Info()")
	log.Error("Testing ", "logger ", "Error()")
	log.Warnf("Testing logger Warnf() %s", "ok")
}

func TestFields(t *testing.T) {
	log, err := NewDefault()
	assert.Equal(t, nil, err)

	fields := make(map[string]interface{})
	fields["key1"] = "value1"
	fields["key2"] = 2

	log.InfoWithFields(fields, "Testing fields")
}

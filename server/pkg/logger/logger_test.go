package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	// --- Arrange ---

	// pre-test setup to ensure logger sets up proper read/write
	originalStderr := os.Stderr
	originalTimeFormat := zerolog.TimeFieldFormat

	r, w, err := os.Pipe()
	require.NoError(t, err, "Failed to create os pipe for test")

	os.Stderr = w

	t.Cleanup(func() {
		os.Stderr = originalStderr
		zerolog.TimeFieldFormat = originalTimeFormat
	})

	// Write first log to above pipe
	deploymentName := "test-deployment"
	_ = NewLogger(deploymentName)
	err = w.Close()
	require.NoError(t, err)

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	require.NoError(t, err)

	// Assert on structure of log message
	// for predictable fields, assert content
	// for unpredictable fields (UUID), assert shape
	var logOutput map[string]interface{}
	err = json.Unmarshal(buf.Bytes(), &logOutput)
	require.NoError(t, err, "Failed to unmarshal log output from captured stderr")

	assert.Equal(t, "info", logOutput["level"])
	assert.Equal(t, "Initializing logger", logOutput["message"])
	assert.Equal(t, deploymentName, logOutput["deployment"])

	assert.Contains(t, logOutput, "time")
	_, isTimeNumber := logOutput["time"].(float64)
	assert.True(t, isTimeNumber, "time field should be a number")

	assert.Contains(t, logOutput, "deployment_id")
	id, isString := logOutput["deployment_id"].(string)
	assert.True(t, isString, "deployment_id field should be a string")
	_, err = uuid.Parse(id)
	assert.NoError(t, err, "deployment_id should be a valid UUID")
}

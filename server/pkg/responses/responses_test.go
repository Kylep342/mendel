package responses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTest creates a gin.Context and an httptest.ResponseRecorder for testing.
func setupTest() (*httptest.ResponseRecorder, *gin.Context) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return w, c
}

func TestRespondData(t *testing.T) {
	t.Run("with struct data", func(t *testing.T) {
		// Arrange
		w, c := setupTest()
		type testStruct struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		sampleData := testStruct{ID: 1, Name: "Test Item"}
		expectedCode := http.StatusOK

		RespondData(c, sampleData, expectedCode)
		assert.Equal(t, expectedCode, w.Code)

		var response map[string]testStruct
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, sampleData, response["data"])
	})

	t.Run("with string data and different status code", func(t *testing.T) {
		w, c := setupTest()
		sampleData := "Item created successfully"
		expectedCode := http.StatusCreated

		RespondData(c, sampleData, expectedCode)
		assert.Equal(t, expectedCode, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, sampleData, response["data"])
	})
}

func TestRespondError(t *testing.T) {
	t.Run("with string error", func(t *testing.T) {
		w, c := setupTest()
		errorMsg := "resource not found"
		expectedCode := http.StatusNotFound

		RespondError(c, errorMsg, expectedCode)
		assert.Equal(t, expectedCode, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, errorMsg, response["error"])
		assert.True(t, c.IsAborted(), "c.AbortWithStatusJSON should abort the context")
	})

	t.Run("with error object", func(t *testing.T) {
		w, c := setupTest()
		errorObj := http.ErrHandlerTimeout
		expectedCode := http.StatusInternalServerError

		RespondError(c, errorObj, expectedCode)
		assert.Equal(t, expectedCode, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, map[string]interface{}{}, response["error"])
		assert.True(t, c.IsAborted(), "c.AbortWithStatusJSON should abort the context")
	})
}

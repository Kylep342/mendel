package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kylep342/mendel/internal/components"
	"github.com/kylep342/mendel/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// testModel is a concrete struct that satisfies the models.Model interface for our tests.
type testModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SetID satisfies the models.Model interface.
func (m *testModel) SetID(id string) {
	m.ID = id
}

// GetID satisfies the models.Model interface.
func (m *testModel) GetID() string {
	return m.ID
}

// newTestModel is the factory function for our test model pointer.
func newTestModel() *testModel {
	return &testModel{}
}

// MockCRUDTable is a mock implementation of the db.CRUDTable interface using testify/mock.
type MockCRUDTable[T any] struct {
	mock.Mock
}

/*

Mock methods below

these mock the HTTP methods
also asserts inner calls to CRUD on [T] have proper args

*/

func (m *MockCRUDTable[T]) GetAll(ctx context.Context) ([]T, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]T), args.Error(1)
}

func (m *MockCRUDTable[T]) Create(ctx context.Context, item *T) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func (m *MockCRUDTable[T]) GetByID(ctx context.Context, id string) (T, error) {
	args := m.Called(ctx, id)
	// We need to assert the type of the first return value to T.
	var zero T
	if args.Get(0) == nil {
		return zero, args.Error(1)
	}
	return args.Get(0).(T), args.Error(1)
}

func (m *MockCRUDTable[T]) Update(ctx context.Context, item *T) error {
	args := m.Called(ctx, item)
	return args.Error(0)
}

func (m *MockCRUDTable[T]) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// setupTest is a helper function to reduce boilerplate in tests.
// It initializes a gin test context, a response recorder, our handler, and the mock table.
func setupTest[T interface{}, PT interface {
	~*T
	components.Model
}](t *testing.T) (*httptest.ResponseRecorder, *gin.Context, *MockCRUDTable[T], *CRUDHandler[T, PT]) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Provides c with Context necessary for tested HTTP methods
	c.Request, _ = http.NewRequest(http.MethodGet, "/", nil)

	mockEnv := &constants.EnvConfig{}
	mockEnv.Server.ReadTimeout = 5 * time.Second
	mockEnv.Server.WriteTimeout = 5 * time.Second

	mockTable := new(MockCRUDTable[T])

	handler := &CRUDHandler[T, PT]{
		Env:   mockEnv,
		Table: mockTable,
		New: func() PT {
			var v T
			return &v
		},
	}

	if any(handler.New()).(any) == any((*testModel)(nil)) {
		handler.New = any(newTestModel).(func() PT)
	}

	return w, c, mockTable, handler
}

// --- Handler Tests ---

func TestCRUDHandler_GetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)

		expectedItems := []testModel{
			{ID: "1", Name: "Plant A"},
			{ID: "2", Name: "Plant B"},
		}

		mockTable.On("GetAll", mock.Anything).Return(expectedItems, nil).Once()
		handler.GetAll(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string][]testModel
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedItems, response["data"])
		mockTable.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)

		dbErr := errors.New("database connection failed")
		mockTable.On("GetAll", mock.Anything).Return(nil, dbErr).Once()
		handler.GetAll(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var errResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errResponse)
		assert.NoError(t, err)
		assert.Contains(t, errResponse["error"], dbErr.Error())
		mockTable.AssertExpectations(t)
	})
}

func TestCRUDHandler_Create(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)

		newItem := testModel{Name: "New Plant"}
		newItemJSON, _ := json.Marshal(newItem)
		c.Request, _ = http.NewRequest(http.MethodPost, "/items", strings.NewReader(string(newItemJSON)))

		mockTable.On("Create", mock.Anything, &newItem).Return(nil).Once()
		handler.Create(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]testModel
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, newItem, response["data"])
		mockTable.AssertExpectations(t)
	})

	t.Run("bad request body", func(t *testing.T) {
		w, c, _, handler := setupTest[testModel, *testModel](t)
		c.Request, _ = http.NewRequest(http.MethodPost, "/items", strings.NewReader(`{"name": "bad json`))

		handler.Create(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var errResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errResponse)
		assert.NoError(t, err)
		assert.NotEmpty(t, errResponse["error"])
	})
}

func TestCRUDHandler_GetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)
		c.Params = gin.Params{gin.Param{Key: "id", Value: "123"}}

		expectedItem := testModel{ID: "123", Name: "Found Plant"}
		mockTable.On("GetByID", mock.Anything, "123").Return(expectedItem, nil).Once()
		handler.GetByID(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]testModel
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedItem, response["data"])
		mockTable.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)
		c.Params = gin.Params{gin.Param{Key: "id", Value: "404"}}

		mockTable.On("GetByID", mock.Anything, "404").Return(testModel{}, sql.ErrNoRows).Once()
		handler.GetByID(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.JSONEq(t, `{"error": "not found"}`, w.Body.String())
		mockTable.AssertExpectations(t)
	})
}

func TestCRUDHandler_Update(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)
		c.Params = gin.Params{gin.Param{Key: "id", Value: "123"}}

		itemToUpdate := &testModel{ID: "123"}
		mockTable.On("Update", mock.Anything, itemToUpdate).Return(nil).Once()
		handler.Update(c)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]testModel
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, *itemToUpdate, response["data"])
		mockTable.AssertExpectations(t)
	})
}

func TestCRUDHandler_Delete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)
		c.Params = gin.Params{gin.Param{Key: "id", Value: "123"}}

		mockTable.On("Delete", mock.Anything, "123").Return(nil).Once()
		handler.Delete(c)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedJSON := fmt.Sprintf(`{"data":"%s"}`, "123")
		assert.JSONEq(t, expectedJSON, w.Body.String())
		mockTable.AssertExpectations(t)
	})

	t.Run("database error", func(t *testing.T) {
		w, c, mockTable, handler := setupTest[testModel, *testModel](t)
		c.Params = gin.Params{gin.Param{Key: "id", Value: "123"}}

		dbErr := errors.New("delete failed")
		mockTable.On("Delete", mock.Anything, "123").Return(dbErr).Once()
		handler.Delete(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var errResponse map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &errResponse)
		assert.NoError(t, err)
		assert.Contains(t, errResponse["error"], dbErr.Error())
		mockTable.AssertExpectations(t)
	})
}

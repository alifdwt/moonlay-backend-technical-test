package handlers

import (
	"backend-technical-test/database"
	"backend-technical-test/models"
	"backend-technical-test/pkg/postgres"
	"backend-technical-test/repository"
	"backend-technical-test/util"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setDatabase(t *testing.T) *gorm.DB {
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_USERNAME", "root")
	t.Setenv("DB_PASSWORD", "secret")
	t.Setenv("DB_NAME", "todo_db_test")
	t.Setenv("APP_TIMEZONE", "Asia/Jakarta")

	postgres.TestDatabase(os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("APP_TIMEZONE"))
	database.RunMigration()
	return postgres.DB
}

func createRandomList(t *testing.T) models.List {
	arg := models.List{
		Title:       util.RandomTitle(),
		Description: util.RandomDescription(),
	}

	requestBody := strings.NewReader(fmt.Sprintf(`{"title": "%s", "description": "%s"}`, arg.Title, arg.Description))

	req := httptest.NewRequest("POST", "/api/lists", requestBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	handler := NewListHandler(repository.NewListRepository(setDatabase(t)))
	err := handler.CreateList(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)

	var res map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &res)
	require.NoError(t, err)

	require.Contains(t, res, "data")
	data, ok := res["data"].(map[string]interface{})
	require.True(t, ok)

	id, ok := data["id"].(float64)
	require.True(t, ok)

	createdList := models.List{
		Id:          int(id),
		Title:       arg.Title,
		Description: arg.Description,
	}

	return createdList
}

func TestCreateList(t *testing.T) {
	createRandomList(t)
}

func TestFindLists(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/lists?page=1&withSublists=true", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	handler := NewListHandler(repository.NewListRepository(setDatabase(t)))
	err := handler.FindLists(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestFindListById(t *testing.T) {
	list := createRandomList(t)
	req := httptest.NewRequest("GET", "/api/v1/lists", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(":id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(list.Id)))

	handler := NewListHandler(repository.NewListRepository(setDatabase(t)))
	err := handler.FindListById(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateList(t *testing.T) {
	list := createRandomList(t)
	title := util.RandomTitle()
	description := util.RandomDescription()

	requestBody := strings.NewReader(fmt.Sprintf(
		`{"title": "%s", "description": "%s"}`, title, description,
	))
	req := httptest.NewRequest("PUT", "/api/lists", requestBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(":id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(list.Id)))

	handler := NewListHandler(repository.NewListRepository(setDatabase(t)))
	err := handler.UpdateList(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteList(t *testing.T) {
	list := createRandomList(t)
	req := httptest.NewRequest("DELETE", "/api/lists", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(":id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(list.Id)))

	handler := NewListHandler(repository.NewListRepository(setDatabase(t)))
	err := handler.DeleteList(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}
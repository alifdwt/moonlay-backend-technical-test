package handlers

import (
	"backend-technical-test/models"
	"backend-technical-test/repository"
	"backend-technical-test/util"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func createRandomSubList(t *testing.T) models.Sublist {
	arg := models.Sublist{
		Title:     util.RandomTitle(),
		Description: util.RandomDescription(),
		ListId: int(util.RandomInt(1, 3)),
	}

	requestBody := strings.NewReader(fmt.Sprintf(`{"title": "%s", "description": "%s", "list_id": %d}`, arg.Title, arg.Description, arg.ListId))

	req := httptest.NewRequest("POST", "/api/sublists", requestBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	handler := NewSublistHandler(repository.NewSublistRepository(setDatabase(t)))
	err := handler.CreateSublist(ctx)
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

	createdSubList := models.Sublist{
		Id:        int(id),
		Title:     arg.Title,
		Description: arg.Description,
		ListId:    arg.ListId,
	}

	return createdSubList
}

func TestCreateSubList(t *testing.T) {
	createRandomSubList(t)
}

func TestFindSublists(t *testing.T) {
	req := httptest.NewRequest("GET", "/api/v1/sublists?page=1", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	handler := NewSublistHandler(repository.NewSublistRepository(setDatabase(t)))
	err := handler.FindSublists(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestFindSubListById(t *testing.T) {
	sublist := createRandomSubList(t)
	req := httptest.NewRequest("GET", "/api/v1/sublists", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(":id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(sublist.Id)))

	handler := NewSublistHandler(repository.NewSublistRepository(setDatabase(t)))
	err := handler.FindSublistById(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestFindSublistWithListId(t *testing.T) {
	sublist := createRandomSubList(t)
	req := httptest.NewRequest("GET", "/api/v1/sublists", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(":id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(sublist.ListId)))

	handler := NewSublistHandler(repository.NewSublistRepository(setDatabase(t)))
	err := handler.FindSublistByListId(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateSublist(t *testing.T) {
	sublist := createRandomSubList(t)
	title := util.RandomTitle()
	description := util.RandomDescription()

	requestBody := strings.NewReader(fmt.Sprintf(
		`{"title": "%s", "description": "%s"}`, title, description,
	))

	req := httptest.NewRequest("PUT", "/api/sublists", requestBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(":id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(sublist.Id)))

	handler := NewSublistHandler(repository.NewSublistRepository(setDatabase(t)))
	err := handler.UpdateSublist(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteSublist(t *testing.T) {
	sublist := createRandomSubList(t)
	req := httptest.NewRequest("DELETE", "/api/sublists", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	ctx.SetPath(":id")
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(sublist.Id)))

	handler := NewSublistHandler(repository.NewSublistRepository(setDatabase(t)))
	err := handler.DeleteSublist(ctx)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, rec.Code)
}
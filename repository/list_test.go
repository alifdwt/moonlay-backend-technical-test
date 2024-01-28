package repository

import (
	"backend-technical-test/database"
	"backend-technical-test/models"
	"backend-technical-test/pkg/postgres"
	"backend-technical-test/util"
	"os"

	"testing"
	"time"

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
		Title:     util.RandomTitle(),
		Description: util.RandomDescription(),
	}

	listRepository := NewListRepository(setDatabase(t))

	list, err := ListRepository.CreateList(listRepository, arg)
	require.NoError(t, err)
	require.NotEmpty(t, list)

	require.Equal(t, arg.Title, list.Title)
	require.Equal(t, arg.Description, list.Description)

	require.NotZero(t, list.Id)
	require.NotZero(t, list.CreatedAt)

	return list
}

func TestCreateList(t *testing.T) {
	createRandomList(t)
}

func TestGetLists(t *testing.T) {
	listRepository := NewListRepository(setDatabase(t))
	lists, err := ListRepository.GetLists(listRepository, 10, 10, false, "%")
	require.NotEmpty(t, lists)

	require.NoError(t, err)
}

func TestGetListWithSublists(t *testing.T) {
	listRepository := NewListRepository(setDatabase(t))
	lists, err := ListRepository.GetLists(listRepository, 10, 10, true, "%")
	require.NotEmpty(t, lists)

	require.NoError(t, err)
}

func TestGetList(t *testing.T) {
	listRepository := NewListRepository(setDatabase(t))
	list1 := createRandomList(t)
	list2, err := ListRepository.FindListById(listRepository, list1.Id)
	require.NoError(t, err)
	require.NotEmpty(t, list2)

	require.Equal(t, list1.Id, list2.Id)
	require.Equal(t, list1.Title, list2.Title)
	require.Equal(t, list1.Description, list2.Description)

	require.WithinDuration(t, list1.CreatedAt, list2.CreatedAt, time.Second)
}

func TestUpdateList(t *testing.T) {
	listRepository := NewListRepository(setDatabase(t))
	list1 := createRandomList(t)

	arg := models.List{
		Id:        list1.Id,
		Title:     util.RandomTitle(),
		Description: util.RandomDescription(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	list2, err := ListRepository.UpdateList(listRepository, arg, list1.Id)
	require.NoError(t, err)
	require.NotEmpty(t, list2)

	require.Equal(t, list1.Id, list2.Id)
	require.Equal(t, arg.Title, list2.Title)
	require.Equal(t, arg.Description, list2.Description)
	
	require.WithinDuration(t, list1.CreatedAt, list2.CreatedAt, time.Second)
}

func TestDeleteList(t *testing.T) {
	list1 := createRandomList(t)
	listRepository := NewListRepository(setDatabase(t))
	_, err := ListRepository.DeleteList(listRepository, list1)
	require.NoError(t, err)

	list2, err := ListRepository.FindListById(listRepository, list1.Id)
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, list2)
}
package repository

import (
	"backend-technical-test/models"
	"backend-technical-test/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func createRandomSublist(t *testing.T) models.Sublist {
	arg := models.Sublist{
		Title:     util.RandomTitle(),
		Description: util.RandomDescription(),
		ListId: int(util.RandomInt(1, 3)),
	}

	sublistRepository := NewSublistRepository(setDatabase(t))

	sublist, err := SublistRepository.CreateSublist(sublistRepository, arg)
	require.NoError(t, err)
	require.NotEmpty(t, sublist)

	require.Equal(t, arg.Title, sublist.Title)
	require.Equal(t, arg.Description, sublist.Description)

	require.NotZero(t, sublist.Id)
	require.NotZero(t, sublist.CreatedAt)

	return sublist
}

func TestCreateSublist(t *testing.T) {
	createRandomSublist(t)
}

func TestGetSublists(t *testing.T) {
	sublistRepository := NewSublistRepository(setDatabase(t))
	sublists, err := SublistRepository.FindSublists(sublistRepository, 10, 10, "%")
	require.NotEmpty(t, sublists)

	require.NoError(t, err)
}

func TestGetSublist(t *testing.T) {
	sublistRepository := NewSublistRepository(setDatabase(t))
	sublist1 := createRandomSublist(t)
	sublist2, err := SublistRepository.FindSublistById(sublistRepository, sublist1.Id)
	
	require.NoError(t, err)
	require.NotEmpty(t, sublist2)

	require.Equal(t, sublist1.Id, sublist2.Id)
	require.Equal(t, sublist1.Title, sublist2.Title)
	require.Equal(t, sublist1.Description, sublist2.Description)

	require.WithinDuration(t, sublist1.CreatedAt, sublist2.CreatedAt, time.Second)
}

func TestGetSublistByListId(t *testing.T) {
	sublistRepository := NewSublistRepository(setDatabase(t))
	// randomInt := int(util.RandomInt(1, 3))
	sublists, err := SublistRepository.FindSublistByListId(sublistRepository, 1, 1, 1, "%")
	require.NoError(t, err)
	require.NotEmpty(t, sublists)

	// sublist2 := createRandomSublist(t)
	// sublist3, err := SublistRepository.FindSublistByListId(sublistRepository, int(sublist2.ListId), 1, 10, string(sublist2.Title))
	
	// require.NoError(t, err)
	// require.NotEmpty(t, sublist3)
	// fmt.Println(sublist2.Id, sublist3[0].Id)
	// require.Equal(t, sublist2.Id, sublist3[0].Id)
	// require.Equal(t, sublist2.Title, sublist3[0].Title)

	// require.WithinDuration(t, sublist2.CreatedAt, sublist3[0].CreatedAt, time.Second)
}

func TestUpdateSublist(t *testing.T) {
	sublistRepository := NewSublistRepository(setDatabase(t))
	sublist1 := createRandomSublist(t)

	arg := models.Sublist{
		Id:        sublist1.Id,
		Title:     util.RandomTitle(),
		Description: util.RandomDescription(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	sublist2, err := SublistRepository.UpdateSublist(sublistRepository, arg, sublist1.Id)
	require.NoError(t, err)
	require.NotEmpty(t, sublist2)

	require.Equal(t, sublist1.Id, sublist2.Id)
	require.Equal(t, arg.Title, sublist2.Title)
	require.Equal(t, arg.Description, sublist2.Description)

	require.WithinDuration(t, sublist1.CreatedAt, sublist2.CreatedAt, time.Second)
}

func TestDeleteSublist(t *testing.T) {
	sublist1 := createRandomSublist(t)
	sublistRepository := NewSublistRepository(setDatabase(t))
	_, err := SublistRepository.DeleteSublist(sublistRepository, sublist1)
	require.NoError(t, err)

	sublist2, err := SublistRepository.FindSublistById(sublistRepository, sublist1.Id)
	require.Error(t, err)
	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
	require.Empty(t, sublist2)
}
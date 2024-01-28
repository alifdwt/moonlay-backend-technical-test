package repository

import (
	"backend-technical-test/models"

	"gorm.io/gorm"
)

type ListRepository interface {
	GetLists(offset int, limit int, withSublists bool, search string) ([]models.List, error)
	FindListById(id int) (models.List, error)
	CreateList(list models.List) (models.List, error)
	UpdateList(list models.List, id int) (models.List, error)
	DeleteList(list models.List) (models.List, error)
}

type listRepository struct {
	db *gorm.DB
}

func NewListRepository(db *gorm.DB) *listRepository {
	return &listRepository{db}
}

func (r *listRepository) GetLists(offset int, limit int, withSublists bool, search string) ([]models.List, error) {
	var lists []models.List
	if withSublists {
		err := r.db.Preload("Sublists").Limit(limit).Offset(offset).Where("title iLIKE ?", "%"+search+"%").Or("description iLIKE ?", "%"+search+"%").Find(&lists).Error
		return lists, err
	}
	err := r.db.Limit(limit).Offset(offset).Where("title iLIKE ?", "%"+search+"%").Or("description iLIKE ?", "%"+search+"%").Find(&lists).Error
	return lists, err
}

func (r *listRepository) FindListById(id int) (models.List, error) {
	var list models.List
	err := r.db.Preload("Sublists").First(&list, id).Error

	return list, err
}

func (r *listRepository) CreateList(list models.List) (models.List, error) {
	err := r.db.Create(&list).Error

	return list, err
}

func (r *listRepository) UpdateList(list models.List, id int) (models.List, error) {
	err := r.db.Model(&list).Where("id = ?", id).Updates(&list).Error

	return list, err
}

func (r *listRepository) DeleteList(list models.List) (models.List, error) {
	err := r.db.Delete(&list).Error

	return list, err
}
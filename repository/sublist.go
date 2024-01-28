package repository

import (
	"backend-technical-test/models"

	"gorm.io/gorm"
)

type SublistRepository interface {
	FindSublists(offset int, limit int, search string) ([]models.Sublist, error)
	FindSublistById(id int) (models.Sublist, error)
	FindSublistByListId(listId int, offset int, limit int, search string) ([]models.Sublist, error)
	CreateSublist(sublist models.Sublist) (models.Sublist, error)
	UpdateSublist(sublist models.Sublist, id int) (models.Sublist, error)
	DeleteSublist(sublist models.Sublist) (models.Sublist, error)
}

type sublistRepository struct {
	db *gorm.DB
}

func NewSublistRepository(db *gorm.DB) *sublistRepository {
	return &sublistRepository{db}
}

func (r *sublistRepository) FindSublists(offset int, limit int, search string) ([]models.Sublist, error) {
	var sublists []models.Sublist
	err := r.db.Limit(limit).Offset(offset).Where("title iLIKE ?", "%"+search+"%").Or("description iLIKE ?", "%"+search+"%").Find(&sublists).Error

	return sublists, err
}

func (r *sublistRepository) FindSublistById(id int) (models.Sublist, error) {
	var sublist models.Sublist
	err := r.db.Preload("List").First(&sublist, id).Error

	return sublist, err
}

func (r *sublistRepository) FindSublistByListId(listId int, offset int, limit int, search string) ([]models.Sublist, error) {
    var sublists []models.Sublist
    query := r.db.Limit(limit).Offset(offset).Where("list_id = ?", listId)

    if search != "%" {
        query = query.Where("title ILIKE ?", "%"+search+"%").Or("description ILIKE ?", "%"+search+"%")
    }

    err := query.Find(&sublists).Error

    return sublists, err
}

func (r *sublistRepository) CreateSublist(sublist models.Sublist) (models.Sublist, error) {
	err := r.db.Create(&sublist).Error

	return sublist, err
}

func (r *sublistRepository) UpdateSublist(sublist models.Sublist, id int) (models.Sublist, error) {
	err := r.db.Model(&sublist).Where("id = ?", id).Updates(&sublist).Error

	return sublist, err
}

func (r *sublistRepository) DeleteSublist(sublist models.Sublist) (models.Sublist, error) {
	err := r.db.Delete(&sublist).Error

	return sublist, err
}
package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Store(Category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) Store(Category *model.Category) error {
	err := c.db.Create(Category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) Update(id int, category model.Category) error {
	err := c.db.Model(&model.Category{}).Where("id = ?", id).Updates(category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) Delete(id int) error {
	err := c.db.Delete(&model.Category{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryRepository) GetByID(id int) (*model.Category, error) {
	var theCategory model.Category
	err := c.db.Where("id = ?", id).First(&theCategory).Error
	if err != nil {
		return nil, err
	}

	return &theCategory, nil
}

func (c *categoryRepository) GetList() ([]model.Category, error) {
	var itsAllCategories []model.Category
	err := c.db.Find(&itsAllCategories).Error
	if err != nil {
		return nil, err
	}

	return itsAllCategories, nil
}
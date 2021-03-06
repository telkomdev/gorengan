package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/musobarlab/gorengan/modules/category/domain"
	"github.com/musobarlab/gorengan/modules/shared"
)

// CategoryRepositoryGorm struct
type CategoryRepositoryGorm struct {
	db *gorm.DB
}

// NewCategoryRepositoryGorm function
func NewCategoryRepositoryGorm(db *gorm.DB) *CategoryRepositoryGorm {
	return &CategoryRepositoryGorm{db: db}
}

// Save function
func (r *CategoryRepositoryGorm) Save(category *domain.Category) shared.Output {
	err := r.db.Save(category).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: category}
}

// Delete function
func (r *CategoryRepositoryGorm) Delete(category *domain.Category) shared.Output {
	err := r.db.Delete(category).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: category}
}

// FindByID function
func (r *CategoryRepositoryGorm) FindByID(id string) shared.Output {
	var category domain.Category

	err := r.db.Where(&domain.Category{ID: id}).Take(&category).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: &category}
}

// FindAll function
func (r *CategoryRepositoryGorm) FindAll(params *shared.Parameters) shared.Output {
	var categories domain.Categories

	err := r.db.Offset(params.Offset).Limit(params.Limit).Order(params.OrderBy).Find(&categories).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: categories}
}

// Count function
func (r *CategoryRepositoryGorm) Count(params *shared.Parameters) shared.Output {
	var count int
	err := r.db.Model(&domain.Category{}).Offset(params.Offset).Limit(params.Limit).Order(params.OrderBy).Count(&count).Error
	if err != nil {
		return shared.Output{Err: err}
	}

	return shared.Output{Result: count}
}

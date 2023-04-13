package db

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Person struct {
	ID uuid.UUID
    Name string
    Age uint8
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	db.AutoMigrate(&Person{})
	return &repository{db}
}

func (r *repository) Insert(p *Person) (*Person, error) {
	result := r.db.Create(&p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (r *repository) Update(p *Person) (*Person, error) {
	result := r.db.Model(&p).Updates(Person{
		Name: p.Name,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (r *repository) Delete(id uuid.UUID) error {
	p := Person{ID: id}

	result := r.db.Delete(&p)

	return result.Error
}

func (r *repository) Paginate() ([]Person, error) {
	var people []Person
	result := r.db.Find(&people)
	if result.Error != nil {
		return nil, result.Error
	}

	return people, nil
}

func (r *repository) FindByID(id uuid.UUID) (*Person, error) {
	var p Person

	result := r.db.First(&p, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &p, nil
}

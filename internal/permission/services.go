package permission

import (
	"account/internal/database"
	"account/internal/pageable"
	"account/internal/util"
	"github.com/google/uuid"
)

type Service struct{}

func (s *Service) AutoMigrate() {
	db := database.GetDB()

	err := db.AutoMigrate(&Model{})
	util.Log("db err: (AutoMigrate)", err)
}

func (s *Service) Save(permission Model) (Model, error) {
	db := database.GetDB()
	err := db.Save(&permission).Error
	return permission, err
}

func (s *Service) Update(permission Model) (Model, error) {
	db := database.GetDB()
	err := db.Model(&permission).Updates(permission).Error

	return permission, err
}

func (s *Service) FindPage(pgb pageable.Pageable, permissions *[]Model) (pageable.Page, error) {
	page := pageable.Page{
		Size: pgb.Size,
		Page: pgb.Page,
		Sort: pgb.Sort,
	}
	db := database.GetDB()

	err := db.Scopes(pageable.Paginate(permissions, pgb, &page, db)).Find(&permissions).Error

	return page, err
}

func (s *Service) FindOne(id uuid.UUID) (Model, error) {
	db := database.GetDB()

	permission := Model{}
	err := db.Where("id = ?", id).Find(&permission).Error

	return permission, err
}

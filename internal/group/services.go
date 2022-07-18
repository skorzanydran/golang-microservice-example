package group

import (
	"account/internal/database"
	"account/internal/pageable"
	"account/internal/permission"
	"account/internal/util"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Service struct{}

func (s *Service) AutoMigrate() {
	db := database.GetDB()

	err := db.AutoMigrate(&Model{})
	util.Log("db err: (AutoMigrate)", err)
}

func (s *Service) Save(group Model) (Model, error) {
	db := database.GetDB()
	err := db.Save(&group).Error
	return group, err
}

func (s *Service) Update(group Model) (Model, error) {
	db := database.GetDB()
	err := db.Preload(clause.Associations).Model(&group).Updates(group).Error

	return group, err
}

func (s *Service) FindPage(pgb pageable.Pageable, groups *[]Model) (pageable.Page, error) {
	page := pageable.Page{
		Size: pgb.Size,
		Page: pgb.Page,
		Sort: pgb.Sort,
	}
	db := database.GetDB()

	err := db.Preload(clause.Associations).Scopes(pageable.Paginate(groups, pgb, &page, db)).Find(&groups).Error

	return page, err
}

func (s *Service) FindOne(id uuid.UUID) (Model, error) {
	db := database.GetDB()

	group := Model{}
	err := db.Preload(clause.Associations).Where("id = ?", id).Find(&group).Error

	return group, err
}

func (s *Service) AddPermission(group Model, perm permission.Model) (Model, error) {
	db := database.GetDB()

	err := db.Model(&group).Association("Permissions").Append(&perm)
	return group, err
}

func (s *Service) RemovePermission(group Model, perm permission.Model) (Model, error) {
	db := database.GetDB()

	err := db.Model(&group).Association("Permissions").Delete(&perm)
	return group, err
}

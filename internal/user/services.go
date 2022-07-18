package user

import (
	"account/internal/database"
	"account/internal/group"
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

func (s *Service) Save(user Model) (Model, error) {
	db := database.GetDB()
	err := db.Save(&user).Error
	return user, err
}

func (s *Service) Update(user Model) (Model, error) {
	db := database.GetDB()
	err := db.Preload(clause.Associations).Preload("Group." + clause.Associations).Model(&user).Updates(user).Error

	return user, err
}

func (s *Service) FindPage(pgb pageable.Pageable, users *[]Model) (pageable.Page, error) {
	page := pageable.Page{
		Size: pgb.Size,
		Page: pgb.Page,
		Sort: pgb.Sort,
	}
	db := database.GetDB()

	err := db.Preload(clause.Associations).Preload("Group." + clause.Associations).Scopes(pageable.Paginate(users, pgb, &page, db)).Find(&users).Error

	return page, err
}

func (s *Service) FindOne(id uuid.UUID) (Model, error) {
	db := database.GetDB()

	users := Model{}
	err := db.Preload(clause.Associations).Preload("Groups."+clause.Associations).Where("id = ?", id).Find(&users).Error

	return users, err
}

func (s *Service) AddPermission(user Model, perm permission.Model) (Model, error) {
	db := database.GetDB()

	err := db.Preload(clause.Associations).Preload("Groups." + clause.Associations).Model(&user).Association("Permissions").Append(&perm)
	return user, err
}

func (s *Service) RemovePermission(user Model, perm permission.Model) (Model, error) {
	db := database.GetDB()

	err := db.Preload(clause.Associations).Preload("Groups." + clause.Associations).Model(&user).Association("Permissions").Delete(&perm)
	return user, err
}

func (s *Service) AddGroup(user Model, group group.Model) (Model, error) {
	db := database.GetDB()

	err := db.Preload(clause.Associations).Preload("Groups." + clause.Associations).Model(&user).Association("Groups").Append(&group)
	return user, err
}

func (s *Service) RemoveGroup(user Model, group group.Model) (Model, error) {
	db := database.GetDB()

	err := db.Preload(clause.Associations).Preload("Groups." + clause.Associations).Model(&user).Association("Groups").Delete(&group)
	return user, err
}

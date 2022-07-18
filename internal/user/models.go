package user

import (
	"account/internal/group"
	"account/internal/permission"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Model struct {
	Id          uuid.UUID          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username    string             `gorm:"unique;not null"`
	Password    string             `gorm:"not null"`
	Groups      []group.Model      `gorm:"many2many:security_user_id_xref_security_group_id;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:group_id"`
	Permissions []permission.Model `gorm:"many2many:security_user_id_xref_security_permission_id;foreignKey:id;joinForeignKey:user_id;References:id;joinReferences:permission_id"`
	Created     time.Time          `gorm:"<-:create;autoUpdateTime"`
	Updated     time.Time          `gorm:"autoCreateTime"`
}

func (Model) TableName() string {
	return "security_users"
}

func (m *Model) SetPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytesPwd := []byte(password)
	byteHashedPwd, _ := bcrypt.GenerateFromPassword(bytesPwd, bcrypt.DefaultCost)
	m.Password = string(byteHashedPwd)
	return nil
}

func (m *Model) checkPassword(password string) error {
	bytePwd := []byte(password)
	byteHashedPwd := []byte(m.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPwd, bytePwd)
}

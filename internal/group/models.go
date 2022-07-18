package group

import (
	"account/internal/permission"
	"github.com/google/uuid"
	"time"
)

type Model struct {
	Id          uuid.UUID          `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string             `gorm:"unique;not null"`
	Permissions []permission.Model `gorm:"many2many:security_group_id_xref_security_permission_id;foreignKey:id;joinForeignKey:group_id;References:id;joinReferences:permission_id"`
	Created     time.Time          `gorm:"<-:create;autoUpdateTime"`
	Updated     time.Time          `gorm:"autoCreateTime"`
}

func (Model) TableName() string {
	return "security_groups"
}

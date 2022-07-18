package permission

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	Id      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name    string    `gorm:"unique;not null"`
	Created time.Time `gorm:"<-:create;autoUpdateTime"`
	Updated time.Time `gorm:"autoCreateTime"`
}

func (Model) TableName() string {
	return "security_permissions"
}

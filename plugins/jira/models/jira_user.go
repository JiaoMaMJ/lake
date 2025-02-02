package models

import (
	"github.com/merico-dev/lake/models"
)

type JiraUser struct {
	models.NoPKModel

	// collected fields
	SourceId    uint64 `gorm:"primarykey"`
	AccountId   string `gorm:"primaryKey"`
	AccountType string
	Name        string
	Email       string
	AvatarUrl   string
	Timezone    string
}

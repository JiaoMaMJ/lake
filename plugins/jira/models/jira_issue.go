package models

import (
	"database/sql"
	"time"

	"github.com/merico-dev/lake/models"
)

type JiraIssue struct {
	models.NoPKModel

	// collected fields
	SourceId                 uint64 `gorm:"primaryKey"`
	IssueId                  uint64 `gorm:"primarykey"`
	ProjectId                uint64
	Self                     string
	Key                      string
	Summary                  string
	Type                     string
	EpicKey                  string
	StatusName               string
	StatusKey                string
	StatusCategory           string
	StoryPoint               float64
	OriginalEstimateMinutes  int64 // user input?
	AggregateEstimateMinutes int64 // sum up of all subtasks?
	RemainingEstimateMinutes int64 // could it be negative value?
	CreatorAccountId         string
	CreatorAccountType       string
	CreatorDisplayName       string
	AssigneeAccountId        string // latest assignee
	AssigneeAccountType      string
	AssigneeDisplayName      string
	PriorityId               uint64
	PriorityName             string
	ParentId                 uint64
	ParentKey                string
	SprintId                 uint64 // latest sprint, issue might cross multiple sprints, would be addressed by #514
	SprintName               string
	ResolutionDate           sql.NullTime
	Created                  time.Time
	Updated                  time.Time

	// enriched fields
	// RequirementAnalsyisLeadTime uint
	// DesignLeadTime              uint
	// DevelopmentLeadTime         uint
	// TestLeadTime                uint
	// DeliveryLeadTime            uint
	SpentMinutes  int64
	LeadTime      uint
	StdStoryPoint uint
	StdType       string
	StdStatus     string

	// internal status tracking
	ChangelogUpdated sql.NullTime
}

package main

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	TeamID    uuid.UUID `json:"teamId" gorm:"type:uuid;default:gen_random_uuid()"`
	TeamName  string    `json:"teamName" gorm:"not null;size:255"`
	OwnerID   uuid.UUID `json:"ownerId" gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TeamMember struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	TeamID    uuid.UUID `json:"teamId" gorm:"type:uuid"`
	MemberID  uuid.UUID `json:"memberId" gorm:"type:uuid"`
	CreatedAt time.Time `json:"createdAt"`
}

type TeamManager struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	TeamID    uuid.UUID `json:"teamId" gorm:"type:uuid"`
	ManagerID uuid.UUID `json:"managerId" gorm:"type:uuid"`
	CreatedAt time.Time `json:"createdAt"`
}

func (tm *TeamMember) BeforeCreate(tx *gorm.DB) error {
	if tm.ID == "" {
		tm.ID = uuid.New().String()
	}
	return nil
}

func (Team) TableName() string {
	return "teams"
}

func (TeamMember) TableName() string {
	return "team_members"
}

func (TeamManager) TableName() string {
	return "team_managers"
}

package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeamCreate struct {
	TeamName string `json:"teamName" binding:"required" validate:"required,min=1"`
}

func CreateTeam(c *gin.Context) {
	var input TeamCreate

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	team := Team{
		TeamName: input.TeamName,
	}

	result := DB.Create(&team)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
		return
	}

	c.JSON(http.StatusCreated, team)
}

type MemberAdd struct {
	MemberID uuid.UUID `json:"memberId"`
}

func AddMember(c *gin.Context) {
	team, err := getTeam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input MemberAdd
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member := TeamMember{
		TeamID:   team.TeamID,
		MemberID: input.MemberID,
	}
	var dbMember TeamMember

	result := DB.Where(member).FirstOrCreate(&dbMember)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add member"})
		return
	}

	c.JSON(http.StatusCreated, dbMember)
}

func RemoveMember(c *gin.Context) {
	member, err := getMember(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := DB.Where(&member).Delete(&TeamMember{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove member"})
		return
	}

	c.Status(http.StatusNoContent)
}

type ManagerAdd struct {
	ManagerID uuid.UUID `json:"managerId" binding:"required"`
}

func AddManager(c *gin.Context) {
	team, err := getTeam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input ManagerAdd
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manager := TeamManager{
		TeamID:    team.TeamID,
		ManagerID: input.ManagerID,
	}

	var dbManager TeamManager
	result := DB.Where(manager).FirstOrCreate(&dbManager)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add manager"})
		return
	}

	c.JSON(http.StatusCreated, dbManager)
}

func RemoveManager(c *gin.Context) {
	manager, err := getManager(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := DB.Where(&manager).Delete(&TeamManager{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove manager"})
		return
	}

	c.Status(http.StatusNoContent)
}

func getTeam(c *gin.Context) (*Team, error) {
	teamID, err := uuid.Parse(c.Param("teamId"))
	if err != nil {
		return nil, errors.New("teamId is not valid UUIDv4")
	}

	var team Team

	err = DB.First(&team, teamID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("Team not found")
	}

	return &team, nil
}

func getMember(c *gin.Context) (*TeamMember, error) {
	team, err := getTeam(c)
	if err != nil {
		return nil, err
	}

	memberID, err := uuid.Parse(c.Param("memberId"))
	if err != nil {
		return nil, errors.New("memberId is not valid UUIDv4")
	}

	member := TeamMember{
		TeamID:   team.TeamID,
		MemberID: memberID,
	}

	var dbMember TeamMember

	err = DB.Where(&member).First(&dbMember).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("member not found")
	}

	return &member, nil
}

func getManager(c *gin.Context) (*TeamMember, error) {
	team, err := getTeam(c)
	if err != nil {
		return nil, err
	}

	memberID, err := uuid.Parse(c.Param("managerId"))
	if err != nil {
		return nil, errors.New("managerId is not valid UUIDv4")
	}

	member := TeamMember{
		TeamID:   team.TeamID,
		MemberID: memberID,
	}

	var dbMember TeamMember

	err = DB.Where(&member).First(&dbMember).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("manager not found")
	}

	return &member, nil
}

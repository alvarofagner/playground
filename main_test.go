package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: v1.23.3
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user1 := User{Name: "Vision"}
	user2 := User{Name: "Wanda"}
	user3 := User{Name: "TChala"}
	team := Team{Name: "Ironman"}

	DB.Create(&user1)
	DB.Create(&user2)
	DB.Create(&user3)
	DB.Delete(&user2)
	DB.Create(&team)

	var user1Result User
	if err := DB.First(&user1Result, user1.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var user3Result User
	if err := DB.First(&user3Result, user3.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var teamResult Team
	if err := DB.First(&teamResult, team.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&teamResult).Association("Users").Append(&user1Result); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var preloadTeam Team
	if err := DB.Where("name = ?", team.Name).Preload("Users").First(&preloadTeam).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	if err := DB.Model(&preloadTeam).Association("Users").Append(&user3Result); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

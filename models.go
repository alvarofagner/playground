package main

import (
	"time"
)

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User struct {
	ID        uint64 `gorm:"primary_key;auto_increment;not null"`
	Name      string
	Age       uint
	Birthday  *time.Time
	CompanyID *int
	ManagerID *uint
	Manager   *User
	Teams     []*Team `gorm:"many2many:team_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Active    bool
}

type Team struct {
	ID        uint64    `gorm:"primary_key;auto_increment;not null"`
	Name      string    `gorm:"size:256"`
	Users     []*User   `gorm:"many2many:team_users;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
}

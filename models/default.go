package models

type Default struct {
	ID   int
	Name string `gorm:"not null;unique;size:255"`
	Age  int    `gorm:"int"`
}



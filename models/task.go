package models

import (
	"time"
)

type Task struct {
	TaskID    uint      `json:"task_id" gorm:"primary_key"`
	Title     string    `json:"title"`
	Time      time.Time `json:"time"`
	ProjectID uint      `json:"-"`
	Project   Project   `json:"project" gorm:"foreignKey:ProjectID"`
	LabelID   uint      `json:"-"`
	Label     Label     `json:"label" gorm:"foreignKey:LabelID"` //gorm:"type:text[]"`
	StatusID  uint      `json:"-"`
	Status    Status    `json:"status" gorm:"foreignKey:StatusID"`
	UserID    uint      `json:"-"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
}
type User struct {
	UserID uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
type Project struct {
	ProjectID    uint   `json:"project_id" gorm:"primaryKey"`
	ProjectTitle string `json:"project_title"`
}

type Label struct {
	LabelID    uint   `json:"label_id" gorm:"primaryKey"`
	LabelTitle string `json:"label_title"`
}

type Status struct {
	StatusID    uint   `json:"status_id" gorm:"primaryKey"`
	StatusTitle string `json:"status_title"`
}

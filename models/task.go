package models

import (
	"time"
)

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

type Task struct {
	TaskID       uint      `json:"task_id" gorm:"primaryKey"`
	Title        string    `json:"title"`
	Time         time.Time `json:"time"`
	ProjectID    int       `json:"-"`
	TitleProject Project   `json:"project" gorm:"foreignKey:ProjectID"`
	LabelID      int       `json:"-"`
	TitleLabel   []Label   `json:"labels" gorm:"type:text" gorm:"foreignKey:LabelID"  ` //gorm:"type:text[]"`
	StatusID     int       `json:"-"`
	Status       Status    `json:"status" gorm:"foreignKey:StatusID"`
}

// type Task struct {
// 	TaskID       uint      `json:"task_id" gorm:"primaryKey"`
// 	Title        string    `json:"title"`
// 	Time         time.Time `json:"time"`
// 	ProjectID    int       `json:"-"`
// 	TitleProject Project   `json:"project" gorm:"foreignKey:ProjectID"`
// 	LableID      int       `json:"-"`
// 	TitleLables  []Label   `json:"labels" gorm:"foreignKey:LabelID"`
// 	StatusID     int       `json:"-"`
// 	TitleStatus  Status    `json:"status" gorm:"foreignKey:StatusID"`
// }

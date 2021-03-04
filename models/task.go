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
	TaskID     int    `json:"-"`
}

type Task struct {
	TaskID       uint      `json:"task_id" gorm:"primaryKey"`
	Title        string    `json:"title"`
	Time         time.Time `json:"time"`
	ProjectID    int       `json:"-"`
	TitleProject Project   `json:"project" gorm:"foreignKey:ProjectID"` //`json:"project" gorm:"foreignKey:ProjectID"`
	TitleLables  []Label   `json:"labels" gorm:"foreignKey:TaskID"`
	Status       string    `json:"status"`
}

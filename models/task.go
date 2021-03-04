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

// type Task struct {
// 	TaskID        uint      `json:"task_id" gorm:"primaryKey"`
// 	Title         string    `json:"title"`
// 	Time          time.Time `json:"time"`
// 	TaskProjectID int       `json:"task_project_id" gorm:"foreignKey:ProjectID"`
// 	TitleProject  string    `json:"project"`
// 	TaskLabelID   int       `json:"task_label_id"  gorm:"foreignKey:LabelID"`
// 	TitleLabel    string    `json:"labels"` //gorm:"type:text[]"`
// 	TaskStatusID  int       `json:"task_status_id" gorm:"foreignKey:StatusID"`
// 	Status        string    `json:"status" `
// }
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

// {
// 	"title": "Create Sketch",
// 	"time": "2021-01-20T21:21:30+00:00",
// 	"project": {
// 		"project_title":
// 		"Banking"
// 	},
// 	"labels": {
// 		"lable_title":
// 		["Email","Design"]
// 	},
// 	"status": {
// 		"status_title": "on progress"
// 	}
// 	}

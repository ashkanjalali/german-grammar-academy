package academy

import "time"

type Lesson struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Level       string    `json:"level" gorm:"index"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug" gorm:"uniqueIndex"`
	Description string    `json:"description"`
	Explanation string    `json:"explanation" gorm:"type:text"`
	Examples    string    `json:"examples" gorm:"type:text"` // JSON array
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Question struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Level         string `json:"level" gorm:"index"`
	LessonID      uint   `json:"lessonId" gorm:"index"`
	Prompt        string `json:"prompt"`
	Options       string `json:"options" gorm:"type:text"` // JSON array
	CorrectAnswer int    `json:"-"` // never expose correct answer before submission
	Explanation   string `json:"explanation" gorm:"type:text"`
}

type Attempt struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"userId" gorm:"index"`
	Level     string    `json:"level" gorm:"index"`
	Score     int       `json:"score"`
	Total     int       `json:"total"`
	Passed    bool      `json:"passed"`
	CreatedAt time.Time `json:"createdAt"`
}

type Progress struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          string    `json:"userId" gorm:"uniqueIndex"`
	CompletedLesson int       `json:"completedLessons"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

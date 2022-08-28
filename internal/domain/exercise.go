package domain

import "time"

type Exercise struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Questions   []Question `json:"questions"`
}

type Question struct {
	ID            int       `json:"id"`
	ExerciseID    int       `json:"-"`
	Body          string    `json:"body"`
	OptionA       string    `json:"option_a"`
	OptionB       string    `json:"option_b"`
	OptionC       string    `json:"option_c"`
	OptionD       *string   `json:"option_d"`
	CorrectAnswer string    `json:"-"`
	Score         int       `json:"score"`
	CreatorID     int       `json:"-"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Answer struct {
	ID         int
	ExerciseID int
	QuestionID int
	UserID     int
	Answer     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewExercise(title string, description string) *Exercise {
	return &Exercise{
		Title:       title,
		Description: description,
	}
}

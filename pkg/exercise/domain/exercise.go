package domain

import "time"

type Exercise struct {
	ID            string
	Audio         string
	Level         string
	Content       string
	Translated    string
	Vocabulary    string
	CorrectAnswer string
	Options       []string
	CorrectStreak int
	IsFavorite    bool
	IsMastered    bool
	UpdatedAt     time.Time
	NextReviewAt  time.Time
}

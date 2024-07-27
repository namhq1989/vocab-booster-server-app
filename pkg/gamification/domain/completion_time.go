package domain

import (
	"time"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type CompletionTimeRepository interface {
	CreateCompletionTime(ctx *appcontext.AppContext, completionTime CompletionTime) error
}

type CompletionTime struct {
	ID        string
	UserID    string
	Action    Action
	Seconds   int
	CreatedAt time.Time
}

func NewCompletionTime(userID, action string, seconds int) (*CompletionTime, error) {
	if !database.IsValidObjectID(userID) {
		return nil, apperrors.User.InvalidUserID
	}

	if seconds <= 0 {
		return nil, apperrors.Gamification.InvalidCompletionTime
	}

	dAction := ToAction(action)
	if !dAction.IsValid() {
		return nil, apperrors.Common.InvalidAction
	}

	return &CompletionTime{
		ID:        database.NewStringID(),
		UserID:    userID,
		Action:    dAction,
		Seconds:   seconds,
		CreatedAt: manipulation.NowUTC(),
	}, nil
}

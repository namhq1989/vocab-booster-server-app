package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ExerciseHub struct {
	client exercisepb.ExerciseServiceClient
}

func NewExerciseHub(client exercisepb.ExerciseServiceClient) ExerciseHub {
	return ExerciseHub{
		client: client,
	}
}

func (r ExerciseHub) GetUserStats(ctx *appcontext.AppContext, userID, timezone string) (*domain.ExerciseUserStats, error) {
	resp, err := r.client.GetUserStats(ctx.Context(), &exercisepb.GetUserStatsRequest{
		UserId:   userID,
		Timezone: timezone,
	})
	if err != nil {
		return nil, err
	}

	return &domain.ExerciseUserStats{
		Mastered:         int(resp.GetMastered()),
		WaitingForReview: int(resp.GetWaitingForReview()),
	}, nil
}

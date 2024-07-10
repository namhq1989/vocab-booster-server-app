package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/exercisepb"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/infrastructure/mapping"
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

func (r ExerciseHub) GetExercises(ctx *appcontext.AppContext, userID, lang string) ([]domain.Exercise, error) {
	resp, err := r.client.GetUserExercises(ctx.Context(), &exercisepb.GetUserExercisesRequest{
		UserId: userID,
		Lang:   lang,
	})
	if err != nil {
		return nil, apperrors.TransformGrpcError(err)
	}

	var (
		result = make([]domain.Exercise, 0)
		mapper = mapping.ExerciseMapper{}
	)

	for _, e := range resp.GetExercises() {
		exercise, _ := mapper.FromGrpcToDomain(e)
		if exercise != nil {
			result = append(result, *exercise)

		}
	}

	return result, nil
}

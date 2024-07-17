package infrastructure

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/exercisepb"
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

func (r ExerciseHub) AnswerExercise(ctx *appcontext.AppContext, payload domain.AnswerExercisePayload) (*domain.AnswerExerciseResult, error) {
	resp, err := r.client.AnswerExercise(ctx.Context(), &exercisepb.AnswerExerciseRequest{
		UserId:     payload.UserID,
		ExerciseId: payload.ExerciseID,
		IsCorrect:  payload.IsCorrect,
	})
	if err != nil {
		return nil, err
	}

	return &domain.AnswerExerciseResult{
		NextReviewAt: resp.GetNextReviewAt().AsTime(),
	}, nil
}

func (r ExerciseHub) GetExercises(ctx *appcontext.AppContext, userID, lang, collectionID string) ([]domain.Exercise, error) {
	resp, err := r.client.GetUserExercises(ctx.Context(), &exercisepb.GetUserExercisesRequest{
		UserId:       userID,
		Lang:         lang,
		CollectionId: collectionID,
	})
	if err != nil {
		return nil, err
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

func (r ExerciseHub) GetReadyForReviewExercises(ctx *appcontext.AppContext, userID, lang string) ([]domain.Exercise, error) {
	resp, err := r.client.GetUserReadyForReviewExercises(ctx.Context(), &exercisepb.GetUserReadyForReviewExercisesRequest{
		UserId: userID,
		Lang:   lang,
	})
	if err != nil {
		return nil, err
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

func (r ExerciseHub) GetExerciseCollections(ctx *appcontext.AppContext, userID, lang string) ([]domain.ExerciseCollection, error) {
	resp, err := r.client.GetExerciseCollections(ctx.Context(), &exercisepb.GetExerciseCollectionsRequest{
		UserId: userID,
		Lang:   lang,
	})
	if err != nil {
		return nil, err
	}

	var (
		result = make([]domain.ExerciseCollection, 0)
		mapper = mapping.ExerciseCollectionMapper{}
	)

	for _, e := range resp.GetCollections() {
		collection, _ := mapper.FromGrpcToDomain(e)
		if collection != nil {
			result = append(result, *collection)

		}
	}

	return result, nil
}

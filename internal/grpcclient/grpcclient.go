package grpcclient

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/exercisepb"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/gamificationpb"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/userpb"
	"github.com/namhq1989/vocab-booster-server-app/internal/genproto/vocabularypb"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(_ *appcontext.AppContext, addr string) (userpb.UserServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return userpb.NewUserServiceClient(conn), nil
}

func NewExerciseClient(_ *appcontext.AppContext, addr string) (exercisepb.ExerciseServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return exercisepb.NewExerciseServiceClient(conn), nil
}

func NewVocabularyClient(_ *appcontext.AppContext, addr string) (vocabularypb.VocabularyServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return vocabularypb.NewVocabularyServiceClient(conn), nil
}

func NewGamificationClient(_ *appcontext.AppContext, addr string) (gamificationpb.GamificationServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return gamificationpb.NewGamificationServiceClient(conn), nil
}

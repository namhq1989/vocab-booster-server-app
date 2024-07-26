package vocabulary

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/grpcclient"
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/application"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/rest"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Module struct{}

func (Module) Name() string {
	return "VOCABULARY"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	vocabularyGRPCClient, err := grpcclient.NewVocabularyClient(ctx, mono.Config().EndpointVocabularyGrpc)
	if err != nil {
		return err
	}

	var (
		vocabularyHub = infrastructure.NewVocabularyHub(vocabularyGRPCClient)

		// app
		app = application.New(vocabularyHub)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	return nil
}

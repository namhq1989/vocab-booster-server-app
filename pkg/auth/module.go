package auth

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/grpcclient"
	"github.com/namhq1989/vocab-booster-server-app/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/application"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/infrastructure"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/rest"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type Module struct{}

func (Module) Name() string {
	return "AUTH"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	userGRPCClient, err := grpcclient.NewUserClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	var (
		cfg = mono.Config()

		authenticationRepository = infrastructure.NewAuthenticationRepository(mono.Authentication())
		jwtRepository            = infrastructure.NewJwtRepository(mono.JWT())
		userHub                  = infrastructure.NewUserHub(userGRPCClient)

		// app
		app = application.New(authenticationRepository, jwtRepository, userHub)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), cfg.IsEnvRelease); err != nil {
		return err
	}

	return nil
}

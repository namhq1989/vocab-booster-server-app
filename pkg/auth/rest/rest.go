package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	appjwt "github.com/namhq1989/vocab-booster-server-app/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/application"
)

type server struct {
	app          application.Instance
	echo         *echo.Echo
	jwt          appjwt.Operations
	isEnvRelease bool
}

func RegisterServer(_ *appcontext.AppContext, app application.Instance, e *echo.Echo, jwt appjwt.Operations, isEnvRelease bool) error {
	var s = server{
		app:          app,
		echo:         e,
		jwt:          jwt,
		isEnvRelease: isEnvRelease,
	}

	s.registerAuthRoutes()

	return nil
}

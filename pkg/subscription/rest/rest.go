package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/core/appcontext"
	appjwt "github.com/namhq1989/vocab-booster-server-app/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-app/pkg/subscription/application"
)

type server struct {
	app  application.App
	echo *echo.Echo
	jwt  appjwt.Operations
}

func RegisterServer(_ *appcontext.AppContext, app application.App, e *echo.Echo, jwt *appjwt.JWT) error {
	var s = server{
		app:  app,
		echo: e,
		jwt:  jwt,
	}

	s.registerSubscriptionRoutes()

	return nil
}

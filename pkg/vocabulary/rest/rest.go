package rest

import (
	"github.com/labstack/echo/v4"
	appjwt "github.com/namhq1989/vocab-booster-server-app/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/application"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type server struct {
	app  application.Instance
	echo *echo.Echo
	jwt  appjwt.Operations
}

func RegisterServer(_ *appcontext.AppContext, app application.Instance, e *echo.Echo, jwt *appjwt.JWT) error {
	var s = server{
		app:  app,
		echo: e,
		jwt:  jwt,
	}

	s.registerVocabularyRoutes()

	return nil
}

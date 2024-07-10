package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s server) registerAuthRoutes() {
	g := s.echo.Group("/api/auth")

	g.POST("/sign-in-with-google", func(c echo.Context) error {
		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.SignInWithGoogleRequest)
		)

		resp, err := s.app.SignInWithGoogle(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.SignInWithGoogleRequest](next)
	})

	g.GET("/access-token", func(c echo.Context) error {
		if s.isEnvRelease {
			return httprespond.R404(c, nil, nil)
		}

		var (
			ctx = c.Get("ctx").(*appcontext.AppContext)
			req = c.Get("req").(dto.GetAccessTokenByUserIDRequest)
		)

		resp, err := s.app.GetAccessTokenByUserID(ctx, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetAccessTokenByUserIDRequest](next)
	})
}

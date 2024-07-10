package appjwt

import (
	"strings"

	"github.com/namhq1989/vocab-booster-server-app/internal/database"

	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"

	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (j JWT) RequireLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			ctx   = c.Get("ctx").(*appcontext.AppContext)
			token = strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
		)

		ctx.SetUserID(database.NewStringID())
		return next(c)

		claims, err := j.ParseAccessToken(ctx, token)
		if claims == nil || err != nil {
			return httprespond.R401(c, apperrors.Common.Unauthorized, nil)
		}

		ctx.SetUserID(claims.UserID)
		return next(c)
	}
}

package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/exercise/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s server) registerExerciseRoutes() {
	g := s.echo.Group("/api/exercise")

	g.GET("", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetExercisesRequest)
			performerID = ctx.GetUserID()
			lang        = ctx.GetLang()
		)

		resp, err := s.app.GetExercises(ctx, performerID, lang, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetExercisesRequest](next)
	})

	g.GET("/ready-for-review", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetReadyForReviewExercisesRequest)
			performerID = ctx.GetUserID()
			lang        = ctx.GetLang()
		)

		resp, err := s.app.GetReadyForReviewExercises(ctx, performerID, lang, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetReadyForReviewExercisesRequest](next)
	})

	g.POST("/:id/answer", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.AnswerExerciseRequest)
			exerciseID  = c.Param("id")
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.AnswerExercise(ctx, performerID, exerciseID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.AnswerExerciseRequest](next)
	})
}

package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/httprespond"
	"github.com/namhq1989/vocab-booster-server-app/internal/utils/validation"
	"github.com/namhq1989/vocab-booster-server-app/pkg/vocabulary/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

func (s server) registerVocabularyRoutes() {
	g := s.echo.Group("/api/vocabulary")

	g.GET("/search", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.SearchVocabularyRequest)
			performerID = ctx.GetUserID()
			lang        = ctx.GetLang()
		)

		resp, err := s.app.SearchVocabulary(ctx, performerID, lang, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.SearchVocabularyRequest](next)
	})

	g.PATCH("/:id/bookmark", func(c echo.Context) error {
		var (
			ctx          = c.Get("ctx").(*appcontext.AppContext)
			req          = c.Get("req").(dto.BookmarkVocabularyRequest)
			performerID  = ctx.GetUserID()
			vocabularyID = c.Param("id")
		)

		resp, err := s.app.BookmarkVocabulary(ctx, performerID, vocabularyID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.BookmarkVocabularyRequest](next)
	})

	g.GET("/bookmarked", func(c echo.Context) error {
		var (
			ctx         = c.Get("ctx").(*appcontext.AppContext)
			req         = c.Get("req").(dto.GetUserBookmarkedVocabulariesRequest)
			performerID = ctx.GetUserID()
		)

		resp, err := s.app.GetUserBookmarkedVocabularies(ctx, performerID, req)
		if err != nil {
			return httprespond.R400(c, err, nil)
		}

		return httprespond.R200(c, resp)
	}, s.jwt.RequireLoggedIn, func(next echo.HandlerFunc) echo.HandlerFunc {
		return validation.ValidateHTTPPayload[dto.GetUserBookmarkedVocabulariesRequest](next)
	})
}

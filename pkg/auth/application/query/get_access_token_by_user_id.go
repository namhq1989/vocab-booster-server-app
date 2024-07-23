package query

import (
	"github.com/namhq1989/vocab-booster-server-app/internal/database"
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/auth/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type GetAccessTokenByUserIDHandler struct {
	jwtRepository domain.JwtRepository
}

func NewGetAccessTokenByUserIDHandler(jwtRepository domain.JwtRepository) GetAccessTokenByUserIDHandler {
	return GetAccessTokenByUserIDHandler{
		jwtRepository: jwtRepository,
	}
}

func (h GetAccessTokenByUserIDHandler) GetAccessTokenByUserID(ctx *appcontext.AppContext, req dto.GetAccessTokenByUserIDRequest) (*dto.GetAccessTokenByUserIDResponse, error) {
	ctx.Logger().Info("[query] new get access token by user id request", appcontext.Fields{"userID": req.UserID})

	ctx.Logger().Text("validate user id")
	if !database.IsValidObjectID(req.UserID) {
		ctx.Logger().Error("invalid user id", nil, appcontext.Fields{})
		return nil, apperrors.User.InvalidUserID
	}

	// set default timezone to UTC +7
	if req.Timezone == "" {
		req.Timezone = "Asia/Ho_Chi_Minh"
	}

	ctx.Logger().Text("generate new access token")
	token, err := h.jwtRepository.GenerateAccessToken(ctx, req.UserID, req.Timezone)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("done get access token by user id request", appcontext.Fields{"token": token})
	return &dto.GetAccessTokenByUserIDResponse{AccessToken: token}, nil
}

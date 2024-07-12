package command

import (
	apperrors "github.com/namhq1989/vocab-booster-server-app/internal/utils/error"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/domain"
	"github.com/namhq1989/vocab-booster-server-app/pkg/user/dto"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type ChangeAvatarHandler struct {
	userRepository domain.UserRepository
}

func NewChangeAvatarHandler(userRepository domain.UserRepository) ChangeAvatarHandler {
	return ChangeAvatarHandler{
		userRepository: userRepository,
	}
}

func (h ChangeAvatarHandler) ChangeAvatar(ctx *appcontext.AppContext, performerID string, req dto.ChangeAvatarRequest) (*dto.ChangeAvatarResponse, error) {
	ctx.Logger().Info("[command] new change avatar request", appcontext.Fields{"performerID": performerID, "avatar": req.Avatar})

	ctx.Logger().Text("find user by id in db")
	user, err := h.userRepository.FindUserByID(ctx, performerID)
	if err != nil {
		ctx.Logger().Error("failed to find user by id in db", err, appcontext.Fields{})
		return nil, err
	}
	if user == nil {
		ctx.Logger().ErrorText("user not found")
		return nil, apperrors.User.UserNotFound
	}

	ctx.Logger().Text("set avatar")
	user.SetAvatar(req.Avatar)

	ctx.Logger().Text("update user in db")
	if err = h.userRepository.UpdateUser(ctx, *user); err != nil {
		ctx.Logger().Error("failed to update user in db", err, appcontext.Fields{})
		return nil, err
	}

	ctx.Logger().Text("done change avatar request")
	return &dto.ChangeAvatarResponse{}, nil
}

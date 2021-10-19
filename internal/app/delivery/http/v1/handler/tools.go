package handler

import (
	"github.com/DmiAS/bd_course/internal/app/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func extractUserInfo(ctx echo.Context) (*models.UserInfo, error) {
	id, ok := ctx.Get(ID).(uuid.UUID)
	if !ok {
		return nil, errors.Errorf("invalid uuid format = %v", ctx.Get(ID))
	}

	role, ok := ctx.Get(Role).(models.Role)
	if !ok {
		return nil, errors.Errorf("invalid role format = %v", ctx.Get(Role))
	}
	return &models.UserInfo{
		ID:   id,
		Role: role,
	}, nil
}

func extractID(ctx echo.Context) (uuid.UUID, error) {
	sid := ctx.Param("id")
	return uuid.Parse(sid)
}

func canManageAccountData(role models.Role, id, targetID uuid.UUID, grantedRoles ...models.Role) error {
	roleIsAccepted := false
	for i := range grantedRoles {
		if role == grantedRoles[i] {
			roleIsAccepted = true
			break
		}
	}
	if !roleIsAccepted && id != targetID {
		return errors.New("access denied")
	}
	return nil
}

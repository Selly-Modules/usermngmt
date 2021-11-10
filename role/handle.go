package role

import (
	"context"

	"github.com/Selly-Modules/usermngmt/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handle struct {
}

// FindByID ...
func (h Handle) FindByID(ctx context.Context, id primitive.ObjectID) (model.DBRole, error) {
	role, err := h.findByID(ctx, id)
	return role, err
}

// Create ...
func (h Handle) Create(payload model.RoleCreateOptions) error {
	// TODO later
	return nil
}

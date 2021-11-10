package role

import (
	"context"

	"github.com/Selly-Modules/usermngmt/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindByID ...
func FindByID(ctx context.Context, id primitive.ObjectID) (model.DBRole, error) {
	role, err := findByID(ctx, id)
	return role, err
}

// Create ...
func Create(payload model.RoleCreateOptions) error {
	// TODO later
	return nil
}

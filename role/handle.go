package role

import (
	"context"

	"github.com/Selly-Modules/usermngmt/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handle struct {
	Col *mongo.Collection
}

// FindByID ...
func (h Handle) FindByID(ctx context.Context, id primitive.ObjectID) (internal.DBRole, error) {
	role, err := h.findByID(ctx, id)
	return role, err
}

// Create ...
func (h Handle) Create(payload internal.RoleCreateOptions) error {
	// TODO later
	return nil
}

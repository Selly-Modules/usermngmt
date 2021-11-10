package role

import (
	"context"

	"github.com/Selly-Modules/usermngmt/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h Handle) findByID(ctx context.Context, id primitive.ObjectID) (internal.DBRole, error) {
	var (
		doc internal.DBRole
	)
	err := h.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

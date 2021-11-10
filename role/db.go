package role

import (
	"context"

	"github.com/Selly-Modules/usermngmt/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h Handle) findByID(ctx context.Context, id primitive.ObjectID) (model.DBRole, error) {
	var (
		doc model.DBRole
	)
	err := h.Col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

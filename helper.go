package usermngmt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//  getUserCollection ...
func (s Service) getUserCollection() *mongo.Collection {
	return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableUser))
}

func (s Service) haveNameOrPhoneExisted(ctx context.Context, phone, email string) bool {
	var (
		col  = s.getUserCollection()
		user = User{}
	)

	// Find
	cond := bson.M{
		"$or": []bson.M{
			{
				"phone": phone,
			},
			{
				"email": email,
			},
		},
	}
	if err := col.FindOne(ctx, cond).Decode(&user); err != nil {
		logger.Error("usermngmt - findByCondition", logger.LogData{
			"condition": cond,
			"err":       err.Error(),
		})
		return true
	}
	return !user.ID.IsZero()
}

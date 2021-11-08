package usermngmt

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//  getUserCollection ...
func (s Service) getUserCollection() *mongo.Collection {
	return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableUser))
}

//  getRoleCollection ...
func (s Service) getRoleCollection() *mongo.Collection {
	return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableRole))
}

func (s Service) isPhoneNumberOrEmailAlreadyExisted(ctx context.Context, phone, email string) bool {
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

func (s Service) isRoleIDAlreadyExisted(ctx context.Context, roleID primitive.ObjectID) bool {
	var (
		col  = s.getRoleCollection()
		role = Role{}
	)

	// Find
	cond := bson.M{
		"_id": roleID,
	}
	if err := col.FindOne(ctx, cond).Decode(&role); err != nil {
		logger.Error("usermngmt - findRoleByCondition", logger.LogData{
			"condition": cond,
			"err":       err.Error(),
		})
		return false
	}
	return !role.ID.IsZero()
}

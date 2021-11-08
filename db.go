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
	if s.TablePrefix != "" {
		return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableUser))
	}
	return s.DB.Collection(tableUser)
}

//  getRoleCollection ...
func (s Service) getRoleCollection() *mongo.Collection {
	if s.TablePrefix != "" {
		s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableRole))
	}
	return s.DB.Collection(tableRole)
}

func (s Service) isPhoneNumberOrEmailExisted(ctx context.Context, phone, email string) bool {
	var (
		col = s.getUserCollection()
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
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - countUserByCondition", logger.LogData{
			"condition": cond,
			"err":       err.Error(),
		})
		return true
	}
	return total != 0
}

func (s Service) isRoleIDExisted(ctx context.Context, roleID primitive.ObjectID) bool {
	var (
		col = s.getRoleCollection()
	)

	// Find
	cond := bson.M{
		"_id": roleID,
	}
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - countRoleByCondition", logger.LogData{
			"condition": cond,
			"err":       err.Error(),
		})
		return false
	}
	return total != 0
}

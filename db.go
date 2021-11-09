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

func (s Service) userCreate(ctx context.Context, doc User) error {
	var (
		col = s.getUserCollection()
	)

	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		logger.Error("usermngmt - Create", logger.LogData{
			"doc": doc,
			"err": err.Error(),
		})
		return fmt.Errorf("error when create user: %s", err.Error())
	}

	return nil
}

func (s Service) userUpdateOneByCondition(ctx context.Context, cond interface{}, payload interface{}) error {
	var (
		col = s.getUserCollection()
	)

	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("usermngmt - Update", logger.LogData{
			"cond": cond,
			"payload": payload,
			"err": err.Error(),
		})
		return fmt.Errorf("error when update user: %s", err.Error())
	}

	return err
}
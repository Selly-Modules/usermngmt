package user

import (
	"context"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/usermngmt/database"
	"github.com/Selly-Modules/usermngmt/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h Handle) isPhoneNumberOrEmailExisted(ctx context.Context, phone, email string) bool {
	var (
		col = database.GetUserCol()
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

func (h Handle) isRoleIDExisted(ctx context.Context, roleID primitive.ObjectID) bool {
	var (
		col = database.GetRoleCol()
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

func (h Handle) roleFindByID(ctx context.Context, id primitive.ObjectID) (model.DBRole, error) {
	var (
		doc model.DBRole
		col = database.GetRoleCol()
	)
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

func (h Handle) create(ctx context.Context, doc model.DBUser) error {
	var (
		col = database.GetUserCol()
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

func (h Handle) updateOneByCondition(ctx context.Context, cond interface{}, payload interface{}) error {
	var (
		col = database.GetUserCol()
	)
	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("usermngmt - Update", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"err":     err.Error(),
		})
		return fmt.Errorf("error when update user: %s", err.Error())
	}

	return err
}

func (h Handle) findByID(ctx context.Context, id primitive.ObjectID) (model.DBUser, error) {
	var (
		doc model.DBUser
		col = database.GetUserCol()
	)
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

func (h Handle) findByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBUser) {
	var (
		col = database.GetUserCol()
	)
	docs = make([]model.DBUser, 0)

	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("usermngmt - All", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &docs); err != nil {
		logger.Error("usermngmt - All - decode", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return
	}
	return
}

// countByCondition ...
func (h Handle) countByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.GetUserCol()
	)
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Count", logger.LogData{
			"err":  err.Error(),
			"cond": cond,
		})
	}
	return total
}
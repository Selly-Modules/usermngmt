package role

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

func findByID(ctx context.Context, id primitive.ObjectID) (model.DBRole, error) {
	var (
		doc model.DBRole
		col = database.GetRoleCol()
	)
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

func create(ctx context.Context, doc model.DBRole) error {
	var (
		col = database.GetRoleCol()
	)
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		logger.Error("usermngmt - Role - InsertOne", logger.LogData{
			"doc": doc,
			"err": err.Error(),
		})
		return fmt.Errorf("error when create role: %s", err.Error())
	}

	return nil
}

func updateOneByCondition(ctx context.Context, cond interface{}, payload interface{}) error {
	var (
		col = database.GetRoleCol()
	)
	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("usermngmt - Role - UpdateOne", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"err":     err.Error(),
		})
		return fmt.Errorf("error when update role: %s", err.Error())
	}

	return err
}

func findByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBRole) {
	var (
		col = database.GetRoleCol()
	)
	docs = make([]model.DBRole, 0)

	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("usermngmt - Role - Find", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &docs); err != nil {
		logger.Error("usermngmt - Role - Decode", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return
	}
	return
}

// countByCondition ...
func countByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.GetRoleCol()
	)
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Role - CountDocuments", logger.LogData{
			"err":  err.Error(),
			"cond": cond,
		})
	}
	return total
}

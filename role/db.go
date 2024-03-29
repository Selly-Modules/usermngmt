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

func create(ctx context.Context, doc model.DBRole) error {
	var (
		col = database.GetRoleCol()
	)
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		logger.Error("usermngmt - Role - InsertOne", logger.LogData{
			Source:  "usermngmt.create",
			Message: err.Error(),
			Data: map[string]interface{}{
				"doc": doc,
			},
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
			Source:  "usermngmt.updateOneByCondition",
			Message: err.Error(),
			Data: map[string]interface{}{
				"cond":    cond,
				"payload": payload,
			},
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
			Source:  "usermngmt.findByCondition",
			Message: err.Error(),
			Data: map[string]interface{}{
				"cond": cond,
				"opts": opts,
			},
		})
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &docs); err != nil {
		logger.Error("usermngmt - Role - Decode", logger.LogData{
			Source:  "usermngmt.findByCondition",
			Message: err.Error(),
			Data: map[string]interface{}{
				"cond": cond,
				"opts": opts,
			},
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
			Source:  "usermngmt.countByCondition",
			Message: err.Error(),
			Data: map[string]interface{}{
				"cond": cond,
			},
		})
	}
	return total
}

func isRoleIDExisted(ctx context.Context, roleID primitive.ObjectID) bool {
	var (
		col = database.GetRoleCol()
	)
	// Find
	cond := bson.M{
		"_id": roleID,
	}
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Role - CountDocuments", logger.LogData{
			Source:  "usermngmt.isRoleIDExisted",
			Message: err.Error(),
			Data: map[string]interface{}{
				"condition": cond,
			},
		})
		return false
	}
	return total != 0
}

func findByID(ctx context.Context, id primitive.ObjectID) (model.DBRole, error) {
	var (
		doc model.DBRole
		col = database.GetRoleCol()
	)
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

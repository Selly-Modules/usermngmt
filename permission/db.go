package permission

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

func findByID(ctx context.Context, id primitive.ObjectID) (model.DBPermission, error) {
	var (
		doc model.DBPermission
		col = database.GetPermissionCol()
	)
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

func create(ctx context.Context, doc model.DBPermission) error {
	var (
		col = database.GetPermissionCol()
	)
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		logger.Error("usermngmt - Permission - InsertOne", logger.LogData{
			"doc": doc,
			"err": err.Error(),
		})
		return fmt.Errorf("error when create permission: %s", err.Error())
	}

	return nil
}

func updateOneByCondition(ctx context.Context, cond interface{}, payload interface{}) error {
	var (
		col = database.GetPermissionCol()
	)
	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("usermngmt - Permission - UpdateOne", logger.LogData{
			"cond":    cond,
			"payload": payload,
			"err":     err.Error(),
		})
		return fmt.Errorf("error when update permission: %s", err.Error())
	}

	return err
}

func deleteOneByCondition(ctx context.Context, cond interface{}) error {
	var (
		col = database.GetPermissionCol()
	)
	_, err := col.DeleteOne(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Permission - DeleteOne", logger.LogData{
			"cond": cond,
			"err":  err.Error(),
		})
		return fmt.Errorf("error when delete permission: %s", err.Error())
	}

	return err
}

func findByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBPermission) {
	var (
		col = database.GetPermissionCol()
	)
	docs = make([]model.DBPermission, 0)

	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("usermngmt - Permission - Find", logger.LogData{
			"cond": cond,
			"opts": opts,
			"err":  err.Error(),
		})
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &docs); err != nil {
		logger.Error("usermngmt - Permission - Decode", logger.LogData{
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
		col = database.GetPermissionCol()
	)
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Permission - CountDocuments", logger.LogData{
			"err":  err.Error(),
			"cond": cond,
		})
	}
	return total
}

func isPermissionIDExisted(ctx context.Context, permissionID primitive.ObjectID) bool {
	var (
		col = database.GetPermissionCol()
	)
	// Find
	cond := bson.M{
		"_id": permissionID,
	}
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Permission - CountDocuments", logger.LogData{
			"condition": cond,
			"err":       err.Error(),
		})
		return false
	}
	return total != 0
}

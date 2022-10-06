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

func isPhoneNumberExisted(ctx context.Context, phone string) bool {
	var (
		col = database.GetUserCol()
	)
	// Find
	cond := bson.M{
		"phone":   phone,
		"deleted": false,
	}
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - User - CountDocuments", logger.LogData{
			Source:  "usermngmt.user.isPhoneNumberExisted",
			Message: err.Error(),
			Data:    cond,
		})
		return true
	}
	return total != 0
}

func isEmailExisted(ctx context.Context, email string) bool {
	var (
		col = database.GetUserCol()
	)
	// Find
	cond := bson.M{
		"email":   email,
		"deleted": false,
	}
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - User - CountDocuments", logger.LogData{
			Source:  "usermngmt.user.isEmailExisted",
			Message: err.Error(),
			Data:    cond,
		})
		return true
	}
	return total != 0
}

func isRoleExisted(ctx context.Context, roleID primitive.ObjectID) bool {
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
			Source:  "usermngmt.user.isRoleExisted",
			Message: err.Error(),
			Data:    cond,
		})
	}
	return total != 0
}

func roleFindByID(ctx context.Context, id primitive.ObjectID) (model.DBRole, error) {
	var (
		doc model.DBRole
		col = database.GetRoleCol()
	)
	err := col.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)
	return doc, err
}

func roleFindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBRole) {
	var (
		col = database.GetRoleCol()
	)
	docs = make([]model.DBRole, 0)

	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("usermngmt - Role - Find", logger.LogData{
			Source:  "usermngmt.user.roleFindByCondition",
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
			Source:  "usermngmt.user.roleFindByCondition",
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

func permissionFindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBPermission) {
	var (
		col = database.GetPermissionCol()
	)
	docs = make([]model.DBPermission, 0)

	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("usermngmt - Permission - Find", logger.LogData{
			Source:  "usermngmt.user.permissionFindByCondition",
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
		logger.Error("usermngmt - Permission - Decode", logger.LogData{
			Source:  "usermngmt.user.permissionFindByCondition",
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

// permissionCountByCondition ...
func permissionCountByCondition(ctx context.Context, cond interface{}) int64 {
	var (
		col = database.GetPermissionCol()
	)
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Permission - CountDocuments", logger.LogData{
			Source:  "usermngmt.user.permissionCountByCondition",
			Message: err.Error(),
			Data:    cond,
		})
	}
	return total
}

func create(ctx context.Context, doc model.DBUser) error {
	var (
		col = database.GetUserCol()
	)
	_, err := col.InsertOne(ctx, doc)
	if err != nil {
		logger.Error("usermngmt - User - InsertOne", logger.LogData{
			Source:  "usermngmt.user.create",
			Message: err.Error(),
			Data:    doc,
		})
		return fmt.Errorf("error when create user: %s", err.Error())
	}

	return nil
}

func updateOneByCondition(ctx context.Context, cond interface{}, payload interface{}) error {
	var (
		col = database.GetUserCol()
	)
	_, err := col.UpdateOne(ctx, cond, payload)
	if err != nil {
		logger.Error("usermngmt - User - UpdateOne", logger.LogData{
			Source:  "usermngmt.user.updateOneByCondition",
			Message: err.Error(),
			Data: map[string]interface{}{
				"cond":    cond,
				"payload": payload,
			},
		})
		return fmt.Errorf("error when update user: %s", err.Error())
	}

	return err
}

func updateManyByCondition(ctx context.Context, cond interface{}, payload interface{}) error {
	var (
		col = database.GetUserCol()
	)
	_, err := col.UpdateMany(ctx, cond, payload)
	if err != nil {
		logger.Error("usermngmt - User - UpdateMany", logger.LogData{
			Source:  "usermngmt.user.updateManyByCondition",
			Message: err.Error(),
			Data: map[string]interface{}{
				"cond":    cond,
				"payload": payload,
			},
		})
		return fmt.Errorf("error when update user: %s", err.Error())
	}

	return err
}

func findByID(ctx context.Context, id primitive.ObjectID) (model.DBUser, error) {
	var (
		doc model.DBUser
		col = database.GetUserCol()
	)
	err := col.FindOne(ctx, bson.M{"_id": id, "deleted": false}).Decode(&doc)
	return doc, err
}

func findByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBUser) {
	var (
		col = database.GetUserCol()
	)
	docs = make([]model.DBUser, 0)

	cursor, err := col.Find(ctx, cond, opts...)
	if err != nil {
		logger.Error("usermngmt - User - Find", logger.LogData{
			Source:  "usermngmt.user.findByCondition",
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
		logger.Error("usermngmt - User - Decode", logger.LogData{
			Source:  "usermngmt.user.findByCondition",
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
		col = database.GetUserCol()
	)
	total, err := col.CountDocuments(ctx, cond)
	if err != nil {
		logger.Error("usermngmt - Count", logger.LogData{
			Source:  "usermngmt.user.countByCondition",
			Message: err.Error(),
			Data: map[string]interface{}{
				"cond": cond,
			},
		})
	}
	return total
}

func findOneByCondition(ctx context.Context, cond interface{}) (model.DBUser, error) {
	var (
		col = database.GetUserCol()
		doc model.DBUser
	)
	err := col.FindOne(ctx, cond).Decode(&doc)
	return doc, err
}

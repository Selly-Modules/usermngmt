package cache

import (
	"context"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/usermngmt/database"
	"github.com/Selly-Modules/usermngmt/model"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func roleFindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBRole) {
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

func permissionFindByCondition(ctx context.Context, cond interface{}, opts ...*options.FindOptions) (docs []model.DBPermission) {
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

package permission

import (
	"context"
	"errors"
	"sync"

	"github.com/Selly-Modules/mongodb"
	"github.com/Selly-Modules/usermngmt/internal"
	"github.com/Selly-Modules/usermngmt/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Create ...
func Create(payload model.PermissionCreateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err := payload.Validate(); err != nil {
		return err
	}

	// New permission data from payload
	doc := newPermission(payload)

	// Create permission
	if err := create(ctx, doc); err != nil {
		return err
	}

	return nil
}

// newPermission ...
func newPermission(payload model.PermissionCreateOptions) model.DBPermission {
	timeNow := internal.Now()
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	return model.DBPermission{
		ID:        mongodb.NewObjectID(),
		Name:      payload.Name,
		Code:      internal.GenerateCode(payload.Name),
		RoleID:    roleID,
		Desc:      payload.Desc,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

// Update ...
func Update(permissionID string, payload model.PermissionUpdateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err := payload.Validate(); err != nil {
		return err
	}

	// Validate permissionID
	id, isValid := mongodb.NewIDFromString(permissionID)
	if !isValid {
		return errors.New("invalid permission id data")
	}

	// Setup condition
	cond := bson.M{
		"_id": id,
	}

	// Setup update data
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	updateData := bson.M{
		"$set": bson.M{
			"name":      payload.Name,
			"code":      internal.GenerateCode(payload.Name),
			"roleId":    roleID,
			"desc":      payload.Desc,
			"updatedAt": internal.Now(),
		},
	}

	// Update
	if err := updateOneByCondition(ctx, cond, updateData); err != nil {
		return err
	}

	return nil
}

// All ...
func All(queryParams model.PermissionAllQuery) (r model.PermissionAll) {
	var (
		ctx  = context.Background()
		wg   sync.WaitGroup
		cond = bson.M{}
	)
	query := model.CommonQuery{
		Page:  queryParams.Page,
		Limit: queryParams.Limit,
		Sort:  bson.M{"createdAt": -1},
	}

	// Assign condition
	query.SetDefaultLimit()

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := findByCondition(ctx, cond, query.GetFindOptionsUsingPage())
		r.List = getResponseList(docs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Total = countByCondition(ctx, cond)
	}()

	wg.Wait()

	return
}

func getResponseList(permissions []model.DBPermission) []model.Permission {
	res := make([]model.Permission, 0)
	for _, permission := range permissions {
		res = append(res, model.Permission{
			ID:        permission.ID.Hex(),
			Name:      permission.Name,
			Code:      permission.Code,
			RoleID:    permission.RoleID.Hex(),
			Desc:      permission.Desc,
			CreatedAt: permission.CreatedAt,
			UpdatedAt: permission.UpdatedAt,
		})
	}

	return res
}

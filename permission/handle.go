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
func Create(payload model.PermissionCreateOptions) (result string, err error) {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err = payload.Validate(); err != nil {
		return
	}

	// New permission data from payload
	doc := newPermission(payload)

	// Create permission
	if err = create(ctx, doc); err != nil {
		return
	}

	result = doc.ID.Hex()
	return
}

// newPermission ...
func newPermission(payload model.PermissionCreateOptions) model.DBPermission {
	timeNow := internal.Now()
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	return model.DBPermission{
		ID:        mongodb.NewObjectID(),
		Name:      payload.Name,
		Code:      payload.Code,
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

	// Find permissionID exists or not
	id, isValid := mongodb.NewIDFromString(permissionID)
	if !isValid {
		return errors.New(internal.ErrorInvalidPermission)
	}
	if !isPermissionIDExisted(ctx, id) {
		return errors.New(internal.ErrorNotFoundPermission)
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
			"code":      payload.Code,
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

// Delete ...
func Delete(permissionID string) error {
	var (
		ctx = context.Background()
	)

	// Find permissionID exists or not
	id, isValid := mongodb.NewIDFromString(permissionID)
	if !isValid {
		return errors.New(internal.ErrorInvalidPermission)
	}
	if !isPermissionIDExisted(ctx, id) {
		return errors.New(internal.ErrorNotFoundPermission)
	}

	// Delete
	if err := deleteOneByCondition(ctx, bson.M{"_id": id}); err != nil {
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
		Page:   queryParams.Page,
		Limit:  queryParams.Limit,
		Sort:   queryParams.Sort,
		RoleID: queryParams.RoleID,
	}

	// Assign condition
	query.SetDefaultLimit()
	query.AssignRoleID(cond)

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

	r.Limit = query.Limit
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

package role

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
func Create(payload model.RoleCreateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err := payload.Validate(); err != nil {
		return err
	}

	// New role data from payload
	doc := newRole(payload)

	// Create role
	if err := create(ctx, doc); err != nil {
		return err
	}

	return nil
}

// newRole ...
func newRole(payload model.RoleCreateOptions) model.DBRole {
	timeNow := internal.Now()
	return model.DBRole{
		ID:        mongodb.NewObjectID(),
		Name:      payload.Name,
		Code:      internal.GetCode(payload.Name),
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
}

// Update ...
func Update(roleID string, payload model.RoleUpdateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err := payload.Validate(); err != nil {
		return err
	}

	// Validate roleID
	id, isValid := mongodb.NewIDFromString(roleID)
	if !isValid {
		return errors.New("invalid role id data")
	}

	// Setup condition
	cond := bson.M{
		"_id": id,
	}

	// Setup update data
	updateData := bson.M{
		"$set": bson.M{
			"name":      payload.Name,
			"code":      internal.GetCode(payload.Name),
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
func All(queryParams model.RoleAllQuery) (r model.RoleAll) {
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

func getResponseList(roles []model.DBRole) []model.Role {
	res := make([]model.Role, 0)
	for _, role := range roles {
		res = append(res, model.Role{
			ID:        role.ID.Hex(),
			Name:      role.Name,
			Code:      role.Code,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
		})
	}

	return res
}

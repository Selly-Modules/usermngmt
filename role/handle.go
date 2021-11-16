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
func Create(payload model.RoleCreateOptions) (result string, err error) {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err = payload.Validate(); err != nil {
		return
	}

	// New role data from payload
	doc := newRole(payload)

	// Create role
	if err = create(ctx, doc); err != nil {
		return
	}

	result = doc.ID.Hex()
	return
}

// newRole ...
func newRole(payload model.RoleCreateOptions) model.DBRole {
	timeNow := internal.Now()
	return model.DBRole{
		ID:        mongodb.NewObjectID(),
		Name:      payload.Name,
		Code:      internal.GenerateCode(payload.Name),
		Level:     payload.Level,
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

	// Find roleID exists or not
	if !isRoleIDExisted(ctx, id) {
		return errors.New("role not found")
	}

	// Setup condition
	cond := bson.M{
		"_id": id,
	}

	// Setup update data
	updateData := bson.M{
		"$set": bson.M{
			"name":      payload.Name,
			"code":      internal.GenerateCode(payload.Name),
			"level":     payload.Level,
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
		Sort:  queryParams.Sort,
	}

	// Assign condition
	query.SetDefaultLimit()

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := findByCondition(ctx, cond, query.GetFindOptionsUsingPage())
		res := make([]model.Role, 0)
		for _, doc := range docs {
			res = append(res, getResponse(doc))
		}
		r.List = res
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Total = countByCondition(ctx, cond)
	}()

	wg.Wait()

	return
}

func getResponse(role model.DBRole) model.Role {
	return model.Role{
		ID:        role.ID.Hex(),
		Name:      role.Name,
		Code:      role.Code,
		Level:     role.Level,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

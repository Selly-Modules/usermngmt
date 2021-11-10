package user

import (
	"context"
	"errors"
	"sync"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"github.com/Selly-Modules/usermngmt/internal"
	"github.com/Selly-Modules/usermngmt/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Handle struct {
}

// Create ...
func (h Handle) Create(payload model.UserCreateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err := payload.Validate(); err != nil {
		return err
	}

	//  Find roleID exists or not
	roleID, isValid := mongodb.NewIDFromString(payload.RoleID)
	if !isValid {
		return errors.New("invalid role id data")
	}
	if !h.isRoleIDExisted(ctx, roleID) {
		return errors.New("role id does not exist")
	}

	// Find phone number,email exists or not
	if h.isPhoneNumberOrEmailExisted(ctx, payload.Phone, payload.Email) {
		return errors.New("phone number or email already existed")
	}

	// New user data from payload
	doc, err := newUser(payload)
	if err != nil {
		return err
	}

	// Create user
	if err = h.create(ctx, doc); err != nil {
		return err
	}

	return nil
}

// newUser ...
func newUser(payload model.UserCreateOptions) (result model.DBUser, err error) {
	timeNow := internal.Now()
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	return model.DBUser{
		ID:             mongodb.NewObjectID(),
		Name:           payload.Name,
		SearchString:   internal.GetSearchString(payload.Name, payload.Phone, payload.Email),
		Phone:          payload.Phone,
		Email:          payload.Email,
		HashedPassword: internal.HashPassword(payload.Password),
		Status:         payload.Status,
		RoleID:         roleID,
		Other:          payload.Other,
		CreatedAt:      timeNow,
		UpdatedAt:      timeNow,
	}, nil
}

// All ...
func (h Handle) All(queryParams model.UserAllQuery) (r model.UserAll) {
	var (
		ctx  = context.Background()
		wg   sync.WaitGroup
		cond = bson.M{}
	)
	query := model.CommonQuery{
		Page:    queryParams.Page,
		Limit:   queryParams.Limit,
		Keyword: queryParams.Keyword,
		RoleID:  queryParams.RoleID,
		Status:  queryParams.Status,
		Sort:    bson.M{"createdAt": -1},
	}

	// Assign condition
	query.SetDefaultLimit()
	query.AssignKeyword(cond)
	query.AssignRoleID(cond)
	query.AssignStatus(cond)

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := h.findByCondition(ctx, cond, query.GetFindOptionsUsingPage())
		r.List = h.getResponseList(ctx, docs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Total = h.countByCondition(ctx, cond)
	}()

	wg.Wait()

	return
}

func (h Handle) getResponseList(ctx context.Context, users []model.DBUser) []model.User {
	res := make([]model.User, 0)

	for _, user := range users {
		roleRaw, _ := h.roleFindByID(ctx, user.RoleID)
		res = append(res, model.User{
			ID:     user.ID.Hex(),
			Name:   user.Name,
			Phone:  user.Phone,
			Email:  user.Email,
			Status: user.Status,
			Role: model.RoleShort{
				ID:      roleRaw.ID.Hex(),
				Name:    roleRaw.Name,
				IsAdmin: roleRaw.IsAdmin,
			},
			Other:     user.Other,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return res
}

// UpdateByUserID ...
func (h Handle) UpdateByUserID(userID string, payload model.UserUpdateOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err := payload.Validate(); err != nil {
		return err
	}

	//  Find roleID exists or not
	roleID, isValid := mongodb.NewIDFromString(payload.RoleID)
	if !isValid {
		return errors.New("invalid role id data")
	}
	if !h.isRoleIDExisted(ctx, roleID) {
		return errors.New("role id does not exist")
	}

	// Find phone number,email exists or not
	if h.isPhoneNumberOrEmailExisted(ctx, payload.Phone, payload.Email) {
		return errors.New("phone number or email already existed")
	}

	// Setup condition
	id, _ := mongodb.NewIDFromString(userID)
	cond := bson.M{
		"_id": id,
	}

	// Setup update data
	updateData := bson.M{
		"$set": bson.M{
			"name":         payload.Name,
			"searchString": internal.GetSearchString(payload.Name, payload.Phone, payload.Email),
			"phone":        payload.Phone,
			"email":        payload.Email,
			"roleId":       roleID,
			"other":        payload.Other,
			"updatedAt":    internal.Now(),
		},
	}

	// Update
	if err := h.updateOneByCondition(ctx, cond, updateData); err != nil {
		return err
	}

	return nil
}

// ChangeUserPassword ...
func (h Handle) ChangeUserPassword(userID string, opt model.ChangePasswordOptions) error {
	var (
		ctx = context.Background()
	)

	// Validate payload
	err := opt.Validate()
	if err != nil {
		return err
	}

	// Validate userID
	if _, isValid := mongodb.NewIDFromString(userID); !isValid {
		logger.Error("usermngmt - ChangePassword: invalid userID data", logger.LogData{
			"payload": opt,
			"userID":  userID,
		})
		return errors.New("invalid user id data")
	}

	// Find user
	id, _ := mongodb.NewIDFromString(userID)
	user, _ := h.findByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New("user not found")
	}

	// Check old password
	if isValid := internal.CheckPasswordHash(opt.OldPassword, user.HashedPassword); !isValid {
		return errors.New("the password is incorrect")
	}

	// Update password
	if err = h.updateOneByCondition(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"hashedPassword": internal.HashPassword(opt.NewPassword),
			"updatedAt":      internal.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}

// ChangeUserStatus ...
func (h Handle) ChangeUserStatus(userID, newStatus string) error {
	var (
		ctx = context.Background()
	)

	// Validate userID
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		return errors.New("invalid user id data")
	}

	// Update status
	if err := h.updateOneByCondition(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"status":    newStatus,
			"updatedAt": internal.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}

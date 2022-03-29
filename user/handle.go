package user

import (
	"context"
	"errors"
	"sync"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"github.com/Selly-Modules/usermngmt/cache"
	"github.com/Selly-Modules/usermngmt/config"
	"github.com/Selly-Modules/usermngmt/internal"
	"github.com/Selly-Modules/usermngmt/model"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create ...
func Create(payload model.UserCreateOptions) (result string, err error) {
	var (
		ctx = context.Background()
	)

	// Validate payload
	if err = payload.Validate(); err != nil {
		return
	}

	//  Find roleID exists or not
	roleID, isValid := mongodb.NewIDFromString(payload.RoleID)
	if !isValid {
		err = errors.New(internal.ErrorInvalidRole)
		return
	}
	if !isRoleExisted(ctx, roleID) {
		err = errors.New(internal.ErrorNotFoundRole)
		return
	}

	// Find phone number,email exists or not
	if config.GetInstance().PhoneNumberIsUnique {
		if isPhoneNumberExisted(ctx, payload.Phone) {
			err = errors.New(internal.ErrorAlreadyExistedPhoneNumber)
			return
		}
	}
	if config.GetInstance().EmailIsUnique {
		if isEmailExisted(ctx, payload.Email) {
			err = errors.New(internal.ErrorAlreadyExistedEmail)
			return
		}
	}

	// New user data from payload
	doc := newUser(payload)

	// Create user
	if err = create(ctx, doc); err != nil {
		return
	}

	result = doc.ID.Hex()
	return
}

// newUser ...
func newUser(payload model.UserCreateOptions) model.DBUser {
	timeNow := internal.Now()
	roleID, _ := mongodb.NewIDFromString(payload.RoleID)
	return model.DBUser{
		ID:                      mongodb.NewObjectID(),
		Name:                    payload.Name,
		SearchString:            internal.GetSearchString(payload.Name, payload.Phone, payload.Email),
		Phone:                   payload.Phone,
		Email:                   payload.Email,
		HashedPassword:          internal.HashPassword(payload.Password),
		RequireToChangePassword: payload.RequireToChangePassword,
		Status:                  payload.Status,
		RoleID:                  roleID,
		Other:                   payload.Other,
		Avatar:                  payload.Avatar,
		CreatedAt:               timeNow,
		UpdatedAt:               timeNow,
	}
}

// FindUser ...
func FindUser(userID string) (r model.User, err error) {
	var (
		ctx = context.Background()
	)

	// Find user exists or not
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		err = errors.New(internal.ErrorInvalidUser)
		return
	}
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		err = errors.New(internal.ErrorNotFoundUser)
		return
	}

	r = getResponse(ctx, user)
	return
}

// FindUserByEmail ...
func FindUserByEmail(email string) (r model.User, err error) {
	var (
		ctx = context.Background()
	)

	// Find user exists or not
	if email == "" {
		err = errors.New(internal.ErrorInvalidEmail)
		return
	}
	user, _ := findOneByCondition(ctx, bson.M{"email": email})
	if user.ID.IsZero() {
		err = errors.New(internal.ErrorNotFoundUser)
		return
	}

	r = getResponse(ctx, user)
	return
}

// GetHashedPassword ...
func GetHashedPassword(userID string) (result string, err error) {
	var (
		ctx = context.Background()
	)

	// Find user exists or not
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		err = errors.New(internal.ErrorInvalidUser)
		return
	}
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		err = errors.New(internal.ErrorNotFoundUser)
		return
	}

	result = user.HashedPassword
	return
}

// All ...
func All(queryParams model.UserAllQuery) (r model.UserAll) {
	var (
		ctx  = context.Background()
		wg   sync.WaitGroup
		cond = bson.M{}
	)
	if queryParams.Cond != nil {
		cond = queryParams.Cond
	}

	query := model.CommonQuery{
		Page:    queryParams.Page,
		Limit:   queryParams.Limit,
		Keyword: queryParams.Keyword,
		RoleID:  queryParams.RoleID,
		Status:  queryParams.Status,
		Sort:    queryParams.Sort,
		Other:   queryParams.Other,
	}

	// Assign condition
	query.SetDefaultLimit()
	query.AssignKeyword(cond)
	query.AssignRoleID(cond)
	query.AssignStatus(cond)
	query.AssignDeleted(cond)
	query.AssignOther(cond)
	cond["deleted"] = false

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := findByCondition(ctx, cond, query.GetFindOptionsUsingPage())
		res := make([]model.User, 0)
		for _, doc := range docs {
			res = append(res, getResponse(ctx, doc))
		}
		r.List = res
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

// GetUsersByPermission ...
func GetUsersByPermission(queryParams model.UserByPermissionQuery) (r model.UserAll) {
	var (
		ctx   = context.Background()
		wg    sync.WaitGroup
		cond  = bson.M{}
		roles = make([]primitive.ObjectID, 0)
	)
	// Validate query
	if err := queryParams.Validate(); err != nil {
		return
	}

	if queryParams.Cond != nil {
		cond = queryParams.Cond
	}

	// Get role by permission
	permissions := permissionFindByCondition(ctx, bson.M{"code": queryParams.Permission})
	for _, value := range permissions {
		roles = append(roles, value.RoleID)
	}

	// Get admin role
	adminRoles := roleFindByCondition(ctx, bson.M{"isAdmin": true})
	for _, value := range adminRoles {
		roles = append(roles, value.ID)
	}

	if len(roles) < 0 {
		return
	}

	query := model.CommonQuery{
		Page:    queryParams.Page,
		Limit:   queryParams.Limit,
		Keyword: queryParams.Keyword,
		Status:  queryParams.Status,
		Sort:    queryParams.Sort,
		Other:   queryParams.Other,
		RoleIDs: roles,
	}

	// Assign condition
	query.SetDefaultLimit()
	query.AssignKeyword(cond)
	query.AssignRoleIDs(cond)
	query.AssignStatus(cond)
	query.AssignDeleted(cond)
	query.AssignOther(cond)
	cond["deleted"] = false

	wg.Add(1)
	go func() {
		defer wg.Done()
		docs := findByCondition(ctx, cond, query.GetFindOptionsUsingPage())
		res := make([]model.User, 0)
		for _, doc := range docs {
			res = append(res, getResponse(ctx, doc))
		}
		r.List = res
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

// Count  ...
func Count(queryParams model.UserCountQuery) int64 {
	var (
		ctx  = context.Background()
		cond = bson.M{}
	)
	query := model.CommonQuery{
		RoleID: queryParams.RoleID,
		Other:  queryParams.Other,
	}

	// Assign condition
	query.AssignRoleID(cond)
	query.AssignOther(cond)

	return countByCondition(ctx, cond)
}

func getResponse(ctx context.Context, user model.DBUser) model.User {
	roleRaw, _ := roleFindByID(ctx, user.RoleID)
	return model.User{
		ID:     user.ID.Hex(),
		Name:   user.Name,
		Phone:  user.Phone,
		Email:  user.Email,
		Status: user.Status,
		Role: model.RoleShort{
			ID:      roleRaw.ID.Hex(),
			Name:    roleRaw.Name,
			Level:   roleRaw.Level,
			IsAdmin: roleRaw.IsAdmin,
		},
		RequireToChangePassword: user.RequireToChangePassword,
		Avatar:                  user.Avatar,
		Other:                   user.Other,
		CreatedAt:               user.CreatedAt,
		UpdatedAt:               user.UpdatedAt,
	}
}

// UpdateByUserID ...
func UpdateByUserID(userID string, payload model.UserUpdateOptions) error {
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
		return errors.New(internal.ErrorInvalidRole)
	}
	if !isRoleExisted(ctx, roleID) {
		return errors.New(internal.ErrorNotFoundRole)
	}

	// Find user exists or not
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		return errors.New(internal.ErrorInvalidUser)
	}
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New(internal.ErrorNotFoundUser)
	}

	// Find phone number,email exists or not
	if config.GetInstance().PhoneNumberIsUnique {
		if user.Phone != payload.Phone {
			if isPhoneNumberExisted(ctx, payload.Phone) {
				return errors.New(internal.ErrorAlreadyExistedPhoneNumber)
			}
		}
	}
	if config.GetInstance().EmailIsUnique {
		if user.Email != payload.Email {
			if isEmailExisted(ctx, payload.Email) {
				return errors.New(internal.ErrorAlreadyExistedEmail)
			}
		}
	}

	// Setup condition
	cond := bson.M{
		"_id": id,
	}

	// Setup Set operator
	setOperator := bson.M{
		"name":         payload.Name,
		"searchString": internal.GetSearchString(payload.Name, payload.Phone, payload.Email),
		"phone":        payload.Phone,
		"email":        payload.Email,
		"roleId":       roleID,
		"updatedAt":    internal.Now(),
	}
	if len(payload.Other) > 0 {
		for key, value := range payload.Other {
			setOperator["other."+key] = value
		}
	}

	// Update
	if err := updateOneByCondition(ctx, cond, bson.M{
		"$set": setOperator,
	}); err != nil {
		return err
	}

	return nil
}

// ChangeUserPassword ...
func ChangeUserPassword(userID string, opt model.ChangePasswordOptions) error {
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
		return errors.New(internal.ErrorInvalidUser)
	}

	// Find user
	id, _ := mongodb.NewIDFromString(userID)
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New(internal.ErrorNotFoundUser)
	}

	// Check old password
	if isValid := internal.CheckPasswordHash(opt.OldPassword, user.HashedPassword); !isValid {
		return errors.New(internal.ErrorIncorrectPassword)
	}

	// Update password
	if err = updateOneByCondition(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"hashedPassword":          internal.HashPassword(opt.NewPassword),
			"requireToChangePassword": false,
			"updatedAt":               internal.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}

// ResetUserPassword ...
func ResetUserPassword(userID string, password string) error {
	var (
		ctx = context.Background()
	)

	// Validate Password
	if password == "" {
		return errors.New(internal.ErrorInvalidPassword)
	}

	// Validate userID
	if _, isValid := mongodb.NewIDFromString(userID); !isValid {
		return errors.New(internal.ErrorInvalidUser)
	}

	// Find user
	id, _ := mongodb.NewIDFromString(userID)
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New(internal.ErrorNotFoundUser)
	}

	// Update password
	if err := updateOneByCondition(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"hashedPassword":          internal.HashPassword(password),
			"requireToChangePassword": false,
			"updatedAt":               internal.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}

// ResetAndRequireToChangeUserPassword ...
func ResetAndRequireToChangeUserPassword(userID string, password string) error {
	var (
		ctx = context.Background()
	)

	// Validate Password
	if password == "" {
		return errors.New(internal.ErrorInvalidPassword)
	}

	// Validate userID
	if _, isValid := mongodb.NewIDFromString(userID); !isValid {
		return errors.New(internal.ErrorInvalidUser)
	}

	// Find user
	id, _ := mongodb.NewIDFromString(userID)
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New(internal.ErrorNotFoundUser)
	}

	// Update password
	if err := updateOneByCondition(ctx, bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"hashedPassword":          internal.HashPassword(password),
			"requireToChangePassword": true,
			"updatedAt":               internal.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}

// ChangeUserStatus ...
func ChangeUserStatus(userID, newStatus string) error {
	var (
		ctx = context.Background()
	)

	// Validate userID
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		return errors.New(internal.ErrorInvalidUser)
	}
	if user, _ := findByID(ctx, id); user.ID.IsZero() {
		return errors.New(internal.ErrorNotFoundUser)
	}

	// Update status
	if err := updateOneByCondition(ctx, bson.M{"_id": id}, bson.M{
		"$set": bson.M{
			"status":    newStatus,
			"updatedAt": internal.Now(),
		},
	}); err != nil {
		return err
	}

	return nil
}

// ChangeAllUsersStatus ...
func ChangeAllUsersStatus(roleID, status string) error {
	var (
		ctx = context.Background()
	)

	// Validate roleID
	id, isValid := mongodb.NewIDFromString(roleID)
	if !isValid {
		return errors.New(internal.ErrorInvalidRole)
	}
	if !isRoleExisted(ctx, id) {
		return errors.New(internal.ErrorNotFoundRole)
	}

	// Setup condition
	cond := bson.M{
		"roleId": id,
	}

	// Setup update data
	updateData := bson.M{
		"$set": bson.M{
			"status":    status,
			"updatedAt": internal.Now(),
		},
	}

	// Update
	if err := updateManyByCondition(ctx, cond, updateData); err != nil {
		return err
	}

	return nil
}

// LoginWithEmailAndPassword ...
func LoginWithEmailAndPassword(email, password string) (result model.User, err error) {
	var (
		ctx = context.Background()
	)

	// Validate email, password
	if email == "" {
		err = errors.New(internal.ErrorInvalidEmail)
		return
	}
	if password == "" {
		err = errors.New(internal.ErrorInvalidPassword)
		return
	}

	// Find user
	user, _ := findOneByCondition(ctx, bson.M{
		"email":   email,
		"deleted": false,
	})
	if user.ID.IsZero() {
		err = errors.New(internal.ErrorNotFoundUser)
		return
	}

	// Check Password
	if !internal.CheckPasswordHash(password, user.HashedPassword) {
		err = errors.New(internal.ErrorIncorrectPassword)
		return
	}

	result = getResponse(ctx, user)
	return
}

// HasPermission ...
func HasPermission(userID, permission string) (result bool) {
	var (
		ctx = context.Background()
	)

	// Validate userID, permission
	if userID == "" || permission == "" {
		logger.Error("usermngmt - HasPermission: email or password cannot be empty", logger.LogData{
			"userID":     userID,
			"permission": permission,
		})
		return
	}
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		logger.Error("usermngmt - HasPermission: invalid user id", logger.LogData{
			"userID":     userID,
			"permission": permission,
		})
		return
	}

	// Find user
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		logger.Error("usermngmt - HasPermission: user not found", logger.LogData{
			"userID":     userID,
			"permission": permission,
		})
		return
	}

	return checkUserHasPermissionFromCache(user.RoleID, permission)
}

func checkUserHasPermissionFromCache(roleID primitive.ObjectID, permission string) bool {
	cachedRole := cache.GetCachedRole(roleID.Hex())

	// Check permission
	if cachedRole.IsAdmin {
		return true
	}
	if _, isValid := funk.FindString(cachedRole.Permissions, func(s string) bool {
		return s == permission
	}); isValid {
		return true
	}

	return false
}

// UpdateAvatar ...
func UpdateAvatar(userID string, avatar interface{}) error {
	var (
		ctx = context.Background()
	)

	if avatar == nil {
		return errors.New(internal.ErrorInvalidAvatar)
	}

	// Find user exists or not
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		return errors.New(internal.ErrorInvalidUser)
	}
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New(internal.ErrorNotFoundUser)
	}

	// Setup condition
	cond := bson.M{
		"_id": id,
	}

	// Setup update data
	updateData := bson.M{
		"$set": bson.M{
			"avatar":    avatar,
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
func Delete(userID string) error {
	var (
		ctx = context.Background()
	)

	// Find user exists or not
	id, isValid := mongodb.NewIDFromString(userID)
	if !isValid {
		return errors.New(internal.ErrorInvalidUser)
	}
	user, _ := findByID(ctx, id)
	if user.ID.IsZero() {
		return errors.New(internal.ErrorNotFoundUser)
	}

	// Setup condition
	cond := bson.M{
		"_id": id,
	}

	// Setup update data
	updateData := bson.M{
		"$set": bson.M{
			"deleted":   true,
			"updatedAt": internal.Now(),
		},
	}

	// Update
	if err := updateOneByCondition(ctx, cond, updateData); err != nil {
		return err
	}

	return nil
}

package cache

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/usermngmt/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Roles ...
func Roles() {
	var (
		ctx  = context.Background()
		wg   sync.WaitGroup
		cond = bson.M{}
	)

	// Find
	roles := roleFindByCondition(ctx, cond)
	permissions := permissionFindByCondition(ctx, cond)

	wg.Add(len(roles))
	for _, value := range roles {
		go func(role model.DBRole) {
			defer wg.Done()
			rolePermissions := make([]string, 0)
			// Get role permissions
			for _, permission := range permissions {
				if permission.RoleID == role.ID {
					rolePermissions = append(rolePermissions, permission.Code)
				}
			}

			// Cache Role
			entry := CachedRole{
				Role:        role.Code,
				IsAdmin:     role.IsAdmin,
				Permissions: rolePermissions,
			}
			if err := SetKeyValue(role.ID.Hex(), entry, 0); err != nil {
				logger.Error("usermngmt - CacheRole", logger.LogData{
					"err": err.Error(),
				})
				return
			}
		}(value)
	}

	wg.Wait()
	return
}

// GetCachedRole ...
func GetCachedRole(key string) CachedRole {
	value, err := GetValueByKey(key)
	if err != nil {
		Roles()
		value, _ = GetValueByKey(key)
	}

	// Unmarshal data
	var cachedRole CachedRole
	if err := json.Unmarshal(value, &cachedRole); err != nil {
		logger.Error("usermngmt - GetCachedRole - Unmarshal", logger.LogData{
			"err": err.Error(),
		})
	}
	return cachedRole
}

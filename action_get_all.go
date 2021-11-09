package usermngmt

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

// AllQuery ...
type AllQuery struct {
	Page    int64
	Limit   int64
	Keyword string
	RoleID  string
	Status  string
}

// All ...
func (s Service) All(queryParams AllQuery) (r UserAll) {
	var (
		ctx  = context.Background()
		wg   sync.WaitGroup
		cond = bson.M{}
	)
	query := commonQuery{
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
		docs := s.userFindByCondition(ctx, cond, query.GetFindOptionsUsingPage())
		r.List = getResponseList(ctx, docs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		r.Total = s.userCountByCondition(ctx, cond)
	}()

	wg.Wait()

	return
}

func getResponseList(ctx context.Context, users []dbUser) []User {
	res := make([]User, 0)

	for _, user := range users {
		role, _ := s.roleFindByID(ctx, user.RoleID)
		res = append(res, User{
			ID:     user.ID.Hex(),
			Name:   user.Name,
			Phone:  user.Phone,
			Email:  user.Email,
			Status: user.Status,
			Role: RoleShort{
				ID:      role.ID.Hex(),
				Name:    role.Name,
				IsAdmin: role.IsAdmin,
			},
			Other:     user.Other,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return res
}

package usermngmt

import (
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commonQuery struct {
	Page    int64
	Limit   int64
	Keyword string
	RoleID  string
	Status  string
	Sort    interface{}
}

// AssignKeyword ...
func (q *commonQuery) AssignKeyword(cond bson.M) {
	if q.Keyword != "" {
		cond["searchString"] = mongodb.GenerateQuerySearchString(q.Keyword)
	}
}

// AssignRoleID ...
func (q *commonQuery) AssignRoleID(cond bson.M) {
	if q.RoleID != "" {
		if id, isValid := mongodb.NewIDFromString(q.RoleID); isValid {
			cond["roleId"] = id
		}
	}
}

// AssignStatus ...
func (q *commonQuery) AssignStatus(cond bson.M) {
	if q.Status != "" {
		cond["status"] = q.Status
	}
}

// GetFindOptionsUsingPage ...
func (q *commonQuery) GetFindOptionsUsingPage() *options.FindOptions {
	opts := options.Find()
	if q.Limit > 0 {
		opts.SetLimit(q.Limit).SetSkip(q.Limit * q.Page)
	}
	if q.Sort != nil {
		opts.SetSort(q.Sort)
	}
	return opts
}

// SetDefaultLimit ...
func (q *commonQuery) SetDefaultLimit() {
	if q.Limit <= 0 || q.Limit > 20 {
		q.Limit = 20
	}
}

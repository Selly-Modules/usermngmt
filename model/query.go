package model

import (
	"github.com/Selly-Modules/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const maxLimit = 500

// CommonQuery ...
type CommonQuery struct {
	Page    int64
	Limit   int64
	Keyword string
	RoleID  string
	Status  string
	Deleted string
	Sort    interface{}
	Other   map[string]interface{}
	RoleIDs []primitive.ObjectID
}

// AssignDeleted ...
func (q *CommonQuery) AssignDeleted(cond bson.M) {
	if q.Deleted == "true" {
		cond["deleted"] = true
	}
	if q.Deleted == "false" {
		cond["deleted"] = false
	}
}

// AssignKeyword ...
func (q *CommonQuery) AssignKeyword(cond bson.M) {
	if q.Keyword != "" {
		cond["searchString"] = mongodb.GenerateQuerySearchString(q.Keyword)
	}
}

// AssignRoleID ...
func (q *CommonQuery) AssignRoleID(cond bson.M) {
	if q.RoleID != "" {
		if id, isValid := mongodb.NewIDFromString(q.RoleID); isValid {
			cond["roleId"] = id
		}
	}
}

// AssignRoleIDs ...
func (q *CommonQuery) AssignRoleIDs(cond bson.M) {
	if len(q.RoleIDs) == 1 {
		cond["roleId"] = q.RoleIDs[0]
	} else if len(q.RoleIDs) > 1 {
		cond["roleId"] = bson.M{
			"$in": q.RoleIDs,
		}
	}
}

// AssignStatus ...
func (q *CommonQuery) AssignStatus(cond bson.M) {
	if q.Status != "" {
		cond["status"] = q.Status
	}
}

// GetFindOptionsUsingPage ...
func (q *CommonQuery) GetFindOptionsUsingPage() *options.FindOptions {
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
func (q *CommonQuery) SetDefaultLimit() {
	if q.Limit <= 0 {
		q.Limit = 20
	}
	if q.Limit > maxLimit {
		q.Limit = 500
	}
}

// AssignOther ...
func (q *CommonQuery) AssignOther(cond bson.M) {
	// Query fields in other object
	if len(q.Other) > 0 {
		for key, value := range q.Other {
			switch v := value.(type) {
			case []primitive.ObjectID:
				cond["other."+key] = bson.M{
					"$in": v,
				}
			case []string:
				cond["other."+key] = bson.M{
					"$in": v,
				}
			default:
				cond["other."+key] = value
			}
		}
	}
}

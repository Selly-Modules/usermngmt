package usermngmt

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

//  getUserCollection ...
func (s Service) getUserCollection() *mongo.Collection {
	return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableUser))
}

//  getRoleCollection ...
func (s Service) getRoleCollection() *mongo.Collection {
	return s.DB.Collection(fmt.Sprintf("%s-%s", s.TablePrefix, tableRole))
}

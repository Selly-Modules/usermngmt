package usermngmt

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

//  getCollectionName ...
func (s Service) getCollectionName(tablePrefix, table string) *mongo.Collection {
	return s.db.Collection(fmt.Sprintf("%s-%s", tablePrefix, table))
}

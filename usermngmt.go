package usermngmt

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/mongodb"
	"github.com/Selly-Modules/usermngmt/role"
	"github.com/Selly-Modules/usermngmt/user"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDBConfig ...
type MongoDBConfig struct {
	Host, User, Password, DBName, Mechanism, Source string
}

// Config ...
type Config struct {
	// MongoDB config, for save documents
	MongoDB MongoDBConfig
	// Table prefix, each service has its own prefix
	TablePrefix string
}

// Handler ...
type Handler struct {
	User user.Handle
	Role role.Handle
}

// Service ...
type Service struct {
	config  Config
	db      *mongo.Database
	handler Handler
}

var s *Service

// Init ...
func Init(config Config) (*Service, error) {
	if config.MongoDB.Host == "" {
		return nil, errors.New("please provide all necessary information for init user")
	}

	// If prefixTable is empty then it is usermngmt
	if config.TablePrefix == "" {
		config.TablePrefix = tablePrefixDefault
	}

	// Connect MongoDB
	db, err := mongodb.Connect(
		config.MongoDB.Host,
		config.MongoDB.User,
		config.MongoDB.Password,
		config.MongoDB.DBName,
		config.MongoDB.Mechanism,
		config.MongoDB.Source,
	)
	if err != nil {
		fmt.Println("Cannot init module User MANAGEMENT", err)
		return nil, err
	}

	s = &Service{
		config: config,
		db:     db,
	}

	// Setup handle
	s.handler = Handler{
		User: user.Handle{
			Col:     s.getCollectionName(config.TablePrefix, tableUser),
			RoleCol: s.getCollectionName(config.TablePrefix, tableRole),
		},
		Role: role.Handle{
			Col: s.getCollectionName(config.TablePrefix, tableRole),
		},
	}

	return s, nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}

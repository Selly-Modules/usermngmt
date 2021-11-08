package usermngmt

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/mongodb"
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

// Service ...
type Service struct {
	Config
	DB *mongo.Database
}

var s *Service

// Init ...
func Init(config Config) (*Service, error) {
	if config.MongoDB.Host == "" {
		return nil, errors.New("please provide all necessary information for init user")
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
		Config: config,
		DB:     db,
	}

	return s, nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}

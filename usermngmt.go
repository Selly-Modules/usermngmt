package usermngmt

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"
	"github.com/Selly-Modules/usermngmt/cache"
	"github.com/Selly-Modules/usermngmt/database"
	"github.com/Selly-Modules/usermngmt/internal"
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
	config Config
}

var s *Service

// Init ...
func Init(config Config) (*Service, error) {
	if config.MongoDB.Host == "" {
		return nil, errors.New("please provide all necessary information for init user")
	}

	// If prefixTable is empty then it is usermngmt
	if config.TablePrefix == "" {
		config.TablePrefix = internal.TablePrefixDefault
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

	logger.Init("usermngmt", "")

	// Set database
	database.Set(db, config.TablePrefix)

	// Init cache
	cache.Init()

	s = &Service{
		config: config,
	}

	return s, nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}

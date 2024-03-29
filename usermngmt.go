package usermngmt

import (
	"errors"
	"fmt"

	"github.com/Selly-Modules/logger"
	"github.com/Selly-Modules/mongodb"

	"github.com/Selly-Modules/usermngmt/cache"
	configMoudle "github.com/Selly-Modules/usermngmt/config"
	"github.com/Selly-Modules/usermngmt/database"
	"github.com/Selly-Modules/usermngmt/internal"
)

// RedisConfig ...
type RedisConfig struct {
	URI, Password string
}

// Config ...
type Config struct {
	// MongoDB config, for save documents
	MongoDB mongodb.Config

	// Redis
	Redis RedisConfig

	// Table prefix, each service has its own prefix
	TablePrefix string

	// Email is unique
	EmailIsUnique bool

	// phone number is unique
	PhoneNumberIsUnique bool
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
	db, err := mongodb.Connect(config.MongoDB)
	if err != nil {
		fmt.Println("Cannot init module User MANAGEMENT", err)
		return nil, err
	}

	logger.Init(fmt.Sprintf("%s-usermngmt", config.TablePrefix), "")

	// Set database
	database.Set(db, config.TablePrefix)

	// Set config module
	configMoudle.Set(&configMoudle.Configuration{
		EmailIsUnique:       config.EmailIsUnique,
		PhoneNumberIsUnique: config.PhoneNumberIsUnique,
	})

	// Init cache
	cache.Init(config.Redis.URI, config.Redis.Password)

	s = &Service{
		config: config,
	}

	return s, nil
}

// GetInstance ...
func GetInstance() *Service {
	return s
}

// GetConnectOptions ...
func GetConnectOptions(Host, DBName string) mongodb.Config {
	return mongodb.Config{
		Host:       Host,
		DBName:     DBName,
		Standalone: &mongodb.ConnectStandaloneOpts{},
		TLS:        &mongodb.ConnectTLSOpts{},
	}
}

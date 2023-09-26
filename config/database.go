package config

import (
	"fmt"
	"sync"
	"time"

	"github.com/fwidjaya20/symphonic/facades"
	"github.com/golang-module/carbon"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	config := facades.Config()

	config.Add("database", map[string]any{
		"connections": map[string]any{
			"postgresql": map[string]any{
				"driver":   "postgresql",
				"host":     config.Get("DB_HOST", "localhost"),
				"port":     config.Get("DB_PORT", 5432),
				"database": config.Get("DB_DATABASE", "forge"),
				"username": config.Get("DB_USERNAME", "postgres"),
				"password": config.Get("DB_PASSWORD", "postgres"),
			},
			"redis": map[string]any{
				"driver":   "redis",
				"host":     facades.Config().Get("REDIS_HOST", "localhost"),
				"port":     facades.Config().Get("REDIS_PORT", 6379),
				"database": facades.Config().Get("REDIS_DATABASE", "0"),
				"password": facades.Config().Get("REDIS_PASSWORD", ""),
			},
		},
		"default":  config.Get("DB_CONNECTION", "postgresql"),
		"timezone": carbon.UTC,
	})
}

var (
	instanceDB *gorm.DB
	syncDB     sync.Once
)

func DB() *gorm.DB {
	syncDB.Do(func() {
		var err error

		config := gorm.Config{
			SkipDefaultTransaction:   true,
			DisableNestedTransaction: true,
			NowFunc: func() time.Time {
				ti, _ := time.LoadLocation(facades.Config().GetString("database.timezone"))
				return time.Now().In(ti)
			},
		}

		if instanceDB, err = gorm.Open(postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			facades.Config().GetString("database.connections.postgresql.host"),
			facades.Config().GetString("database.connections.postgresql.username"),
			facades.Config().GetString("database.connections.postgresql.password"),
			facades.Config().GetString("database.connections.postgresql.database"),
			facades.Config().GetString("database.connections.postgresql.port"),
			facades.Config().GetString("database.timezone"),
		)), &config); nil != err {
			facades.Logger().Panic(err.Error())
		}

		if "production" != facades.Config().GetString("app.env") {
			instanceDB = instanceDB.Session(&gorm.Session{
				Logger: instanceDB.Logger.LogMode(logger.Info),
			})
		}
	})

	return instanceDB
}

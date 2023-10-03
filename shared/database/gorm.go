package database

import (
	"fmt"
	"sync"
	"time"

	"github.com/fwidjaya20/symphonic/facades"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	instanceDB *gorm.DB
	syncDB     sync.Once
)

func Gorm() *gorm.DB {
	syncDB.Do(func() {
		var err error

		config := gorm.Config{
			SkipDefaultTransaction:   true,
			DisableNestedTransaction: true,
			NowFunc: func() time.Time {
				ti, _ := time.LoadLocation(facades.Config().GetString("repository.timezone"))
				return time.Now().In(ti)
			},
		}

		if instanceDB, err = gorm.Open(postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			facades.Config().GetString("repository.connections.postgresql.host"),
			facades.Config().GetString("repository.connections.postgresql.username"),
			facades.Config().GetString("repository.connections.postgresql.password"),
			facades.Config().GetString("repository.connections.postgresql.repository"),
			facades.Config().GetString("repository.connections.postgresql.port"),
			facades.Config().GetString("repository.timezone"),
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

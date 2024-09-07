package configs

import (
	"aichat/common/util"
	"aichat/configs/initstruct"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func dbInit(Cfg initstruct.DataBaseConfig) {

	connString := "host=" + Cfg.Host +
		" port=" + Cfg.Port +
		" user=" + Cfg.User +
		" dbname=" + Cfg.Name +
		" sslmode=disable" +
		" password=" + Cfg.Password

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level(这里记得根据需求改一下)
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	// Error
	if connString == "" || err != nil {
		util.Log().Error("pgsql lost: %v", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		util.Log().Error("pgsql lost: %v", err)
		panic(err)
	}

	//设置连接池
	//空闲
	sqlDB.SetMaxIdleConns(Cfg.MaxIdleConns)
	//打开
	sqlDB.SetMaxOpenConns(Cfg.MaxOpenConns)

	DB = db

}

func GetDatabase() *gorm.DB {
	if DB == nil {
		panic("database not initialized")
	}
	return DB
}

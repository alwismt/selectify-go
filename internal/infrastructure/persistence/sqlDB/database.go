package sqldb

import (
	"os"
	"strconv"

	"github.com/alwismt/selectify/internal/infrastructure/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDBwithPool() {
	// get DB_DRIVER from env
	driver := os.Getenv("DB_DRIVER")
	if driver == "" || (driver != "pgsql" && driver != "mysql") {
		panic("DB_DRIVER not set")
	}
	dsn, err := utils.ConnectionURLBuilder(driver)
	if err != nil {
		panic("failed to get url")
	}
	var db *gorm.DB
	switch driver {
	case "pgsql":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			PrepareStmt: true,
		})
		if err != nil {
			panic("failed to connect database")
		}
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			PrepareStmt: true,
		})
		if err != nil {
			panic("failed to connect database")
		}
	}
	maxIdle := os.Getenv("DB_MAX_IDLE_CONNS")
	idle, err := strconv.Atoi(maxIdle)
	if err != nil {
		panic("DB_MAX_IDLE_CONNS not set")
	}
	maxOpen := os.Getenv("DB_MAX_OPEN_CONNS")
	open, err := strconv.Atoi(maxOpen)
	if err != nil {
		panic("DB_MAX_OPEN_CONNS not set")
	}
	// Set connection pool settings
	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(idle)
	sqldb.SetMaxOpenConns(open)

	DB = db
}

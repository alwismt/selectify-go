package sqldb

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestConnectToDBwithPool(t *testing.T) {
	err := godotenv.Load("./../../../../.env")
	if err != nil {
		t.Errorf("error loading .env file %s", err)
		return
	}

	ConnectToDBwithPool()

	// Ensure DB connection was established
	assert.NotNil(t, DB)

	// Ensure DB connection is of type *gorm.DB
	if DB == nil {
		t.Errorf("DB connection is nil")
		return
	}

	// Try to ping the database
	sql, err := DB.DB()
	if err != nil {
		t.Errorf("sql connection error: %v", err)
		return
	}
	err = sql.Ping()
	if err != nil {
		t.Errorf("failed to ping database: %v", err)
		return
	}

	maxOpen := os.Getenv("DB_MAX_OPEN_CONNS")
	open, err := strconv.Atoi(maxOpen)
	if err != nil {
		t.Errorf("DB_MAX_OPEN_CONNS not set")
		return
	}
	// assert.Equal(t, idle, sql.Stats().Idle)
	assert.Equal(t, open, sql.Stats().MaxOpenConnections)
}

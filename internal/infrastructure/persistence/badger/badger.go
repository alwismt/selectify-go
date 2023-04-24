package badger

import (
	"time"

	"github.com/gofiber/storage/badger"
)

var Badger *badger.Storage

func ConnectToBadger() {
	badger := badger.New(badger.Config{
		Database:   "./data/badger/admin/session",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	Badger = badger
}

func AdminCsrf() *badger.Storage {
	storage := badger.New(badger.Config{
		Database:   "./data/badger/admin/Csrf",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	return storage
}

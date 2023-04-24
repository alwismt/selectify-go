package session

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type Session interface {
	NewSession(c *fiber.Ctx)
	Set(c *fiber.Ctx, key string, val interface{})
	Get(c *fiber.Ctx, key string) interface{}
	Delete(c *fiber.Ctx, key string)
	Destroy(c *fiber.Ctx)
	ID(c *fiber.Ctx) string
	Keys(c *fiber.Ctx) []string
}

type session_s struct {
	store *session.Store
}

func NewSession(storage *session.Store) Session {
	if storage == nil {
		storage = store
	}
	return &session_s{store: storage}
}

func (s *session_s) NewSession(c *fiber.Ctx) {
	sess, err := s.store.Get(c)
	if err != nil {
		log.Println(err)
	}
	err = sess.Regenerate()
	if err != nil {
		log.Println(err)
	}
	// save
	if err := sess.Save(); err != nil {
		log.Println(err)
	}
}

func (s *session_s) Set(c *fiber.Ctx, key string, val interface{}) {
	sess, err := s.store.Get(c)
	if err != nil {
		log.Println(err)
	}
	// Set key value
	sess.Set(key, val)
	// save
	if err := sess.Save(); err != nil {
		log.Println(err)
	}
}

func (s *session_s) Get(c *fiber.Ctx, key string) interface{} {

	sess, err := s.store.Get(c)
	if err != nil {
		log.Println(err)
	}

	val := sess.Get(key)

	return val
}

func (s *session_s) Delete(c *fiber.Ctx, key string) {

	sess, err := s.store.Get(c)
	if err != nil {
		log.Println(err)
	}
	// Delete key
	sess.Delete(key)

	// save
	if err := sess.Save(); err != nil {
		log.Println(err)
	}
}

func (s *session_s) Destroy(c *fiber.Ctx) {

	sess, err := s.store.Get(c)
	if err != nil {
		log.Println(err)
	}
	// Destroy session
	if err := sess.Destroy(); err != nil {
		panic(err)
	}
}

func (s *session_s) ID(c *fiber.Ctx) string {
	sess, err := s.store.Get(c)
	if err != nil {
		log.Println(err)
	}
	id := sess.ID()

	return id
}

func (s *session_s) Keys(c *fiber.Ctx) []string {
	sess, err := s.store.Get(c)
	if err != nil {
		log.Println(err)
	}
	keys := sess.Keys()

	return keys
}

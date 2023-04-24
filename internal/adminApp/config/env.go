package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func CheckEnv() error {
	err := godotenv.Load("./.env")
	if err != nil {
		return fmt.Errorf("error loading .env file %s", err)
	}

	ck_enc := os.Getenv("COOKIE_ENCRYPTION")
	if ck_enc == "true" {
		ck_enc_key := os.Getenv("COOKIE_ENCRYPTION_KEY")
		if ck_enc_key == "" {
			return fmt.Errorf("cookie encryption key not found")
		}
		if len(ck_enc_key) != 32 {
			return fmt.Errorf("cookie encryption key must be 32 bytes")
		}
	}
	return nil
}

package utils

import (
	"context"
	"fmt"
	"os"
	"time"
)

// CreateContext func for create a new context with timeout of X seconds
func CreateContextWithTimeout() (context.Context, context.CancelFunc) {
	timeOut := os.Getenv("SERVER_TIMEOUT")
	// if timeOut is not set, set it to 5 seconds
	if timeOut == "" {
		timeOut = "5"
	}
	seconds, err := time.ParseDuration(timeOut + "s")
	if err != nil {
		fmt.Println("Failed to parse timeout duration", err)
		seconds = 5 * time.Second
	}
	return context.WithTimeout(context.Background(), seconds)
}

// ctx.IP() returns the IP address of the client making the request.
// ctx.Get("User-Agent") returns the value of the "User-Agent" header, which contains information about the browser and operating system being used.
// ctx.Cookies("cookie_name") returns the value of a specific cookie.

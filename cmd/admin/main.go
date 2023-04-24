package main

import (
	"log"

	"github.com/alwismt/selectify/internal/adminApp/config"
	"github.com/alwismt/selectify/internal/infrastructure/web"
)

func main() {

	// Load env and check
	err := config.CheckEnv()
	if err != nil {
		log.Fatalf("%s", err)
	}

	web.AdminServerStart()

	// app.Listen(":8088")
}

// func install() (bool, error) {
// 	fmt.Println("Installing...")
// 	fmt.Println("Enter your database host: ")
// 	var dbHost string
// 	fmt.Scanln(&dbHost)
// 	fmt.Println("Enter your database port: ")
// 	var dbPort string
// 	fmt.Scanln(&dbPort)
// 	fmt.Println("Enter your database name: ")
// 	var dbName string
// 	fmt.Scanln(&dbName)
// 	fmt.Println("Enter your database username: ")
// 	var dbUsername string
// 	fmt.Scanln(&dbUsername)
// 	fmt.Println("Enter your database password: ")
// 	var dbPassword string
// 	fmt.Scanln(&dbPassword)
// 	fmt.Println("Enter your database driver: ")
// 	var dbDriver string
// 	fmt.Scanln(&dbDriver)

// 	// Create .env file
// 	file, err := os.Create(".env")
// 	if err != nil {
// 		return false, err
// 	}
// 	defer file.Close()
// 	// Write to .env file
// 	_, err = file.WriteString(
// 		"DB_HOST=" + dbHost + "\n" + "DB_PORT=" + dbPort + "\n" + "DB_NAME=" + dbName + "\n" + "DB_USERNAME=" + dbUsername + "\n" + "DB_PASSWORD=" + dbPassword + "\n" + "DB_DRIVER=" + dbDriver + "\n")
// 	if err != nil {
// 		return false, err
// 	}
// 	fmt.Println("Installation complete")
// 	return true, nil
// }

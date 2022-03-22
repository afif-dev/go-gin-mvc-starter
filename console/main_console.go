package console

import (
	"os"
	"strconv"
)

func InitConsole() {

	// # Uncomment this code to load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	var cron_set, _ = strconv.ParseBool(os.Getenv("CRON_ENABLE"))

	//run console 
	if cron_set {
		CronTest()
	}
}
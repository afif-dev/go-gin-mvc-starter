package console

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/robfig/cron"
)

func CronTest() {
	//	set log daily
	//	f, _ := os.OpenFile("logs/console.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	logfile, _ := os.Create("logs/console_"+time.Now().Format("2006-01-02")+".log")
	log.SetOutput(io.MultiWriter(logfile))

	c := cron.New()
	c.AddFunc("@every 3s", func() {
		log.Printf("Console Cron: Run every 3 seconds")
	})
	c.Start()
}
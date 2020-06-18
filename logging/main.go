package main

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	filename := filepath.Base(os.Args[0])
	syslog, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, filename)
	if err != nil {
		log.Fatalln(err, "error")
	} else {
		log.SetOutput(syslog)
	}

	log.Println("LOG_MAIL: Google is a good")
	log.Println("Google is a good")

}

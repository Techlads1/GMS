package main

import (
	"fmt"
	"gateway/package/log"
	"gateway/webserver"
	"os"
	"os/signal"
	"syscall"

	"gateway/package/config"
)

func init() {
	path := config.LoggerPath()
	fmt.Println(path)
	log.SetOptions(
		log.Development(),
		log.WithCaller(true),
		log.WithLogDirs(path),
	)
}
func main() {
	go webserver.StartWebserver()

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	<-stop
	log.Infoln("GRM is shutting down...  👋 !")
	fmt.Println("GRM is shutting down .... 👋 !")
	//database.Close()

	go func() {
		<-stop
		log.Fatalln("GRM is terminating...")
	}()

	defer os.Exit(0)
}

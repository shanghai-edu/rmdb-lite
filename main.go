package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shanghai-edu/rmdb-lite/controller"
	"github.com/shanghai-edu/rmdb-lite/g"
	"github.com/shanghai-edu/rmdb-lite/models"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	init := flag.String("i", "", "init db with csv file, -i routers.csv")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	if *init != "" {
		err := models.InitData(*init)
		if err != nil {
			log.Fatalf("Init DB Failed: %s", err)
		}
		os.Exit(0)
	}
	g.InitLog(g.Config().LogLevel)
	g.InitDB()

	srv := controller.InitGin(g.Config().Http.Listen)

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: %s", err)
	}
	log.Println("Server exit")

}

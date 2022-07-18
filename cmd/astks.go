package main

import (
	"flag"
	"github.com/yddeng/astk/astks"
	"github.com/yddeng/astk/pkg/util"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	file = flag.String("file", "./astks.json", "config file")
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	var err error
	var cfg astks.Config
	if err = util.DecodeJsonFromFile(&cfg, *file); err != nil {
		panic(err)
	}

	if err = astks.Start(cfg); err != nil {
		panic(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("receive signal:%v to stopping. ", <-sigChan)
	astks.Stop()
	log.Println("stopped. ")
}

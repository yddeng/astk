package main

import (
	"flag"
	"github.com/yddeng/astk/astke"
	"github.com/yddeng/astk/pkg/util"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	file = flag.String("file", "./astke.json", "config file")
)

func main() {
	flag.Parse()

	rand.Seed(time.Now().UnixNano())

	var err error
	var cfg astke.Config
	if err = util.DecodeJsonFromFile(&cfg, *file); err != nil {
		panic(err)
	}

	if err = astke.Start(cfg); err != nil {
		panic(err)
	}

	go func() {
		http.ListenAndServe("10.128.2.123:40160", nil)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("receive signal:%v to stopping. ", <-sigChan)
	astke.Stop()
	log.Println("stopped. ")
}

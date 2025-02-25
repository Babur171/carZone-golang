package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Babur171/carZone-golang/config"
)

func main() {

	config.LoadConfig()
	fmt.Println("server is runinggg", config.AppConfig.BaseURL)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("its workingggggggg"))
	})
	addr := ":" + config.AppConfig.Port
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := http.ListenAndServe(addr, nil)

		if err != nil {
			log.Fatal("server feailed", err)
		}
	}()

	<-done
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// htt

	slog.Info("Shutdon the system")

}

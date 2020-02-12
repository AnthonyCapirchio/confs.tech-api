package main

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/gol4ng/signal"

	"squidward.confs.tech/conference"
	"squidward.confs.tech/importer"
	"squidward.confs.tech/server"
)

func main() {
	idleConnsClosed := make(chan struct{})

	conferences := InitConferencesStore()
	httpServer := server.NewServer(":9997", conferences)

	defer signal.SubscribeWithKiller(func(signal os.Signal) {
		log.Println(signal.String(), "signal received : gracefully stopping application")
		close(idleConnsClosed)
	}, os.Interrupt, syscall.SIGTERM, syscall.SIGSTOP)()

	go func() {
		if err := httpServer.Start(); err != nil {
			log.Fatal(fmt.Sprintf("Error from API server %v", err))
		}
	}()

	<-idleConnsClosed

	// startDate, _ := time.Parse("2006-01-02", "2020-01-01")
	// endDate, _ := time.Parse("2006-01-02", "2020-12-31")

	// usaJava := store.
	// 	FilterByCountry("France").
	// 	FilterByCategory("general").
	// 	FilterByDateRange(startDate, endDate, 5)

	// b, _ := json.Marshal(usaJava.Entries)

	// fmt.Println(string(b))
}

func InitConferencesStore() *conference.Store {
	imp := importer.LocalFileImporter{
		BasePath: "./data",
		Years:    []string{"2020", "2021"},
	}

	store, _ := imp.Import()

	log.Println("Store loaded")

	return store
}

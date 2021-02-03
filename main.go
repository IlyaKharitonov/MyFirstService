package main

import (
	"context"
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"time"
	"os"
	"os/signal"
	"syscall"
	
)

func main() {
	addr := flag.String("addr", ":8080", "Server address")
	addrDB := flag.String("addrDB", "root:1643@(0.0.0.0:3306)/usersdb", "Database address")

	flag.Parse()

	db, err := sql.Open("mysql", *addrDB)
	if err != nil {
		log.Fatal("Error connecting to the database when starting the service")
	}
	var storage Storager = Storage{
		database: db,
	}

	srv := http.Server{Addr: *addr}

	server := Server{
		httpsrv: srv,
		strg:    storage,
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	go func() {
		<-sigChan
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.httpsrv.Shutdown(ctx); err != nil {
			log.Fatal("server stopped by signal")
		}
	}()

	server.Run()

}

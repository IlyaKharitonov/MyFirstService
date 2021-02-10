package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//ждем пока развернется база
	<-time.After(12 * time.Second)

	addr := flag.String("addr", ":8080", "Server address")
	addrDB := flag.String("addrDB", "root:1643@(mysqldb)/", "Database address")
	flag.Parse()

	db, err := sql.Open("mysql", *addrDB)
	if err != nil {
		log.Fatal("Error connecting to the database when starting the service")
	}
	//создаем базу и таблицу
	_, err = db.Exec("create database usersdb")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("use usersdb")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("create table users(id int auto_increment primary key, name varchar(30) not null, age int not null) ")
	if err != nil {
		log.Fatal(err)
	}
	//создаем объект хранилище
	var storage Storager = Storage{
		database: db,
	}
	//создаем объект сервер и передаем в него хранилище
	srv := http.Server{Addr: *addr}
	server := Server{
		httpsrv: srv,
		strg:    storage,
	}
	//обработка сигнала
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	go func() {
		<-sigChan
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err := server.httpsrv.Shutdown(ctx); err != nil {
			log.Fatal("server stopped by signal")
		}
	}()
	//пуск
	server.Run()
}

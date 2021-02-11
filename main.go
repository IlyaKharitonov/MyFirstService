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

func DB(data string) (*sql.DB, error) {
	db, err := sql.Open("mysql", data)
	if err != nil {
		log.Fatal("Error connecting to the database when starting the service")
	}
	//создаем базу и таблицу
	_, err = db.Exec("create database if not exists usersdb")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("use usersdb")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("create table if not exists users(id int auto_increment primary key, name varchar(30) not null, age int not null) ")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func main() {
	port := flag.String("port", ":8080", "Server address")
	login := flag.String("user", "root", "Login")
	pass := flag.String("pass", "1643", "Password")
	addr := flag.String("addr", "0.0.0.0:3306", "Address")
	flag.Parse()

	data := *login + ":" + *pass + "@(" + *addr + ")/"

	db, err := DB(data)
	if err != nil {
		log.Fatal(err)
	}
	//создаем объект хранилище
	var storage Storager = Storage{
		database: db,
	}
	//создаем объект сервер и передаем в него хранилище
	srv := http.Server{Addr: *port}
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

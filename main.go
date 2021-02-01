package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"net/http"
	"flag"
)

func main(){
	addr:=flag.String("addr",":8080","Адрес сервера")
	addrDB:=flag.String("addrDB", "root:1643@(0.0.0.0:3306)/usersdb", "Адрес базы" )

	flag.Parse()


	db,err := sql.Open("mysql", *addrDB)
		if err != nil{
			log.Fatal("Ошибка подключения к базе")
	}
	var storage Storager = Storage{
		database: db,
	}

	srv := http.Server{Addr:*addr}

 	server := Server{
		 httpsrv: srv,
		 strg: storage,
	}

 	server.Run()

}


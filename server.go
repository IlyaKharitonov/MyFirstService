package main

import (
	"fmt"
	"net/http"
	"strconv"
	"encoding/json"
	// "io"
	"os"
	"os/signal"
	"syscall"
	"log"
	"context"
	"time"
)

type Server struct {
	httpsrv http.Server
	strg Storager
}

//обрабатывает запрс в зависимости от его содержимого
func (s Server) handlerAdd(w http.ResponseWriter, r *http.Request) {
	
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))

	if name == "" || age == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Данные не переданы. Ожидаемый формат Add?name=Borat&Age=68")
	} else {
		err := s.strg.Add(name, age)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Fprint(w, "Данные добавлены")
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))

	if id == 0 || name == "" || age == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Данные не переданы. Ожидаемый формат Update?id=69&name=Borat&Age=68")
	} else {
		err := s.strg.Update(id, name, age)
		if err != nil {
			fmt.Fprint(w, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, "Данные изменены")
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerCount(w http.ResponseWriter, r *http.Request) {
	count, err := s.strg.Count()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	respJSON,err := json.Marshal(count)
	if err != nil{
		fmt.Fprint(w, "Ошибка json.Marshal handlerCount")
	}
	fmt.Fprint(w,string(respJSON))
	w.WriteHeader(http.StatusOK)
}

func (s Server) handlerGet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Данные не получены. Ожидаемый формат Get?id=69")
	} else {
		resReq, err := s.strg.Get(id)
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
		}
		respJSON,err := json.Marshal(resReq)
		if err != nil{
			fmt.Fprint(w, "Ошибка json.Marshal handlerGet")
		}
		fmt.Fprint(w,string(respJSON))
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerGetByName(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Данные не получены. Ожидаемый формат GetByName?name=Borat")
	} else {
		resReq, err := s.strg.GetByName(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		respJSON,err := json.Marshal(resReq)
		if err != nil{
			fmt.Fprint(w, "Ошибка json.Marshal handlerGetByName")
		}
		fmt.Fprint(w,string(respJSON))
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerGetByAge(w http.ResponseWriter, r *http.Request) {

	age, _ := strconv.Atoi(r.FormValue("age"))
	if age == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Данные не получены. Ожидаемый формат GetByAge?age=68")
	} else {
		resReq, err := s.strg.GetByAge(age)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		respJSON,err := json.Marshal(resReq)
		if err != nil{
			fmt.Fprint(w, "Ошибка json.Marshal handlerGetByAge")
		}
		fmt.Fprint(w,string(respJSON))
		w.WriteHeader(http.StatusOK)

	}
}


// func (s Server) NotFound(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w,"Страница не найдена")
// 	w.WriteHeader(http.StatusNotFound)
// }

func (s Server) Run() {
	
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,syscall.SIGTERM)
	go func() {
	<-sigChan
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := s.httpsrv.Shutdown(ctx); err != nil {
		log.Fatal("Сервер остановлен")
	}
	}()
	
	http.HandleFunc("/Add", s.handlerAdd)
	http.HandleFunc("/Update", s.handlerUpdate)
	http.HandleFunc("/Count", s.handlerCount)
	http.HandleFunc("/Get", s.handlerGet)
	http.HandleFunc("/GetByName", s.handlerGetByName)
	http.HandleFunc("/GetByAge", s.handlerGetByAge)

	fmt.Println("Запустил")
	/* http.ListenAndServe(":8080", nil) */
	s.httpsrv.ListenAndServe()
}

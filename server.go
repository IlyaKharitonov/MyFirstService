package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	httpsrv http.Server
	strg    Storager
}

//обрабатывает запрс в зависимости от его содержимого
func (s Server) handlerAdd(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))

	if name == "" || age == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No data transferred. Expected format Add?name=Borat&Age=68")
	} else {
		err := s.strg.Add(name, age)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		fmt.Fprint(w, "Data added ")
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	name := r.FormValue("name")
	age, _ := strconv.Atoi(r.FormValue("age"))

	if id == 0 || name == "" || age == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No data transferred. Expected format Update?id=69&name=Borat&Age=68")
	} else {
		err := s.strg.Update(id, name, age)
		if err != nil {
			fmt.Fprint(w, err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprint(w, "Data changed")
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerCount(w http.ResponseWriter, r *http.Request) {
	count, err := s.strg.Count()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	respJSON, err := json.Marshal(count)
	if err != nil {
		fmt.Fprint(w, "Error json.Marshal handlerCount")
	}
	fmt.Fprint(w, string(respJSON))
	w.WriteHeader(http.StatusOK)
}

func (s Server) handlerGet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))

	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No data received. Expected format Get?id=69")
	} else {
		resReq, err := s.strg.Get(id)
		if err != nil {

			w.WriteHeader(http.StatusInternalServerError)
		}
		respJSON, err := json.Marshal(resReq)
		if err != nil {
			fmt.Fprint(w, "Error json.Marshal handlerGet")
		}
		fmt.Fprint(w, string(respJSON))
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerGetByName(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No data received. Expected format GetByName?name=Borat")
	} else {
		resReq, err := s.strg.GetByName(name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		respJSON, err := json.Marshal(resReq)
		if err != nil {
			fmt.Fprint(w, "Error json.Marshal handlerGetByName")
		}
		fmt.Fprint(w, string(respJSON))
		w.WriteHeader(http.StatusOK)
	}
}

func (s Server) handlerGetByAge(w http.ResponseWriter, r *http.Request) {

	age, _ := strconv.Atoi(r.FormValue("age"))
	if age == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No data received. Expected format GetByAge?age=68")
	} else {
		resReq, err := s.strg.GetByAge(age)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		respJSON, err := json.Marshal(resReq)
		if err != nil {
			fmt.Fprint(w, "Error json.Marshal handlerGetByAge")
		}
		fmt.Fprint(w, string(respJSON))
		w.WriteHeader(http.StatusOK)

	}
}


func (s Server) Run() {

	http.HandleFunc("/Add", s.handlerAdd)
	http.HandleFunc("/Update", s.handlerUpdate)
	http.HandleFunc("/Count", s.handlerCount)
	http.HandleFunc("/Get", s.handlerGet)
	http.HandleFunc("/GetByName", s.handlerGetByName)
	http.HandleFunc("/GetByAge", s.handlerGetByAge)

	fmt.Println("Запустил на ",s.httpsrv.Addr )
	s.httpsrv.ListenAndServe()
}

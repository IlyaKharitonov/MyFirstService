package main

import (
	// "database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	testreq    string
	response   string
	statuscode int
}

type mockStorage struct{

}

func(m mockStorage)Add(name string, age int)(err error){
	return nil
}

func(m mockStorage)Update(id int, newName string, newAge int)(err error){
	return nil
}

func(m mockStorage)Count()(count int,err error){
	return 10,nil
}

func(m mockStorage)Get(id int)(user Data, err error){
	user = Data{
		ID: id,
		Name: "Khabib",
		Age: 26,
	}

	return user, nil
}

func(m mockStorage)GetByName(name string)(users []Data, err error){
	users = []Data{
		Data{
		ID: 1,
		Name: name,
		Age: 26,
		},
		Data{
		ID: 1,
		Name: name,
		Age: 26,
		},
	}
	return users, nil
}

func(m mockStorage)GetByAge(age int)(users []Data, err error){

	users = []Data{
		Data{
		ID: 1,
		Name: "Khabib",
		Age: age,
		},
		Data{
		ID: 1,
		Name: "Khabib",
		Age: age,
		},
	}
	return users, nil
}



func database() Server { //
	// dbtest, err := sql.Open("mysql", "root:1643@(0.0.0.0:3306)/usersdb")
	// if err != nil {
	// 	panic(err)
	// }
	var storage Storager = mockStorage{

	}
	// storage := Storage{
	// 	database: dbtest,
	// }
	server := Server{
		strg: storage,
	}
	return server
}



func TestHandlerAdd(t *testing.T) {

	server := database()
	cases := []testCase{
		testCase{
			testreq:    "?name=&age=105",
			response:   "Данные не переданы. Ожидаемый формат Add?name=Borat&Age=68",
			statuscode: http.StatusBadRequest,
		},
		testCase{
			testreq:    "?name=Borat&age=105",
			response:   "Данные добавлены",
			statuscode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://test.com/Add" + item.testreq
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		server.handlerAdd(w, req)
		if w.Code != item.statuscode {
			t.Errorf("Фактический статускод не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		if string(body) != item.response {
			t.Errorf("Фактический ответ не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
	}
}

func TestHandlerUpdate(t *testing.T) {

	server := database()
	cases := []testCase{
		testCase{
			testreq:    "?name=&age=105",
			response:   "Данные не переданы. Ожидаемый формат Update?id=69&name=Borat&Age=68",
			statuscode: http.StatusBadRequest,
		},
		testCase{
			testreq:    "?id=39&name=Borat&age=105",
			response:   "Данные изменены",
			statuscode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://test.com/Update" + item.testreq
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		server.handlerUpdate(w, req)
		if w.Code != item.statuscode {
			t.Errorf("Фактический статускод не совпал с ожидаемым. Кейс номер %d ", caseNum)
		}
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		if string(body) != item.response {
			fmt.Println(string(body))
			t.Errorf("Фактический ответ не совпал с ожидаемым. Кейс номер %d ", caseNum)
		}
	}
}

func TestHandlerCount(t *testing.T) {

	server := database()
	cases := []testCase{
		testCase{
			testreq:    "",
			response:   `10`,
			statuscode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://test.com/Count"
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		server.handlerCount(w, req)
		if w.Code != item.statuscode {
			fmt.Println(w.Code)
			t.Errorf("Фактический статускод не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		if string(body) != item.response {
			fmt.Println(string(body))
			t.Errorf("Фактический ответ не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
	}
}

func TestHandlerGet(t *testing.T) {

	server := database()
	cases := []testCase{
		testCase{
			testreq:    "?name=&age=105",
			response:   "Данные не получены. Ожидаемый формат Get?id=69",
			statuscode: http.StatusBadRequest,
		},
		testCase{
			testreq:    "?id=34",
			response:  `{"ID":34,"Name":"Khabib","Age":26}`,
			statuscode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://test.com/Get" + item.testreq
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		server.handlerGet(w, req)
		if w.Code != item.statuscode {
			t.Errorf("Фактический статускод не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		if string(body) != item.response {
			fmt.Println(string(body))
			t.Errorf("Фактический ответ не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
	}
}

func TestHandlerGetByName(t *testing.T) {

	server := database()
	cases := []testCase{
		testCase{
			testreq:    "?name=&age=105",
			response:   "Данные не получены. Ожидаемый формат GetByName?name=Borat",
			statuscode: http.StatusBadRequest,
		},
		testCase{
			testreq:    "?name=Khabib",
			response:   `[{"ID":1,"Name":"Khabib","Age":26},{"ID":1,"Name":"Khabib","Age":26}]`,
			statuscode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://test.com/GetByName" + item.testreq
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		server.handlerGetByName(w, req)
		if w.Code != item.statuscode {
			t.Errorf("Фактический статускод не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		if string(body) != item.response {
			fmt.Println(string(body))
			t.Errorf("Фактический ответ не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
	}
}

func TestHandlerGetByAge(t *testing.T) {

	server := database()
	cases := []testCase{
		testCase{
			testreq:    "?age=",
			response:   "Данные не получены. Ожидаемый формат GetByAge?age=68",
			statuscode: http.StatusBadRequest,
		},
		testCase{
			testreq:    "?age=105",
			response:   `[{"ID":1,"Name":"Khabib","Age":105},{"ID":1,"Name":"Khabib","Age":105}]`,
			statuscode: http.StatusOK,
		},
	}
	for caseNum, item := range cases {
		url := "http://test.com/GetByAge" + item.testreq
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		server.handlerGetByAge(w, req)
		if w.Code != item.statuscode {
			t.Errorf("Фактический статускод не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
		response := w.Result()
		body, _ := ioutil.ReadAll(response.Body)
		if string(body) != item.response {
			fmt.Println(string(body))
			t.Errorf("Фактический ответ не совпал с ожидаемым. Кейс номер %d", caseNum)
		}
	}
}

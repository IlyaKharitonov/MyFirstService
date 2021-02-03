 package main 

 import (
	 "testing"
	 "math/rand"
	//  "fmt"
	 "net/http"
	 "log"
	 "strconv"
	 "time"
	 "sync"
	 "database/sql"
	 _ "github.com/go-sql-driver/mysql"
 )

 	var client = &http.Client{Timeout: 2*time.Second}

 type Request struct{
	ID int
	Name string
	Age int
	Method string
	ServAdress string
}

func testRequest(sr Request)(error){

	url := sr.ServAdress+"/"+sr.Method+"?id="+strconv.Itoa(sr.ID)+"&name="+sr.Name+"&age="+strconv.Itoa(sr.Age)
	searchReq,err:= http.NewRequest("GET",url,nil)
	if err != nil{
		log.Fatal(err)
	}
	//отправляем запрос на сервер
	_, err = client.Do(searchReq)
	if err != nil{
		log.Fatal(err)
	}
	// defer client.CloseIdleConnections()
	return  err

}

func TestProductivity(t *testing.T){
	
	db, err := sql.Open("mysql", "root:1643@(0.0.0.0:3306)/usersdb")
	if err != nil {
		log.Fatal("Error connecting to the database when starting the service")
	}
	_, err = db.Exec("delete from usersdb.users")
    if err != nil{
        log.Fatal(err)
    }	

	var mu sync.Mutex

	mu.Lock()
	requests := []Request{
	   
			Request{
				ID: 150,
				Name: "Putin",
				Age: 55,
				Method: "Add",
				ServAdress: "http://localhost:8080",
		 },
			Request{
				ID: 150,
				Name: "Stepan",
				Age: 23,
				Method: "Add",
				ServAdress: "http://localhost:8080",
		},
			Request{
				ID: 150,
				Name: "Konstantin",
				Age: 29,
				Method: "Add",
				ServAdress: "http://localhost:8080",
		},
			Request{
				ID: 150,
				Name: "NAVALNY",
				Age: 22,
				Method: "Add",
				ServAdress: "http://localhost:8080",
		},
	}
	lenSlice := len(requests)
	mu.Unlock()
	
	var wg sync.WaitGroup 
	
	counter := 40000
	numGoroutine := 2
	expectedRow := counter * numGoroutine

	for i:=0; i<counter; i++{
		
		wg.Add(numGoroutine)
		//выбираем случайный запрос из слайса
		go func(){
		rand := requests[rand.Intn(lenSlice)]
		err := testRequest(rand)
			if err != nil{
				t.Error("The request failed")
			}
			wg.Done()
		}()	
			go func(){
			rand2 := requests[rand.Intn(lenSlice)]
			err := testRequest(rand2)
				if err != nil{
					t.Error("The request failed")
				}	
				wg.Done()
			}()
				
		wg.Wait()
		
	}
	
	
	var count int
	err = db.QueryRow("select count(*) from usersdb.users").Scan(&count)
	if count != expectedRow{
		t.Error("Test failed")
	}
}
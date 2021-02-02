 package main 

 import (
	 "testing"
	 "math/rand"
	 "fmt"
	 "net/http"
	 "log"
	 "strconv"
	 "time"
	 "sync"
 )

 var client = &http.Client{Timeout: time.Second}

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
	_ err := client.Do(searchReq)
	if err != nil{
		log.Fatal(err)
	}
	defer client.CloseIdleConnections()
	return  err

}

 func TestProductivity(t *testing.T){
	//слайс запросов, отправляемых на тестируемый сервер

	var mu sync.Mutex

	mu.Lock()
	requests := []Request{
	   
			Request{
				ID: 150,
				Name: "Putin",
				Age: 450,
				Method: "GetByName",
				ServAdress: "http://localhost:8080",
		},
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
	 mu.Unlock()
	
	lenSlice := len(requests)

	var wg sync.WaitGroup 
	
	for i:=0; i<50000; i++{
		
		wg.Add(9)
		//выбираем случайный запрос из слайса
		go func(){
		rand := requests[rand.Intn(lenSlice)]
		err := testRequest(rand)
			if err != nil{
				fmt.Print("The request failed")
			}
			// fmt.Println(1)
			wg.Done()
		}()	
			go func(){
			rand2 := requests[rand.Intn(lenSlice)]
			err := testRequest(rand2)
				if err != nil{
					fmt.Print("The request failed")
				}	
				// fmt.Println(2)
				wg.Done()
			}()
				go func(){
				rand3 := requests[rand.Intn(lenSlice)] 
				err := testRequest(rand3)
					if err != nil{
						fmt.Println("The request failed")
					}
					// fmt.Print(3)
					wg.Done()
				}()	
					go func(){
					rand4 := requests[rand.Intn(lenSlice)]
					err := testRequest(rand4)
						if err != nil{
							fmt.Println("The request failed")
						}
						// fmt.Print(4)
						wg.Done()
					}()	
						go func(){
						rand5 := requests[rand.Intn(lenSlice)]
						// mu.Unlock() 
						err := testRequest(rand5)
							if err != nil{
								fmt.Println("The request failed")
							}
							// fmt.Print(5)
							wg.Done()
						}()	
							
						go func(){
							rand6 := requests[rand.Intn(lenSlice)]
							// mu.Unlock() 
							err := testRequest(rand6)
								if err != nil{
									fmt.Print("The request failed")
								}
								// fmt.Println(6)
								 wg.Done()
							}()	

							go func(){
								rand7 := requests[rand.Intn(lenSlice)]
								// mu.Unlock() 
								err := testRequest(rand7)
									if err != nil{
										fmt.Print("The request failed")
									}
									// fmt.Println(6)
									 wg.Done()
								}()	
								go func(){
									rand8 := requests[rand.Intn(lenSlice)]
									// mu.Unlock() 
									err := testRequest(rand8)
										if err != nil{
											fmt.Print("The request failed")
										}
										// fmt.Println(6)
										 wg.Done()
									}()	
									go func(){
										rand9 := requests[rand.Intn(lenSlice)]
										// mu.Unlock() 
										err := testRequest(rand9)
											if err != nil{
												fmt.Print("The request failed")
											}
											// fmt.Println(6)
											 wg.Done()
										}()	
		wg.Wait()
		// fmt.Println("request number",i)
	}
}
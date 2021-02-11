package main 

 import (
	 "fmt"
	 "net/http"
	 "log"
// 	 "strconv"
	 "time"
	 "io/ioutil"
	 "encoding/json"
	 "flag"
	
 )

var client = &http.Client{Timeout: time.Second}


type User struct{
	Id int
	Name string
	Age int
 }

 type searchRequest struct{
	ID string
	Name string
	Age string
	Method string
	ServAdress string
	Message string
}

func Request(sr searchRequest)(User, error){

	url := sr.ServAdress+"/"+sr.Method+"?id="+sr.ID+"&name="+sr.Name+"&age="+sr.Age

	searchReq,err:= http.NewRequest("GET",url,nil)
	if err != nil{
		log.Fatal(err)
	}
	response, err := client.Do(searchReq)
	if err != nil{
		log.Fatal(err)
	}
	if sr.Method == "Get"||sr.Method == "GetByName" || sr.Method == "GetByAge"{
		body, err := ioutil.ReadAll(response.Body)
		if err != nil{
			log.Fatal(err)
		}
		var res User
		err = json.Unmarshal(body, &res)
		if err != nil{
			log.Fatal(err)
		}
		return res, err
	}
	emptystruct := User{
        Name:sr.Method,
    }
	return emptystruct,err

}
func main(){
	id := flag.String("id","","")
	name := flag.String("name","Putin","")
	age := flag.String("age","","")
	method := flag.String("method","Add","")
	addr := flag.String("addr","http://localhost:8080","")

	flag.Parse()
		req := searchRequest{
		ID: *id,
		Name: *name,
		Age: *age,
		Method: *method,
		ServAdress: *addr,
	}
	resResp, _ := Request(req)
	fmt.Println(resResp)
} 


package main 

 import (
	 "fmt"
	 "net/http"
	 "log"
	 "strconv"
	 "time"
	 "io/ioutil"
	 "encoding/json"
 )

var client = &http.Client{Timeout: time.Second}

type User struct{
	Id int
	Name string
	Age int
 }

 type searchRequest struct{
	ID int
	Name string
	Age int
	Method string
	ServAdress string
}

func Request(sr searchRequest)(User,error){

	url := sr.ServAdress+"/"+sr.Method+"?id="+strconv.Itoa(sr.ID)+"&name="+sr.Name+"&age="+strconv.Itoa(sr.Age)

	searchReq,err:= http.NewRequest("GET",url,nil)
	if err != nil{
		log.Fatal(err)
	}
	response, err := client.Do(searchReq)
	if err != nil{
		log.Fatal(err)
	}
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
func main(){
	id := flag.Int("id","","")
	name := flag.String("name","","")
	age := flag.Int("age","","")
	method := flag.String("method","","")
	addr := flag.String("addr", "")

		req := searchRequest{
		ID: 150,
		Name: "Denis",
		Age: 450,
		Method: "Get",
		ServAdress: "http://localhost:8080",
	}
	resResp, _ := Request(req)
	fmt.Println(resResp)
} 

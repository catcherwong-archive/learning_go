package main

import(
	"fmt"
	"net/http"
	"github.com/gorilla/mux"	
	"encoding/json"
	"strconv"
)

func main(){	
	r := NewRouter()
	http.Handle("/", r)
	fmt.Println("Now listening on http://localhost:8000...")
	http.ListenAndServe(":8000", r)	
}

type Student struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.Methods("GET").Path("/students").HandlerFunc(GetAllStudentsHandler)
	r.Methods("GET").Path("/students/{id}").HandlerFunc(GetAStudentsHandler)
	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func GetAllStudentsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json;charset=UTF-8")	
	list := [...]Student{ Student{1,"catcher wong"},Student{2,"james"}}	
	result, _ := json.Marshal(list)
	w.Write(result)
}

func GetAStudentsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json;charset=UTF-8")	
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	sid, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {		
		w.Write([]byte(`{"message": "Error"}`))
		return
	}
	s := Student{sid,"catcher wong"}	
	result, _ := json.Marshal(s)
	w.Write(result)
}
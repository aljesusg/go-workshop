package main

import (
	"net/http"
	"encoding/xml"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)



type Etsiitiano struct {
	ID        string   `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Hobbies []string  `json:"hobbies,omitempty"`
	Address Address   `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var students = []Etsiitiano{}

func GetOutput(result interface{}, w http.ResponseWriter, typ string){
	w.Header().Set("Content-Type", "application/" + typ)
	log.Print("Returned in " + typ + " format")
	var response []byte;
	var err error;

	switch typ {
	case "xml":
		response, err = xml.MarshalIndent(result, "", "  ")
	case "json":
		response, err = json.Marshal(result)
	default:
		log.Print("Error no type " + typ)
		return
	}

	if err != nil {
		log.Fatal("Error in marshal")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	GetOutput(students, w, r.Header.Get("Type"))
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range students {
		if item.ID == params["id"] {
			GetOutput(item, w, r.Header.Get("Type"))
			return
		}
	}
	GetOutput(&Etsiitiano{}, w, r.Header.Get("Type"))

}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Etsiitiano
	error := json.NewDecoder(r.Body).Decode(&person)
	if error != nil{
		log.Fatal(error)
	}
	person.ID = params["id"]
	students = append(students, person)
	GetOutput(students, w, r.Header.Get("Type"))
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for index, item := range students {
		if item.ID == params["id"] {
			students = append(students[:index], students[index+1:]...)
			break
		}
		GetOutput(students, w, r.Header.Get("Type"))
	}
}


func main() {
	fmt.Println("Hello,go developers\n")
	router := mux.NewRouter()
	router.HandleFunc("/students",GetStudents).Methods("GET")
	router.HandleFunc("/student/{id}",GetStudent).Methods("GET")
	router.HandleFunc("/student/{id}",CreateStudent).Methods("POST")
	router.HandleFunc("/student/{id}",DeleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

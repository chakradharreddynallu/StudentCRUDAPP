//  Author: Manda Supraja
// email: mandasupraja1365139@gmail.com

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/k3a/html2text"
)

// ----------Student struct-----------
type student struct {
	Rollno  string   `rollno`
	Name    string   `name`
	Course  string   `Course`
	Subject *subject `subject`
}

// -----------subject struct------------
type subject struct {
	Sub1 string `sub1`
	Sub2 string `sub2`
}

// -----------Slice struct--------------
var Student []student

// -----------Main Function -------------------
func main() {
	fmt.Println("This iss a Golang Application")
	r := mux.NewRouter()

	Student = append(Student, student{Rollno: "1", Name: "ramesh", Course: "IT", Subject: &subject{Sub1: "Golang", Sub2: "Python"}})
	Student = append(Student, student{Rollno: "2", Name: "Supraja", Course: "IT", Subject: &subject{Sub1: "Golang", Sub2: "Java"}})
	Student = append(Student, student{Rollno: "3", Name: "Vilash", Course: "IT", Subject: &subject{Sub1: "Docker", Sub2: "Python"}})

	r.HandleFunc("/home", home).Methods("GET")
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/student", newStudent).Methods("POST")
	r.HandleFunc("/getStudent/{rollno}", getStudent).Methods("GET")
	r.HandleFunc("/deleteStudent/{rollno}", deleteStudent)
	r.HandleFunc("/update/{rollno}", updateStudent).Methods("PUT")

	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

// --------------Home Function-----------------
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.Error(w, "Error in path", http.StatusNotFound)
	}
	html := `<h1> Hello this is Home page</h1>`
	plain := html2text.HTML2Text(html)
	fmt.Fprintln(w, plain)

}

// ------------GetAllStudents--------------------
func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(Student)
}

// -----------DeleteStudentById---------------------
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, items := range Student {
		if items.Rollno == params["rollno"] {
			Student = append(Student[:index], Student[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Student)
}

// ------------------GetStudentById------------------
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for _, items := range Student {
		if items.Rollno == params["rollno"] {
			json.NewEncoder(w).Encode(items)
			return
		}
	}

}

// -----------------InserNewtStudent----------------------
func newStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var stud student
	_ = json.NewDecoder(r.Body).Decode(&stud)
	stud.Rollno = strconv.Itoa((rand.Intn(100000)))
	Student = append(Student, stud)
	json.NewEncoder(w).Encode(stud)
}

// -------------------UpdateStudent-------------------------

func updateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range Student {
		if item.Rollno == params["rollno"] {
			Student = append(Student[:index], Student[index+1:]...)
			var stud student
			_ = json.NewDecoder(r.Body).Decode(&stud)
			stud.Rollno = strconv.Itoa((rand.Intn(100000)))
			Student = append(Student, stud)
			json.NewEncoder(w).Encode(stud)
		}
	}

}

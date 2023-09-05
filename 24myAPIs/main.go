package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// fake model of data 😂 //i.e this goes in seperate -file
type Course struct {
	CourseId    string  `json:"couseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB 🤣 i.e this goes in seperate -file. we are creating slice to store data for now
var courses []Course

//lets create middle ware where we verify, our courseId in not nil and courseName is not nil. this is also called as helper method

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("welcome to api creation in golang tutorial")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ractjs", CoursePrice: 569, Author: &Author{FullName: "abhay", Website: "www.abhay.com"}})
	courses = append(courses, Course{CourseId: "5", CourseName: "java", CoursePrice: 769, Author: &Author{FullName: "bijay", Website: "www.bijay.com"}})

	//lets create al routes here
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("Delete")

	//listining port
	log.Fatal(http.ListenAndServe(":4000", r))
}

//lets create controller. yes, this also goes in seperate file in big projects

// serve home route // this route is when localhost:4000/ is called than there should come someting in the page.
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hi, i am abhay. This is my first api in golang.</h1>"))
}

// w of type http.ResponseWriter: This is used to write the response back to the client.
// r of type *http.Request: This represents the incoming HTTP request from the client.
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// let find one course by it id
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get once cource")
	w.Header().Set("Content-Type", "Application/json")

	//grab id from request
	param := mux.Vars(r)

	//loop though the courses
	for _, course := range courses { //courses is db
		if course.CourseId == param["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

// here a post request for creating one course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	parms := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == parms["id"] {
			_ = json.NewDecoder(r.Body).Decode(&course)
		}
	}

	// TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	// for _,course := range courses{
	// 	if course.CourseName ==
	// }

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100)) // convert random integer into a string. i.e Itoa
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// updateone cource
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content=Type", "Application/json")

	//first grab id from request
	parms := mux.Vars(r)

	//loop,id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == parms["id"] {
			courses = append(courses[:index], courses[index+1:]...) //remove data base on id
			var cource Course
			_ = json.NewDecoder(r.Body).Decode(&cource) //get body from request and decode it.
			cource.CourseId = parms["id"]
			courses = append(courses, cource)
			json.NewEncoder(w).Encode(cource) //here we send responce that course is updated. or we can send message , that course updated sucessfully
			return
		}

	}
	//TODO When id is not found
	json.NewEncoder(w).Encode("No course found with given id")
	return

}

// delete one course
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete one course from data base")
	w.Header().Set("Content-Type", "application/json")

	parms := mux.Vars(r)

	//loop, id , remove(index, index+1)

	for index, course := range courses {
		if course.CourseId == parms["id"] {
			courses := append(courses[:index], courses[index+1:]...)
			fmt.Println("course succsfully deleted", courses)
			json.NewEncoder(w).Encode("data with the given course id: " + course.CourseId + "  deleted successfully")
			break
		} else {
			fmt.Println("course not found with given id")
			json.NewEncoder(w).Encode("couse not found with the given id.")
			return

		}

	}

}

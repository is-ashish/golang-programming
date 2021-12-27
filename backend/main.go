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

// Model for courses -- file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// face DB
var courses []Course

// middleware, helper -- file

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	router := mux.NewRouter()
	courses = append(courses, Course{
		CourseId:    "1",
		CourseName:  "ReactJS",
		CoursePrice: 200,
		Author: &Author{
			FullName: "Ashish",
			Website:  "ashish.jain.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "Golang",
		CoursePrice: 600,
		Author: &Author{
			FullName: "Manish Jain",
			Website:  "manish.jain.com",
		},
	})
	router.HandleFunc("/", homeRoute).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Course Management Portal :)</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
	return
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")
	// grab ID from Request
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found for Given id" + params["id"])
}
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")
	// if Body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request Body is Empty")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data Inside JSON")
		return
	}
	// Generate Unique ID as String
	// append course into courses
	rand.Seed(time.Now().UnixNano())
	courseID := strconv.Itoa(rand.Intn(100))
	course.CourseId = courseID
	fmt.Println(course)
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Successfully added course id" + courseID)
	return
}
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One Course")
	w.Header().Set("Content-Type", "application/json")
	// if Body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request Body is Empty")
	}
	params := mux.Vars(r)
	for index, course := range courses {
		if params["id"] == course.CourseId {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("ID not found")
	return
}
func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Successfully deleted Given ID" + params["id"])
			return
		}
	}
	json.NewEncoder(w).Encode("ID Not Found")
	return
}

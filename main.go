package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Blog struct {
	Title     string
	Startdate string
	Enddate   string
	Content   string
	Check     []string
}

var Blogs = []Blog{
	{
		Title:     "test",
		Startdate: "2022-11-24",
		Enddate:   "2022-11-25",
		Content:   "Test",
		Check:     []string{"NodeJS", "ReactJS", "PHP", "Javascript"},
	},
}

func main() {
	route := mux.NewRouter()

	//route untuk menginisialisai folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/formblog", formblog).Methods("GET")
	route.HandleFunc("/blog-detail/{id}", blogDetail).Methods("GET")
	route.HandleFunc("/addblog", addblog).Methods("POST")
	route.HandleFunc("/delete-blog/{index}", deleteBlog).Methods("GET")
	route.HandleFunc("/update-blog/{index}", updateBlog).Methods("POST")
	route.HandleFunc("/update-blog/{index}", getUpdateBlog).Methods("GET")

	fmt.Println("Server berjalan pada port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}
	respData := map[string]interface{}{
		"Blogs": Blogs,
	}
	tmpt.Execute(w, respData)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, tmpt)
}

func formblog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/addblog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, tmpt)
}

func addblog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	startdate := r.PostForm.Get("std")
	enddate := r.PostForm.Get("etd")
	//get value checkbox
	check := r.Form["check"]

	var newBlog = Blog{
		Title:     title,
		Content:   content,
		Startdate: startdate,
		Enddate:   enddate,
		Check:     check,
	}
	// fmt.Println(newBlog)

	Blogs = append(Blogs, newBlog)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// get update blog[index]
func getUpdateBlog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/update-blog.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["index"])

	var UpdateDetail = Blog{}

	for index, data := range Blogs {
		if index == id {
			UpdateDetail = Blog{
				Title:     data.Title,
				Startdate: data.Startdate,
				Enddate:   data.Enddate,
				Content:   data.Content,
				Check:     data.Check,
			}
		}
	}

	Detail := map[string]interface{}{
		"Blogs": UpdateDetail,
	}

	tmpt.Execute(w, Detail)
}

// update blog berdasarkan id

func updateBlog(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	id, _ := strconv.Atoi(mux.Vars(r)["index"])

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	startdate := r.PostForm.Get("std")
	enddate := r.PostForm.Get("etd")
	//get value checkbox
	check := r.Form["check"]

	var newBlog = Blog{
		Title:     title,
		Content:   content,
		Startdate: startdate,
		Enddate:   enddate,
		Check:     check,
	}

	Blogs[id] = newBlog
	//println newblog
	// fmt.Println(newBlog)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func blogDetail(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var BlogDetail = Blog{}

	for index, data := range Blogs {
		if index == id {
			BlogDetail = Blog{
				Title:     data.Title,
				Startdate: data.Startdate,
				Enddate:   data.Enddate,
				Content:   data.Content,
				Check:     data.Check,
			}
		}
	}

	Detail := map[string]interface{}{
		"Blogs": BlogDetail,
	}

	tmpt.Execute(w, Detail)
}

func deleteBlog(w http.ResponseWriter, r *http.Request) {

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	Blogs = append(Blogs[:index], Blogs[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

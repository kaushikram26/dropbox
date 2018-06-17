package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}



func (p *Page) save() error {
	filename := p.Title + ".txt"
	//last arguement is file permission
	//create a file and store the body into it
	return ioutil.WriteFile(filename,p.Body,0600)
}

func loadPage(title string) (*Page,error) {
	filename := title + ".txt"
	body,err := ioutil.ReadFile(filename)
	if err != nil{
		return nil,err
	}
	return &Page{Title: title,Body: body},nil
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hi this is kaushik, test string %s",r.URL.Path)
}

func viewHandler(w http.ResponseWriter , r *http.Request){
	title := r.URL.Path[len("/view/"):]
	p,_ := loadPage(title)
	fmt.Fprintf(w,"<h1>%s</h1><div>%s</div>",p.Title,p.Body)
}




func main(){
	http.HandleFunc("/view/",viewHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}


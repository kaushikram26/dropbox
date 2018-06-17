package main

import (
	"fmt"
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

func main(){
	p1 := &Page{Title: "kaushik_page",Body: []byte("This is a kaushik_page")}
	p1.save()
	p2,_ := loadPage("kaushik_page")
	fmt.Println(string(p2.Body))
}
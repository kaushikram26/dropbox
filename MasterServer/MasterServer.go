//package server
package main

import(
	"fmt"
	"encoding/gob"
	//"bytes"
	//"io/ioutil"
	"os"
	"log"
)


//for now username and password user details
type UserDetails struct{
	UserName string
	Password string
}

const FileStorage string = "UserData.bkp"

func ReadIntoMap() map[string]string{

	m := make(map[string]string)
	//read from file
	m["test1"] = "test_value1"
	m["test2"] = "test_value2"
	m["test3"] = "test_value3"

	return m
}

func ErrorHandler(e error){

	if e != nil{
		fmt.Println("Error problem",e)
		log.Fatal(e)
	}
}

func StoreMap(object interface{}) {
	
	fd, err := os.OpenFile(FileStorage,os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	ErrorHandler(err)
	
	encoder := gob.NewEncoder(fd)
	encoder.Encode(object)
	fmt.Println("value :",encoder)
	fd.Close()
	//ErrorHandler(err)
	

}

func InitializeServer() map[string]string{
	//Initialize the server. Read from hashmap to memory
	//var m map[string]string
	m := ReadIntoMap()
	return m
}


func New(username string,password string) UserDetails {
	userdetails := UserDetails { UserName: username , Password: password }
	return userdetails
}



func SaveUserData(userdata UserDetails){
	
	//open file and save in file

	//add to hashmap

}


func TestMap(m map[string]string){
	//m["test4"] = "test_value4"
	//m["test5"] = "test_value5"
	//m["test6"] = "test_value6"

	for key,value := range m{
		fmt.Println("Key : ",key," value: ",value )
	}
}

func ReadMap(m interface{}) error{
	fd,err := os.Open(FileStorage)
	ErrorHandler(err)
	//decode the file 
	decoder := gob.NewDecoder(fd)
	decoder.Decode(m)
	fd.Close()
	fmt.Println("value :",decoder)
	return err
}

func main(){
	m := InitializeServer()
	
	StoreMap(m)
	TestMap(m)
	//err := ReadMap(m)
	//if err==nil{
	//	for key,value := range m{
	//		fmt.Println("Key : ",key," value: ",value )
	//	}
	//}

}



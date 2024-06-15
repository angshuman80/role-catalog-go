package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)



func main() {
  // http.HandleFunc("/",roleCatagHandler)
   r := http.NewServeMux()
   r.HandleFunc("/",roleCatagHandler)
   r.HandleFunc("/employee/{id}",dynamicHandler)
   r.HandleFunc("/roleCatalog",roleCatalogJsonHandler)
   log.Fatal(http.ListenAndServe(":8090",r))
}

func roleCatagHandler(w http.ResponseWriter , r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func dynamicHandler(w http.ResponseWriter , r *http.Request){
	pathValue := r.PathValue("id")
	fmt.Println("Path Value ",pathValue," ",r.RequestURI)
	fmt.Fprintf(w," Hello world with id")
}

func roleCatalogJsonHandler(w http.ResponseWriter , r *http.Request){
      w.WriteHeader(http.StatusCreated)
	  w.Header().Add("Content-Type","application/json")
	  //rCat := model.populateRoleCatalog()
	  //resp , err := json.Marshal(rCat)
	  //if err!=nil{
	//	fmt.Fprintf(w, "something bad happened")
	//	return
	//  }
	 // w.Write(resp)
	  return
	 // rolCat := model.populateRoleCatalog()

}
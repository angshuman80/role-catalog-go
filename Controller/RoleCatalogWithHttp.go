package Controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"role-catalog-go/model"
)

func InitializeHttpRoute() error {
	r := http.NewServeMux()
	r.HandleFunc("/roleCatalog/{id}", roleCatalogdynamicHandler)
	r.HandleFunc("/roleCatalog", roleCatalogJsonHandler)
	log.Fatal(http.ListenAndServe(":8090", r))
	return nil
}

func roleCatalogdynamicHandler(w http.ResponseWriter, r *http.Request) {
	pathValue := r.PathValue("id")
	fmt.Println("Path Value ", pathValue, " ", r.RequestURI)
	result := "Role Catalog" + pathValue
	_, err := fmt.Fprintf(w, result)
	if err != nil {
		return
	}
}

func roleCatalogJsonHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	rCat := model.PopulateRoleCatalog()
	resp, err := json.Marshal(rCat)
	if err != nil {
		_, err := fmt.Fprintf(w, "something bad happened")
		if err != nil {
			return
		}
	}
	_, err = w.Write(resp)
	if err != nil {
		return
	}
	return
}

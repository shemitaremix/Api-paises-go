package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Bienvenido a esta api papu: con /countries puedes ver los paises")
		return
	}

	fmt.Fprintf(w, "Bienvenido a esta api papu: con /countries puedes ver los paises: con metodo GET puedes ver los paises y con metodo POST puedes agregar un pais")
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)

}

func addCountry(w http.ResponseWriter, r *http.Request) {

	country := &Country{}

	err := json.NewDecoder(r.Body).Decode(&country)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "$v", err)

		return
	}

	countries = append(countries, *country)
	fmt.Fprintf(w, "pais agregado")
}

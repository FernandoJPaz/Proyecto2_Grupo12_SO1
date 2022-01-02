package main

import (
	//"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v7"
	"strconv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-redis/redis/v7"
)

type Persona struct {
	Name          string `json:"name"`
	Location      string `json:"location"`
	Age           int `json:"age"`
	Vaccine_type string `json:"vaccine_type"`
	NDose         int `json:"n_dose"`
}

func main() {

	cliente := redis.NewClient(&redis.Options{

		Addr: "35.223.95.97:6379",
		Password: "sopes1pass123",
		DB:       0,
	})

	http.HandleFunc("/api/persona", func(res http.ResponseWriter, req *http.Request) {

		enableCors(&res)
		if req.Method == "POST" {
			fmt.Println("Recibido");
			if err := req.ParseForm(); err != nil {
				fmt.Printf("ParseForm() err: %v", err)
				return
			}

			var p Persona
			err := json.NewDecoder(req.Body).Decode(&p)
			if err != nil {
				//http.Error(res, err.Error(), http.StatusBadRequest)
				fmt.Print("error")
				//return
			}
			cliente.Publish("canal1", p.Name+","+p.Location+","+strconv.Itoa(p.Age)+","+p.Vaccine_type+","+strconv.Itoa(p.NDose))
			cliente.LPush("milista", "{'name':'"+p.Name+"','location':'"+p.Location+"','age':"+strconv.Itoa(p.Age)+",'vaccine_type':'"+p.Vaccine_type+"','n_dose':'"+strconv.Itoa(p.NDose)+"'}")
		} else {

			io.WriteString(res, "hola mundo2")
		}

	})

	puerto := ":3000"
	fmt.Println("Servidor escuchando en ..." + puerto)
	http.ListenAndServe(puerto, nil)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

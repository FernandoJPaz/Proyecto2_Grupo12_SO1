package main

import (
	"log"
	"strconv"
	//"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"

	"encoding/json"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/fernandopaz/proyecto2/Ruta1/vacuna"
)


type casoS struct {
		Name string             `json:"name"`
		Location  string        `json:"location"`
		Age    int32            `json:"age"`
		VaccineType string 	`json:"vaccine_type"`
		NDose int32			`json:"n_dose"`
	
}

var edad string
var edadI int64
var err error
//InsercionHandler funcion
func InsercionHandler (w http.ResponseWriter, request *http.Request) {
	enableCors(&w)
	//datos,err := ioutil.ReadAll(request.Body)

	errParse := request.ParseForm()
	/*log.Println(datos)*/
	if errParse != nil {
		w.Write([]byte("Error al validar los datos enviados"))
		return
	}
	//var casso casoS;
	var cassoc vacuna.Registro
	err := json.NewDecoder(request.Body).Decode(&cassoc)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
	edad = request.PostForm.Get("age")
	edadI ,err =  strconv.ParseInt(edad,10,32)
	if err != nil{
		edadI = -1
	}

	//log.Println(cassoc)
	//log.Printf("Recibido:%s",casoRecibido.Name)
	go ReenvioCaso(&cassoc)
	w.Write([]byte("{Enviado a Insertar"))
}

//DefaultHandler funcion
func DefaultHandler (w http.ResponseWriter, request *http.Request) {
	enableCors(&w)
	casoPrueba := vacuna.Registro{
		Name: "Marta Kent",
		Location: "Guatemala City",
		Age: 50,
		VaccineType: "Sputnik V",
		NDose: 2,
	}
	go ReenvioCaso(&casoPrueba)
	w.Write([]byte("Se agrego un registro default"))
}

var conn *grpc.ClientConn
var ctxWithW2 context.Context
var c vacuna.VacunaServiceClient
//var cancel CancelFunc
func main(){
	
	conn,err := grpc.Dial("grpc-server:9000",grpc.WithInsecure())
	//172.17.0.2
	//conn,err := grpc.Dial("172.17.0.2:9000",grpc.WithInsecure())
	if err != nil{
		log.Fatalf("No se pudo conectar al servidor: %s",err)

	}
	defer conn.Close()

	//c = covid.NewCovidServiceClient(conn)
	c = vacuna.NewVacunaServiceClient(conn)
	log.Println("Cliente listo!!")
	
	//servidor web
	//router para los endpoit
	router := mux.NewRouter()

	//endpoints

	router.HandleFunc("/insertar", InsercionHandler).Methods("POST")
	router.HandleFunc("/api/persona", InsercionHandler).Methods("POST")
	router.HandleFunc("/default", DefaultHandler).Methods("GET")

	//configuración de listener
	log.Fatal(http.ListenAndServe(":3000", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//ReenvioCaso reenvia el caso a traves de la conexion gRPC al servidor, de manera asyncrona
func ReenvioCaso(caso *vacuna.Registro){
	//log.Println(caso)
	_, err := c.AgregarRegistro(context.Background(),caso)
	if err != nil {
		log.Printf("Error al llamar a agregar Caso: %s",err)
	}
}
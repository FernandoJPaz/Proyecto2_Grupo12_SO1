package vacuna

import (
	"log"
	//"time"
	"golang.org/x/net/context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	

	"github.com/go-redis/redis/v7"
)


//Server es un struct
type Server struct {
	Nombre string
	ConnMongo * MongoConnection
	UnimplementedVacunaServiceServer
}

//MongoConnection struct con info de base de datos
type MongoConnection struct{
	CovidCollection      *mongo.Collection
    Cliente *mongo.Client
    //logger  *logrus.Logger
}

//CasoB struct con la información del caso
type CasoB struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Location  string             `bson:"location"`
	Age    int32             `bson:"age"`
	VaccineType string 	`bson:"vaccine_type"`
	NDose int32		`bson:"n_dose"`
}

//AgregarCaso metodo que agrega un caso a la base de datos
func (s *Server) AgregarRegistro(ctx context.Context, in *Registro) (*Respuesta, error) {
	//insercion asincrona
	go InsercionAsincrona(s,in)
	//retorno de valor
	return &Respuesta{Resultado:1}, nil
}

//InsercionAsincrona es una funcion
func InsercionAsincrona(s *Server,caso *Registro){
	registrorecivido := CasoB{
		Name : caso.GetName(),
		Location: caso.GetLocation(),
		Age: caso.GetAge(),
		VaccineType: caso.GetVaccineType(),
		NDose: caso.GetNDose(),
	}
	//insercion de objeto
	//log.Println(caso)
	_ , err2 := s.ConnMongo.CovidCollection.InsertOne(context.TODO(),registrorecivido)
	if err2 != nil {
		log.Println(err2)
	}
	

	cliente := redis.NewClient(&redis.Options{

		Addr: "35.223.95.97:6379",
		Password: "sopes1pass123",
		DB:       0,
	})
	//casoJson := fmt.
cliente.LPush("milista", "{'name':'"+registrorecivido.Name+"','location':'"+registrorecivido.Location+"','age':"+fmt.Sprint(registrorecivido.Age)+",'vaccine_type':'"+registrorecivido.VaccineType+"','n_dose':'"+fmt.Sprint(registrorecivido.NDose)+"'}")
cliente.Close()
}
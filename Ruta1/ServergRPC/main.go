package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"github.com/fernandopaz/proyecto2/Ruta1/vacuna"
	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	//mongodb+srv://myUserAdmin:sopes1pass123@cluster0.w5dzu.mongodb.net/Vacunas?retryWrites=true&w=majority
	clientOptions := options.Client().ApplyURI("mongodb://myUserAdmin:sopes1pass123@35.223.95.97:27017/")
    //clientOptions := options.Client().ApplyURI("mongodb+srv://myUserAdmin:sopes1pass123@cluster0.w5dzu.mongodb.net/Vacunas?retryWrites=true&w=majority")
    
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil { log.Fatal(err) }
	coleccion := client.Database("Vacunas").Collection("Registro")
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Se ha establecido una conexion con MongoDB")
	mongoConn := vacuna.MongoConnection{
		CovidCollection : coleccion,
		Cliente : client,
	} 

	s:= vacuna.Server{
		Nombre: "Server gRPC SO1-Proyecto2",
		ConnMongo : &mongoConn,
	}

	grpcServer := grpc.NewServer()

	vacuna.RegisterVacunaServiceServer(grpcServer, &s)

	
	log.Println("El servidor esta listo para recibir peticiones")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %s", err)
	}
	client.Disconnect(context.TODO())

}
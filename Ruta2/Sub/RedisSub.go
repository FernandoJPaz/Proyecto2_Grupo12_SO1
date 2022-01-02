package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	cliente := redis.NewClient(&redis.Options{
		Addr: "35.223.95.97:6379",
		Password: "sopes1pass123",
		DB:       0,
	})

	sub := cliente.Subscribe("canal1")
	fmt.Println("Suscrito al canal y listo para recibir mensajes.....")
	for {
		
		time.Sleep(100 * time.Millisecond)
		msg, err := sub.ReceiveMessage()

		if err != nil {

			panic(err)
		}
		//fmt.Println(msg.Payload);
		Vcadena := strings.Split(msg.Payload, ",")

		/*fmt.Println(Vcadena[0])
		fmt.Println(Vcadena[1])
		fmt.Println(Vcadena[2])
		fmt.Println(Vcadena[3])
		fmt.Println(Vcadena[4])
		*/
		//******Mongo
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://myUserAdmin:sopes1pass123@35.223.95.97:27017/"))
		//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://myUserAdmin:sopes1pass123@cluster0.w5dzu.mongodb.net/Vacunas?retryWrites=true&w=majority"))
		
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal(err)
		}
		defer client.Disconnect(ctx)
		base_casos := client.Database("Vacunas")
		collection_casos := base_casos.Collection("Registro")

		//*******
		_, err = collection_casos.InsertOne(ctx, bson.D{
			{Key: "name", Value: Vcadena[0]},
			{Key: "location", Value: Vcadena[1]},
			{Key: "age", Value: Vcadena[2]},
			{Key: "vaccine_type", Value: Vcadena[3]},
			{Key: "n_dose", Value: Vcadena[4]},
		})

		//fmt.Println(resInsercion)
		if err != nil {
			log.Fatal(err)
		}
	}

}

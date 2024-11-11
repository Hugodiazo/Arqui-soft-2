// cursos-api/db/mongodb.go
package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

// Conectar a MongoDB
func ConnectMongoDB(uri, dbName string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	// Verifica la conexión
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error al verificar la conexión con MongoDB: %v", err)
	}

	MongoDB = client.Database(dbName)
	fmt.Println("Conexión a MongoDB establecida con éxito")
}

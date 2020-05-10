package arango

import (
	"context"
	"github.com/arangodb/go-driver"
	arangoHttp "github.com/arangodb/go-driver/http"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var arangoClient driver.Client
var arangoDbInstance driver.Database

func InitializeArangoDatabase() {
	once := sync.Once{}
	once.Do(func() {
		conn := createArangoConnection()
		arangoClient = createArangoClient(conn)
		arangoDbInstance = createDbInstance(arangoClient)
	})
}

func createArangoConnection() driver.Connection {
	host := viper.GetString("arango.host")
	port := viper.GetString("arango.port")
	if host == "" || port == "" {
		log.Fatalln("Arango host/port config not found")
	}

	conn, err := arangoHttp.NewConnection(arangoHttp.ConnectionConfig{
		Endpoints: []string{"http://" + host + ":" + port},
	})
	if err != nil {
		log.Fatalln("Error in Arango Client Creation::", err.Error())
	}
	log.Println("Arango connected")
	return conn
}

func createArangoClient(conn driver.Connection) driver.Client {
	username := viper.GetString("arango.username")
	password := viper.GetString("arango.password")
	if username == "" || password == "" {
		log.Fatalln("Arango username/password config not found")
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(username, password),
	})
	if err != nil {
		log.Fatalln("Error creating Arango Client:: ", err.Error())
	}
	log.Println("Arango client created")
	return client
}

func createDbInstance(client driver.Client) driver.Database {
	dbName := "rb-wheels"
	ctx := context.Background()
	if _, err := client.DatabaseExists(ctx, dbName); err != nil {
		log.Fatalf("error accessing db: %s :: %s\n", dbName, err.Error())
	}

	dbInstnce, err := client.Database(ctx, dbName)
	if err != nil {
		log.Fatalf("error getting %s db instance:: %s\n", dbName, err.Error())
	}

	return dbInstnce
}

func getDbInstance() driver.Database {
	return arangoDbInstance
}

func getCollectionInstance(collectionName string) driver.Collection {
	ctx := context.Background()
	db := getDbInstance()

	exist, err := db.CollectionExists(ctx, collectionName)
	if err != nil || !exist {
		log.Fatalln(collectionName, " collection not found:: ", err.Error())
	}

	coll, dbErr := db.Collection(ctx, collectionName)
	if dbErr != nil {
		log.Fatalln("get collection error: ", collectionName, " :: ", dbErr.Error())
	}
	return coll
}

package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"samt/api"
	"samt/db/makedb"
	db "samt/db/sqlc"
)

func main() {
	MakeDir()
	err := godotenv.Load("config.env")

	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
		return
	}

	serveraddress := os.Getenv("SERVERADDRESS")

	conn, err := makedb.SetUpDataBase()

	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		os.Exit(0)
	}

	mystore := db.NewStore(conn)

	myserver, err := api.NewServer(mystore)

	if err != nil {
		log.Fatal("can not create server = ", err)
		return
	}
	myserver.Start(serveraddress)
}
func MakeDir() {
	if _, err := os.Stat("./images"); os.IsNotExist(err) {
		err := os.Mkdir("./images", 0755)
		if err != nil {
			log.Fatal("can not create image dir")
			return
		}
	}
}

package main

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	keyvaluestore "keyvaluestore.com/hrs/client/com.key.value.grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:7501", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	c := keyvaluestore.NewKeyValueServiceClient(cc)

	// Contact the server and print out its response.
	name := "dbdbdbdbdb"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateDatabase(ctx, &keyvaluestore.CreateDatabase{Database: name})
	if err != nil {
		log.Fatalf("could not create database: %v", err)
	}
	log.Printf("CreateDatabase: %s", r.Message)

}

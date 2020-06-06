package f

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func initDb(ctx context.Context) *mongo.Database {
	var newDb *mongo.Database
	var clientOptions *options.ClientOptions
	var dbURI string
	nodeEnv := os.Getenv("NODE_ENV")
	dbEnv := os.Getenv("DB_ENV")

	if dbEnv == "prd" {
		log.Printf("Target: PRD Cluster")
		dbURI = os.Getenv("PRD_URI")
		clientOptions = options.Client().ApplyURI(os.Getenv("PRD_URI"))
	} else if dbEnv == "tst" || dbEnv == "dev" {
		log.Printf("Target: Local DB")
		dbURI = os.Getenv("DEV_URI")
		clientOptions = options.Client().ApplyURI(os.Getenv("DEV_URI"))
	} else {
		log.Fatalf("[CRITICAL] Undefined phase!\n")
	}

	newClient, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("[CRITICAL] Error initialising client. URI: %s\n", dbURI)
	}

	// To be removed when another DB URI is used
	if nodeEnv == "prd" {
		newDb = newClient.Database("games_stats_app")
		log.Printf("Target: PRD collection\n")
	} else if nodeEnv == "dev" || nodeEnv == "tst" {
		newDb = newClient.Database("games_stats_app_TST")
		log.Printf("Target: DEV collection\n")
	} else {
		log.Fatalf("[CRITICAL] Undefined phase!\n")
	}

	err = newClient.Connect(ctx)
	if err != nil {
		log.Fatalf("[CRITICAL] error connecting client. %s\n", err)
	}

	return newDb
}

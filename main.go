package main

import (
	"context"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	ResourceGroup = ""
	Server        = ""
	Database      = ""
)

var (
	newDBPlan = "S4"
)

func main() {

	// Create new authorization from env variables. see file: .env.example
	auth, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	// CLIENT to work with SQL REST API
	dbClient := sql.NewDatabasesClient(os.Getenv("AZURE_SUBSCRIPTION_ID"))
	dbClient.Authorizer = auth

	ctx := context.Background()

	// UPDATE the SKU/plan
	future, err := dbClient.Update(
		ctx,
		ResourceGroup,
		Server,
		Database,
		sql.DatabaseUpdate{
			Sku: &sql.Sku{
				Name: &newDBPlan,
			},
		},
	)

	if err != nil {
		log.Fatal("Failed to update DB", err)
	}

	log.Println(future.Status())

	// otherwise wait for scale operation to complete
	err = future.WaitForCompletionRef(ctx, dbClient.Client)
	if err != nil {
		log.Fatal("Scaling operation to S4 failed", err)
	}

	log.Println("DB scale operation to", newDBPlan, future.Status())
}

package loucore

import (
	"context"
	"github.com/joaopedrosgs/loucore/ent"
	"github.com/joaopedrosgs/loucore/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var client *ent.Client

func SetupStorage() error {
	var err error
	client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if err != nil {
		return err
	}
	return client.Schema.Create(context.Background(), migrate.WithGlobalUniqueID(true))

}
func SetupStorageCustom(driverName string, dataSourceName string) {
	var err error
	client, err = ent.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}

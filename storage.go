package loucore

import (
	"log"
	"lou-core/ent"
)

var client *ent.Client

func SetupStorage() {
	var err error
	client, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
func SetupStorageCustom(driverName string, dataSourceName string) {
	var err error
	client, err = ent.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}

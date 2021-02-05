package models

import (
	"fmt"

	"github.com/couchbase/gocb"
)

var EventBucket *gocb.Bucket
var ReadBucket *gocb.Bucket

func ConnectToCouchBase(host, user, password string) {
	fmt.Println(host)
	// fmt.Println("Waiting for couchBase")
	// time.Sleep(12 * time.Second)
	fmt.Println("Connecting to CouchBase")

	cluster, couchConErr := gocb.Connect(fmt.Sprintf("couchbase://%s", host))

	if couchConErr != nil {
		fmt.Println("CouchBase Connect error")
		panic(couchConErr)
	}

	authErr := cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: user,
		Password: password,
	})

	if authErr != nil {
		fmt.Println("CouchBase Auth error")
		panic(authErr)
	}

	tempEventBucket, eventConErr := cluster.OpenBucket("event-store", "")

	if eventConErr != nil {
		fmt.Println("Couch Event bucket error")

		panic(eventConErr)
	}

	tempReadBucket, readConErr := cluster.OpenBucket("read-models", "")

	if readConErr != nil {
		fmt.Println("Couch Read bucket error")

		panic(readConErr)
	}

	if readConErr != nil {
		fmt.Println("Couch Event bucket error")
		panic(readConErr)
	}

	fmt.Println("CouchBase successfully connected!")

	EventBucket = tempEventBucket
	ReadBucket = tempReadBucket
}

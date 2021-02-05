package main

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
	"github.com/joho/godotenv"
)

func Initialize(couchHost, couchUser, couchPw string) {
	models.ConnectToCouchBase(couchHost, couchUser, couchPw)
}

func main() {
	env, _ := godotenv.Read(".env")

	Initialize(env["COUCH_HOST"], env["COUCH_USER"], env["COUCH_PASSWORD"])
}

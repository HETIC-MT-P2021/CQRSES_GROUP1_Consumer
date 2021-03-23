package main

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/rabbit"
	"github.com/joho/godotenv"
)

func Initialize(env map[string]string) {
	models.ConnectToCouchBase(env["COUCH_HOST"], env["COUCH_USER"], env["COUCH_PASSWORD"])

	rabbit.ConnectToRabbit(
		env["RABBIT_HOST"],
		env["RABBIT_PORT"],
		env["RABBIT_USER"],
		env["RABBIT_PASSWORD"])

	forever := make(chan bool)

	rabbit.ReceiveCreatePost()
	rabbit.ReceiveUpdatePost()

	<-forever
}

func main() {
	env, _ := godotenv.Read(".env")

	Initialize(env)
}

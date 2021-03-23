# CQRSES_GROUP1_Consumer


The goal of this project is to process commands from the [main app](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1).
We used Golang, A RabbitMQ and a CouchBase Database
This repository hosts the consumer application.
The goal of this project is to process commands from the [main app](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1).
We used Golang, RabbitMQ and a CouchBase Database

You can find the main project and full information about the project in the main app.


## Starting the project

After cloning the repo, `cd` into the project, create the .env according to .env.example.
Start the main project first, then, on localhost, decomment the lign 16 of docker-compose.
Then run the following commands

```bash
docker-compose up --build
```

Your consumer is ready to go !



- Go version: `1.15`

## Contributing

See [CONTRIBUTING.MD](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1/blob/main/CONTRIBUTING.MD)



### Authors 🏄 


- [Tsabot](https://github.com/Tsabot)
- [myouuu](https://github.com/myouuu)
- [acauchois](https://github.com/acauchois)
- [gensjaak](https://github.com/gensjaak)

### Tasks management 🎨
We like to iterate quickly, setting up an agile board helps everyone stay on task.  With Trello, we Created lists for backlogged items, what’s being worked on in the current sprint, and (most importantly) what’s completed.

You can take a look at our organization of tasks on [Trello](https://trello.com/b/uY6KOh4i/go-cqrs)


### License 🔖

The code is available under the MIT license.

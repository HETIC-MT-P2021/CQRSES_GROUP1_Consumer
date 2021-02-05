# CQRSES_GROUP1_Consumer

The goal of this project is to process commands from the [main app](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1).
We used Golang, A RabbitMQ and a CouchBase Database

## Starting the project

After cloning the repo, `cd` into the project, create the .env according to .env.example.
Start the main project first, then, on localhost, decomment the lign 16 of docker-compose.
Then run the following commands

```bash
docker-compose up --build
```

Your consumer is ready to go !

### Documentation

You can find the api doc by clicking on the link below :

[Swagger](https://app.swaggerhub.com/apis-docs/acauchois/GoTemplate/1.0.0)

### Technical Choices

Feel free to discuss with any contributor about the technical choices that were made.

- Go version: `1.15`

## Contributing

See [CONTRIBUTING.MD](https://github.com/HETIC-MT-P2021/CQRSES_GROUP1/blob/main/CONTRIBUTING.MD)

### Authors

- [Tsabot](https://github.com/Tsabot)
- [myouuu](https://github.com/myouuu)
- [acauchois](https://github.com/acauchois)
- [gensjaak](https://github.com/gensjaak)

### License

The code is available under the MIT license.

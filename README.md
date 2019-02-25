# game
Mango API: Game

## Run with Docker
*$ go build
*$ docker build -t avosa/game:dev .
*$ docker rm gameDEV
*$ docker run -d --network host --name gameDEV avosa/game:dev 
*$ docker logs gameDEV

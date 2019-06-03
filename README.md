# game
Mango API: Game
Helps to apply discretion.
https://www.slideshare.net/AbhishekJain189/gamification-in-enterprise-software
https://uxmag.com/articles/five-steps-to-enterprise-gamification

## Run with Docker
* $ docker build -t avosa/game:dev .
* $ docker rm GameDEV
* $ docker run -d -p 8100:8100 --network mango_net --name GameDEV avosa/game:dev
* $ docker logs GameDEV
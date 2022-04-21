# Price Oscilation Bot

## Hi there 👋

Connect public ticker and retrieve the Currency Pair rate every fetch interval that the user whants. Each time you retrieve a new rate, the bot compare it with the first one and decide if it should alert of an oscillation. 

## Installation

You need go version 1.17 or greater.

Clone the repo 

```bash
  git clone github.com/alex-necsoiu/uphold-bot.git
  
```

  Open the following folder and execute the bash file 

```bash
  cd uphold-bot/bot/pkg 
  ./d build
``` 
## Running Tests

To run tests, run the following bash command

```bash
  cd uphold-bot/bot/pkg 
  ./d test
```

  
## Demo

To run the demo, run the following command:

```bash
  cd uphold-bot/bot/pkg 
  ./d run
```

## Dockerize the app

To dockerize the app, run the following command:

```bash
  cd uphold-bot/bot/pkg 
  ./d docker-build
```

## Run Docker app

To dockerize the app, run the following command:

```bash
  cd uphold-bot/bot/pkg 
  ./d docker-run
```

## Authors

- [@alex-necsoiu](https://www.github.com/alex-ncsoiu)

  
  
## License

[MIT](https://choosealicense.com/licenses/mit/)

# Price Oscilation Bot

## Hi there 👋

Connect to Uphold public ticker and retrieve the Currency Pair rate every 5 seconds. Each time you retrieve a new rate, the bot must compare it with the first one and decide if it should alert of an oscillation. For the purpose of this exercise we want to be alerted if the price changes 0.01 percent in either direction (price goes up or down).

## Installation

Clone the repo 

```bash
  git clone github.com/alex-necsoiu/uphold-bot.git
  
```

  Open the following folder and execute the  

```bash
  cd uphold-bot/bot/pkg 
  ./d build
``` 
## Running Tests

To run tests, run the following command

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

## Authors

- [@alex-necsoiu](https://www.github.com/alex-ncsoiu)

  
  
## License

[MIT](https://choosealicense.com/licenses/mit/)
# Tonycat
A telegram chat robot   

### How to set config  
You can modify env.go file, set config default value as your bot.    
``` 
var BOT_TOKEN = utils.GetEvnWithDefaultVal("BOT_TOKEN", "")            //set your telegram bot token, you can reference [here](https://core.telegram.org/bots/api)  
var WEBHOOK_URL = utils.GetEvnWithDefaultVal("WEBHOOK_URL", "")        //set bot callback webhook address  
var BOT_USER_NAME = utils.GetEvnWithDefaultVal("BOT_USER_NAME", "")   //set your bot username which from telegram bot profile   
var TURING_APIKEY = utils.GetEvnWithDefaultVal("TURING_APIKEY", "")   //set turing APIKey, You can register one turing account by the [link](http://www.turingapi.com/), after that, you can create a app and get a apikey from this app  
var TURING_USERID = utils.GetEvnWithDefaultVal("TURING_USERID", "")   //set turing user id, You can set this at will  
var TURING_OPEN_API = utils.GetEvnWithDefaultVal("TURING_OPEN_API", "http://openapi.tuling123.com/openapi/api/v2")  //set turing open api url    
```

### How to build  
To run this command: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -a -o tonycat .   

### How to deploy   
1, To create a docker-compose.yml in this project.   
You can reference below this:  
```
version: '3.5'

x-logging:
  &default-logging
  driver: "json-file"
  options:
    max-size: "100m"
    max-file: "10"

services:
  tonycat:
    build: .
    image: xxx.com/robot/tonycat:latest
    restart: always
    logging: *default-logging
    environment:
      GIN_MODE: 'release'
      WEBHOOK_URL: ''
      BOT_USER_NAME: ''
      TURING_APIKEY: ''
      TURING_USERID: ''
      TURING_OPEN_API: ''
      
    ports:
      - "127.0.0.1:1300:1300"
```  
2, To run docker-compose build and docker-compose push . You done !   

3, Run this image in your server, your bot is online.  



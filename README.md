# stock-market
An experimental application trying to realize polyglot programming.

<img src="http://i66.tinypic.com/i5wrno.png" style="width: 100%;">

### Technical Spec
* [Python](www.python.org) / as CSV parser
* [Go](www.golang.org) / as RESTful API & Server-Sent Events combiner
* [Lua](www.lua.org) / as Redis publish & subscribe listener & monitor tool
* [Node.js](www.nodejs.org) / as Javascript hero
* [Babel (ES6)](babel.io) / as ES6 dialect
* [Webpack](webpack.github.io) / as bundling takinardi
* [Redis](redis.io) / as key-value hulk
* [Server-Sent Events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events) / as events sending from server-side
* [React](https://facebook.github.io/react/) / as client-side view and state control
* [React Sparklines](https://github.com/borisyankov/react-sparklines) / as line chart taranta

### Setup

##### Install
* [Python](https://www.python.org/downloads/)
* [LuaRocks](https://github.com/keplerproject/luarocks/wiki/Installation-instructions-for-Mac-OS-X)
* [Lua](http://www.lua.org/download.html)
* [Redis](http://redis.io/download)
* [Go](https://golang.org/dl/)
* [Node](https://nodejs.org/en/download/)

##### Install Node Dependencies
```javascript
npm install
```

##### App Lifecycle
* Start Redis
```
redis-server
```

##### Data Layer
* csvReader, parses the data from the web specified with the links in urlDispatcher.py and notifies the redis instance when the new data has come. The data is parsed every minute by requesting the endpoint.

* Start crawling by make csvReader.py started
Description: Every minute, the python script fetches the CSV data and passes it through Redis instance. The keys are specified as the company names like 'facebook', 'google', 'apple' and 'yahoo' with their full of values. It also publishes an event over (again) Redis like this dialect: "data:facebook:stock_market'". 
```python
python csvReader.py
```

##### Monitoring Layer
* eventManager.lua, follow the tracks of the updated data on the working redis instance and make it print into the terminal for monitoring pusposes. This can also be implemented with another UI tool to wath what is happening right there 

* Start monitoring by make eventManager.lua started
Description: This Lua script is an experimental thing. It listens the events published through Redis instance and pushes the whole data every minute when the events fired onto the terminal.
```lua
lua eventManager.lua
```

##### Server Layer
* sseBroker.go, is an event generator to be sent to the client-side with an event-stream model structured on Server-Sent Events. When the data comes from the publish/subscribe channel, it makes the client-side recognized that the data is currently available to use.
* server.go is a simple static file server to run the index.html file to represent the data and the user interface together.

* Start serving the local clientLayer/index.html page by make static.go started
Description: Here we have two important file to activate server and server-sent event supported API endpoint. Broker here passes data with "event-stream" value of "Content-Type" when every event fired in the Redis instance to the client-side. This also serves a static file called index.html containing the server-sent event listener as an event binding.
```go
go run sseBroker.go
```

##### Client Layer
* index.html is the main file to show the charts. React and React Sparklines are used to show the data's history in the view side.
* client.js, contains the js code for the components of React

* Start bundling tool to compile the client-side files
Description: Here I use webpack to bundle the client files and its dialects. ES6 and its features used with React and React Sparklines libraries. It generates a bundle file called client.bundle.js.
```javascript
npm run compile
```

* Start bundling server to serve the client-side starter file, clientLayer/index.html.
Description: It will generate a bundle file called clientLayer/client.bundle.js
```go
npm run start
```

* Open the browser and hit:
```
localhost:8000/app/
```


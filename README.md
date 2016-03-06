# stock-market
An experimental application trying to realize polyglotism

<img src="http://i66.tinypic.com/i5wrno.png" style="width: 100%;">

###Technical Spec
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

###Technical Story

####Data Layer
* CvsNotifier, parses the data from the web specified with the links in urlDispatcher.py and notifies the redis instance when the new data has come. The data is parsed every minute by requesting the endpoint.

####Persistance Layer
* eventManager.lua, follow the tracks of the updated data on the working redis instance and make it print into the terminal for monitoring pusposes. This can also be implemented with another UI tool to wath what is happening right there 

####Server Layer
* sseBroker.go, is an event generator to be sent to the client-side with an event-stream model structured on Server-Sent Events. When the data comes from the publish/subscribe channel, it makes the client-side recognized that the data is currently available to use.
* server.go is a simple static file server to run the index.html file to represent the data and the user interface together.

###Client Layer
*index.html is the main file to show the charts. React and React Sparklines are used to show the data's history in the view side.


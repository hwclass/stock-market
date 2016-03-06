package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/garyburd/redigo/redis"
)

//	"encoding/json"

type Broker struct {

	// Events are pushed to this channel by the main events-gathering routine
	Notifier chan []byte

	// New client connections
	newClients chan chan []byte

	// Closed client connections
	closingClients chan chan []byte

	// Client connections registry
	clients map[chan []byte]bool
}

/*
func NewRedis() {
	redisClient := &redis.Options{
		Addr:     "localhost:6379",
	    Password: "",
	    DB:       0, 
	}
	return
}
*/

func NewServer() (broker *Broker) {
	// Instantiate a broker
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}

	// Set it running - listening and broadcasting events
	go broker.listen()

	return
}

func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	// Make sure that the writer supports flushing.
	//
	flusher, ok := rw.(http.Flusher)

	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	// Each connection registers its own message channel with the Broker's connections registry
	messageChan := make(chan []byte)

	// Signal the broker that we have a new connection
	broker.newClients <- messageChan

	// Remove this client from the map of connected clients
	// when this handler exits.
	defer func() {
		broker.closingClients <- messageChan
	}()

	// Listen to connection close and un-register messageChan
	notify := rw.(http.CloseNotifier).CloseNotify()

	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	for {

		// Write to the ResponseWriter
		// Server Sent Events compatible
		fmt.Fprintf(rw, "data: %s\n\n", <-messageChan)

		// Flush the data immediatly instead of buffering it for later.
		flusher.Flush()
	}

}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:

			// A new client has connected.
			// Register their message channel
			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))
		case s := <-broker.closingClients:

			// A client has dettached and we want to
			// stop sending them messages.
			delete(broker.clients, s)
			log.Printf("Removed client. %d registered clients", len(broker.clients))
		case event := <-broker.Notifier:

			// We got a new event from the outside!
			// Send event to all connected clients
			for clientMessageChan, _ := range broker.clients {
				clientMessageChan <- event
			}
		}
	}

}

type Stock struct {
  Id int64 `db:"id" json:"id"`
  Firm string `db:"firm" json:"firm"`
  Data string `db:"data" json:"data"`
} 

func main() {

	broker := NewServer()

	/*
	redisClient := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    */

	//log.Println(redisClient)

	//redisClient.Subscribe('data:stock_market')
	//pubsub, err := redis.NewTCPClient(":6379","",-1).PubSubClient()
	redisClient, err := redis.Dial("tcp", ":6379")
	if err != nil {
        log.Println(err)
    }
	defer redisClient.Close()
	//psc := redisClient.PubSubConn{}
	//pubsub, err := redisClient.Subscribe("data:stock_market")

	//redisClient.Send("SUBSCRIBE", "data:stock_market")
	//redisClient.Flush()

	psc := redis.PubSubConn{redisClient}
	psc.Subscribe("data:stock_market")

	go func() {
		for {
			time.Sleep(time.Second * 2)
			type Stocks []Stock
			
			for {
			    switch v := psc.Receive().(type) {
			    case redis.Message:
			        //fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
			        //log.Println(v)
			        stocks := v.Data
			        //result, err := json.Marshal(stocks)
			        if err != nil {
			        	fmt.Println("error:", err)
			        }
			        //result2 := fmt.Sprintf(string(result))
			        log.Println("Receiving event")
			        broker.Notifier <- []byte(stocks)
			    case redis.Subscription:
			        fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		      	}
			}

		}
	}()

	//http.Handle("/", http.FileServer(http.Dir("../clientLayer")))
    //http.ListenAndServe(":3000", nil)

	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", broker))

}
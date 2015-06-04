package main

import (   
    "log"  
    "encoding/json"
    "os"
    
    "github.com/krustnic/runner/amqp"
    "github.com/krustnic/runner/worker"
    "github.com/krustnic/runner/config"
)

var workingQueue chan map[string]interface{}

func init() {
    settings, err := config.LoadRemoteConfig()
    
    if err != nil {
        os.Exit(1)
    }
    
    log.Printf("Remote config is %+v", settings)
}

func main() {
    workingQueue = make( chan map[string]interface{} )
    
    for i := 0; i<config.Config.ThreadsCount; i++ {
        w := worker.CreateWorker( i, workingQueue )
        go w.Run()    
    }    
    
    queue := config.Config.Queue
    
	c, err := amqp.NewConsumer(queue.Uri, queue.Exchange, queue.ExchangeType, queue.Queue, queue.BindingKey, queue.ConsumerTag, dispatcher)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("running forever")
	select {}
	
	log.Printf("shutting down")

	if err := c.Shutdown(); err != nil {
		log.Fatalf("error during shutdown: %s", err)
	}
}

func dispatcher( messageString string ) {
    var message map[string]interface{}
    
    if err := json.Unmarshal( []byte(messageString), &message); err != nil {
        log.Printf("Bad message: '%s'", messageString)
        return
    }
        
    workingQueue <- message
}


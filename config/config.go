package config

import (
    "log"
    "encoding/json"
    
    "github.com/krustnic/runner/api"
)

type RemoteConfig struct {
    // Number of workers
    ThreadsCount int
    
    // RabbitMQ Server Settings
    Queue struct {
        Uri string
        Exchange string
        ExchangeType string
        Queue string
        BindingKey string
        ConsumerTag string
    }
}

var Config *RemoteConfig

func LoadRemoteConfig() (RemoteConfig, error) {
    remoteConfig := RemoteConfig{}
            
    log.Printf("Start requesting remote configuration")
    
    remoteConfigString := api.RequestRemoteConfiguration()
        
    log.Printf("Parsing remote configuration from json to struct")
    
    if err := json.Unmarshal([]byte(remoteConfigString), &remoteConfig); err != nil {
        log.Printf("Error parsing remote settings file from JSON to struct: %s", err)
        return remoteConfig, err
    }
    
    log.Printf("Remote settings successefully unmarshaled to struct: %+v", remoteConfig)    
    
    // Export configuration
    Config = &remoteConfig
    
    return remoteConfig, nil
}
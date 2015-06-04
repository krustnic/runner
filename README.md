# Golang simple task Runner example for RabbitMQ

___Created for personal use only (not common solution).___ What it does:
* Loads local configuration file (config.json)
* Based on local configuration requests remote configuration:
    * sends current runtime information (memory, cpu, network settings collected by [runner-info](https://github.com/krustnic/runtime-info.git))
    * receivse additional settings for current machine (number of workers and RabbitMQ connection information)    
* Runs workers
* Connects to RabbitMQ server
* On each new task (JSON message) pushed to queue run job on available worker

To achieve needed functionality `job` package should be rewriten

## Local configuration file format

```json
{
    "Username"      : "",
    "Password"      : "",
    "ApiHost"       : "",
    "ConfigApiPath" : ""
}
```

## Remote configuration file format
```json
{
    "ThreadsCount" : 3,
    "Queue":{
        "Uri"          : "amqp:\/\/test:test@example.com:5672",
        "Exchange"     : "",
        "ExchangeType" : "direct",
        "Queue"        : "test",
        "BindingKey"   : "",
        "ConsumerTag"  : ""
    }
}
```
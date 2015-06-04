# Golang simple task Runner for RabbitMQ

Created for personal use only (not common solution). What does it do:
* Load local configuration file (config.json)
* Based on local configuration request remote configuration:
    * send current runtime information (memory, cpu, network settings collected by [runner-info](https://github.com/krustnic/runtime-info.git))
    * receive additional settings for current machine (number of workers and RabbitMQ connection information)    
* Run workers
* Connect to RabbitMQ server
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
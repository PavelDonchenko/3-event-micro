1.First microservice will start working on port "9090" inside our localhost
that microservice works at:
> POST Endpoint: /tweet

its body consist of:
> POST body: { creator: String, body: String}
>
Then it pushes received information to a RabbitMQ Queue

2.That microservice subscribes to the queue from RabbitMQ and processes the message. Processing messages means saving the message to Redis.

3.Third microservice will start working on port "9091" inside our localhost:
```
err := router.Run("localhost:9091")
```
that microservice works at:
> GET Endpoint: /message/list

that retrieving an array of objects with the creator, and tweet body, content that was stored in the Redis database in chronologically descending order
Also, that microservice works by passing a parameter to the URL to retrieve the object of a specific creator in chronologically descending order: creator: String
```
router.GET("/tweet/list/:creator", getDataFromRedisDatabaseByParameter)
```
# TODO

- add `ConnectWithAppSecret`
- add `ConnectAsAdmin`
- wrap `EmitWithAck` - ack with body statusCode should capture and emit error. example:
```
6:::2+[{"body":{"error":"Validation Error.","validationErrors":{"to":["to is a missing required parameter."],"message":["message is a missing required parameter."]}},"headers":{"Request-Id":"5bce7c0a-fad9-4ec8-a3cf-74568228c5fd","RateLimit-Limit":10,"RateLimit-Time-Units":"seconds","RateLimit-Remaining":18},"statusCode":400}]
```
- unit tests

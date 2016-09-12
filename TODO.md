# TODO

- implement `ConnectWithAppSecret`
- implement `ConnectAsAdmin`
- wrap socket method `On` on main `respoke.Client`
- wrap `EmitWithAck` - ack with body statusCode should capture and emit error. example:
```
6:::2+[{"body":{"error":"Validation Error.","validationErrors":{"to":["to is a missing required parameter."],"message":["message is a missing required parameter."]}},"headers":{"Request-Id":"5bce7c0a-fad9-4ec8-a3cf-74568228c5fd","RateLimit-Limit":10,"RateLimit-Time-Units":"seconds","RateLimit-Remaining":18},"statusCode":400}]
```
- groups: implement `.Leave()`, `.GetGroupMembers()`
- presence: implement `.SubscribePresence()`, `.GetPresence()`, `.SetPresence()`
- messaging: wrap `pubsub` and `message` respoke events
- implement REST API requests for `apps`
- implememt REST `.GetRoles(appID string)`
- implement `CreateSessionForEndpoint(appID string, roleID string, endpointID string)` (starting the auth dance)
- unit tests

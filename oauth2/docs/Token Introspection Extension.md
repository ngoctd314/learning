# Token Introspection Extension

Resource server needs some way to verify the access token. Require coordination between resource and authorization servers. With small services, both endpoints are part of the same system, and can share token internally. In large systems where the two endpoints are on different servers, this has led to non-standard protocols for communicating between the two servers.

The OAuth 2.0 Token Introspection extension defines a protocol that returns information about an access token, intended to be used by resource servers or other internal servers.

### Introspection Endpoint

Return information about a token

### Token Information Request

It is expected that this endpoint is not made publicly available to developers. Applications should not be allowed to use this endpoint since the response may contain privileged information that developers should not have access to.

```txt
POST /token_info HTTP/1.1
Host: authorization-server.com
Authorization: Basic Y4NmE4MzFhZGFkNzU2

token=cq234jksdf41M1aX4a...
```
### Token Information Response

**active**

Required. Presented token is currently active

**scope**

Optional. A JSON string containing list of scopes associated with this token

**client_id**

Optional. The client identifier for the OAuth 2.0 client

**username**

Optional. A human-readable identifier for the user who authorized this token.

**exp**

Optional. token expire time

```json
{
    "active": true,
    "scope": "read write email",
    "client_id": "J8NFmU4tJVgDxKaJFmXTWvaHO",
    "username": "aaronpk",
    "exp": 1437275311
}
```

### Error Response

Considered a "inactive" token:

- The token requested does not exist or is invalid
- The token expired
- The token was issued to a different client than is making this request

```json
{
    "active": false
}
```

## Security Considerations

Using a token introspection endpoint means that any resource server will be relying on the endpoint to determine whether an access token is currently active or not.

## Token Fishing

## Caching

Never cache the value beyond the expiration time of the token

## Limiting Information
# Design database for authentication

References: https://medium.com/@vietmaihoang/thi%E1%BA%BFt-k%E1%BA%BF-db-cho-vi%E1%BB%87c-qu%E1%BA%A3n-l%C3%BD-authentication-f3d1a6b909f6

## How to store Password


Criteria for store password algorithm:

- With 2 two input => 2 hash
- Difficult to decode 

Some algorithm: SHA-2, SHA-3, MD5, and SHA-1. But MD5 and SHA-1 can violate two above criteria.

How to compare password without store it?

Solution: We compare hash. 

```go
var inputHashed = hashFunc(inputPassword)
if inputHashed == storeHashedPassword {
    return VALID_PASSWORD
}
return INVALID_PASSWORD
```

Problem with brute force attack:

|User|Plain-text Password|Hashed Password|
|-|-|-|
|user1|abc|asdfwerewrERafd|
|user2|abc|asdfwerewrERafd|

Two users with the same password => same hash

Solution: salting

```go
// store password
var salt = saltFunc() // random
var hashedPassword = salt + "_" + hashFunc(salt + plainPassword)
save(user, hashedPassword)

// compare password
var salt = getSaltFromPassword(hashedPassword)
var inputHashed = salt + "_" + hashFunc(salt + inputPassword)
if inputHashed == hashedPassword {
    return VALID_PASSWORD
}
return INVALID_PASSWORD
```

## Email verify and reset password

**Why we need to verify email to complete register process?**

Two reasons:

- Make sure email is provided by user is real email, and user has permission to access into that email.
- Prevent user use email of another people to register.

**How to verify email process work**

System create an active token (unique with each account). This active token will be store in database and then send an link to user's email:

Example:
```txt
https://domain.com/active?code=HbhUPq3i8w90Kdv4QtwiT2cVk3YoLq
```

When user click into link, user is verified.

Active token usually has count down time.

**How to reset password process work**

Reset password allow users update their password in case they forgot password to login into system.

Similar to verify email, a password_reset_token will be generate, store in DB, and then be send to user's email. This token will expired in limited time (example 10 mins to avoid brute force attack).

Example:
```txt
https://domain.com/reset_password?code=b6dcQVNw6REQ8rUs6TyKYtU48plQ9o
```
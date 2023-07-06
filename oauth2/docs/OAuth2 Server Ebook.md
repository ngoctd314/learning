# OAuth 2 servers ebook

Reference: https://www.oauth.com/

## 1. Getting Ready

Things you need to know when building an app that talks to an existing OAuth 2.0 API.

### 1.1. Creating an Application

The registration process typically involves creating a developer account on the service's website, then entering basic information about the application such as the name, website, logo. After registering the application, you'll be given a client_id and a client_secret that you'll use when your app interacts with the service.

Important things: register one or more redirect URLs the application will use. Redirect URLs are where the OAuth 2.0 service will return the user to after they have authorized the application.

### 1.2. Redirect URLs and State

To prevent redirection attacks where an authorization code or access token can be intercepted by an attacker. Some services may allow you to register multiple redirect URLs, which can help when your web app may be running on serveral different subdomains.

Redirect URL must be an https endpoint to prevent the authorization code from being intercepted during the authorization process. If your redirect URL is not https, then an attacker may be able to intercept the authorization code and use it to hijack a session.

Some application may have multiple places they want to start the OAuth process from, such as a login link on a home page as well as a login link when viewing public item. For these applications, it may be tempting to try to register multiple redirect URLs, or you may think you need to be able to vary the redirect URL per request. Instead, OAuth2 2.0 provides a mechnism for this, the "state" parameter.

The "state" parameter can be used to encode application state, but it must also include some amount of random data if you're not also including PKCE parameters in the request. The state parameter is a string that is opaque to the OAuth 2.0 service, so whatever state value you pass in during the initial authorization request will be returned after the user authorizes the application.

## 2. Accessing Data in an OAuth Server

How to access your data at an existing OAuth 2.0 server.

Steps

1. Create an application
2. Setting up the Environment
3. Authorization Request
4. Obtaining an Access Token
5. Making API requests

## 3. Singing in with Google

Despite OAuth being an authorization protocol rather than an authentication protocol, it is often used as the basis for authentication workflows anyway.

Authenticating: who the user is.
Authorizing: users is trying to gain access or modify something that belong to the user.

OAuth was designed as an authorization protocol, so the end result of every OAuth flow is the app obtains an access token in order to be able to access or modify something about the user's account. The access token itself says nothing about who the user is.

Call GET /user-info with access token to get user info 

1. Create an Application

Need to create an app in the Google API console in order to get a **client ID** and **client secret**, and register the **redirect URL**.

2. Setting up the Environment

```go
var (
    googleClientID = "client_id"
    googleClientSecret = "client_secret"
    // authorization gateway
    authorizeURL = "https://accounts.google.com/o/oauth2/v2/auth"
    // Google's OpenID connect token endpoint
    tokenURL = "https://www.googleapis.com/oauth2/v4/token"
    // redirect url
    baseURL = "https://servername.com"
)
```

3. Authorization Request

4. Getting an ID Token

When the user is redirected back to our app, there will be a code and state parameter in the query string. The state parameter will be the same as the one we set in the initial authorization request. This helps protect our app from CSRF attacks.

5. Verifying the User Info

## 4. Server-Side Apps

Server-side apps are the most common type of application encountered when dealing with OAuth servers. These apps run on a web server when the source code of the app is not available to the public, so they can maintain the confidentiality of their client secret.

<div style="text-align:left"><img src="../assets/oauth_server_app.png" width=500 /></div>
The app's server communicates with the API

### 4.1. Authorization Code Grant

The authorization code is a temporary code that the client will exchange for an access token.

```txt
https://authorization-server.com/oauth/authorize?client_id=client_id&response_code=code&state=state_code&redirect_uri=http://redirect.com/auth/redirect&scope=photos
```

You most need to register your redirect URL at the service before it will be accepted. This also means you can't change your redirect URL per request. Instead, you can use the state statement to customize the request.

The user clicks "approve" the server will redirect back to the app, with a "code" and the same "state" parameter you provided in the query string parameter. It is important to note that this is not an access token. The only thing you can do with the authorizatino code is to make a request to get an access token.

### 4.2. Authorizationi Request Parameters

- reponse_type = code
response_type is set to code indicating that you want an authorization code as the response

- client_id
The client_id is the identifier for your app

- redirect_uri (optional)
This must match the redirect URL that you have previously registered with the service

- scope (optional)

Include one or more scope values (space-separated) to request additional levels of access.

- state

The state parameter serves two functions. When the user is redirected back to your app, whatever value you include as the state will also be included in the redirect. Using state parameter as session key. Used to indicate what action in the app to perform after authorization is complete.

The state parameter also serves as a CSRF protection mechnism if it contains a random value per request.

### 4.3. Exchange the authorization code for an access token

To exchange the authorization code for an access token, the app makes a POST request to the service's token endpoint.

- grant_type (required)
The grant_type parameter must be set to "authorization_code"

- code (required)
This parameter is for the authorization code received from the authorization server which will be in the query string parameter "code" in this request.

- redirect_uri (possibly)

- Client Authentication (required)

client_id, client_secret

The service will require the client authentication itself when making the request for an access token. Typically services support client authentication via HTTP Basic Auth with the client's client_id and client_secret. However, some services support authentication by accepting the client_id and client_secret as POST body parameters.

### 4.4. PKCE Verifier

If the service supports PKCE for web server apps, then the client will need to inlucde the followup PKCE parameter when exchanging the authorization code as well.

### 4.5. Example Flow

Using the authorization code flow with PKCE

- 1. Create a log-in link with the app's client ID, redirect URL, state and PKCE code challenge parameters

```html
<a href="https://authorization-server.com/oauth/authorize?response_type=code&client_id=mRkZGFjM&state=5ca75bd30&scope=photos&code_challenge_method=S256&code_challenge=kPaerowervf_wera">Account</a>
```

- 2. The user sees the authorization prompt and approves the request

- 3. The service is redirected back the app's server with an auth code

The service sends a redirect header redirecting the user's browser back to the app that made the request. The redirect will include a code in the URL and the original state

https://example-app.com/cb?code=YasdferNDfew&state=5ca75bd30

- 4. The app exchanges the auth code for an access token

**The app initiates the authorization request**

### 4.6. Possible Errors

### 4.7. User Experience Considerations

In order for the authorization code grant to be secure, the authorization page must appear in a web browser the user is familar with, and must not be embedded in an iframe popup or an embedded browser in a mobile app. If it could be embedded in another website, the user would have no way of verifying it is the legitimate service and is not a phishing attempt.

If an app wants to use the authorization code grant but can't protect its secret (native mobile app or single-page app), then the client secret is not required when making a request to exchange the auth code for an access token, and PKCE must be used as well. However, some services still do not support PKCE, so it may not be possbile to perform and authorization flow from the single-page app itself, and the client-side JS code may need to have a companion sever-side component that performs the OAuth flow instead.

## 5. Single-Page Apps

<div style="text-align:left"><img src="../assets/user_browser_server.png" width=500 /></div>

Single-page apps run entirely in the browser after loading the JS and HTML source code from a web page. Since the entire source is available to the browser, they cannot maintain the confidentiality of a client secret, so a secret is not used for these apps. Because they can't use a client secret, the best option is to use the PKCE extension to protect the authorization code in the redirect. This is similar to the solution for mobile apps which also can't use a client secret.

### 5.1. Authorization

The authorization code is a temporary code that the client will exchange for an access token. After the user visits the authorization page, the service shows the user an explanation of the request, including application name, scope, etc. If the user clicks approve the server will redirect back to the website, with an authorization code and the state value in the URL query string. 

### 5.2. Authorization Grant Parameters

### 5.3. Security Considerations for Single-Page Apps **


### 5.4. Implicit Flow

Some services use the alternative Implicit Flow for single-page apps, rather than allow the app to use the Authorization Code flow with no secret.

The Implicit Flow bypasses the code exchange step, and instead the access token is returned in the query string fragement to the client immediately.

In order for a single-page up to  use the Authorization Code flow, is must be able to make a POST request to the authorization server. This means if the authorization server is on a different domain, the server will need to support the appropriate CORS headers. It supporting CORS headers is not an option, then the service may use the Implicit Flow instead.

In any case, with both the Implicit Flow as well as the Authorization Code Flow with no secret, the server must require registration of the redirect URL in order to maintain the security of the flow.

### 5.5. Security Considerations

The only way the authorization code grant with no client secret can be secure is by using the "state" parameter and restricting the redirect URL to trusted clients. Since the secret is not used, there is no way to verify the identity of the client other than by using a registered redirect URL. This is why you need to pre-register your redirect URL with the OAuth 2.0 service.

### 5.6. Implicit Flow for Single-Page Apps

### 5.7. Security Considerations for Single-Page Apps

**Refresh Tokens**

If the authorization server wishes to allow JavaScript apps to use refresh tokens, then they must also follow the best practices. Specially, refresh tokens must be valid for only one use, and the authorization server must issue a new refresh token each time a new access token is issued in response to a refresh token grant. This provides the authorization server a way to detect if a refresh token has been copied and used by an attacker

**Storing Tokens**

There are currently no general-purpose secure storage mechanism in browsers.

Generally, the browser's LocalStorage API is the best place to store this data as it provides the easiest API to store and retrieve data and is about as secure as you can get in a browser. The downside is that any scripts on the page, even from different domains such as your analytics or ad network, will be able to access the LocalStorage. This means anything you store in LocalStorage is potentially visible to thirf-party scripts on your page.

Because of the risks of data leakage from 3rd scripts, it is extremely important to have a good Content-Security Policy configured for your app so that you can more confident that arbitrary scripts aren't able to run in the application.

## 6. Mobile and Native apps

Like single-page apps, mobiles apps also cannot maintain the confidentially of a client secret. Mobile app must also use an OAuth flow that does not require a client secret. The current best practice is to use the Authorization Flow with PKCE.

### 6.1. Authorization

**Initiate the authorization request**

To begin the authorization process, the app should have a "sign in button. The link should be constructed as a full URL to the service's authorization endpoint.

The client first creates that is known as  a PKCE "code verifier".

Once the app has generated the code verifier, it used that to create the code challenge. The code challenge is a Base64-URL encoded string of the SHA256 hash of the code verifier. This hashed value is sent in the authorization request. So that the original random string is never exposed to anything outside the app. So that the original random string is never exposed to anything outside the app.

**The user approves the request**

**The service redirects the user back to the app**

**Exchange the authorization code for an access token**
To exchange the authorization code for an access token, the app makes a POST request to the service's token endpoint.

|Parameter|Description|
|-|-|
|grant_type(required)|must be set to authorization_code|
|code(required)|authorization code received from the authorization server|
|redirect_uri(possibly required)||
|code_verifier(required)||
|client_id(required)||

### 6.2. Security Considerations

Always use the secure embedded browser APIs or launch a native browser

## 7. Making Authenticated Requests

Regardless of which grant type you used or whether you used a client secret, you now have an OAuth 2.0 Bearer Token you can use with the API.

The access token is sent to the service in the HTTP Authorization header prefixed by the text Bearer.

```txt
POST /resource/1/update HTTP/1.1
Authorization: Bearer RsT5OjbzRn430zqMLgV3Ia
```
Your app only use it to make API requests.

The access token is not indended to be parsed or understood by your application. The only thing your application should do with it is used it to make API requests. Some services will use structured tokens like JWTs as their access tokens, but the client does not need to worry about decoding the token in this case.

Server makes no guarantees that access tokens will always continue to be in the same format. It's entirely possible that the next time you get an access token from the service, it will be in a different format. Access tokens are opaque to the client, and should only be used to make API requests and not interpreted themselves.

If you are trying to find out whether your access token has expired, you can either store the expiration lifetime that was returned when you first got the access token, or just try to make the request anyway, and get a new access token if the current one has expired. While preemptively refreshing the access token can save an HTTP request, you still need to handle the case when an API call reports an expired token you were expecting it to expire, since the access tokens can expire for many reasons beyond just their expected lifetime.

### 7.1. Refresh Tokens 

```json
{
    "access_token": "AYjcyMzY3ZDhiNmJkNTY",
    "refresh_token": "RjY2NjM5NzA2OWJjuE7c",
    "token_type": "Bearer",
    "expires": 3600 // second
}
```
You can use expires time to preemptively refresh your access tokens instead of waiting for a request with an expired token to fail. Some people like to get a new access token shortly before the current one will expire in order to save an HTTP request of an API call falling. Access tokens can expire for many reasons, such as the user revoking an app, or if the authorization server expires all tokens when a user changes their password.

Access token request

```json
{
    "grant_type": "refresh_token",
    "refresh_token": "",
    "client_id": "",
    "client_secret": ""
}
```

Access token response

```json
{
    "access_token": "AYjcyMzY3ZDhiNmJkNTY",
    "refresh_token": "RjY2NjM5NzA2OWJjuE7c",
    "token_type": "Bearer",
    "expires": 3600 // second
}
```

If you do not get back a new refresh token, then it means your existing refresh token will continue to work when the new access token expires.

**When the refresh token changes after each use, if the authorization server ever detects a refresh token was used twice, it means it has likely been copied and is being used by an attacker, and the authorization server can revoke all access tokens and refresh tokens associated with it immediately.**

Keep in mind that at any point the user can revoke an application, so your application needs to be able to handle the case when using the refresh token also fails. At that point, you will need to prompt the user for authorization again, beginning a new OAuth flow from scratch.

expires_in property refers to the access token, not the refresh token. The expiration time of the refresh token is intentionally never communicated to the client.

If a refresh token expires for any reason, then the only action the application can take is to ask the user to log in again, starting a new OAuth flow from scratch, which will issue a new access token and refresh token to the application.

## 8. Registering a New Application

When a developer comes to your website, they will need a way to create a new application and obtain credentials. 

### 8.1. The Client ID and Secret

The client_id is a public identifier for apps. If the clientID is guessable, it makes slightly easier to craft phishing attacks against arbitrary applications.

If the developer is creating a "public" app (a mobile or single-page app), then you should not issue a client_secret to the app at all. **If it doesn't exist, it can't be leaked.**

The client_secret is a secret known only to the application and the authorization server. It is essential the application's own password. It must be sufficiently random to not be guessable.

### 8.2. Storing and Displaying the client ID and secret

client_id and client_secret equivalent to a username and password.

### 8.3. Deleting Applications and Revoking Secrets

## 9. Authorization

Clients may use either the authorization code grant type or the implicit grant. Along with the type of grant specified by the response_type parameter, the request will use a number of other parameters to indicate the specifics of the request.

### 9.1. The Authorization Request

### 9.2. Requiring User Login

### 9.3. The Authorization Interface

### 9.4. The Authorization Response

### 9.5. Security Considerations  

**1. Phishing Attacks**

**2. Clickjacking**

In a clickjacking attack, the attacker creates a malicious website in which it loads the authorization server URL in a transparent iframe above the attacker's web page. The attacker's web page is stacked below the iframe, and has some innocuous-looking buttons or links, placed very carefully to be directly under the authorization server's confirmation button. When the user clicks the misleading visible button, they are actually clicking the invisible button on the authorization page, there by granting access to the attacker's application. This allows the attacker to trick the user into granting access without their knowledge.

**3. Redirect URL Manipulation**

An attacker can construct an authorization URL using a client ID that belongs to a known good application, but set the redirect URL to a URL under the control of the attacker. If the authorization server does not validate redirect URLs, and the attacker uses the token response type, the user will be returned to the attacker's application with the access token in the URL. If the client is a public client, and the attacker intercepts the authorization code, then the attacker can also exchange the code for an access token.

## 10. Scope

Scope is a way to limit an app's access to a user's data. A way to request a more limited scope of what they are allowed to do on behalf of a user. Scope is a way to control access and help the user identify the permissions they are granting to the application.

Scope is a way to limit what an application can do within the context of what a user can do. For example, if you have a user in the customer group and the application is requesting the admin scope, the OAuth server is not going to create an access token with the "admin" scope, because that user is not allowed to use that scope themselves.

Scope should be thought of as the application requesting permission from the user who's using th app.

### 10.1. Defining Scopes

The challenge when defining scopes for you service is to not get carried away with defining too many scopes. Users need to be able to understand what level of access they are granting to the application.

**Read vs. Write**

Read access to a user's private profile information is treated with separate access control from apps wanting to update the profile information.

## 11. Redirect URLs

Redirect URLs are a critical part of the OAuth flow. Because the redirect URL will contain sensitive information, it is critical that the service doesn't redirect the user to arbitrary locations.

### 11.1. Redirect URL Registration

In order to avoid exposing users to open redirector attacks, you must require developers register one or more redirect URL for the application. The authorization server must never redirect to any other location. 

If an attacker can manipulate the redirect URL before the user reaches the authorization server, they could cause the server to redirect the user to a malicious server which would send the authorization code to the attacker. This is one way attackers can try to intercept an OAuth exchange and steal access tokens. If the authorization endpoint does not limit the URLs that it will redirect to, then it's considered an "open redirector" and can be use in combination with other things to launch attacks.

Registered redirect URLs may contain query string parameters, but must not contain anything in the fragment. The registration server should reject the request if the developer tries to register a redirect URL that contains a fragment.

### 11.2. Redirect URLs for Native Apps

For example, an iOS application may register a custom protocol such as myapp:// and the use a redirect_uri of myapp://callback

**App-Claimed hgttps URL Redirection**

Some platforms, (Android, and iOS as of iOS 9), allow the app to override specific URL patterns to launch the native application instead of a web browser. For example, an application could register https://app.example.com/auth and whenever the web browser attempts to reidrect to that URL, the operating system launches the native app instead.

### 11.3. Redirect URL Validation

There are three cases when you'll need to validate redirect URLs

- When the developer registers the redirect URL as part of creating an application
- In the authorization request (both authorization code and implicit grant types)
- When the application exchanges an authorization code for an access token

## 12. Access Token

Access tokens are the thing that applications use to make API requests on behalf of a user. The access token represents the authorization of a specific application to access specific parts of a user's data.

With client, the access token is an opaque string, and it will take whatever the string is and use it in HTTP request. The resource server will need to understand what the access token means and how to validate it.

### 12.1. Authorization Code Request

The authorization code grant is used when an application exchanges an authorization code for an access token.

## 13. Listing Authorizations

Once uses have begun to authorize multiple applications, giving many apps acces to their account, it becomes necessary to provide a way to allow the user to manage the apps that have access. 

## 14. The Resource Server

The resource server is the OAuth 2.0 term for your API server. The resource server handles authenticated requested after the application has obtained an access token.

Large scale developments may have more than one resource server. Each of these resource servers are distinctly separate but they all share the same authorization server.

### 14.1. Verifying Access Tokens

The resource server will be getting requests from applications with an HTTP Authorization header containing an access token. The resource server needs to verify the access token.

If you're using self-encoded access tokens, then verifying the tokens can be done entirely.

If your tokens are stored in a database, then verifying the token is simply a database lookup on the token table.

Another option is to use the Token Introspection spec to build an API to verify access tokens (use can encapsulate all of the logic of access tokens in a single server, exposing the information via an API to other parts of the system). The token instropection endpoint is intended to be used only internally, so you will want to protect it with some internal authorization, or only enable it on a server with the firewall of the system.

### 14.2. Verify Scope

The resource server needs to know the list of scopes that are associated with the access token. The server is responsible for denying the request if the scopes in the access token do not include the required scope to perform the designed action.

### 14.3. Expired Tokens

If your service uses short-lived access tokens with long-lived refresh tokens, then you'll need to make sure to return the proper error response when an application makes a request with an expired token.

```json
HTTP/1.1 401 Unauthorized
WWW-Authenticate: Bearer error="invalid_token"
                  error_description="The access token expired"

Content-type: application/json

{
    "error": "invalid_token",
    "error_description": "The access token expired"
}
```
Thống kê lượng sms theo client_id, template_id (slow query)


## 15. OAuth for Native Apps

Browser-based apps, native apps can't use a client secret, as that would require that the developer ship the secret in their binary distribution of the application. It has been proven to be relatively easy to decompile and extract the secret. As such, native apps must use and OAuth flow that does not require a preregistered client secret.

The current industry best practice is to use the Authorization Flow along with the PKCE extension, omitting the client secret from the request, and to use an external user agent to complete the flow.

## 18. Token Introspection Endpoint

When an OAuth 2.0 client makes a request to the resource server, the resource server needs some way to verify the access token.

The OAuth2 Token Introspection extension defines a protocol that returns information about an access token, intended to be used by resource servers or other internal servers.

An alternative to token introspection is to use a structured token format that is recognized by both the authorization server and resource server. The JWT Profile for OAuth 2.0 Access Token is a recent RFC that describes a standardized format for access tokens using JWTs. This enables a resource server to validate access tokens without a newwork call, by validating the signature and parsing the claims within the structured token itself.

### 18.1. Introspection Endpoint

### 18.2. Token Information Request

The request will be a POST request containing just a parameter named "token". It is expected that this endpoint is not made publicly available to developers. Applications should not be allowed to use this endpoint since the response may contain privileged information that developers should not have access to. One way to protect endpoint is to put it on an internal server that is not accessible from the outside world, or it could be protected with HTTP basic auth.

```sh
POST /token_info HTTP/1.1
Host: authorization-server.com
Authorization: Basic Y4NmE4MzFhZGFkNzU2YWRhN

token=c1MGYwNDJiYmYxNDFkZjVkOGI0MSAgLQ
```

### 18.3. Token Information Response 

```txt
active
    Requried.
scope
client_id
username
exp
```

## 20. Terminology Reference

Roles

## 22. OpenID Connect

The OAuth 2.0 framework explicitly does not provide and information about the user that has authorized an application. OAuth 2.0 is delegation framework, allowing third-party applications to act on behalf of a user, without the application needing to know the identify of the user.

OIDC takes the OAuth 2.0 framework and adds an identity layer on top. It provides information about the user, as well as anables clients to establish login sessions. While this chater is not meant to be a complete guide to OpenID Connect.

**1. Authorization vs Authentication**

OAuth 2.0 is called an authorization "framework" rather than a "protocol".  OAuth 2.0 does not provide a mechanism to say who a user is or how they authenticated, it just says that a user delegated and application to act on their behalf. The OAuth 2.0 framework provides this delegation in the form of the user.

When you check in to a hotel, you get a key card which you can use to enter your assigned room. You can think of the key card as an access token. The key card says nothing about who you are or how you were authenticated at the front desk, but you can use th card to access your hottel room for the duration of your stay.

**2. Building an Authentication Framework**

It is quite possible to use the OAuth 2.0 framework as the basic for building an authentication and identity protocol.

To use OAuth 2.0 as the basis of an authenticatino protocol, you will need to do at least a few things.

- Define an endpoint to return attributes about a user

```sh
GET /api/v1/me
```

- Define one or more scopes that the third-party applications can use to request identity information from the user

```sh
GET /api/v1/scopes
```

**3. ID Tokens**

The core of OpenID Connect is based on a concept called "ID Tokens." This is a new token that the authorization server will return which encodes the user's authentication information. In contrast to access tokens, which are only intended to be understood by the resource server, ID tokens are intended to be understood by the OAuth client. When the client makes an OpenID Connect request, it can request an ID token along with an access token.

OpenID Connect's ID Tokens take the form of a JWT (JSON Web Token), which is a JSON payload that is signed with the private key of the issuer, and can be parsed and verified by the application.

Inside the JWT are a handful or defined property names that provide information to the application.

- sub short for subject (a unique identifier for the user) 
- iss the identifier for the server that issued the token (issue)
- aud the identifier for the client that requested this token (audience)

```json
{
    "iss": "https://server.example.com",
    "sub": "24400320",
    "aud": "s6BhdRkqt3",
    "nonce": "n-0S6_WzA2Mj",
    "exp": 1311281970,
    "iat": 1311280970,
    "auth_time": 1311280969
}
```

## IndieAuth

IndieAuth is decentralized identity protocol built on OAuth 2.0, which works using URLs to identify users and applications. It allows people to use a domain under their control as their identity while signing in and authorizing applications using that identity.

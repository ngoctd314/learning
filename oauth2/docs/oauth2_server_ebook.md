# OAuth 2 servers ebook

References: https://www.oauth.com/

## Getting Ready

Things you need to know when building an app that talks to an existing OAuth 2.0 API.

### Creating an Application

- given a client_id (client_secret in some cases)

Important things: register one or more redirect URLs the application will use. Redirect URLs are where the OAuth 2.0 service will return the user to after they have authorized the application.

### Redirect URLs and State

To prevent redirection attacks where an authorization code or access token can be intercepted by an attacker. Some services may allow you to register multiple redirect URLs, which can help when your web app may be running on serveral different subdomains.

Redirect URL must be an https endpoint to prevent the authorization code from ebing intercepted during the authorization process. 

Some apps may have multiple places they want to start the OAuth process from, such as a login link on a home page as well as a login link when viewing some public item. Try to register multiple redirect URls. OAuth 2.0 provides a mechanism for this, the "state" parameter.

The "state" parameter can be used to encode application state. The state parameter is a string that is opaque to the OAuth 2.0 service

## Accessing Data in an OAuth Server

Steps

1. Create an application
2. Setting up the Environment
3. Authorization Request
4. Obtaining an Access Token
5. Making API requests

## Singing in with Google

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

## Server-Side Apps

Server-side apps are the most common type of application encountered when dealing with OAuth servers. These apps run on a web server when the source code of the app is not available to the public, so they can maintain the confidentiality of their client secret.

### Authorization Code Grant

The authorization code is a temporary code that the client will exchange for an access token.

```txt
https://authorization-server.com/oauth/authorize?client_id=client_id&response_code=code&state=state_code&redirect_uri=http://redirect.com/auth/redirect&scope=photos
```

You most need to register your redirect URL at the service before it will be accepted. This also means you can't change your redirect URL per request. Instead, you can use the state statement to customize the request.

The user clicks "approve" the server will redirect back to the app, with a "code" and the same "state" parameter you provided in the query string parameter. It is important to note that this is not an access token. The only thing you can do with the authorizatino code is to make a request to get an access token.

### Authorizationi Request Parameters

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

### Exchange the authorization code for an access token

To exchange the authorization code for an access token, the app makes a POST request to the service's token endpoint.

- grant_type (required)
The grant_type parameter must be set to "authorization_code"

- code (required)
This parameter is for the authorization code received from the authorization server which will be in the query string parameter "code" in this request.

- redirect_uri (possibly)

- Client Authentication (required)

client_id, client_secret

### PKCE Verifier

If the service supports PKCE for web server apps, then the client will need to inlucde the followup PKCE parameter when exchanging the authorization code as well.

### Example Flow

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

### Possible Errors

### User Experience Considerations

## Single-Page Apps
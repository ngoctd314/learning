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

Single-page apps ruun entirely in the browser after loading the JS and HTMl source code from a web page. Since the entire source is available to the browser, they cannot maintain the confidentiality of a client secret, so a secret is not used for these apps. Because they can't use a client secret, the best option is to use the PKCE extension to protect the authorization code in the redirect. This is similar to the solution for mobile apps which also can't use a client secret.

**Deprecation Notice**

A common historical pattern for single-page apps was to use the Implicit flow to receive an access token in the redirect without the intermediate authorization code exchange step. This has a number of security issues and should no longer be used.

### Implicit Flow

Some services use the alternative Implicit Flow for single-page apps, rather than allow the app to use the Authorization Code flow with no secret.

The Implicit Flow bypasses the code exchange step, and instead the access token is returned in the query string fragement to the client immediately.

In order for a single-page up to  use the Authorization Code flow, is must be able to make a POST request to the authorization server. This means if the authorization server is on a different domain, the server will need to support the appropriate CORS headers. It supporting CORS headers is not an option, then the service may use the Implicit Flow instead.

In any case, with both the Implicit Flow as well as the Authorization Code Flow with no secret, the server must require registration of the redirect URL in order to maintain the security of the flow.

### Security Considerations

The only way the authorization code grant with no client secret can be secure is by using the "state" parameter and restricting the redirect URL to trusted clients. Since the secret is not used, there is no way to verify the identity of the client other than by using a registered redirect URL. This is why you need to pre-register your redirect URL with the OAuth 2.0 service.

### Implicit Flow for Single-Page Apps

### Security Considerations for Single-Page Apps

**Refresh Tokens**

If the authorization server wishes to allow JavaScript apps to use refresh tokens, then they must also follow the best practices. Specially, refresh tokens must be valid for only one use, and the authorization server must issue a new refresh token each time a new access token is issued in response to a refresh token grant. This provides the authorization server a way to detect if a refresh token has been copied and used by an attacker

**Storing Tokens**

There are currently no general-purpose secure storage mechanism in browsers.

Generally, the browser's LocalStorage API is the best place to store this data as it provides the easiest API to store and retrieve data and is about as secure as you can get in a browser. The downside is that any scripts on the page, even from different domains such as your analytics or ad network, will be able to access the LocalStorage. This means anything you store in LocalStorage is potentially visible to thirf-party scripts on your page.

## Mobile and Native apps

Like single-page apps, mobiles apps also cannot maintain the confidentially of a client secret. Mobile app must also use an OAuth flow that does not require a client secret. The current best practice is to use the Authorization Flow with PKCE.

## Making Authenticated Requests

Regardless of which grant type you used or whether you used a client secret, you now have an OAuth 2.0 Bearer Token you can use with the API.

The access token is sent to the service in the HTTP Authorization header prefixed by the text Bearer.

```txt
POST /resource/1/update HTTP/1.1
Authorization: Bearer RsT5OjbzRn430zqMLgV3Ia
```

Your app only use it to make API requests.

Server makes no guarantees that access tokens will always continue to be in the same format. It's entirely possible that the next time you get an access token from the service, it will be in a different format. Access tokens are opaque to the client, and should only be used to make API requests and not interpreted themselves.

## Refresh Tokens 

## Registering a New Application

When a developer comes to your website, they will need a way to create a new application and obtain credentials.

### The Client ID and Secret

The client_id is a public identifier for apps. If the clientID is guessable, it makes slightly easier to craft phishing attacks against arbitrary applications.

If the developer is creating a "public" app (a mobile or single-page app), then you should not issue a client_secret to the app at all. If it doesn't exist, it can't be leaked.

The client_secret is a secret known only to the application and the authorization server. It is essential the application's own password. It must be sufficiently random to not be guessable.

### Storing and Displaying the client ID and secret

client_id and client_secret equivalent to a username and password.

### Deleting Applications and Revoking Secrets

## Authorization

### The Authorization Request

### Requiring User Login

### The Authorization Interface

### The Authorization Response

Once the user has finished logging in and approving the request, the authorization server is ready to redirect the user back to the application. 

### Security Considerations

**1. Phishing Attacks**

**2. Clickjacking**

In a clickjacking attack, the attacker creates a malicious website in which it loads the authorization server URL in a transparent iframe above the attacker's web page. The attacker's web page is stacked below the iframe, and has some innocuous-looking buttons or links, placed very carefully to be directly under the authorization server's confirmation button. When the user clicks the misleading visible button, they are actually clicking the invisible button on the authorization page, there by granting access to the attacker's application. This allows the attacker to trick the user into granting access without their knowledge.

**3. Redirect URL Manipulation**

An attacker can construct an authorization URL using a client ID that belongs to a known good application, but set the redirect URL to a URL under the control of the attacker. If the authorization server does not validate redirect URLs, and the attacker uses the token response type, the user will be returned to the attacker's application with the access token in the URL. If the client is a public client, and the attacker intercepts the authorization code, then the attacker can also exchange the code for an access token.

## Scope

Scope is a way to limit an app's access to a user's data. A way to request a more limited scope of what they are allowed to do on behalf of a user. Scope is a way to control access and help the user identify the permissions they are granting to the application.

### Defining Scopes

The challenge when defining scopes for you service is to not get carried away with defining too many scopes. Users need to be able to understand what level of access they are granting to the application.

**Read vs. Write**

Read access to a user's private profile information is treated with separate access control from apps wanting to update the profile information.

## Redirect URLs

Redirect URLs are a critical part of the OAuth flow. Because the redirect URL will contain sensitive information, it is critical that the service doesn't redirect the user to arbitrary locations.

### Redirect URL Registration
### Redirect URLs for Native Apps

### Redirect URL Validation

## Access Token

Access tokens are the thing that applications use to make API requests on behalf of a user.

With client, the access token is an opaque string, and it will take whatever the string is and use it in HTTP request. The resource server will need to understand what the access token means and how to validate it.

### Authorization Code Request

## Token Introspection Endpoint

## 13. Listing Authorizations

Once uses have begun to authorize multiple applications, giving many apps acces to their account, it becomes necessary to provide a way to allow the user to manage the apps that have access. 

## 14. The Resource Server

The resource server is the OAuth 2.0 term for your API server. The resource server handles authenticated requested after the application has obtained an access token.
Large scale developments may have more than one resource server. Each of these resource servers are distinctly separate but they all share the same authorization server.

**1. Verifying Access Tokens**

The resource server will be getting requests from applications with an HTTP Authorization header containing an access token. The resource server needs to verify the access token. If your tokens are stored in a database, then verifying the token is simply a database lookup on the token table.

Another option is to use the Token Introspection spec to build an API to verify access tokens (use can encapsulate all of the logic of access tokens in a single server, exposing the information via an API to other parts of the system). The token instropection endpoint is intended to be used only internally, so you will want to protect it with some internal authorization, or only enable it on a server with the firewall of the system.

**2. Verify Scope**

The resource server needs to know the list of scopes that are associated with the access token.

**3. Expired Tokens**

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

**4. OAuth for Native Apps**

Browser-based apps, native apps can't use a client secret, as that would require that the developer ship the secret in their binary distribution of the application. It has been proven to be relatively easy to decompile and extract the secret. As such, native apps must use and OAuth flow that does not require a preregistered client secret.

The current industry best practice is to use the Authorization Flow along with the PKCE extension, omitting the client secret from the request, and to use an external user agent to complete the flow.

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

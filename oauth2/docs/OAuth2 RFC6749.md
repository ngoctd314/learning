# The OAuth 2.0 Authorization Framework

Reference: https://datatracker.ietf.org/doc/html/rfc6749

OAuth 2.0 enables a third-party application to obtain limited access to an HTTP service.

## 1. Introduction

The traditional client-server authentication model

<div style="text-align:left"><img src="../assets/traditional_flow.png" width=500 /></div>

**This creates several problems and limitations:**

- Third-party applications are required to store the resource owner's credentials for future use, typically a password in clear-text.
- Servers are required to support password authentication, despite the security weakness inherent in passwords.
- Third-party applications gain overly broad access to the resource ownner's protected resources, leaving resource owners without any ability to restrict duration or access to a limited subset of resources.
- Resource owners cannot revoke access to an individual third party without revoking access to all third parties, and must do so by chaning the third party's password.

### 1.1. Roles

OAuth2 defines four roles:

| Roles                | Description                                                                                                                        |
| -------------------- | ---------------------------------------------------------------------------------------------------------------------------------- |
| Resource owner       | An entity capable of granting access to a protected resource. When resource owner is a person, it is referred to as an end-user    |
| Resource server      | The server hosting the protected resources, capable of accepting and responding to protected resource requests using access tokens |
| Client               | An application making protected resource requests on behalf of the resource owner and with its authorization.                      |
| Authorization server | The server issuing access tokens to the client after successfully authenticating the resource owner and obtaining authorization    |

The authorization server may be the same server as the resource server or a separate entity. A single authorization server may issue access tokens accepted by multiple resource servers.

### 1.2. Protocol Flow

<div style="text-align:left"><img src="../assets/protocol_flow.png" width=500 /></div>

| Step | Action                                                                                                                 |
| ---- | ---------------------------------------------------------------------------------------------------------------------- |
| 1    | Client requests authorization from the resource owner.                                                                 |
| 2    | Client receives an authorization grant, a credential represeting the resource owner's authorization.                   |
| 3    | Client requests an access token by authenticating with the authorization server and presenting the authorization grant |
| 4    | Authorization server authenticates the client and validates the authorization grant, issues an access token            |
| 5    | Client requests the protected resource from the resource server and authenticates by presenting the access token       |
| 6    | The resource server validates the access token, and serves the request                                                 |

### 1.3. Authorization Grant

An authorization grant is a credential representing the resource owner's authorization (to access its protected resources) used by the client to obtain an access token.

<div style="text-align:left"><img src="../assets/authorization_grant.png" width=500 /></div>

#### 1.3.1. Authorization Code

Using an authorization server as an intermediary between the client and resource owner, the client directs the resource owner to an authorization server, which in turn directs the resource owner back to the client with the authorization code. Before directing the resource owner back to the client with the authorization code, the authorization server authenticates the resource owner and obtains authorization. Because the resource owner only authenticates with the authorization server, the resource owner's credentials are never shared with the client.

#### 1.3.2. Implicit

#### 1.3.3. Resource Owner Password Credentials

The resource owner password credentials (username and password) can be used directly as an authorization grant to obtain an access token. The credentials should only be used when there is a high degree of trust between the resource owner and the client (the client is part of the device operating system or a highly privileged application), and when other authorization grant types are not available (such as an authorization code).

Even though this grant type requires direct client access to the resource owner credentials, the resource owner credentials are used for a single request and are exchanged for an access token. This grant type can eliminate the need for the client to store the resource owner credentials for future use, by exchangine the credentails with a long-lived access token or refresh token.

#### 1.3.4. Client Credentials

### 1.4. Access Token

Access tokens are credentials used to access protected resources.

<div style="text-align:left"><img src="../assets/access_token.png" width=500 /></div>

### 1.5. Refresh Token

Refresh tokens are credentials used to obtain access tokens. Refresh tokens are issued to the client by the authorization server and are used to obtain a new access token when the current access token becomes invalid or expires.

<div style="text-align:left"><img src="../assets/refresh_token.png" width=500 /></div>

### 1.6. TLS Version

### 1.7. HTTP Redirections

The client or the authorization server directs the resource owner's user-agent to another destination.

## 2. Client Registration

When registering a client, the client developer SHALL:

- Specify the client type
- Provide its client redirection URIs
- Include any other information required by the authorizatino server (app name, website, description...)

### 2.1. Client Types

Based on their ability to authenticate securely with the authorization server (i.e ability to maintain the confidentiality of their client credentials):

<div style="text-align:left"><img src="../assets/client_type.png" width=500 /></div>

The client type designation is based on the authorization server's definition of secure authentication and its acceptable exposure levels of client credentials. The authorization server SHOULD NOT make assumptions about the client type.

**Web application**

<div style="text-align:left"><img src="../assets/web_app_oauth2.png" width=500 /></div>

**User-agent-based application**

<div style="text-align:left"><img src="../assets/user_agent_based_app.png" width=500 /></div>

**Native application**

Is a public client installed and executed on the device used by the resource owner. Protocol data and credentials are accessible to the resource owner. It is assumed that any client authentication credentials included in the application can be extracted.

### 2.2. Client Identifier

<div style="text-align:left"><img src="../assets/client_identifier.png" width=500 /></div>

### 2.3. Client Authentication

If the client type is confidential, the client and authorization server establish a client authentication method suitable for the security requirements of the authorization server.

Confidential clients are typically issued (or establish) a set of client credentials used for authenticating with the authorization server (password, public/private key pair).

#### 2.3.1. Client Password \*\*

Clients in possession of a client password MAY use the HTTP basic authentication scheme as defined in to authenticate with the authorization scheme to authenticate with the authorization server. The client identifier is encoded using the "application/x-www-form-urlencoded" encoding algorithm. The authorization server MUST support the HTTP Basic authentication scheme for authenticating clients that were is issued a client password.

Authorization: Basic czZCaGRSa3F0Mzo3RmpmcDBaQnIxS3REUmJuZlZkbUl3

The authorization server MAY support including the client credentials in the request-body

```txt
client_id: REQUIRED. The client identifier issued to the client during the registration process.
client_secret: REQUIRED. The client secret. The client MAY omit the parameter if the client secret is an empty string.
```

#### 2.3.2. Other Authentication Methods

### 2.4. Unregistered Clients

## 3. Protocol Endpoints

The authorization process has two authorization server endpoints (HTTP resources):

- Authorization endpoint - used by the client to obtain authorization from the resource owner via user-agent redirection.
- Token endpoint - used by the client to exchange an authorization grant for an access token, typically with client authorization.
- Redirection endpoint - used by the authorization server to return responses containing authorization credentials to the client via the resource owner user-agent.

### 3.1. Authorization Endpoint

The authorization endpoint is used to interact with the resource owner and obtain and authorization grant. The authorization server MUST first verify the identity of the resource owner. The way in which the authorization server authenticates the resource owner (e.g username and password login, session cookies) is beyond the scope of this specification.

#### 3.1.1. Response Type

The authorization endpoint is used by the authorization code grant type and implicit grant type flows.

```txt
response_type
    REQUIRED. The value MUST be one of "code" for requesting an authorization code,
    "token" for requesting an access token (implicit grant)
```

If an authorization request is missing the "response_type" parameter, or if the response type is not understood, the authorization server MUST return an error response.

#### 3.1.2. Redirection Endpoint

After completing its interaction with the resource owner, the authorization server directs the resouce owner's user-agent back to the client.

##### 3.1.2.1. Endpoint Request Confidentiality

The redirection request will result in the transmission of sensitive credentials over an open network. Lack of transport-layer security can have a severe impact on the security of the client and the protected resources it is authorized to access.

##### 3.1.2.2. Registration Requirements

The authorization server MUST require the following clients to register their redirection endpoint:

- Public clients
- Confidential clients utilizing the implicit grant type

The authorization server SHOULD requrie all clients to register their redirection endpoint prior to utilizing the authorization endpoint.

The authorization server MAY allow the client to register multiple redirection endpoints

Lack of a redirection URI registration requirement can enable an attacker to use th authorization endpoint as an open redirector

##### 3.1.2.3. Dynamic Configuration

If multiple redirection URIs have been registered, the client MUST include a redirection URI with the authorization request using the "redirect_uri" request parameter.

When a redirection URI is included in an authorization request, the authorization server MUST compare and match the value received against at lest one of the registered redirection URIs.

##### 3.1.2.4. Invalid Endpoint

Inform the resource owner of the error

##### 3.1.2.5. Endpoint Content

The redirection request to the client's endpoint typically results in an HTML document response, processed by the user-agent.

### 3.2. Token Endpoint

The token endpoint is used by the client to obtain an access token by presenting its authorization grant or refresh token. The token endpoint is used with every authorization grant except for the implicit grant type (since an access token is issued directly).

Since requests to the token endpoint result in the transmission of clear-text credentials (in the HTTP request and response), the authorization server MUST request the use of TLS.

#### 3.2.1. Client Authentication \*\*

### 3.3. Access Token Scope

The authorization and token endpoints allow the client to specify the scope of the access request using the "scope" request parameter.

The value of the scope parameter is expressed as a list of space-delimited, case-sensitive strings. The strings are defined by the authorization server.

The authorization server MAY fully or partitionally ignore the scope requested by the client, based on the authorization server policy or the resource owner's instructions. If the issued access token scope is different from the one requested by the client, the authorization server MUST include the "scope" response parameter to inform the client of the actual scope granted.

If the client omits the scope parameter when requesting authorization, the authorization server MUST either process the request using a pre-defined default value or fail the request indicating an invalid scope.

## 4. Obtaining Authorization

### 4.1. Authorization Code Grant

<div style="text-align:left"><img src="../assets/authorization_code_grant.png" width=500 /></div>

- (A) The client initiates the flow by directing the resource owner's user-agent to the authorization endpoint. The client includes its client identifier, requested scope, local state, and a redirection URI to which the authorization server will send the user-agent back once access is granted (or denied).

- (B) The authorization server authenticates the resource owner (via the user-agent) and establishes whether the resource owner grants or denies the client's access request

- (C) Assuming the resource onwer grants access, the authorization server redirects the user-agent back to the client using the redirection URI provided earilier (in the request or during the client registration). The redirection URI includes an authorization code and any local state provided by the client earlier.

- (D) The client requests an access token from the authorization server's token endpoint by including the authorization code received in the previous step.

- (E) The authorization server authenticates the client, validates the authorization code, and ensures that the redirection URI received matches the URI used to redirect the client in step (C). If valid, the authorization server responds back with an access token and, optionally, a refresh token.

#### 4.1.1. Authorization Request

The client constructs the request URI by adding the following params, using the "application/x-www-form-urlencoded" format

```txt
response_type
    REQUIRED. Value MUST be set to "code"
client_id
    REQUIRED. The client identifier
redirect_uri
    OPTIONAL.
scope
    OPTIONAL
state
    RECOMMENDED. An opaque value used by the client to maintain state between the request and
    callback. The authorization server includes this value when redirecting the user-agent
    back to the client. The parameter SHOULD be used for preventing CSRF
```

#### 4.1.2. Authorization Response

If the resource owner grants the access request, the authorization server issues an authorization code and delivers it to the client by adding the following parameters to the query component of the redirection URI using the "application/x-www-form-urlencoded" format

```txt
code
    REQUIRED. The authorization code generated by the authorization server. The authorization code MUST expire shortly after it issued to mitigate the risk of leaks.

state
    REQUIRE if the "state" parameter was present in the client authorization request. The exact value received from the client.
```

##### 4.1.2.1. Error Response \*\*

#### 4.1.3. Access Token Request

The client makes a request to the token endpoint by sending the following parameters using the "application/x-www-form-urlencoded" format with a character encoding of UTF-8 in the HTTP request entity-body

```txt
grant_type
    REQUIRED. Value MUST be set to "authorization_code"

code
    REQUIRED. The authorization code received from the authorization server

redirect_uri
    REQUIRED if the "redirect_uri" was included in the authorization request

client_id
    REQUIRED, if the client is not authenticating with the authorization server
```

The authorization server must

- Require client authentication for confidential clients or for any client that was issued client credentails
- Authenticate the client if the authentication is included
- ...

#### 4.1.4. Access Token Response

If the access token is valid and authorized, the authorization server issues an access token and optional refresh token. If the request client authentication failed or is invalid, the authorization server returns an error response.

### 4.2. Implicit Grant \*\*

### 4.3. Resource Owner Password Credentials Grant

The resource owner password credentials grant type is suitable in cases where the resource owner has a trust relationship with the client, such as the device operating system or a highly privileged application.

<div style="text-align:left"><img src="../assets/resource_owner_pw.png" width=500 /></div>

(A) The resource owner provides the client with its username and password

(B) The client requests an access token from the authorization server's token endpoint by including the credentials received from the resource owner. When making the request, the client authenticates with the authorization server.

(C) The authorization server authenticates the client and validates the resource owner credentails, if valid, issues an access token.

#### 4.3.1. Acces Token Request

```txt
grant_type
    REQUIRED. Value MUST be set to "password"

username
    REQUIRED. The resource owner username

password
    REQUIRED. The resource owner password

scope
    OPTIONAL. The scope of the access request as described by
```

If the client type is condidential or the client was issued client credentials (or assigned other authentication requirements), the client MUST authenticate with the authorization server.

The authorization server MUST:

- require client authentication for confidential clients or for any client that was issued credentials (or with other authentication requirement)
- authenticate the client if the client authentication is included
- validate the resource owner password credentials using its existing password validation algorithm

Since this access token request utilizes the resource owner's password, the authorization server MUST protect the endpoint against brute force attacks (rate-limit)

#### 4.3.2. Access Token Response

If the access token request is valid and authorized, the authorization server issues ac access token and optional refresh token token.

### 4.4. Client Credentials Grant \*\*

### 4.5. Extension Grants

The client uses an extension grant type by specifying the grant type using an absolute URI (defined by the authorization server) as the value of the "grant_type" parameter of the token endpoint, by adding any additional parameters necessary.

```txt
grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Asaml2-bearer&assertion=PEFzc2VydGlvbiBJc3N1ZUluc3RhbnQ9IjIwMTEtMDU
```

## 5. Issuing an Access Token

If is valid request, the authorization server issues an access token and optional refresh token.

### 5.1. Successful Response

The authorization server issues an access token and option refresh token, and constructs the response

```txt
access_token
    REQUIRED. The access token issued by the authorization

token_type
    REQUIRED. The type of the access token issued

expires_in
    RECOMMENDED. The lifetime in seconds of the access token.

refresh_token
    OPTIONAL. The refresh token, which can be used to obtain new access tokens using the same authorization grant.

scope
    OPTIONAL, if identical to the scope requested by the client;
```

The authorization server MUST include the HTTP "Cache-Control" response header field with a value of "no-store" in any response containing tokens, credentials, or other sensitive information

### 5.2. Error Response

The authorization server responds with an HTTP 400 (Bad Request) status code and includes

```txt
error
    REQUIRED. A single ASCII error code from the following:
    invalid_request
        The request is missing a required parameter, includes an unsupported parameter value (other than grant type)
    invalid_client
        Client authentication failed (e.g unknown client, no client authentication included, or unsupported authentication method)
    invalid_grant
        The provided authorization grant (authorization_code, resource owner credentials) or refresh token is invalid, expired, revoked,
        does not match the redirection URI used in authorization request, or was issued to another client.
    unauthorized_client
        The authenticated client is not authorized to use this authorization grant type.
    unsupported_grant_type
        The authorization grant type is not supported by the authorization server
    invalid_scope
        The requested scope is invalid, unknown, malformed, or exceeds the scope granted by the resource owner.


error_description
    OPTIONAL

error_uri
    OPTIONAL
```

## 6. Refreshing an Access Token

If the authorization server issued a refresh token to the client, the client makes a refresh request to the token endpoint

```txt
grant_type
    REQUIRED. Value MUST be set to "refresh_token"

refresh_token
    REQUIRED. The refresh token issued to the client

scope
    OPTIONAL.
```

Because refresh tokens are typically long-lasting credentials used to request additional access tokens, the refresh token is bound to the client to which it was issued. If the client type is confidential or the client was issued client credentials (or assigned other authentication requirements).

The authorization server MAY issue a new refresh token, in which case the client MUST discard the old refresh and replace it with the new refresh token. The authorization server MAY revoke the old refresh token after issuing a new refresh token to the client. If a new refresh token is issued, the refresh token scope MUST be identical to that of the refresh included by the client in the request.

## 7. Accessing Protected Resources

The client accesss protected resources by presenting the access token to the resource server. The resource server MUST validate the access token and ensure that it has expired and that its scope covers the requested resource. The methods used by the resource server to validate the access token (as well as any error responses) are beyond the scope of this specification but generally involve an interaction or coordination between the resource server and the authorization server.

### 7.1. Access Token Types

The access token type provides the client with the information required to successfully utilize the access token to make a protected resource request. The client MUST NOT use an access token if it does not understand the token type.

The "bearer" token type is utilized by simply including the access token string in the request

While th "mac" token type is utilized by issuing a Message Authentication Code (MAC) key together with the access token that is used to sign certain components of the HTTP requests.

### 7.2. Error Response

## 8. Extensibility

### 8.1. Defining Access Token Types

Access token types can be defined in one of two ways: registered in the Access Token Types registry

### 8.2. Defining New Endpoint Parameters

### 8.3. Defining New Authorization Grant Types

### 8.4. Defining New Authorization Endpoint Response Types

### 8.5. Defining Additional Error Codes

## 9. Native Applications \*\*

Native applications are clients installed and executed on the device used by the resource owner (desktop application, native mobile application). Native applications require special consideration related to security, platform capabilities, and overall end-user experience.

The authorization endpoint requires interaction between the client and the resource owner's user-agent. Native applications can invoke an external user-agent or embed a user-agent within the application. Native applications can invoke an external user-agent or embed a user-agent within the application.

- External user-agent - the native application can capture the response from the authorization server using a redirection URI

## 10. Security Considerations

Security guidelines focused on the three client profiles: web application, user-agent-based application, and native application.

### 10.1. Client Authentication

The authorization server establishes client credentials with web application clients for the purpose of client authentication. Web application clients MUST ensure confidentiality of client passwords and other client credentials.

The authorization server MUST NOT issue client passwords or other client credentials to native application or user-agent based application clients for the purpose of client authentication. The authorization server MAY issue a client password or other credentials for a specific installation a native application client on a specific device.

When client authentication is not possible, the authorization server SHOLD employ other means to validate the client's identity.

### 10.2. Client Impersonation

A malicious client can impersonate another client and obtain access to protected resources if the impersonated client fails to, or is unable to, keep its client credentials confidential.

The authorization server MUST authenticate the client whenever possible. If the authorization server cannot authenticate the client due to the client's nature, the authorization server MUST require the registration of any redirection URI used for receiving authorization responses and SHOULD utilize other means to protect resource owners from such potentially malicious clients.

### 10.3. Access Tokens

Access Token credentials (as well as any confidential access token attributes) MUST be kept confidential in transit and storage.

The client SHOULD request access tokens with the minimal scope necessary. The authorization server SHOULD take the client identity into account whenc choosing how to honor the requested scope and MAY issue an access token with less rights that requested.

### 10.4. Refresh Tokens

Refresh tokens MUST be kept confidential in transit and storage, and shared only among the authorization server and the client to who the refresh tokens were issued.

The authorization server could employ refresh token rotation in which a new refresh token is issued with every access token refresh response.

### 10.5. Authorization Codes

The transmission of authorization codes SHOULD be made over a secure channel, and the client SHOULD require the use of TLS with its redirection URI if the URI identifies a network resource.

Authorization codes MUST be short lived and single-use. If the authorization server observes multiple attempts to exchange an authorization code for an access token, the authorization server SHOULD attempt to revoke all access tokens already granted based on the compromised authorization code.

### 10.6. Authorization Code Redirection URI Manipulation

When requesting authorization using the authorization code grant type, the client can specify a redirection URI via the "redirect_uri" parameter. If an attacker can manipulate the value of the redirection URI, it can cause the authorization server to redirect the resource owner user-agent to a URI under the control of the attacker with the authorization code.

### 10.7. Resource Owner Password Credentials

### 10.8. Request Confidentiality

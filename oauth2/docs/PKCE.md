# Proof Key For Code Exchange by OAuth Public Clients 

References: https://datatracker.ietf.org/doc/html/rfc7636

## Introduction

OAuth 2.0 public clients are susceptible to the authorization code interception attack.

In this attack, the attacker intercepts the authorization code returned from the authorization endpoint within a communication path not protected by TLS, such as inter-application communication within the client's os.

Once the attacker has gained access to the authorization code, it can use it to obtain the access token.

[malicious](../assets/malicious.png)

In step (1), the native application running on the end device, such as smartphone, issues an OAuth 2.0 Authorization Request via the browser/operating system. Step (1) happens through a secure API that cannot be intercepted, though it may potentially be observed in advanced attack scenarios. The request then gets forward to the OAuth 2.0 authorization server in step (2). Because OAuth requires the use of TLS, this communication is protected by TLS and cannot be intercepted. The authorization server returns the authorization code in step (3). In step (4), the Authorization Code is returned to the requester via the Redirect Endpoint URI that was provided in the step(1).

Not that is possible for a malicious app to register ifself as a handler for the custom scheme in addition to the legitimate OAuth 2.0 app.

To mitigate this attack, this extension utilizes a dynamically created cryptographically random key called "code verifier". A unique code verifier is created for every authorization request, and its transformed value, called "code challenge", is sent to the authorization server to obtain the authorization code. The authorization code obtained is then sent to the token endpoint with the "code verifier", and the server compares it with the previously received request code so that it can perform the proof of possession of the "code verifier" by the client. This works as the mitigation since the attacker would not know this one-time key, since it is sent over TLS and cannot be intercepted.

### 1.1. Protocol Flow

[](../assets/pkce1.png)

A. The client creates and records a secret named the "code_verifier" and derives a transformed version "t(code_verifier)" (referred to as the "code_challenge"), which is sent in the OAuth 2.0 Authorization Request along with the transformation method "t_m".
B. The Authorization Endpoint responds as usual but records "t(code_verifier)" and then transformation method.
C. The client then sends the authorization code in the Access Token Request as usual but includes the "code_verifier" secret generated at (A)
D. The authorization server transform "code_verifier" and compares it to "t(code_verifier)" from (B). Access is denied if they are not equal.

## 3. Terminology

In addition in the terms defined in OAuth 2.0, this specification defines the following terms:

**code verifier**
A cryptographically random string that is used to correlate the authorization request to the token request.

**code challenge**
A challenge derived from the code verifier that is sent in the authorization request, to be verified against later

**code challenge method**
A method that was used to derive code challenge

**Base64url Encoding**

## 4. Protocol

### 4.1. Client creates a code verifier

The client first creates a code verifier, "code_verifier" for each OAuth 2.0 Authorization Request

### 4.2. Client creates the code challenge

The client then creates a code challenge derived from the code verifier by using one of the following transformation on the code verifier:

**plain**
    code_challenge = code_verifier

**s256**
    code_challenge = BASE64URL-ENCODE(SHA256(ASCII(code_verifier)))

### 4.3 Client sends the code challenge with the Authorization Request

The client sends the code challenge as part of the OAuth 2.0 Authorization Request using the following additional parameters:

code_challenge
    REQUIRED. Code challenge.

code_challenge_method
    OPTIONAL, defaults to "plain" if not present in the request. Code verifier transformation method is "S256" or "plain"

### 4.4 Server returns the code

When the server issues the authorization code in the authorization response, it MUST associate the "code_challenge" and "code_challenge_method" values with the authorization code so it can be verified later.

Typically the "code_challenge" and "code_challenge_method" values are stored in encrypted form in the "code" itself but could alternatively be store on the server associated with the code. The server MUST NOT include the "code_challenge" value in client requests in a form that other entities can extract. 

code_code_challenge_code_challenge_method
xyz
123_asdf41_sha

=> 123_asdf41_sha, xyz: sha(xyz) = adsf41 => matched

#### 4.4.1. Error Response

If the server requires Proof Key for Code Exchange (PKCE) by OAuth public clients and the client does not send the "code_challenge" in the request, the authorization endpoint MUST return the unauthorization error response with the "error" value set to "invalid_request".

### 4.5. Client sends the Authorization Code and the Code Verifier to the Token Endpoint

Upon receipt of the Authorization Code, the Client sends the Access Token Request to the token endpoint. In addition to the parameters defined in the OAuth 2.0 Access Token Request it send the following parameters:

code_verifier
    REQUIRED. Code verifier

The "code_challenge_method" is bound to the Authorization Code when the Authorization Code is issued. That is the method that the token endpoint MUST use to verify the "code_verifier".

### 4.6. Server verifies code_verifier before returning the Tokens

Upon receipt of the request at the token endpoint, the server verifies it by caculating the code challenge from the received "code_verified" and comparing it with the previously associated "code_challenge", after first transforming it according to the "code_challenge_method" method specified by the client.

### 4.6. Server Verifies code_verifier before Returning the Tokens
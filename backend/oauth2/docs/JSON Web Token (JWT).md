# JSON Web Token (JWT)

Reference: https://datatracker.ietf.org/doc/html/rfc7519

## 1. Introduction

JSON Web Token (JWT) is a compact, URL-safe means of reprenting claims to be transferred between two parties. The claims in a JWT are encoded as a JSON object that is used as the payload of a JSON Web Signature (JWS) structure or as the plaintext of a JSON Web Encryption (JWT) structure, enabling the claims to be digitally signed or integrity protected with a Message Authentication Code signed or integrity protected with a Message Authentication Code.

## 2. Terminology

JSON Web Token (JWT)
A string representing a set of claims as a JSON object that is encoded in a JWS or JWE, enabling the claims to be digitally signed or MACed and/or encrypted.

JWT Claims Set
A JSON object that contains the claims conveyed by the JWT.

Claim
A piece of information asserted about a subject. A claim is represented as a name/value pair consisting of a Claim Name and a Claim value.

Claim Name
The name portion of a claim representation. A Claim name is always a string.

Claim Value
The value portion of a claim representation. A Claim value can be any JSON value.

Nested JWT
A JWT in which nested signing and/or encryption are employed

Unsecured JWT
A JWT whose claims are not integrity protected or encrypted

Collision-Resistant Name

StringOrURI
A JSON string value, with the additional requirement that while arbitrary string values MAY be used, any value containing a ":"

NumericDate

## 3. JSON Web Token (JWT) Overview

Header

```json
{
    "typ": "JWT",
    "alg": "HS256"
}
```

## 4. JWT Claims

The JWT Claims Set represents a JSON object whose members are the claims conveyed by the JWT. The Claim Names within a JWT Claims Set MUST be unique

There are three classes of JWT Claim Names: Registered Claim Names, Public Claim Names, and Private Claim Names.

### 4.1. Registered Claim Names

The following Claim Names are registered in the IANA "JSON Web Token Claims". None of the claims defined below are intended to be mandatory to use or implement in all cases, but rather they provide starting point for a set of useful, interoperable claims. 

#### 4.1.1. "iss" (Issuer) Claim

The "iss" (issuer) claim identifies the principal that issued the JWT. The processing of this claim is generrally application specific. The "iss" value is a case-sensitive string containing a StringOrURI value.

#### 4.1.2. "sub" (Subject) Claim

The "sub" (subject) claim identifies the princical that is the subject of the JWT. The claims in a JWT are normally statements about the subject. The subject value MUST either be scoped to the locally unique in the context of the issuer or be globally unique.

#### 4.1.3. "aud" (Audience) Claim

The "aud" (audience) claim identifies the recipients that the JWT is intended for. Each principal intended to process the JWT MUST identify itself with a value in the audience claim. If the principal processing the claim does not identify itself with a value in the "aud"

#### 4.1.4. "exp" (Expiration Time) Claim

The "exp" (expiration time) claim identifies the expiration time on or after which the JWT MUST NOT be accepted for processing.

#### 4.1.5 "nbf" (Not Before) Claim

The "nbf" (not before) claim identifies the time before which the JWT MUST NOT be accepted for processing. The processing of the "nbf" claim requires that the current date/time MUST be after or equal to the not-before date/time listed in the "nbf" claim.

#### 4.1.6. "iat" (Issued At) Claim

The "iat" (issued at) claim identifies the time at which the JWT was issued. This claim can be used to determine the age of the JWT. Its value MUST be a number containing a NumericDate value. Use of this claim is OPTIONAL.

#### 4.1.7. "jti" (JWT ID) Claim

The "jti" (JWT ID) claim provides a unique identifier for the JWT. The identifier value MUST be assigned in a manner that ensures that there is a negligible probability that the same value will be accidentally assigned to a different data object;

### 4.2. Public Claim Names


### 4.3. Private Claim Names

A producer and consumer of a JWT MAY agree to use Claim Nams that are Private Nams: names that are not Registered Claim Names or Public Claim Names. Unlike Public Claims Names, Private Claim Names are subject to collision and should be used with caution.

## JOSE Header
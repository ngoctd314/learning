# JSON Web Token (JWT) Profile for OAuth 2.0 Access Tokne

## Abstract

This specification defines a profile for issuing OAuth 2.0 access tokens in JSON Web Token (JWT) format. Authorization servers and 

## 1. Introduction

The original OAuth 2.0 Authorization Framework specification does not mandate any specific format for access tokens. In-market use has shown that many commercial OAuth 2.0 implementations elected to issue access tokens using a format that can be parsed and validated by resource servers directly, without further authorization server involvement. The approach is particularly common in topologies where the authorization server and resource server are not co-located, are not run by the same entity, or are otherwise separated by some boundary.

JWT access tokens share the same functional layout, using JWT claims to convey the information needed to support a common set of usecases: token validation, transporting authorization information in the form of scopes and entitlements, carrying identify informatino about the subject, and so on.

JWT access token: An OAuth 2.0 access token encoded in JWT format and complying with the requirements described in this specification.

## 2. JWT access token header and data structure

### 2.1. Header

JWT access tokens MUST be signed. Although JWT access tokens can use any signing algorithm, use of asymmetric cryptography is RECOMMENDED as it simplifies the process of acquiring validation information for resource server.

JWT access token MUST include at+jwt in the "type" header parameter to explicitly declare that the JWT represents an access token complying with this profile. The "typ" value used SHOULD be "at+jwt".

### 2.2. Data Structure

|Field|Severity|
|-|-|
|iss|REQUIRED|
|exp|REQUIRED|
|aud|REQUIRED|
|sub|REQUIRED|
|client_id|REQUIRED|
|iat|REQUIRED|
|jtt|REQUIRED|

#### 2.2.1. Authentication Information Claims

|Field|Severify|
|-|-|
|auth_time|OPTIONAL|
|acr|OPTIONAL|
|amr|OPTIONAL|

### 2.2.2. Identity Claims

Resource servers can consume token directly for authorization or other purposes without any further round trips to introspection or UserInfo endpoints. 
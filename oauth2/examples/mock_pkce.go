package main

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type pkceClient struct {
	verifyCode          string
	codeChallengeMethod string
	server              pkceAuthzServer
}

func (client pkceClient) getAuthzCode() string {
	h := sha256.New()
	h.Write([]byte(client.verifyCode))
	codeChanlenge := ""
	if client.codeChallengeMethod == "S256" {
		codeChanlenge = base64.StdEncoding.EncodeToString(h.Sum(nil))
	} else {
		codeChanlenge = client.verifyCode
	}

	return client.server.genAuthzCode(codeChanlenge, client.codeChallengeMethod)
}

func (client pkceClient) getAccessToken(authzCode string) string {
	accessToken, err := client.server.genAccessToken(authzCode, client.verifyCode)
	if err != nil {
		log.Println(err)
	}

	return accessToken
}

type pkceAuthzServer struct{}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func (server pkceAuthzServer) genAuthzCode(codeChallenge string, codeChallengeMethod string) string {
	return fmt.Sprintf("%d__%s__%s", rand.Intn(1e9), codeChallenge, codeChallengeMethod)
}

func (server pkceAuthzServer) genAccessToken(authzCode string, verifyCode string) (string, error) {
	// verify code challenge
	authz := strings.Split(authzCode, "__")
	code, codeChallenge, codeChallengeMethod := authz[0], authz[1], authz[2]
	if codeChallengeMethod == "S256" {
		h := sha256.New()
		h.Write([]byte(verifyCode))

		if base64.StdEncoding.EncodeToString(h.Sum(nil)) != codeChallenge {
			return "", errors.New("invalid code challenge")
		}
		return fmt.Sprintf("access_token:%d:%s", rand.Intn(1e9), code), nil

	}

	// plain strategy
	if verifyCode != codeChallenge {
		return "", errors.New("invalid code")
	}

	return fmt.Sprintf("access_token:%d:%s", rand.Intn(1e9), code), nil
}

type maliciousPKCEClient struct {
	verifyCode          string
	codeChallengeMethod string
	server              pkceAuthzServer
}

func (client maliciousPKCEClient) getAuthzCode() string {
	h := sha256.New()
	h.Write([]byte(client.verifyCode))
	codeChanlenge := ""
	if client.codeChallengeMethod == "S256" {
		codeChanlenge = base64.StdEncoding.EncodeToString(h.Sum(nil))
	} else {
		codeChanlenge = client.verifyCode
	}

	return client.server.genAuthzCode(codeChanlenge, client.codeChallengeMethod)
}

func (client maliciousPKCEClient) getAccessToken(authzCode string) string {
	accessToken, err := client.server.genAccessToken(authzCode, client.verifyCode)
	if err != nil {
		log.Println(err)
	}

	return accessToken
}

func mockPKCE() {
	pkceClient := pkceClient{
		verifyCode:          "valid_client",
		codeChallengeMethod: "S256",
		server:              pkceAuthzServer{},
	}
	authzCode := pkceClient.getAuthzCode()
	fmt.Println(pkceClient.getAccessToken(authzCode))

	maliciousPKCEClient := maliciousPKCEClient{
		verifyCode:          "malicious_client",
		codeChallengeMethod: "S256",
		server:              pkceAuthzServer{},
	}
	fmt.Println(maliciousPKCEClient.getAccessToken(authzCode))
}

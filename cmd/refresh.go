// Copyright © 2018 Justin Lee <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	// "bytes"
	// "fmt"
	"io/ioutil"
	// "crypto/tls"
	// "encoding/json"
	// "encoding/base64"

	"log"
	// "strings"
	// "net/http"
	// "time"
	// "errors"

	"github.com/spf13/cobra"
	// "github.com/dgrijalva/jwt-go"
)

// type loginResponse struct {
// 	Token string `json:"token"`
// }

// func getKey(token *jwt.Token) (interface{}, error) {

// 	jwksEndpoint := "https://" + master + "/acs/api/v1/auth/jwks"
// 	// TODO: cache response so we don't have to make a request every time 
// 	// we want to verify a JWT
// 	set, err := jwk.FetchHTTP(jwksEndpoint)
// 	if err != nil {
// 			return nil, err
// 	}

// 	keyID, ok := token.Header["kid"].(string)
// 	if !ok {
// 			return nil, errors.New("expecting JWT header to have string kid")
// 	}

// 	if key := set.LookupKeyID(keyID); len(key) == 1 {
// 			return key[0].Materialize()
// 	}

// 	return nil, errors.New("unable to find key")

// }

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh Service Authentication Token if it has expired",
	Long: `Refresh Service Authentication Token if it has expired`,

	Run: func(cmd *cobra.Command, args []string) {

		if privateKeyFile == "" || uid == "" || outputFile == "" {
			log.Fatal("Must provide at least a private key (-k) and a uid (-u) and token file (-o)")
		}

		refresh := false
		tokenString, err := ioutil.ReadFile(outputFile)
		
		if err != nil {
			// File does not exist; let's write to it
			refresh = true
		} else {
			refresh, err = checkExpired(string(tokenString))
			if err != nil {
				log.Fatal(err)
			}
		}

			// jwksEndpoint = "https://" + master + "/acs/api/v1/auth/jwks"
			
			// set, err := jwk.FetchHTTP(jwksEndpoint)

			// if err != nil {
			// 	return 
			// }

			// token, err := jwt.Parse(string(tokenString), getKey)
			// if err != nil {
			// 	fmt.Println(err)
			// }
			// claims := token.Claims.(jwt.MapClaims)
			// fmt.Println(claims)

			// token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
			// 	// We aren't actually validating the token, just whether it thinks it has expired (for now)
			// 	// TODO actually validate the token
			// 	return []byte("YES"), nil
			// })

			// // File exists but is not a token; probably don't want to touch it
			
			// claims := token.Claims.(jwt.MapClaims)

			
			// expTime := claims["exp"].(float64)

			// fmt.Println(minValidTime)
			// fmt.Println(expTime)

			// if err != nil {
			// 	fmt.Println(err)
			// 	// log.Fatal("Token file is not parseable as token; not modifying")
			// }
			
			// if expTime < minValidTime {
			// 	refresh = true
			// }
		
		if refresh {
			// fmt.Println("Let's refresh")
	
			privateKey, err := ioutil.ReadFile(privateKeyFile)
			if err != nil {
				log.Fatal(err)
			}
	
			// Returns a []byte
			loginObject, err := generateServiceLoginObject(privateKey, uid, validTime)
			if err != nil {
				log.Fatal(err)
			}
			
			authToken, err := login(master, loginObject)
	
			err = output([]byte(authToken))
			if err != nil {
				log.Fatal(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(refreshCmd)
}


// func generateServiceLoginToken(privateKey byte[], uid string, validTime int) (jwt string, err error) {

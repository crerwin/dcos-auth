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
	"fmt"
	"io/ioutil"
	"time"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
)

// signJwtCmd represents the signJwt command
var signJwtCmd = &cobra.Command{
	Use:   "signJwt",
	Short: "Sign a JWT token using a private key and uid",
	Long: `Sign a JWT token using a private RSA key and uid.
	
	Designed for Enterprise DC/OS, and outputs in format accepted by the
/acs/api/v1/auth/login endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		if privateKeyFile == "" || uid == "" {
			log.Fatal("Must provide at least a private key (-k) and a uid (-u)")
		}

		privateKey, err := ioutil.ReadFile(privateKeyFile)
		if err != nil {
			log.Fatal(privateKeyFile + " does not exist or is not readable.")
		}

		key, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
		if err != nil {
			log.Fatal(privateKeyFile + "is not an RSA private key")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"uid": uid,
			"exp": time.Now().Add(time.Second * time.Duration(validTime)).Unix(),
		})

		tokenString, err := token.SignedString(key)

		if err != nil {
			log.Fatal("Unable to generate token")
		}

		// Not sure if it's more efficient (spacewise) to do this or add the json package.
		fmt.Println("{\"token\": \"" + tokenString + "\", \"uid\": \"" + uid + "\"}")
	},
}

func init() {
	rootCmd.AddCommand(signJwtCmd)
}

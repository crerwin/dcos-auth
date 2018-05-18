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
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Refresh Service Authentication Token if it has expired",
	Long:  `Refresh Service Authentication Token if it has expired`,

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

		if refresh {
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

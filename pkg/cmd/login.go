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
	"log"

	"github.com/crerwin/dcosauth/pkg/dcosauth"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in and generate Service Authentication Token",
	Long:  `Log in and generate Service Authentication Token`,

	Run: func(cmd *cobra.Command, args []string) {

		if privateKeyFile == "" || uid == "" {
			log.Fatal("Must provide at least a private key (-k) and a uid (-u)")
		}

		privateKey, err := dcosauth.Input(privateKeyFile)

		dcosauther := dcosauth.New(master, uid, string(privateKey))

		authToken, err := dcosauther.Token()

		err = dcosauth.Output([]byte(authToken), outputFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

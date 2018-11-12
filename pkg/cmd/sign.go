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
	// "fmt"
	"io/ioutil"
	"log"

	"github.com/crerwin/dcosauth/pkg/dcosauth"
	"github.com/spf13/cobra"
)

// signJwtCmd represents the signJwt command
var genSLTCmd = &cobra.Command{
	Use:   "genSLT",
	Short: "Generate a Service Login Token",
	Long:  `Generate a Service Login Token, in JWT format`,

	Run: func(cmd *cobra.Command, args []string) {

		if privateKeyFile == "" || uid == "" {
			log.Fatal("Must provide at least a private key (-k) and a uid (-u)")
		}

		privateKey, err := ioutil.ReadFile(privateKeyFile)
		if err != nil {
			log.Fatal(err)
		}

		loginToken, err := dcosauth.GenerateServiceLoginToken(privateKey, uid, validTime)
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Println(string(loginObject))
		err = dcosauth.Output([]byte(loginToken), outputFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var genSLOCmd = &cobra.Command{
	Use:   "genSLO",
	Short: "Generate JSON-formatted Service Login Token",
	Long:  `Generate JSON-formatted Service Login Token, formatted for DC/OS ACS API`,

	Run: func(cmd *cobra.Command, args []string) {

		if privateKeyFile == "" || uid == "" {
			log.Fatal("Must provide at least a private key (-k) and a uid (-u)")
		}

		privateKey, err := ioutil.ReadFile(privateKeyFile)
		if err != nil {
			log.Fatal(err)
		}

		loginObject, err := dcosauth.GenerateServiceLoginObject(privateKey, uid, validTime)
		if err != nil {
			log.Fatal(err)
		}

		err = dcosauth.Output(loginObject, outputFile)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(genSLTCmd)
	rootCmd.AddCommand(genSLOCmd)
}

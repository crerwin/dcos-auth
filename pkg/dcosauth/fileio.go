package dcosauth

import (
	"fmt"
	"io/ioutil"
)

// Input reads a file (to be used for reading private key files)
func Input(inputFilePath string) ([]byte, error) {
	return ioutil.ReadFile(inputFilePath)
}

// Output writes given content to a given filepath
func Output(content []byte, outputFilePath string) (err error) {
	err = nil
	if outputFilePath != "" {
		err = ioutil.WriteFile(outputFilePath, []byte(content), 0600)
	} else {
		fmt.Println(string(content))
	}
	return err
}

package main

import (
	"fmt"

	"github.com/prashant182/goutils"
)

func main() {

	err := goutils.ExtractFromRemoteFile("/tmp", "http://localhost:8080/nginx.tar.gz")
	if err != nil {
		fmt.Errorf("Error in doing this", err)
	}
}

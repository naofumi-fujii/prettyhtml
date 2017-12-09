package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yosssi/gohtml"
)

func main() {
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gohtml.Format(string(body)))
}

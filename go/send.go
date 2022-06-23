package main

import (
	// "log"
	// "ioutil"
	"github.com/Microsoft/go-winio"
	"fmt"
	"reflect"
)

func main() {
	conn, err := winio.DialPipe(`\\.\pipe\pipe1`, nil)
	if err != nil {
		// <handle error>
	}
	fmt.Println(reflect.TypeOf(conn))
	fmt.Fprintf(conn, "Hi server!\n")
	fmt.Println("sent")
	// msg, err := bufio.NewReader(conn).ReadString('\n')
}
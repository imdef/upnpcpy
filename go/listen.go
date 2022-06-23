package main

import (
	// "log"
	// "ioutil"
	"github.com/Microsoft/go-winio"
	// "reflect"
	"fmt"
	"net"
	"bufio"
	"io"
)

func handleConnection(conn1 net.Conn, conn2 net.Conn) {
	fmt.Println("got from: ",conn1.RemoteAddr())
	fmt.Println("got from: ",conn2.RemoteAddr())

	// msg, err := bufio.NewReader(conn).ReadByte()
	// msg2, err := bufio.NewReader(conn).ReadByte()
	// if (err != nil){fmt.Println(err); return}
	// fmt.Println(reflect.TypeOf(msg))
	// fmt.Println("=====================================")
	// fmt.Println(msg2)

	// ===========Byte by Byte==============
	// for {
	// 	msg, err := bufio.NewReader(conn).ReadByte()
	// 	if (err != io.EOF){fmt.Println("Byte end EOF (still confused)")}
	// 	if (err != nil){fmt.Println("Overwrite Check/End of ffmpeg Stream: ",err); return}
	// 	fmt.Println(msg)
	// }

	// ========== Bytes string =============
	// N := 4096
	// buf := make([]byte,N)
	// for {
		// msg, err := bufio.NewReader(conn).ReadBytes(io.EOF)
		bufreader := bufio.NewReader(conn1)
		bufwriter := bufio.NewWriter(conn2)
		// n, err := io.ReadFull(bufreader, buf)
		n, err := io.Copy(bufwriter, bufreader)

		// if (err != io.EOF){fmt.Println("Bytes Block end EOF (still confused)")}
		if (err != nil){fmt.Println("Overwrite Check/End of ffmpeg Stream: ",err)}
		
		if (n > 0) {
			fmt.Println("Success")
		}
		fmt.Println(conn1.Close())
		fmt.Println(conn2.Close())

	// }

}

func main() {

	ln1, err := winio.ListenPipe(`\\.\pipe\pipe1`, nil)
	ln2, err := winio.ListenPipe(`\\.\pipe\pipe2`, nil)
	fmt.Println("hoop")
	if err != nil {
		// handle error
	}

	for {
		conn1, err := ln1.Accept()
		// conn1, err = ln1.Accept()

		// fmt.Println(err)
		conn2, err := ln2.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn1, conn2)
	}
}


// conn is of type *winio.win32Pipe and net.Conn , ln is *winio.win32PipeListener and net.Listener 
// (https://pkg.go.dev/github.com/microsoft/go-winio@v0.3.4#ListenPipe)

// one server two clients (one for ffmpeg, one for dlna)
package main

import (
	"log"
	// "ioutil"
	"github.com/Microsoft/go-winio"
	// ffmpeg "github.com/floostack/transcoder/ffmpeg"
	// "reflect"
	"fmt"
	"net"
	"bufio"
	"io"
	"net/http"
	"time"
	"os/exec"
)

func makeServer() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello, %q", "go")
		http.ServeFile(w, r, `\\.\pipe\pipe2`)
	})

	log.Fatal(http.ListenAndServe(":1226", nil))
}

func handleConnection(conn1 net.Conn, conn2 net.Conn) {
	fmt.Println("got from: ",conn1.RemoteAddr())
	fmt.Println("got from: ",conn2.RemoteAddr())

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
}


func main() {

	// go makeServer()

	// ln1, err := winio.ListenPipe(`\\.\pipe\pipe1`, nil)
	// // ln2, err := winio.ListenPipe(`\\.\pipe\pipe2`, nil)
	// fmt.Println("hoop")
	// if err != nil {
	// 	// handle error
	// }

	// conn1, err := ln1.Accept()
	// conn1, err = ln1.Accept()
	// bufreader := bufio.NewReader(conn1)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		pipename := "pipe1"
		ln1, err := winio.ListenPipe(`\\.\pipe\` + pipename, nil)
		fmt.Println(ln1)
		fmt.Println("hoop")
		if err != nil {
			// handle error
		}

		go func(pipe string){
			time.Sleep(time.Second * 1)
			fmt.Println("inside ffmpeg call")
			err := exec.Command("ffmpeg", "-y", "-f", "gdigrab", "-r", "30" , "-i", "desktop", "-vcodec", "libx264","-qp", "22", "-preset" ,"medium", "-threads", "0", "-vf", "format=yuv420p", "-f", "matroska", "\\\\.\\pipe\\pipe1").Run()
			fmt.Println(err)
			// exec.Command("ffmpeg -y -f gdigrab -r 30 -i desktop -vcodec libx264 -qp 22 -preset medium -threads 0 -vf format=yuv420p -f matroska \\\\.\\pipe\\pipe1").Run()
		}(pipename)

		fmt.Println("ffmpeg called")
		fmt.Println(ln1.Addr())
		conn1, err := ln1.Accept()
		fmt.Println("Stream got")
		bufreader := bufio.NewReader(conn1)
		n, err := io.Copy(w, bufreader)
		if (err != nil){ fmt.Println("Overwrite Check/End of ffmpeg Stream: ",err) }

		if (n > 0) {
			fmt.Println("Success")
			ln1.Close()
		}
	})

	log.Fatal(http.ListenAndServe(":1226", nil))
	
}

// 30 fps on the tvpc
package main

import (
	"log"
	// "ioutil"
	ffmpeg "github.com/floostack/transcoder/ffmpeg"
	// "github.com/Microsoft/go-winio"
	"github.com/Microsoft/go-winio"
	// "reflect"
	"fmt"
	"sync"
	// "io"
	"bufio"
	"time"
)

func listen(wg *sync.WaitGroup) {
	defer wg.Done()

	var wg1 sync.WaitGroup
	wg1.Add(2)

	
	fmt.Println("listen")
	server_ln, err := winio.ListenPipe(`\\.\pipe\pipe1`, nil) // nil means what maybe read only???
	if err != nil {fmt.Println("error")}
	go opff(&wg1)
	output_conn, err := server_ln.Accept()
	op_buf := bufio.NewReadWriter(bufio.NewReader(output_conn), bufio.NewWriter(output_conn))

	fmt.Println("listen1")
	if err != nil {fmt.Println("error", output_conn)}
	fmt.Println("outputff Connection", output_conn.RemoteAddr())
	go ipff(&wg1)
	// input_conn, err := server_ln.Accept()
	// ip_buf := bufio.NewReadWriter(bufio.NewReader(input_conn), bufio.NewWriter(input_conn))

	fmt.Println("listen2")
	// if err != nil {fmt.Println("error", input_conn)}
	// fmt.Println("inputff Connection", input_conn.RemoteAddr())

	for {
		b, err := op_buf.ReadByte()
		if err != nil { fmt.Println(err) }
		if (err == nil) {
			fmt.Println(b)
			// input_conn.Write(b)
			// b = nil
		}
		// fmt.Println(n, err)
		time.Sleep(time.Second * 1)
	}

	wg1.Wait()
	fmt.Println("listen end")
}

func opff(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("opff")
	format := "pipe"
	overwrite := true
	opts := ffmpeg.Options{
			OutputFormat: &format,
			Overwrite:    &overwrite,
		}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   "C:\\Users\\User\\Desktop\\tv\\ffmpeg\\bin\\ffmpeg.exe",
		FfprobeBinPath:  "C:\\Users\\User\\Desktop\\tv\\ffmpeg\\bin\\ffprobe.exe",
		ProgressEnabled: false,
	}

	_,err := ffmpeg.
		New(ffmpegConf).
		Input("C:\\Users\\User\\Videos\\2020-11-25 14-25-31.mkv").
		Output("\\\\.\\pipe\\pipe1").
		WithOptions(opts).
		Start(opts)

	if err != nil { fmt.Println("err in opff") }
	fmt.Println("Started opff")

}

func ipff(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("ipff")
	format := "mp4"
	ip_format := "pipe"
	overwrite := true
	opts := ffmpeg.Options{
			InputFormat: &ip_format,
			OutputFormat: &format,
			Overwrite:    &overwrite,
		}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   "C:\\Users\\User\\Desktop\\tv\\ffmpeg\\bin\\ffmpeg.exe",
		FfprobeBinPath:  "C:\\Users\\User\\Desktop\\tv\\ffmpeg\\bin\\ffprobe.exe",
		ProgressEnabled: false,
	}

	progress, _ := ffmpeg.
		New(ffmpegConf).
		Input("\\\\.\\pipe\\pipe1").
		Output("C:\\Users\\User\\Desktop\\tv\\ffmpeg\\bin\\go\\a.mp4").
		WithOptions(opts).
		Start(opts)

	fmt.Println("Started ipff")

	for msg := range progress {
		log.Printf("%+v", msg)
	}
	// if err != nil { fmt.Println("err in ipff") }
	
}

func main() {

	// client_conn, err := winio.DialPipe(`\\.\pipe\pipe1`, nil)
	// client_reader := io.PipeReader(client_conn)
	// client_writer := io.PipeWriter(client_conn)

	// FFMPEG

	
	fmt.Println("main")

	var wg sync.WaitGroup
	wg.Add(1)

	go listen(&wg)

	wg.Wait()
	fmt.Println("main end")
	

	

}


/*
SO,
ffmpeg detects overwrite on the first connection and then connects again to put the data -_-
*/
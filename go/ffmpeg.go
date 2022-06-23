package main

import (
	"log"
	// "ioutil"
	ffmpeg "github.com/floostack/transcoder/ffmpeg"
	// "github.com/Microsoft/go-winio"
)

// ffmpeg 

func main() {

	format := "mp4"
	overwrite := true

	opts := ffmpeg.Options{
		OutputFormat: &format,
		Overwrite:    &overwrite,
	}

	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath:   "C:\\Users\\imDef\\Desktop\\tv\\ffmpeg\\bin\\ffmpeg.exe",
		FfprobeBinPath:  "C:\\Users\\imDef\\Desktop\\tv\\ffmpeg\\bin\\ffprobe.exe",
		ProgressEnabled: true,
	}

	progress, err := ffmpeg.
		New(ffmpegConf).
		Input("C:\\Users\\imDef\\Videos\\2020-11-07 10-45-39.mkv").
		Output("\\\\.\\pipe\\pipe1").
		WithOptions(opts).
		Start(opts)

	if err != nil {
		log.Fatal(err)
	}

	for msg := range progress {
		log.Printf("%+v", msg)
	}

}


// Use ioutil for now, if doesn't work use bufio
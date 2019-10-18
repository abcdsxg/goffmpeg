package test

import (
	"log"
	"testing"

	"github.com/abcdsxg/goffmpeg/transcoder"
)

func TestInputNotFound(t *testing.T) {

	var inputPath = "/Users/shingle/Desktop/1570593272272801.mp4"
	var audioPath="/Users/shingle/Desktop/3.mp3"
	var outputPath = "/Users/shingle/Desktop/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath,audioPath}, outputPath)
	if err != nil {
		t.Fatal(err)
		return
	}
	trans.MediaFile().SetShortest(true)
	trans.MediaFile().SetVideoCodec("copy")
	m:=map[int]string{
		0:"v:0",
		1:"a:0",
	}
	trans.MediaFile().SetMapIds(m)
	done := trans.Run(true)

	progress := trans.Output()

		for msg := range progress {
			log.Println("proc:",msg.Progress)
		}

	err = <-done
	if err != nil {
		t.Fatal(err)
	}
}

func TestTranscoding3GP(t *testing.T) {

	var inputPath = "/data/test3gp"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingAVI(t *testing.T) {

	var inputPath = "/data/testavi"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingFLV(t *testing.T) {

	var inputPath = "/data/testflv"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingMKV(t *testing.T) {

	var inputPath = "/data/testmkv"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingMOV(t *testing.T) {

	var inputPath = "/data/testmov"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingMPEG(t *testing.T) {

	var inputPath = "/data/testmpeg"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingOGG(t *testing.T) {

	var inputPath = "/data/testogg"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingWAV(t *testing.T) {

	var inputPath = "/data/testwav"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingWEBM(t *testing.T) {

	var inputPath = "/data/testwebm"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingWMV(t *testing.T) {

	var inputPath = "/data/testwmv"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(false)
	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

func TestTranscodingProgress(t *testing.T) {

	var inputPath = "/data/testavi"
	var outputPath = "/data/testmp4.mp4"

	trans := new(transcoder.Transcoder)

	err := trans.Initialize([]string{inputPath}, outputPath)
	if err != nil {
		t.Error(err)
		return
	}

	done := trans.Run(true)
	for val := range trans.Output() {
		if &val != nil {
			break
		}
	}

	err = <-done
	if err != nil {
		t.Error(err)
		return
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hajsf/mjpeg-streamer"
	"gocv.io/x/gocv"
)

func main() {
	deviceID := 0
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}

	// create the mjpeg stream
	stream := mjpeg.NewStream()

	// start capturing
	go func(webcam *gocv.VideoCapture, stream *mjpeg.Stream) {
		defer webcam.Close()

		window := gocv.NewWindow("Capture Window")
		defer window.Close()

		img := gocv.NewMat()
		defer img.Close()

		fmt.Printf("Start reading device: %v\n", deviceID)
		for {
			if ok := webcam.Read(&img); !ok {
				fmt.Printf("Device closed: %v\n", deviceID)
				return
			}
			if img.Empty() {
				continue
			}
			buf, _ := gocv.IMEncode(".jpg", img)
			stream.UpdateJPEG(buf)
			window.IMShow(img)
			if window.WaitKey(1) == 27 { // 27 => Esc
				break
			}
		}
	}(webcam, stream)

	http.Handle("/", stream)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

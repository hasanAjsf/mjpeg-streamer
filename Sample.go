func ExampleUpdateJPEG() {
  // Create new buffer for the image
  buf, _ := gocv.IMEncode(".jpg", img)
  // Push the new buffer back to the stream
	stream.UpdateJPEG(buf)
  // Output:
  // New stream is pushed to the server
}

func ExampleUpdateJPEG_options() {
  // Output:
}

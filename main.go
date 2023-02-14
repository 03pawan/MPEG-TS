package main

func main(){

	filename := "~/HDD/Downloads/sample_1280x720_surfing_with_audio.ts"
	byteStream := ReadFile(filename) // returns bytes of File

	transportStreamPackets := formPackets(byteStream)

	headers := getHeader(transportStreamPackets)

	PIDS := mapPIDS(headers)

	PIDS.print()

}
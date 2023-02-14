package main

import(
	"fmt"
	"os"
)

const transportStreamPacketLength = 188
const headerLength = 4

type transportStreamPacket []byte
type mapPID map[int]bool

func ReadFile(filename string) []byte{

	file, _ := os.OpenFile(filename,os.O_RDONLY,222)
	fileStat, _ := file.Stat()
	
	byteStream := make([]byte,fileStat.Size())
	file.Read(byteStream)

	return byteStream
}

func formPackets(byteStream []byte) []transportStreamPacket{

	transportStreamPackets := []transportStreamPacket{}

	newPacket := make(transportStreamPacket,transportStreamPacketLength)
	for i, bytes := range byteStream{
		if i % transportStreamPacketLength-1 == 0 && i != 0{
			transportStreamPackets = append(transportStreamPackets, newPacket)
		}
		newPacket[i%transportStreamPacketLength-1] = bytes
	}

	return transportStreamPackets
}

func getHeader(packet []transportStreamPacket) []transportStreamPacket{

	headers := make([]transportStreamPacket,headerLength)
	for _, pack := range packet{
		headers = append(headers, pack[:4])
	}
	return headers
}

func mapPIDS(header []transportStreamPacket) mapPID{
	var PIDS map[int]bool 

	for _,bits := range header{
		pid := byteToIntPID(bits)
		PIDS[pid] = true
	}
	
	return PIDS
}

func (m mapPID) print(){
	for key := range m{
		fmt.Println("PID: ",key)
	}
}

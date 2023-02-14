package main

func byteToIntPID(bytes []byte) int{
	var bitRep []int
	
	pidByteOne := toBit((int)(bytes[1]))
	pidByteTwo := toBit((int)(bytes[2]))

	extractPID(pidByteOne,pidByteTwo,bitRep)
	return toInt(bitRep)
}


func extractPID(pidByteOne []int,pidByteTwo []int, bitrep []int){
	bitrep = append(bitrep,pidByteOne[3:]...)
	bitrep = append(bitrep, pidByteTwo...)
}

func toBit(dec int) []int{
	i:=0
	var res []int
	for i<8 {
		res = append([]int{dec&1}, res...) 	// pre-pending from front
		dec >>= 1
	}
	return res
}

func toInt(bits []int) int{
	i := 1
	res := 0
	for _,bit := range bits{
		res += (i*bit)
		i*=2
	} 

	return res
}
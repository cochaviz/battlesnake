package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	Turn   []int
	Length []int
}

// TODO Make this just better
func (data Data) appendMeasuremen(measurement Measurement) {
	data.Turn = append(data.Turn, measurement.Turn)
	data.Length = append(data.Length, measurement.Length)
}

func (data Data) WriteToFile(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err, "Trying to create file")
		f, err = os.Create(filename)
	}
	if err != nil {
		log.Println(err)
	}
	jsonEncoded, _ := json.MarshalIndent(data, "", " ")
	fmt.Fprintln(f, string(jsonEncoded), ",")
}

type Measurement struct {
	Turn   int
	Length int
}

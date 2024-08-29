package main

import (
	"go-face-test/service"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	// -------------------------- INIT START--------------------------
	service.NewRecognise()
	// -------------------------- INIT END--------------------------
	name := "周杰伦"
	imgData, _ := ioutil.ReadFile("./jay.jpg")
	recImg, _ := ioutil.ReadFile("./jay1.jpg")
	err := service.EnterFaceData(name, imgData)
	if err != nil {
		log.Fatalln("EnterFaceData failed: ", err)
	}
	start := time.Now()
	err_ := service.RecogniseFace(recImg)
	if err_ != nil {
		log.Fatalln("RecogniseFace err", err_)
	}
	elapsed := time.Since(start)
	log.Printf("RecogniseFace took %s", elapsed)
}

package service

import (
	"errors"
	"fmt"
	"go-face-test/define"
)

// EnterFaceData 录入人脸数据
func EnterFaceData(name string, imgData []byte) (err error) {
	face, err := Rec.RecognizeSingle(imgData)
	if err != nil {
		return
	}
	if face == nil {
		err = errors.New("未检出到人脸数据")
		return
	}
	// 判断人脸数据是否存在
	id := Rec.ClassifyThreshold(face.Descriptor, define.Tolerance)
	if id > 0 {
		err = errors.New("数据已存在，无需重复录入")
		return
	}
	// 录入人脸数据
	FaceData.Samples = append(FaceData.Samples, face.Descriptor)
	FaceData.Ids = append(FaceData.Ids, int32(len(FaceData.Ids)+1))
	FaceData.Names = append(FaceData.Names, name)
	Rec.SetSamples(FaceData.Samples, FaceData.Ids)
	fmt.Println("数据录入成功")
	return
}

// RecogniseFace 人脸识别
func RecogniseFace(imgData []byte) (err error) {
	face, err := Rec.RecognizeSingle(imgData)
	if err != nil {
		fmt.Println("Rec.RecognizeSingle fail", err)
		return
	}
	if face == nil {
		err = errors.New("未检出到人脸数据")
		return
	}
	id := Rec.ClassifyThreshold(face.Descriptor, define.Tolerance)
	if id < 0 {
		err = errors.New("人脸数据不存在")
		return
	}
	fmt.Println("识别成功, 你好", FaceData.Names[id-1])
	return
}

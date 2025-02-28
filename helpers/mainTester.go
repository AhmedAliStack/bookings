package helpers

import (
	"errors"
	"log"
)

func mainTester() {

	result, err := Deviding(10,0)
	
	if err != nil {
		log.Println("There is error", err.Error())
		return
	}

	log.Println("Devide Result is", result)
}

func Deviding(firstNum float32, secondNum float32) (result float32, err error) {
	if secondNum == 0 {
		err = errors.New("devide by 0 is forbidden")
		return 0, err
	}

	result = firstNum / secondNum
	return result, nil
}
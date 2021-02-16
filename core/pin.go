package core

import (
	"strconv"
)

type Pin struct {
	data   []byte
}

func (pin *Pin) CreatePin(pinString string) bool {
	if len(pinString) < 4 && len(pinString) > 16 {
		return false
	}
	_, err := strconv.Atoi(pinString)
	if err != nil {
		return false
	}
	pin.data = []byte(pinString)
	return true
}

func (pin *Pin) GetBytes() []byte {
	return pin.data
}
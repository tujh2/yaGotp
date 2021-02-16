package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"math"
	"time"
)

const timePeriodSec = 30
const otpLength = 8

func toBigEndianArray(value int64) []byte {
	var res [8]byte
	for i := 7; i >= 0; i-- {
		res[i] = byte(value & 0xFF)
		value >>= 8
	}

	return res[:]
}

func fromBigEndianArray(startIndex int, array []byte) int64 {
	var value int64 = 0

	for i := startIndex; i <= startIndex + 7; i++ {
		value = value << 8 | int64(array[i])
		//println(value)
	}
	return value
}

func toBase26(value int64, length int) string {
	var chars = make([]rune, length, length)
	var index = length - 1

	for index >= 0 {
		chars[index] = 'a' + rune(value%26)
		value /= 26
		index--
	}
	return string(chars)
}

func ComputeOtp(pin Pin, secret Secret) string {
	keyHashTmp := sha256.Sum256(append(pin.data, secret.data...))
	var keyHash []byte
	if keyHashTmp[0] == 0 {
		keyHash = keyHashTmp[1:]
	} else {
		keyHash = keyHashTmp[:]
	}
	var curPeriod = time.Now().UTC().Unix() / timePeriodSec

	hash := hmac.New(sha256.New, keyHash[:])
	hash.Write(toBigEndianArray(curPeriod))
	periodHash := hash.Sum(nil)
	periodHashSlice := fromBigEndianArray(int(periodHash[len(periodHash) - 1] & 15), periodHash) & int64(^uint64(0) >> 1)

	limitValue := int64(math.Pow(26, otpLength))
	return toBase26(periodHashSlice % limitValue, otpLength)
}
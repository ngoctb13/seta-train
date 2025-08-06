package utils

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

var uuidFunc = uuid.NewRandom
var timeNowFunc = time.Now

func Generate() string {
	uid, err := uuidFunc()
	if err != nil {
		return strconv.Itoa(int(timeNowFunc().Unix()))
	}

	return uid.String()
}

package util

import (
	"github.com/bobgo0912/b0b-common/pkg/log"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

func GetLastPrizeNumber(prizeNumber string) (int, error) {
	if prizeNumber == "" {
		return 0, errors.New("bad prizeNumber")
	}
	p := strings.Split(prizeNumber, ",")
	if len(p) > 0 {
		s := p[len(p)-1]
		i, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return 0, errors.Wrap(err, "ParseInt fail")
		}
		return int(i), nil
	}
	return 0, errors.New("bad PrizeNumber1")
}

// Retry Retry1
func Retry[T any](tryTimes int, sleep time.Duration, callback func() (*T, error)) (*T, error) {
	for i := 1; i <= tryTimes; i++ {
		ret, err := callback()
		if err == nil {
			return ret, nil
		}
		if i == tryTimes {
			log.Error("Retry err=", err.Error())
			return nil, errors.New("Retry over")
		}
		time.Sleep(sleep)
	}
	return nil, nil
}

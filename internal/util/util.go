package util

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
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

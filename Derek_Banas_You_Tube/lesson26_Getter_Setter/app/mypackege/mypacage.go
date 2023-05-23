package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Rasulaka"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, i := range intArr {
		strArr = append(strArr, strconv.Itoa(i))
	}
	return strArr
}

type Date struct {
	day   int
	month int
	year  int
}

func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("incorrect dey value")
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(m int) error {
	if (m < 1) || (m > 12) {
		return errors.New("incorrect month value")
	}
	d.month = m
	return nil
}

func (d *Date) SetYear(Y int) error {
	if (Y < 1874) || (Y > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	d.year = Y
	return nil
}

func (d *Date) Day() int { return d.day }

func (d *Date) Month() int { return d.month }

func (d *Date) Year() int { return d.year }

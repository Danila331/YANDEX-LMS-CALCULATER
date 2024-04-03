package pkg

import "github.com/Danila331/YAP-2/internal/models"

func CalcEndTime(expressionString string, timeNow int) (int, int, error) {
	var time models.Time
	time, err := time.ReadAll()
	if err != nil {
		return 0, 0, err
	}
	var countPulse, countMinus, countProz, countDel, timeProcessing int
	for _, el := range expressionString {
		if string(el) == "+" {
			countPulse++
		}
		if string(el) == "-" {
			countMinus++
		}
		if string(el) == "*" {
			countProz++
		}
		if string(el) == "/" {
			countDel++
		}
	}
	timeProcessing += (countPulse*time.TimePulse + countMinus*time.TimeMinus + countProz*time.TimeProz + countDel*time.TimeDel)
	timeEnd := timeNow + timeProcessing + 500
	return timeEnd, timeProcessing, nil
}

package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	goErr "github.com/ralstan-vaz/go-errors"
)

const (
	Year   = "year"
	Month  = "month"
	Day    = "day"
	Hour   = "hour"
	Minute = "minute"
	Second = "second"
)

// Bind ... uses mapstructure internally to bind input to output (does not use tags)
func Bind(input interface{}, output interface{}) error {
	err := mapstructure.Decode(input, output)
	if err != nil {
		return goErr.NewInternalError(err).SetCode("PKG.UTILS.DECODE_ERROR")
	}

	return nil
}

func MapToString(ip map[string]interface{}) string {
	jsonStr, err := json.Marshal(ip)
	if err != nil {
		return ""
	}
	return string(jsonStr)
}

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func StringToInt(req string) (res int) {
	res, err := strconv.Atoi(req)
	if err != nil {
		return 0
	}
	return res
}

func StringToEpoch(req string) (res int64) {
	t, _ := time.Parse(time.RFC1123, req)
	res = t.Unix() * 1e3
	return res
}

func StringToTime(req string) (res time.Time) {
	res, _ = time.Parse(time.RFC1123, req)
	return res
}

func GetTimeUnixMilli(t time.Time) int64 {
	return t.Unix() * 1e3
}

func ConvertDurationToTime(t *time.Time, duration *string, period *string) (res *time.Time, err error) {
	if period == nil || duration == nil {
		return nil, errors.New("Invalid Request Duration Or Input")
	}

	num, err := strconv.Atoi(*duration)
	if err != nil {
		return nil, err
	}

	var date time.Time
	timePeriod := strings.ToLower(*period)
	if strings.Contains(timePeriod, Year) {
		date = t.AddDate(num, 0, 0).UTC()
	} else if strings.Contains(timePeriod, Month) {
		date = t.AddDate(0, num, 0).UTC()
	} else if strings.Contains(timePeriod, Day) {
		date = t.AddDate(0, 0, num).UTC()
	} else if strings.Contains(timePeriod, Hour) {
		date = t.Add(time.Hour * time.Duration(num)).UTC()
	} else if strings.Contains(timePeriod, Minute) {
		date = t.Add(time.Minute * time.Duration(num)).UTC()
	} else if strings.Contains(timePeriod, Second) {
		date = t.Add(time.Second * time.Duration(num)).UTC()
	} else {
		return nil, errors.New("Invalid Request Duration Format")
	}

	return &date, nil
}

// Gets The Encoded Value
func Encode(input int64) string {
	encodeChars := "0123456789ABCDEFGHJKMNPQRSTUWXYZ"
	encodeLen := int64(len(encodeChars))
	input %= 1412010

	strData1 := ToBaseX(encodeChars, encodeLen, input/encodeLen)
	strData2 := ToBaseX(encodeChars, encodeLen, input%encodeLen)

	// Returns
	return strData1 + strData2
}

// Gets The Base Encoded Value
func ToBaseX(encodeChars string, encodeLen int64, encodeValue int64) string {
	var data1, data2, res string
	if encodeValue < 10 {
		return strconv.FormatInt(encodeValue, 10)
	} else if encodeValue < encodeLen {
		return encodeChars[encodeValue : encodeValue+1]
	} else {
		data1 = ToBaseX(encodeChars, encodeLen, encodeValue/encodeLen)
		data2 = ToBaseX(encodeChars, encodeLen, encodeValue%encodeLen)
		res = data1 + data2
	}

	// Returns
	return res
}

// GetPointerToString ...
func GetPointerToString(x string) *string {
	return &x
}

// GetPointerToTime ...
func GetPointerToTime(x time.Time) *time.Time {
	return &x
}

func GetDurationInSeconds(duration string) (float64, error) {
	if duration != "" {
		durationSplits := strings.Split(duration, ":")
		durationFormat := fmt.Sprintf("%sh%sm", durationSplits[0], durationSplits[1])
		if len(durationSplits) == 3 {
			durationFormat = fmt.Sprintf("%sh%sm%ss", durationSplits[0], durationSplits[1], durationSplits[2])
		}

		h, err := time.ParseDuration(durationFormat)
		if err != nil {
			return 0, err
		}
		return h.Seconds(), nil
	}
	return 0, nil
}

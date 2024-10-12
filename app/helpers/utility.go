package helpers

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func BoolToInt(payload bool) int {
	if payload {
		return 1
	} else {
		return 0
	}
}

func IntToBool(payload int) bool {
	if payload == 1 {
		return true
	} else {
		return false
	}
}

func StirngToInt8(payload string) int8 {
	value, err := strconv.ParseInt(payload, 10, 8)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	// Typecast to int8
	int8Value := int8(value)

	return int8Value
}

func ByteToKiloByte(bytes int64) float64 {
	return float64(bytes) / 1024.0
}

func AnyToString(payload interface{}) string {
	switch v := payload.(type) {
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func StringToDate(datestring string) (time.Time, error) {
	parserDate, err := time.Parse(time.RFC3339, datestring)

	if err != nil {
		return time.Time{}, err
	}

	return parserDate, nil
}

func DateFormatFromString(datetimestring string) string {
	parseDate, err := time.Parse(time.RFC3339, datetimestring)

	if err != nil {
		return "NA"
	}

	return parseDate.Format("02-01-2006")
}

func DateFormatFromStringToIso(datetimestring string) string {
	parseDate, _ := time.Parse(time.RFC3339, datetimestring)

	return parseDate.Format("2006-01-02")
}

func DateFormatFromIndToIso(datetime string) string {
	var year string = "1900"
	var month string = "01"
	var date string = "01"
	if datetime != "-" {
		year = datetime[6:10]
		month = datetime[3:5]
		date = datetime[0:2]
	}

	return year + "-" + month + "-" + date

}

func ConvertStringArrayToArray(datas string) []int {
	tempdata := datas
	tempdata = strings.Trim(tempdata, "[]")
	splitdatas := strings.Split(tempdata, ",")

	var datasarray []int

	for _, splitdata := range splitdatas {
		id, _ := strconv.Atoi(strings.TrimSpace(splitdata))
		datasarray = append(datasarray, id)
	}

	return datasarray
}

func GeneratRegisterNumber() int {
	rand.Seed(time.Now().UnixNano())

	min := 10000000
	max := 99999999

	randomNum := min + rand.Intn(max-min+1)

	return randomNum
}

func GeneratePassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789"

	//convert chartset to slice
	charsetRunes := []rune(charset)

	//create slice to randum string
	result := make([]rune, length)

	for i := range result {
		result[i] = charsetRunes[rand.Intn(len(charsetRunes))]
	}

	return string(result)
}

func SetMeta(page int, itemsPerPage int, total int64) MetaResponse {
	meta := MetaResponse{
		Total:       total,
		PerPage:     itemsPerPage,
		CurrentPage: page,
		FirstPage:   page,
		LastPage:    math.Ceil(float64(total/int64(itemsPerPage)) + 1),
	}

	return meta
}

func CalcultaeAge(birthdate string) int {

	birthday, _ := time.Parse(time.RFC3339, birthdate)
	now := time.Now()
	years := now.Year() - birthday.Year()

	// Adjust the age if the birthday has not occurred yet this year
	if now.Month() < birthday.Month() || (now.Month() == birthday.Month() && now.Day() < birthday.Day()) {
		years--
	}

	return years
}

func CalculateMonthAge(birthdate string) int {
	birthday, _ := time.Parse(time.RFC3339, birthdate)
	now := time.Now()
	years := now.Year() - birthday.Year()
	months := int(now.Month() - birthday.Month())

	if now.Day() < birthday.Day() {
		months--
	}

	totalMonths := years*12 + months
	return totalMonths
}

func CreateSlug(payload string) string {
	// Convert the input string to lowercase
	slug := strings.ToLower(payload)

	// Replace all spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove all non-alphanumeric characters (except hyphens)
	re := regexp.MustCompile(`[^\w-]+`)
	slug = re.ReplaceAllString(slug, "")

	// Trim any extra hyphens from the start or end
	slug = strings.Trim(slug, "-")

	return slug
}

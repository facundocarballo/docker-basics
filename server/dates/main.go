package dates

import (
	"fmt"
	"time"
)

func GetDateStringFromUint(bytesFecha []byte) *time.Time {
	strFecha := string(bytesFecha)

	fecha, err := time.Parse("2006-01-02 15:04:05", strFecha)
	if err != nil {
		fmt.Println("Error parsing the date. ", err)
		return nil
	}

	return &fecha
}

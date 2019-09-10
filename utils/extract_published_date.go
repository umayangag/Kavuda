package utils

import (
	"fmt"
	"time"
)

/*
return date from string or if the date layout is different return current date
 */
func ExtractPublishedDate(layout string, timeString string) time.Time {
	t, err := time.Parse(layout, timeString)
	if err != nil {
		fmt.Println("error in date", err)
		loc, _ := time.LoadLocation("UTC")
		t = time.Now().In(loc)
	}
	return t
}

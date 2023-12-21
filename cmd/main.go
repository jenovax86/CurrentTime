package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	time := time.Now()
	gmtFormat := time.Format(http.TimeFormat)
	gmt := strings.Split(gmtFormat, " ")[5]
	zone, _ := time.Zone()
	currentTime := fmt.Sprintf("%s %s %d %d %d:%d:%d %s %s (Iran standard time)",
		time.Weekday(), time.Month(), time.Day(), time.Year(), time.Hour(), time.Minute(), time.Second(), gmt, zone)
	fmt.Println(currentTime)
}

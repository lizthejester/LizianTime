package main

import (
	"fmt"
	"time"

	"github.com/Lizthejester/LizianTime/pkg/ltime"
)

func main() {
	currentYear, currentMonth, currentDay := time.Now().Date()
	lizMonth, lizDay := ltime.GetDayMonth(currentYear, currentMonth.String(), currentDay)
	fmt.Println("Current date:", lizMonth, lizDay, ltime.GetDayOfWeek(time.Now().YearDay()))

	testMonth, testDay := ltime.GetDayMonth(2023, "January", 10)
	fmt.Println("Test date:", testMonth, testDay, ltime.GetDayOfWeek(10, lizMonth))
}

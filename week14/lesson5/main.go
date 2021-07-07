package main

import (
	"fmt"
	"math/rand"
	"time"
)

const HOURS_IN_DAY = 24
const MINUTES_IN_HOUR = 60
const SECONDS_IN_MINUTE = 60

func dayToHours(days_amount int) int {
	return days_amount * HOURS_IN_DAY
}

func hoursToMinutes(hours_amount int) int {
	return hours_amount * MINUTES_IN_HOUR
}

func minutesToSeconds(minutes_amount int) int {
	return minutes_amount * SECONDS_IN_MINUTE
}

func main() {
	fmt.Println("Данная программа определяет, сколько часов в n днях")

	rand.Seed(time.Now().UnixNano())

	days := rand.Intn(100)
	fmt.Println("Количество дней:", days)
	hours := dayToHours(days)
	minutes := hoursToMinutes(hours)
	seconds := minutesToSeconds(minutes)
	fmt.Println("Количество часов:", seconds)
}
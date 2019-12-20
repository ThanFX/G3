package models

import (
	"database/sql"
	"log"
)

var currentDay int

type TimePeriod struct {
	Name   string
	Min    int
	Max    int
	InDays int
}

var Calendar map[string]TimePeriod

func SetCalendar() {
	Calendar = make(map[string]TimePeriod)
	Calendar["ten"] = TimePeriod{"Декада", 1, 3, 10}
	Calendar["month"] = TimePeriod{"Месяц", 1, 12, 30}
	Calendar["year"] = TimePeriod{"Год", 1, 10000, 360}
}

func ReadDate(DB *sql.DB) {
	var date int
	err := DB.QueryRow("select value from params where key='date'").Scan(&date)
	if err != nil {
		log.Fatal("ошибка парсинга записи времени : ", err)
	}
	SetDate(date)
}

func GetDate() int {
	return currentDay
}

func SetDate(date int) {
	currentDay = date
}

func IncDate() {
	currentDay++
}

func GetCalendarDate() map[string]int {
	c := make(map[string]int)
	curDate := GetDate() - 1
	//fmt.Println(curDate)
	c["year"] = int(curDate/Calendar["year"].InDays) + Calendar["year"].Min
	//fmt.Println(c["year"])
	curDate -= (c["year"] - Calendar["year"].Min) * Calendar["year"].InDays
	//fmt.Println(curDate)
	c["month"] = int(curDate/Calendar["month"].InDays) + Calendar["month"].Min
	curDate -= (c["month"] - Calendar["month"].Min) * Calendar["month"].InDays
	//fmt.Println(curDate)
	c["ten"] = int(curDate/Calendar["ten"].InDays) + Calendar["ten"].Min
	curDate -= (c["ten"] - Calendar["ten"].Min) * Calendar["ten"].InDays
	c["day"] = curDate + 1
	return c
}

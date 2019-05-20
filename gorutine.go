package main

import (
	"strconv"
	"time"
)

type Person struct {
	PersonId   int         `json: "personId"`
	Name       string      `json: "name"`
	State      string      `json: "state"`
	Health     float32     `json: "health"`
	Fatigue    float32     `json: "fatigue"`
	Hunger     float32     `json: "hunger"`
	Thirst     float32     `json: "thirst"`
	Somnolency float32     `json: "somnolency"`
	in chan string
	out chan string
}

type Fish struct {
	Id int
	StartSize int
	SpeedSize int
}

//func (p *Person) Start() (chan string, chan string){
//	return
//}


func tickerTime(time_out chan int) {
	ticker := time.NewTicker(time.Second)
	var i int = 0
	for range ticker.C {
		time_out <- i
		i++
	}
}

func get_ticks(n int, in chan int, out chan string) {
	out := make(chan string)
	for {
		select {
		case tick := <-in:
			out <- "Номер: " + strconv.Itoa(n) + ", тик: " + strconv.Itoa(tick)
		}
	}
}

func main() {
	time_out := make(chan int)
	var get_ticks_ch [10]chan int
	var get_ticks_ch_out [10]chan string

	for i := 0; i < 10; i++ {
		get_ticks_ch[i] = make(chan int)
		get_ticks_ch_out[i] = make(chan string)
		go get_ticks(i+1, get_ticks_ch[i])
	}

	go tickerTime(time_out)

	go func() {
		for cur_time := range time_out {
			for i := 0; i < 10; i++ {
				get_ticks_ch[i] <- cur_time
			}
		}
	}

	for ch_out := range get_ticks_ch_out {

	}
}
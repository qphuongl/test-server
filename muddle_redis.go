package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/func25/mathfunc/mathfunc"
)

var alphabet = []string{
	"Alain",
	"Bernard",
	"Celeste",
	"Didier",
	"Emile",
	"Francois",
	"Gaston",
	"Henri",
	"Isabelle",
	"Jean",
	"Kylian",
	"Louis",
	"Michel",
	"Nicolas",
	"Olivier",
	"Pierre",
	"Quentin",
	"Rene",
	"Stephane",
	"Thierry",
	"Ursule",
	"Valerie",
	"William",
	"Xavier",
	"Yvonne",
	"Zoe",
	"Amazing",
	"Magnetic",
	"Fire",
	"Magnificent",
	"Dynamic",
	"Doctor",
	"Wonder",
	"Super",
	"Awesome",
	"Kick-ass",
	"Bad-ass",
	"Alabaster",
	"Chocolate",
	"Vanilla",
	"Strawberry",
	"Astro",
	"Secret",
	"Space",
	"Orange",
	"Night",
	"Purple",
	"Jugemental",
	"Sticky",
	"Pretentious",
	"D-cup",
	"Bodacious",
}

func muddleRedis() {
	go incRedis("keyne")
	for i := 0; i < 10; i++ {
		go incRedis("keyne" + strconv.Itoa(i))
	}

	for {
		sleepTime, _ := mathfunc.RandInt(int(time.Millisecond), int(500*time.Millisecond))
		time.Sleep(time.Duration(sleepTime))
		readRedis()
		sleepTime, _ = mathfunc.RandInt(int(time.Millisecond), int(500*time.Millisecond))
		time.Sleep(time.Duration(sleepTime))
		writeRedis()
	}
}

func readRedis() {
	if _, err := redisClient.Get(muddleStrings()).Result(); err == nil {
	}
}

func writeRedis() {
	redisClient.Set(muddleStrings(), muddleStrings(), 5*time.Hour).Result()
}

func muddleStrings() string {
	s := ""
	lenStr, _ := mathfunc.RandInt(1, len(alphabet))
	for i := 0; i < lenStr; i++ {
		s += alphabet[i]
	}
	return s
}

func incRedis(keyne string) {
	i := 0
	for ; ; i++ {
		s := redisClient.Set(keyne, i, 0)
		if s.Err() != nil {
			fmt.Println("something wrong with redis", s.Err())
		}
		sleepTime, _ := mathfunc.RandInt(int(time.Millisecond), int(time.Second))
		time.Sleep(time.Duration(sleepTime))
	}
}

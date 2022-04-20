package main

import (
	"strconv"
	"time"

	"github.com/func25/batchlog"
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
	if _, err := redisClient.Get(muddleStrings()).Result(); err != nil {
	}
}

func writeRedis() {
	_, err := redisClient.Set(muddleStrings(), muddleStrings(), 5*time.Hour).Result()
	if err != nil {
		logger.Error().BatchErr(err).BatchMsg("[write-redis]" + err.Error())
	}

}

func muddleStrings() string {
	s := ""
	lenStr, _ := mathfunc.RandInt(1, len(alphabet))
	for i := 0; i < lenStr; i++ {
		randID, _ := mathfunc.Random0ToInt(len(alphabet))
		s += alphabet[randID]
	}
	return s
}

var logger = batchlog.NewLogger(batchlog.OptTimeout(time.Minute), batchlog.OptWait(5*time.Second), batchlog.OptMaxRelativeBatch(100))

func incRedis(keyne string) {
	i := 0
	for ; ; i++ {
		s := redisClient.Set(keyne, i, 0)
		if s.Err() != nil {
			logger.Error().Err(s.Err()).BatchMsg("[redis]" + s.Err().Error())
		}
		sleepTime, _ := mathfunc.RandInt(int(time.Millisecond), int(time.Second))
		time.Sleep(time.Duration(sleepTime))
	}
}

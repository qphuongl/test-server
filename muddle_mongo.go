package main

import (
	"context"
	"time"

	"github.com/func25/mathfunc/mathfunc"
	"github.com/func25/mongofunc/mongoquery"
	"github.com/func25/mongofunc/mongorely"
)

type Hero struct {
	Name  string `bson:"name"`
	Alias string `bson:"alias"`
}

func (Hero) GetMongoCollName() string {
	return "Heroes"
}

func muddleMongo() {
	for {
		sleepTime, _ := mathfunc.RandInt(int(time.Millisecond), int(500*time.Millisecond))
		time.Sleep(time.Duration(sleepTime))
		readMongo()
		sleepTime, _ = mathfunc.RandInt(int(time.Millisecond), int(500*time.Millisecond))
		time.Sleep(time.Duration(sleepTime))
		writeMongo()
		sleepTime, _ = mathfunc.RandInt(int(time.Millisecond), int(500*time.Millisecond))
		time.Sleep(time.Duration(sleepTime))
		updateMongo()
	}
}

func readMongo() {
	var x Hero
	filter := mongoquery.Init(
		mongoquery.Equal("name", muddleStrings()),
	)
	mongorely.Find(context.Background(), Hero{}.GetMongoCollName(), &x, filter)
}

func writeMongo() {
	_, err := mongorely.Create(context.Background(), Hero{
		Name:  muddleStrings(),
		Alias: muddleStrings(),
	})
	if err != nil {
		logger.Error().BatchErr(err).BatchMsg("[write-mongo]" + err.Error())
	}
}

func updateMongo() {
	filter := mongoquery.Init(
		mongoquery.Equal("name", muddleStrings()),
	)
	update := mongoquery.Init(
		mongoquery.Set(mongoquery.PairSetter{
			FieldName: "alias",
			Value:     muddleStrings(),
		}),
	)
	if _, err := mongorely.UpdateMany(context.Background(), Hero{}.GetMongoCollName(), filter, update); err != nil {
		logger.Error().BatchErr(err).BatchMsg("[update-mongo]" + err.Error())
	}
}

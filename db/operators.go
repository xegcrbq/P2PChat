package db

import "math/rand"

var names []string = []string{"Vasya228", "MadFox", "Cryptoinvestor3000"}

func GetRandomName() string {
	return names[rand.Int()%len(names)]
}

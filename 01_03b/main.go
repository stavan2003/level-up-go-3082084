package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("unable to read json file, err : ", err)
	}
	var res []raffleEntry
	if err = json.Unmarshal(data, &res); err != nil {
		log.Fatal("unmarshal error, err : ", err)
		return nil
	}
	return res
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}

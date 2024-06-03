package data

import (
	"os"

	"github.com/cassandramcc/songpoll/model"
	"github.com/gocarina/gocsv"
)

// GetDataFromCSV reads data from a CSV file and returns a slice of Artist structs
func GetDataFromCSV(csv string) []*model.Artist {
	in, err := os.Open(csv)
	if err != nil {
		panic(err)
	}
	defer in.Close()

	var artists []*model.Artist

	if err := gocsv.UnmarshalFile(in, &artists); err != nil {
		panic(err)
	}

	return artists
}

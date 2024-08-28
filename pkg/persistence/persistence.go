package persistence

import (
	"log"
	"os"
	"todo/pkg/task"

	"github.com/gocarina/gocsv"
)

//go:generate mockery --name Persistence
type Persistence interface {
	Save(records []task.Task) error
	Load() ([]task.Task, error)
}

type CSVPersistence struct {
}

func (persistence *CSVPersistence) Save(records []task.Task) error {
	file, err := os.OpenFile("data/data.csv", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = gocsv.MarshalFile(records, file)
	if err != nil {
		panic(err)
	}

	return err
}

func (persistence *CSVPersistence) Load() ([]task.Task, error) {
	file, err := os.Open("data/data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records := []task.Task{}
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		panic(err)
	}

	return records, err
}

//what about tests of Save() and Load()

// records = []task.Task{
// 	{Id: "2345", Description: "pick up junk", Done: false},
// 	{Id: "3469", Description: "eat", Done: false},
// }

package store

import "fmt"

// global var DB []Entry with key and value is struct with
//value and username

var DB = map[string]Entry{}

type Entry struct {
	Value    string
	Username string
}

// instantiate global struct

var NewEntry Entry

// define and initialize global channel i.e. queue
// for each func to put jobs onto
// can use closures to get function argument values into JobList

var JobList = make(chan func())

// Remove entry from DB helper func
func RemoveIndex(s []Entry, index int) []Entry {
	return append(s[:index], s[index+1:]...)
}

// Have getlistening func that pulls things off the channel. Channel is like a queue

func StartListening() {

	go func() {

		for {
			f := <-JobList
			f()
		}
	}()
}

// Get

func GetEntry(key string) string {

	// log 2 - this has been called and with what key file

	// need to push func onto JobList aka global channel/queue.
	// result channel gets result back into func
	result := make(chan string, 1)
	f := func() {

		result <- DB[key].Value

	}

	JobList <- f
	return <-result

	// when returns and what it returned with log file
}

// Get All entries needs to return a slice of entry structs
func GetAllEntries() map[string]Entry {

	// need to iterate through db and return key and values, not username

	result := make(chan map[string]Entry, 1)

	f := func() {

		// 	for k, v := range DB {
		// 		fmt.Printf(k, v.Value)
		// 	}
		// 	result <- k, v
		// }

		result <- DB
	}

	JobList <- f
	return <-result
}

//Add

func AddEntry(key string, value string, username string) {
	fmt.Println("Reached add entry handler")

	f := func() {

		for k, v := range DB {
			if k == key && v.Username == username {
				fmt.Println("Update DB")
				NewEntry.Value = value
				NewEntry.Username = username
				DB[key] = NewEntry
			} else {
				NewEntry.Value = value
				NewEntry.Username = username
				DB[key] = NewEntry
			}
		}
		// account for empty database
		if len(DB) == 0 {
			NewEntry.Value = value
			NewEntry.Username = username
			DB[key] = NewEntry
		}

	}
	JobList <- f
}

// Delete
func DeleteEntry(key string, username string) {

	f := func() {

		for k, v := range DB {
			if k == key && v.Username == username {
				fmt.Println("Deleting entry")
				delete(DB, key)
			} else {
				fmt.Printf("Permission denied")
				fmt.Println("Permission denied")

			}
		}

	}
	JobList <- f
}

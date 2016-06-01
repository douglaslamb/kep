package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/urfave/cli"
)

type Config struct {
	KepFile string `json:"kepfile"`
}

type Contact struct {
	LastName  string `json:"l"`
	FirstName string `json:"f"`
	Address   string `json:"a"`
	Phone     string `json:"p"`
	City      string `json:"c"`
	State     string `json:"s"`
	Country   string `json:"co"`
	Email     string `json:"e"`
	Note      string `json:"n"`
}

type ContactArray []Contact

func (c ContactArray) Len() int {
	return len(c)
}

func (c ContactArray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ContactArray) Less(i, j int) bool {
	return c[i].LastName < c[j].LastName
}

func main() {
	// 1. slurp config file
	jsonFile := loadConfig()
	jsonFile = strings.Replace(jsonFile, "~", os.Getenv("HOME"), -1)
	dat, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Println(jsonFile + " not found. Create " + jsonFile + " with the following form and put your contacts in the file. It's an array of JSON objects.\n" + `
	  [
	  {
	    "l": "Doe",
	    "f": "John",
	    "e": "jdoe@hell.com",
	    "a": "347 Bort Street",
	    "p": "555 555 5555",
	    "c": "Hot Shower Town",
	    "s": "Texiss",
	    "co": "Bungholia",
	    "n": "Nice boy with a big mouth."
	  }
	  ]`)
		return
	}
	contacts := []Contact{}
	if err := json.Unmarshal(dat, &contacts); err != nil {
		panic(err)
	}
	sort.Sort(ContactArray(contacts))

	app := cli.NewApp()
	app.Name = "Lek"
	app.Usage = "Manage contacts"
	app.Action = func(c *cli.Context) error {
		out := ""
		switch c.Args()[0] {
		case "all", "a":
			out = allContacts(contacts)
		case "fname", "f":
			out = byFirstName(contacts, c.Args()[1])
		case "lname", "l":
			out = byLastName(contacts, c.Args()[1])
		case "note", "n":
			out = byNote(contacts, c.Args()[1])
		}
		fmt.Print(out)
		return nil
	}

	app.Run(os.Args)
}

func allContacts(contacts []Contact) string {
	out := ""
	for i, contact := range contacts {
		out = out + fmt.Sprintf("%v  %v", i+1, formatContact(&contact))
	}
	return out
}

func byFirstName(contacts []Contact, firstName string) string {
	out := ""
	count := 1
	for _, contact := range contacts {
		if strings.ToLower(contact.FirstName) == strings.ToLower(firstName) {
			out = out + fmt.Sprintf("%v  %v", count, formatContact(&contact))
			count = count + 1
		}
	}
	return out
}

func byLastName(contacts []Contact, lastName string) string {
	out := ""
	count := 1
	for _, contact := range contacts {
		if strings.ToLower(contact.LastName) == strings.ToLower(lastName) {
			out = out + fmt.Sprintf("%v  %v", count, formatContact(&contact))
			count = count + 1
		}
	}
	return out
}

func byNote(contacts []Contact, note string) string {
	out := ""
	count := 1
	for _, contact := range contacts {
		if strings.Contains(strings.ToLower(contact.Note), strings.ToLower(note)) {
			out = out + fmt.Sprintf("%v  %v", count, formatContact(&contact))
			count = count + 1
		}
	}
	return out
}

func formatContact(contact *Contact) string {
	var out string = ""
	if contact.FirstName != "" {
		out = out + contact.FirstName + "\n   "
	}
	if contact.LastName != "" {
		out = out + contact.LastName + "\n   "
	}
	if contact.Email != "" {
		out = out + contact.Email + "\n   "
	}
	if contact.Address != "" {
		out = out + contact.Address + "\n   "
	}
	if contact.Phone != "" {
		out = out + contact.Phone + "\n   "
	}
	if contact.City != "" {
		out = out + contact.City + "\n   "
	}
	if contact.State != "" {
		out = out + contact.State + "\n   "
	}
	if contact.Country != "" {
		out = out + contact.Country + "\n   "
	}
	if contact.Note != "" {
		out = out + contact.Note + "\n   "
	}
	out = out + "\n"
	return out
}

func loadConfig() string {
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.keprc")
	if err != nil {
		fmt.Println("~/.keprc not found. Creating ~/.keprc with default settings.\n")
		err := ioutil.WriteFile(os.Getenv("HOME")+"/.keprc", []byte("{\n\"kepfile\": \"~/.kep.json\"\n}"), 0644)
		if err != nil {
			panic(err)
		}
		dat, err = ioutil.ReadFile(os.Getenv("HOME") + "/.keprc")
		if err != nil {
			panic(err)
		}
	}
	config := Config{}
	json.Unmarshal(dat, &config)
	return config.KepFile
}

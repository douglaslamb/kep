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

// alright so I want to make a thing that slurps some json and then gives you some options re sorting them like contacts. like I want to be able to look up a person by like last name, like with "kt find last-name first-name" or kt name last-name first-name or kt name last-name or kit knt knt knt knt knit knt kt knt kt I dunno.
// so I want
// kt name last-name
// kt name first-name
// kt addr address
// kt city city (also accepts a zip)
// kt state state
// kt cntry country
// kt note
// kt email email-addrese
// kt phone

// parts of this program are
// 0.5 need to slurp the config file first
// 1. slurping json
// 2. filtering the json array
// 3. printing the contents of the array to stdout

type Config struct {
	KtFile string `json:"ktfile"`
}

type Contact struct {
	LastName  string `json:"lastname"`
	FirstName string `json:"firstname"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	Email     string `json:"email"`
	Note      string `json:"note"`
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
		panic(err)
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
		case "all":
			out = allContacts(contacts)
		case "fname":
			out = byFirstName(contacts, c.Args()[1])
			// print only with matching first name
		case "lname":
			out = byLastName(contacts, c.Args()[1])
			// print only with matching last name
		case "note":
			// print only with matching note content
			// TO BE WRITTEN
			// probably make this a regex search
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
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ktrc")
	if err != nil {
		panic(err)
	}
	config := Config{}
	json.Unmarshal(dat, &config)
	return config.KtFile
}

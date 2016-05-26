package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	app := cli.NewApp()
	app.Name = "Lek"
	app.Usage = "Manage contacts"
	app.Action = func(c *cli.Context) error {
		if c.Args()[0] == "all" {
			// print all contacts option
			printAll(contacts)
		}
		return nil
	}

	app.Run(os.Args)
}

func printAll(contacts []Contact) {
	for _, contact := range contacts {
		fmt.Println(contact.FirstName)
		fmt.Println(contact.LastName)
		fmt.Println()
	}
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

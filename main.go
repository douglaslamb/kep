package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	// 1. slurp config file
	jsonFile := loadConfig()
	fmt.Println(jsonFile)
	dat, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}
	contacts := string(dat)
	//contactsArr, file := filepath.Split(contacts)
	//fmt.Println(contactsArr)
	fmt.Println(os.Getenv("HOME"))
	/*
		app := cli.NewApp()
		app.Name = "boom"
		app.Usage = "make a fart"
		app.Action = func(c *cli.Context) error {
			fmt.Println("farts are retarded")
			return nil
		}

		app.Run(os.Args)
	*/
}

func loadConfig() string {
	dat, err := ioutil.ReadFile(os.Getenv("HOME") + "/.ktrc")
	if err != nil {
		panic(err)
	}
	config := Config{}
	json.Unmarshal(dat, &config)
	//fmt.Println(string(dat))
	return config.KtFile
}

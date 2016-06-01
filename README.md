# kep

Kep is a command-line tool for managing contacts written in Go by Douglas Lamb. You maintain JSON array of objects in a text file, and kep enables you to look up one or more of your contacts and print their information to the console.

## Install from source

1. Install Golang 
2. Set your GOPATH and put GOPATH in your PATH. Search Google for help if needed.
2. Clone repo `git clone https://github.com/douglaslamb/kep`
3. Install the CLI library `go get github.com/urfave/cli` Hopefully it is still around.
4. Navigate to the kep repo you cloned and `go install`

## Setup

Run `kep` with no arguments. Kep will create .keprc in your home folder. The default .keprc specifies ~/.kep.json as your contact file. kep will instruct you to create ~/.kep.json and will print the JSON schema to the console. Create ~/.kep.json with at least one contact. You do not need to fill in every property. Just the first and last name will suffice.

Here is how it should look:

```
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
]
```

## Commands 

### kep a

`kep a` prints all of the contacts in alphabetical order.

### kep f aFirstName
`kep f aFirstName` prints all of the contacts with the first name aFirstName in alphabetical order. For example, `kep f Sara` prints all of the contacts with the first name "Sara." It is not case-sensitive. The input string cannot have any spaces.

### kep l aLastName
`kep l aLastName` prints all of the contacts with the last name aLastName in alphabetical order. For example, `kep f Tundly` prints all of the contacts with the last name "Tundly." It is not case-sensitive. The input string cannot have any spaces.

### kep n aWordInNote
`kep n aWordInNote` prints all of the contacts with the aWordInNote in their note in alphabetical order. For example, `kep n butthole` prints all of the contacts with "butthole" in their notes. It is not case-sensitive. The input string cannot have any spaces.

## Attributions

This software uses Go https://golang.org and cli https://github.com/urfave/cli. 

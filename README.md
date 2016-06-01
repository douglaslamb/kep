# Kep

Kep is a contact management command-line tool written in Go.

## 
This is a REST API written in Go by Douglas Lamb. It responds to http GET requests for a list of businesses with pagination and for individual businesses specified by an ID.

## Run

First create a folder called "resources" in the root of the repo. Place "engineering_project_businesses.csv" in the "resources" folder. Run `go run main.go structs.go` from the root of the repo. This is not necessary if you intend to install. Please see example_engineering_project_businesses.csv for an example of valid input.

## Install

First create a folder called "resources" in the root of the repo. Place "engineering_project_businesses.csv" in the "resources" folder. Run `go install` from the root of the repo. Execute the program by running `ownlocal-api` from anywhere within the system. Ensure that GOPATH/bin has been added to your PATH. Please see example_engineering_project_businesses.csv for an example of valid input.

## Usage

As written it hosts the API at http://localhost:8080 . It responds to two types of GET requests.

### GET /businesses

Running http://localhost:8080/businesses from the server will return a JSON object containing an array of businesses and a "links" object. The businesses are paginated. The default page number is 1 and the default page size is 50. Hence http://localhost:8080/businesses will return the first 50 businesses.

The client may specify the page and/or page size. For example, http://localhost:8080?page=20&perPage=60 . This request will return the 20th page with 60 businesses per page.

The "links" object contains four strings which correspond to the first, previous, next, and last pages respectively. The client may use these links to traverse the collection.

### GET /business/{id}

Running http://localhost:8080/business/{id} from the server will return a JSON object representing the business with id {id}. For example http://localhost:8080/business/584 returns the business with id 584. Ids must be integers, zero or greater.

## Testing

Start the server by following the instructions in the run or install section above. Then run `go test` from the root of the repo.

## Attributions

This software uses Go (https://golang.org) and httprouter (https://github.com/julienschmidt/httprouter). This software compiles on OSX but has not been tested on Linux.

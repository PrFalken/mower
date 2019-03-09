package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "file", "input.txt", "handle instructions via file")
	doAPI := flag.Bool("api", false, "handle instructions via API")
	flag.Parse()

	if *doAPI {
		router := mux.NewRouter()
		router.HandleFunc("/", handleMowerJob).Methods("POST")
		router.HandleFunc("/", hello).Methods("GET")

		log.Fatal(http.ListenAndServe(":8000", router))
	}

	handleFile(fileName)

}

type instructions struct {
	Commands string `json:"commands,omitempty"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func handleMowerJob(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var instructions instructions
	err = json.Unmarshal(b, &instructions)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := executeMowers(strings.NewReader(instructions.Commands))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte(output))
}

func handleFile(f string) {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	output, err := executeMowers(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

func executeMowers(r io.Reader) (result string, err error) {
	lawn := lawn{}
	err = lawn.parseInput(r)
	lawn.mow()
	for _, mower := range lawn.mowers {
		result += fmt.Sprintf("%v %v %s\n", mower.xPos, mower.yPos, mower.orientation)
	}
	return result, nil

}

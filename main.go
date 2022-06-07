package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
    http.HandleFunc("/parse", handler)
    http.HandleFunc("/tokenize", handler2)
    http.HandleFunc("/genast", handler3)
    http.HandleFunc("/codegen", handler4)
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal(err)
        return
    }
}
func handler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	cmd := exec.Command("./c9", "-cst", "--json")

	cmd.Stdin = strings.NewReader(string(reqBody))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		w.Header().Set("Content-Type", "plain/text")
		fmt.Fprint(w, "error")
	}

	w.Header().Set("Content-Type", "plain/text")
	fmt.Fprint(w, out.String())
}

func handler2(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	cmd := exec.Command("./c9", "-lex", "--json")

	cmd.Stdin = strings.NewReader(string(reqBody))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		w.Header().Set("Content-Type", "plain/text")
		fmt.Fprint(w, "error")
	}

	w.Header().Set("Content-Type", "plain/text")
	fmt.Fprint(w, out.String())
}

func handler3(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	cmd := exec.Command("./c9", "-ast", "--json")

	cmd.Stdin = strings.NewReader(string(reqBody))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		w.Header().Set("Content-Type", "plain/text")
		fmt.Fprint(w, "error")
	}

	w.Header().Set("Content-Type", "plain/text")
	fmt.Fprint(w, out.String())
}

func handler4(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	cmd := exec.Command("./c9", "--codegen", "--json")

	cmd.Stdin = strings.NewReader(string(reqBody))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		w.Header().Set("Content-Type", "plain/text")
		fmt.Fprint(w, "error")
	}

	w.Header().Set("Content-Type", "plain/text")
	fmt.Fprint(w, out.String())
}

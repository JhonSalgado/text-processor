package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	originFolder      = "./stopwords/"
	destinationFolder = "./processor/stopwords/"
)

func mainw() {

	// open origin folder
	files, err := ioutil.ReadDir(originFolder)
	if err != nil {
		log.Fatal(err)
	}

	// create init file
	initFile, err := os.Create(destinationFolder + "init.go")
	if err != nil {
		log.Fatal(err)
	}
	initWriter := bufio.NewWriter(initFile)
	defer initFile.Close()

	// write init
	initWriter.WriteString("package stopwords\n\n")
	initWriter.WriteString("var Stopwords map[string][]string = make(map[string][]string)\n\n")
	initWriter.WriteString("func init() {\n")

	// for each file in the folder
	for _, f := range files {

		// open origin file
		originFile, err := os.Open(originFolder + f.Name())
		if err != nil {
			log.Fatal(err)
		}
		fileScanner := bufio.NewScanner(originFile)
		defer originFile.Close()

		// create and open destination file
		destinationFileName := strings.Replace(f.Name(), "txt", "go", 1)
		destinationFile, err := os.Create(destinationFolder + destinationFileName)
		if err != nil {
			log.Fatal(err)
		}
		w := bufio.NewWriter(destinationFile)
		defer destinationFile.Close()

		// start to write the go file
		lang := strings.Replace(f.Name(), ".txt", "", 1)
		w.WriteString("package stopwords\n\n")

		w.WriteString(fmt.Sprintf("var %s = []string{\n", lang))
		for fileScanner.Scan() {
			line := fileScanner.Text()
			w.WriteString(fmt.Sprintf("\t\"%s\",\n", line))
		}
		fmt.Printf("%s written\n", f.Name())
		w.WriteString("}\n")
		w.Flush()

		// write the stopword list in init.go
		initWriter.WriteString(fmt.Sprintf("\tStopwords[\"%s\"] = %s\n", lang, lang))
	}
	initWriter.WriteString("}\n")
	initWriter.Flush()
}

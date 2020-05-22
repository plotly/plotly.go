package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/plotly/plotly.go"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app       = kingpin.New("json_plotter", "plot json files from the command line")
	inputFile = app.Flag("input", "Input json file. Defaults to STDIN.").Short('i').String()
	name      = app.Flag("name", "The filename for the plot in plotly. Include any folders.").Short('n').Default(fmt.Sprint(time.Now().Unix())).String()
	download  = app.Flag("download", "Download the plot automatically.").Short('d').Bool()
	outFile   = app.Flag("out", "File name for the downloaded image. Defaults to the same as plotly name in current folder.").Short('o').String()
	private   = app.Flag("private", "Make this plot private.").Short('p').Bool()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	var inputReader io.Reader
	var figure *grob.Fig
	if *inputFile == "" {
		inputReader = os.Stdin
	} else {
		fileReader, err := os.Open(*inputFile)
		check(err, "Could not open input file: "+*inputFile)
		inputReader = fileReader
	}
	inputData, err := ioutil.ReadAll(inputReader)
	check(err, "Error while reading data.")
	err = json.Unmarshal(inputData, &figure)

	//tempLayout, _ := json.Marshal(figure.Layout)
	//figure.Layout = string(tempLayout)

	check(err, "Error while processing data. Should contain a 'data' and 'layout' element only.")
	log.Printf("%#v", figure)
	isPublic := !*private
	url, err := plotly.Create(*name, figure, isPublic)
	check(err, "Error while POSTing to plot.ly.")
	fmt.Println(url)

	outF := *outFile
	if *download {
		if outF == "" {
			outF = path.Base(*name) + ".png"
		}
		err = plotly.Save(url.Id(), outF)
		check(err, "Error while downloading image.")
	}
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

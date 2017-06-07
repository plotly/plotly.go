package main

import (
	"encoding/json"
	"fmt"
	"github.com/plotly/go-api/plotly"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"
)

var inputFile string
var name string
var download bool
var fileName string
var public bool

var plotCmd = &cobra.Command{
	Use:   "json_plotter",
	Short: "Plot the given json data on plot.ly",
	Long:  "Please see the documentation for plot.ly for the specifications of the json. And https://github.com/baruchlubinsky/go-plotly for instructions on how to connect to your plot.ly account.",
	Run:   plot,
}

func init() {
	flags := plotCmd.Flags()
	flags.StringVarP(&inputFile, "input", "i", "", "Input json file. Defaults to STDIN.")
	flags.StringVarP(&name, "name", "n", fmt.Sprint(time.Now().Unix()), "The filename for the plot in plotly. Include any folders.")
	flags.BoolVarP(&download, "download", "d", false, "Download the plot automatically.")
	flags.StringVarP(&fileName, "image", "o", "", "File name for the downloaded image. Defaults to the same as plotly name in current folder.")
	flags.BoolVarP(&public, "public", "p", true, "Make this plot publicly visible. Free accounts limit the number of private plots.")
}

func plot(cmd *cobra.Command, args []string) {
	var inputReader io.Reader
	var figure plotly.Figure
	if inputFile == "" {
		inputReader = os.Stdin
	} else {
		fileReader, err := os.Open(inputFile)
		check(err, "Could not open input file: "+inputFile)
		inputReader = fileReader
	}
	inputData, err := ioutil.ReadAll(inputReader)
	check(err, "Error while reading data.")
	err = json.Unmarshal(inputData, &figure)

	//tempLayout, _ := json.Marshal(figure.Layout)
	//figure.Layout = string(tempLayout)

	check(err, "Error while processing data. Should contain a 'data' and 'layout' element only.")
	url, err := plotly.Create(name, figure, public)
	check(err, "Error while POSTing to plot.ly.")
	fmt.Println(url)
	if download {
		if fileName == "" {
			fileName = path.Base(name) + ".png"
		}
		err = plotly.Save(url.Id(), fileName)
		check(err, "Error while downloading image.")
	}
}

func main() {
	plotCmd.Execute()
}

func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

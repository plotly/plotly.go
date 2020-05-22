package main

import (
	"fmt"
	"strings"

	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
	"github.com/plotly/plotly.go"
)

// This program creates a plot on plotly, retrieves it, and saves it as an image.
func main() {
	f := &grob.Fig{
		Data: grob.Traces{
			&grob.Scatter{
				Type: grob.TraceTypeScatter,
				X:    []float64{4.54, 3, 34, 35, 362},
				Y:    []int{1, 2, 3, 4, 5},
			},
		},
		Layout: &grob.Layout{},
	}
	result, err := plotly.SaveFig(f, "new golang file")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Successfully created plot!\nFilename: %v\nURL: %v\n", result.Filename, result.Url)
	}

	fields := strings.Split(result.Url, "/")
	id := fields[4]
	response, err := plotly.Get(id)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Downloaded plot")
	}
	err = plotly.Download(response.Payload.Figure, "image.png")
	if err != nil {
		fmt.Println(err)
	}
}

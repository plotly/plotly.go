package plotly

import (
	grob "github.com/MetalBlueberry/go-plotly/graph_objects"
)

func SaveFig(f *grob.Fig, filename string) (*PostResponse, error) {
	req := NewRequest()
	req.Filename = filename
	req.Figure = f
	req.Origin = "plot"
	res, err := Post(req)
	return res, err
}

package plotly

import (
	"encoding/json"
	"fmt"
)

func Create(filename string, figure Figure, public bool) (url Url, err error) {
	request := NewRequest()
	request.Origin = "plot"
	args, err := json.Marshal(figure.Data)
	if err != nil {
		return
	}
	request.Args = string(args)
	request.Kwargs = fmt.Sprintf(`{"filename":"%v",
        "fileopt":"overwrite",
        "world_readable":%v,
        "layout":%v
  }`, filename, public, figure.Layout)
	result, err := Post(request)
	if err != nil {
		return
	}
	if result.Url == "" {
		return Url(""), result
	}
	return Url(result.Url), nil
}

func Save(id string, filename string) error {
	response, err := Get(id)
	if err != nil {
		return err
	} else if response.Payload.Figure.Data == nil {
		return response
	}
	err = Download(response.Payload.Figure, filename)
	return err
}

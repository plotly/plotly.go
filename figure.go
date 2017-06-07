package plotly

import (
	"encoding/json"
	"fmt"
)

type Figure struct {
	Layout Layout  `json:"layout,omitempty"`
	Data   []Trace `json:"data,omitempty"`
}

func (f *Figure) Save(filename string) (*PostResponse, error) {
	req := NewRequest()
	req.Filename = filename
	req.Figure = f
	req.Origin = "plot"
	res, err := Post(req)
	return res, err
}
func (f *Figure) Overwrite(fid string) error {
	return nil
}

type Layout struct {
	Title  *string `json:"title,omitempty"`
	Width  *int64  `json:"width,omitempty"`
	Height *int64  `json:"width,omitempty"`
}

// S returns a pointer to the specified string
func S(in string) *string {
	return &in
}

// I returns a pointer to the specified int64
func I(in int64) *int64 {
	return &in
}

// F returns a pointer to the specified float64
func F(in float64) *float64 {
	return &in
}

type Data []*Trace

type Trace struct {
	Type   string   `json:"type"`
	X      Array    `json:"x,omitempty"`
	Y      Array    `json:"y,omitempty"`
	Z      Array    `json:"z,omitempty"`
	R      Array    `json:"r,omitempty"`
	T      Array    `json:"t,omitempty"`
	Mode   *string  `json:"mode,omitempty"`
	Name   *string  `json:"name,omitempty"`
	Text   Array    `json:"text,omitempty"`
	ErrorX *ErrorBar `json:"error_x,omitempty"`
	ErrorY *ErrorBar `json:"error_y,omitempty"`
	Marker *Marker   `json:"marker,omitempty"`
	// 	YAxis  string        `json:"yaxis"`
	//Line   Line          `json:"line,omitempty"`
	TextPosition *string `json:"textposition,omitempty"`
	//TextFont Font `json:"textfont,omitempty"`
	Fill      *string `json:"fill,omitempty"`
	FillColor *string `json:"fillcolor,omitempty"`
}

type ErrorBar struct {
	Type       *string  `json:"type,omitempty"`
	Symmetric  *bool    `json:"symmetric,omitempty"`
	Array      Array    `json:"array,omitempty"`
	Value      *float64 `json:"value,omitempty"`
	ArrayMinus Array    `json:"arrayminus,omitempty"`
	ValueMinus *float64 `json:"valueminus,omitempty"`
	Color      *string  `json:"color,omitempty"`
	Thickness  *float64 `json:"thickness,omitempty"`
	Width      *float64 `json:"width,omitempty"`
	Opacity    *float64 `json:"opacity,omitempty"`
	CopyYStyle *bool    `json:"copy_ystyle,omitempty"`
	Visible    *bool    `json:"visible,omitempty"`
}

type Marker struct {
	Color      StringOrList `json:"color,omitempty"`
	Size       FloatOrList  `json:"size,omitempty"`
	Symbol     StringOrList `json:"symbol,omitempty"`
	Line       Line         `json:"line,omitempty"`
	Opacity    *float64     `json:"opacity,omitempty"`
	SizeRef    *float64     `json:"sizeref,omitempty"`
	SizeMode   string       `json:"sizemode,omitempty"`
	ColorScale ColorScale   `json:"colorscale,omitempty"`
	CAuto      *bool        `json:"cauto,omitempty"`
	CMin       *float64     `json:"cmin,omitempty"`
	CMax       *float64     `json:"cmax,omitempty"`
}

type Font struct {
}

type Line struct {
	Color *string `json:"color,omitempty"`
}

// type Trace struct {
// 	X      []interface{} `json:"x"`
// 	Y      []interface{} `json:"y,omitempty"`
// 	Name   string        `json:"name"`
// 	YAxis  string        `json:"yaxis"`
// 	Type   string        `json:"type"`
// 	Marker Marker        `json:"marker,omitempty"`
// 	Line   Line          `json:"line,omitempty"`
// 	Fill   string        `json:"fill,omitempty"`
// }

// }

type StringOrList struct {
	String string
	List   []string
}

func (v *StringOrList) MarshalJSON() ([]byte, error) {
	if v.String != "" {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(v.List)
	}
}

func (v *StringOrList) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return nil
	}

	firstChr := string(input[0])
	if firstChr == "[" {
		return json.Unmarshal(input, v.List)
	} else if firstChr == "n" { // probably "null"
		return nil
	} else {
		return json.Unmarshal(input, v.String)
	}
	// Doesn't handle `true`, `false` or nested objects.
}

type FloatOrList struct {
	Float float64
	List  []float64
}

func (v *FloatOrList) MarshalJSON() ([]byte, error) {
	if v.List == nil {
		return json.Marshal(v.Float)
	} else {
		return json.Marshal(v.List)
	}
}

func (v *FloatOrList) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return nil
	}

	firstChr := string(input[0])
	if firstChr == "[" {
		return json.Unmarshal(input, v.List)
	} else if firstChr == "n" { // probably "null"
		return nil
	} else {
		return json.Unmarshal(input, v.Float)
	}
	// Doesn't handle `true`, `false` or nested objects.
}

type ColorScale struct {
	Preset string
	Custom []ColorStop
}

// ColorStop is a stop-point in a ColorScale (under field Custom)
type ColorStop struct {
	Position float64
	Color    string
}

func (v *ColorScale) MarshalJSON() ([]byte, error) {
	if v.Custom == nil {
		return json.Marshal(v.Preset)
	} else {
		return json.Marshal(v.Custom)
	}
}

func (v *ColorScale) UnmarshalJSON(input []byte) error {
	if len(input) == 0 {
		return nil
	}

	firstChr := string(input[0])
	if firstChr == "[" {
		// TODO: unmarshal to a list of lists, and go through that
		// to create ColorStop objects
		var dump [][]interface{}
		if err := json.Unmarshal(input, dump); err != nil {
			return fmt.Errorf("plotly.ColorScale: %s", err)
		}

		for _, subList := range dump {
			if len(subList) != 2 {
				return fmt.Errorf("ColorScale::UnmarshalJSON error: an element doesn't have exactly two elements: %#v", subList)
			}
			pos, ok := subList[0].(float64)
			if !ok {
				return fmt.Errorf("ColorScale::UnmarshalJSON error: first part of a custom stop isn't numeric: %#v", subList[0])
			}

			color, ok := subList[1].(string)
			if !ok {
				return fmt.Errorf("ColorScale::UnmarshalJSON error: first part of a custom stop isn't numeric: %#v", subList[0])
			}

			v.Custom = append(v.Custom, ColorStop{pos, color})
		}

		return nil

	} else if firstChr == "n" { // probably "null"

		return nil

	} else {

		return json.Unmarshal(input, v.Preset)

	}
	// Doesn't handle `true`, `false` or nested objects.
}

package plotly

import (
	"fmt"
	"sort"
)

// Create a stacked bar plot.
// Data contains one entry for each point on the x-axis, refering to a slice with
// values for each of the categories. Categories are sorted.
func StackedBarPlot(categories []string, colors []string, data map[string][]interface{}, filename string, title string, xTitle string, yTitle string, public bool) (Url, error) {
	traces := make([]Trace, 0, len(data))
	for i, category := range categories {
		x := make([]interface{}, 0)
		y := make([]interface{}, 0)
		keys := make([]string, 0)
		// for key, values := range data {
		// 	x = append(x, key)
		// 	y = append(y, values[i])
		// }
		for key := range data {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			x = append(x, key)
			y = append(y, data[key][i])
		}
		traces = append(traces, Trace{
			X:    x,
			Y:    y,
			Name: &category,
			Type: "bar",
			Marker: &Marker{
				Color: StringOrList{String: colors[i]},
			},
		})
	}
	_ = fmt.Sprintf(`{
    "title":"%v",
    "barmode":"stack",
    "yaxis":{
      "title":"%v"
    },
    "xaxis":{
      "title":"%v",
      "type":"category"
    }
  }
  `, title, yTitle, xTitle)
	figure := Figure{
		Data: traces,
	}
	return Create(filename, figure, public)
}

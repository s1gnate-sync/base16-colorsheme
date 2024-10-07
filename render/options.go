package render

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"colorscheme/hsluv"
)

func makeVars(name string, colors []string, sel int) map[string]string {
	vars := map[string]string{
		"name": name,
	}

	for index, hex := range colors {
		vars[fmt.Sprintf("base%d", index)] = hex

		if index == sel {
			vars["index"] = fmt.Sprintf("%d", index)
			vars["color"] = hex
		}
	}

	return vars
}

var tpl *template.Template

func hsl(hh, ss, ll float64, hex string) string {
	h, s, l := hsluv.HsluvFromHex("#" + hex)

	h += hh
	if h > 255 {
		h = 255
	} else if h < 0 {
		h = 0
	}

	s += ss
	if s > 255 {
		s = 255
	} else if s < 0 {
		s = 0
	}

	l += ll

	if l > 255 {
		l = 255
	} else if l < 0 {
		l = 0
	}

	return strings.TrimPrefix(hsluv.HsluvToHex(h, s, l), "#")
}

func Init(str string) {
	if tpl != nil {
		return
	}

	var err error
	tpl, err = template.New("render").Funcs(map[string]interface{}{
		"hsl": hsl,
		"s": func(v float64, hex string) string {
			return hsl(0, v, 0, hex)
		},
		"l": func(v float64, hex string) string {
			return hsl(0, 0, v, hex)
		},
		"h": func(v float64, hex string) string {
			return hsl(v, 0, 0, hex)
		},
		"nl": func() string {
			return "\n"
		},
		"tab": func() string {
			return "\t"
		},
	}).Delims("{", "}").Parse(str)
	if err != nil {
		panic(err)
	}
}

func Render(name string, colors []string, sel int) {
	tpl.Execute(os.Stdout, makeVars(name, colors, sel))
}

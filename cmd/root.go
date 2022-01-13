package cmd

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/infrable-io/logo/svg"
	"github.com/spf13/cobra"
)

var invert bool
var size int

func init() {
	rootCmd.Flags().BoolVarP(&invert, "invert", "i", false, "whether to invert the colorscheme")
	rootCmd.Flags().IntVarP(&size, "size", "s", 500, "size of the generated logo")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "logo",
	Short: "Generates the Infrable logo.",
	Long: `
	Logo is a simple CLI that programmatically generates the Infrable logo.
	The output is an SVG file, which can be converted to any other image file
	format.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		//           bg
		//      -------------
		//      \___     ___/
		//      fg -|-> |
		//           \  |
		//            + |
		//             \|<- stroke
		//  stoke ->|\
		//          | +
		//          |  \
		//          | <-|- fg
		//       ---     ---
		//      /____________\
		//

		// set colorscheme
		bg := "#ffffff"
		fg := "#000000"
		stroke := "#ffffff"

		if invert == true {
			bg = "#000000"
			fg = "#ffffff"
			stroke = "#000000"
		}

		// set attributes
		strokeWidth := float64(size) / 30.0

		s := svg.SVG{
			XMLName:      "svg",
			XMLNamespace: "http://www.w3.org/2000/svg",
			Version:      "2.0",
			Width:        float64(size),
			Height:       float64(size),
		}

		// generate background
		s.AppendElement(svg.Rect{
			XMLName: "rect",
			X:       0,
			Y:       0,
			Width:   float64(size),
			Height:  float64(size),
			Fill:    bg,
		})

		//             |--- size --|
		//    c        |- a -|- d -|
		//  - -  -------------
		//  | _  \___     ___/       Slope (m): 2
		//  b        |   |           Ratio (a/b): 0.8
		//  |         \  |           a = size / 1 + sqrt(5)
		//  -          + |           b = a / 0.8
		//              \|           c = b / 2.5
		//
		//       | c | c | c |
		//
		// Note: The golden ratio, (1 + sqrt(5)) / 2, is used to determine a.

		center := float64(size) / 2.0
		a := float64(size) / (1.0 + math.Sqrt(5.0))
		b := a / 0.8
		c := b / 2.5
		ppts1 := []svg.Point{
			P1(a, b, center, false),
			P2(a, b, center, false),
			P3(a, b, c, center, false),
			P4(a, b, c, center, false),
			P5(a, b, c, center, false),
			P6(a, b, c, center, false),
			P7(a, b, c, center, false),
			P8(a, b, c, center, false),
		}

		//           |\
		//           | +
		//           |  \
		//           |   |
		//        ---     ---
		//       /____________\

		ppts2 := []svg.Point{
			P1(a, b, center, true),
			P2(a, b, center, true),
			P3(a, b, c, center, true),
			P4(a, b, c, center, true),
			P5(a, b, c, center, true),
			P6(a, b, c, center, true),
			P7(a, b, c, center, true),
			P8(a, b, c, center, true),
		}

		p1 := svg.Polygon{
			XMLName:     "polygon",
			Fill:        fg,
			Stroke:      stroke,
			StrokeWidth: strokeWidth,
		}

		p2 := svg.Polygon{
			XMLName:     "polygon",
			Fill:        fg,
			Stroke:      stroke,
			StrokeWidth: strokeWidth,
		}

		for _, pt := range ppts1 {
			p1.AppendPoint(pt)
		}

		for _, pt := range ppts2 {
			p2.AppendPoint(pt)
		}

		s.AppendElement(p2)
		s.AppendElement(p1)

		out, err := svg.Marshal(s)
		if err != nil {
			fmt.Printf("An error occurred: %v", err)
			return
		}

		err = ioutil.WriteFile("logo.svg", out, 0644)
		if err != nil {
			fmt.Printf("An error occurred.")
			return
		}
	},
}

func P1(a, b, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
	}
	return svg.Point{X: center - a, Y: center - b}
}

func P2(a, b, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
	}
	return svg.Point{X: center + a, Y: center - b}
}

func P3(a, b, c, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
		c = -c
	}
	return svg.Point{X: center + a - 0.5*c, Y: center - b + c}
}

func P4(a, b, c, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
		c = -c
	}
	return svg.Point{X: center + a - 1.5*c, Y: center - b + c}
}

func P5(a, b, c, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
		c = -c
	}
	return svg.Point{X: center + a - 1.5*c, Y: center + c}
}

func P6(a, b, c, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
		c = -c
	}
	return svg.Point{X: center - a + 1.5*c, Y: center - c}
}

func P7(a, b, c, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
		c = -c
	}
	return svg.Point{X: center - a + 1.5*c, Y: center - b + c}
}

func P8(a, b, c, center float64, mirror bool) svg.Point {
	if mirror {
		a = -a
		b = -b
		c = -c
	}
	return svg.Point{X: center - a + 0.5*c, Y: center - b + c}
}

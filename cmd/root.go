package cmd

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/infrable-io/logo/svg"
	"github.com/spf13/cobra"
)

var size int

func init() {
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
	Short: "Generates the Infrable.io logo.",
	Long: `
	Logo is a simple CLI that programmatically generates the Infrable.io logo.
	The output is an SVG file, which can be converted to any other image file
	format.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// set colorscheme
		base03 := "#002b36"
		base02 := "#073642"
		base01 := "#586e75"
		red := "#dc322f"

		// set attributes
		strokeWidth := float64(size) / 70.0
		radius := float64(size) / 70.0 // radius of circles

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
			Fill:    base03,
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
			Fill:        base03,
			Stroke:      red,
			StrokeWidth: strokeWidth,
		}

		p2 := svg.Polygon{
			XMLName:     "polygon",
			Fill:        base02,
			Stroke:      base01,
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

		cpts1 := [6]svg.Point{
			C1(ppts1[0], float64(size)/15.0, false),
			C2(ppts1[1], float64(size)/15.0, false),
			C3(ppts1[2], float64(size)/22.5, false),
			C5(ppts1[4], float64(size)/7.5, false),
			C6(ppts1[5], float64(size)/20.0, false),
			C8(ppts1[7], float64(size)/22.5, false),
		}

		cpts2 := [6]svg.Point{
			C1(ppts2[0], float64(size)/15.0, true),
			C2(ppts2[1], float64(size)/15.0, true),
			C3(ppts2[2], float64(size)/22.5, true),
			C5(ppts2[4], float64(size)/7.5, true),
			C6(ppts2[5], float64(size)/20.0, true),
			C8(ppts2[7], float64(size)/22.5, true),
		}

		for _, pt := range cpts1 {
			s.AppendElement(svg.Circle{
				XMLName: "cirle",
				Cx:      pt.X,
				Cy:      pt.Y,
				R:       radius,
				Fill:    red,
			})
		}

		for _, pt := range cpts2 {
			s.AppendElement(svg.Circle{
				XMLName: "cirle",
				Cx:      pt.X,
				Cy:      pt.Y,
				R:       radius,
				Fill:    base01,
			})
		}

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

func C1(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	x := c.X + rad*math.Cos(math.Atan(2.0)/2.0)
	y := c.Y + rad*math.Sin(math.Atan(2.0)/2.0)
	return svg.Point{X: x, Y: y}
}

func C2(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	x := c.X - rad*math.Cos(math.Atan(2.0)/2.0)
	y := c.Y + rad*math.Sin(math.Atan(2.0)/2.0)
	return svg.Point{X: x, Y: y}
}

func C3(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	x := c.X - rad*math.Cos((math.Atan(1.0/2.0)+math.Pi/2.0)/2.0)
	y := c.Y - rad*math.Sin((math.Atan(1.0/2.0)+math.Pi/2.0)/2.0)
	return svg.Point{X: x, Y: y}
}

func C4(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	x := c.X - rad*math.Cos(math.Atan(1.0))
	y := c.Y - rad*math.Sin(math.Atan(1.0))
	return svg.Point{X: x, Y: y}
}

func C5(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	x := c.X - rad*math.Sin(math.Atan(1.0/2.0)/2.0)
	y := c.Y - rad*math.Cos(math.Atan(1.0/2.0)/2.0)
	return svg.Point{X: x, Y: y}
}

func C6(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	// Note: This is a hack to make C5 and C6 vertically aligned.
	x := c.X + rad*(20.0/7.5)*math.Sin(math.Atan(1.0/2.0)/2.0)
	y := c.Y - rad*math.Cos((math.Atan(2.0)+math.Pi/2.0)/2.0)
	return svg.Point{X: x, Y: y}
}

func C7(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	x := c.X + rad*math.Cos(math.Atan(1.0))
	y := c.Y - rad*math.Sin(math.Atan(1.0))
	return svg.Point{X: x, Y: y}
}

func C8(c svg.Point, rad float64, mirror bool) svg.Point {
	if mirror {
		rad = -rad
	}
	x := c.X + rad*math.Cos((math.Atan(1.0/2.0)+math.Pi/2.0)/2.0)
	y := c.Y - rad*math.Sin((math.Atan(1.0/2.0)+math.Pi/2.0)/2.0)
	return svg.Point{X: x, Y: y}
}

// Copyright 2023 Infrable. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package svg

import (
	"encoding/xml"
	"fmt"
)

// For documentation on Scalable Vector Graphics (SVG), see the following:
//   - https://www.w3.org/TR/SVG

// An SVG document fragment consists of any number of SVG elements contained
// within an 'svg' element.
//
// SVG 2.0: https://www.w3.org/TR/SVG/struct.html#SVGElement
type SVG struct {
	XMLName      string  `xml:"svg"`
	XMLNamespace string  `xml:"xmlns,attr"`
	Version      string  `xml:"version,attr"`
	Width        float64 `xml:"width,attr"`
	Height       float64 `xml:"height,attr"`
	Fill         string  `xml:"fill,attr"`
	Stroke       string  `xml:"stroke,attr"`
	StrokeWidth  float64 `xml:"stroke-width,attr"`
	Elements     []*Element
}

type Element interface{}

func (s *SVG) AppendElement(e Element) {
	s.Elements = append(s.Elements, &e)
}

// The 'rect' element defines a rectangle which is axis-aligned with the
// current user coordinate system. Rounded rectangles can be achieved by
// setting non-zero values for the rx and ry geometric properties.
//
// SVG 2.0: https://www.w3.org/TR/SVG/shapes.html#RectElement
type Rect struct {
	XMLName     string  `xml:"rect"`
	X           float64 `xml:"x,attr"`
	Y           float64 `xml:"y,attr"`
	Width       float64 `xml:"width,attr"`
	Height      float64 `xml:"height,attr"`
	Rx          float64 `xml:"rx,attr"`
	Ry          float64 `xml:"ry,attr"`
	Fill        string  `xml:"fill,attr"`
	Stroke      string  `xml:"stroke,attr"`
	StrokeWidth float64 `xml:"stroke-width,attr"`
}

// The 'circle' element defines a circle based on a center point and a radius.
//
// SVG 2.0: https://www.w3.org/TR/SVG/shapes.html#CircleElement
type Circle struct {
	XMLName     string  `xml:"circle"`
	Cx          float64 `xml:"cx,attr"`
	Cy          float64 `xml:"cy,attr"`
	R           float64 `xml:"r,attr"`
	Fill        string  `xml:"fill,attr"`
	Stroke      string  `xml:"stroke,attr"`
	StrokeWidth float64 `xml:"stroke-width,attr"`
}

// The 'polygon' element defines a closed shape consisting of a set of
// connected straight line segments.
//
// SVG 2.0: https://www.w3.org/TR/SVG/shapes.html#PolygonElement
type Polygon struct {
	XMLName     string  `xml:"polygon"`
	Points      string  `xml:"points,attr"`
	Fill        string  `xml:"fill,attr"`
	Stroke      string  `xml:"stroke,attr"`
	StrokeWidth float64 `xml:"stroke-width,attr"`
}

func (p *Polygon) AppendPoint(pt Point) {
	if p.Points == "" {
		p.Points = fmt.Sprintf("%s", pt.String())
		return
	}
	p.Points = fmt.Sprintf("%s, %s", p.Points, pt.String())
}

type Point struct {
	X float64
	Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("%.f,%.f", p.X, p.Y)
}

func Marshal(v interface{}) ([]byte, error) {
	dat, err := xml.MarshalIndent(v, " ", " ")
	if err != nil {
		return nil, err
	}
	return dat, err
}

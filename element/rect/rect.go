package rect

import "encoding/xml"

type Rect struct {
	elem svgElem
}

type svgElem struct {
	XMLName xml.Name `xml:"rect"`
	X       uint32   `xml:"x,attr"`
	Y       uint32   `xml:"y,attr"`
	Width   uint32   `xml:"width,attr"`
	Height  uint32   `xml:"height,attr"`
}

func New(x, y, w, h uint32) *Rect {
	rect := &Rect{
		elem: svgElem{
			X:      x,
			Y:      y,
			Width:  w,
			Height: h,
		},
	}
	return rect
}

func (r *Rect) Stroke() {

}

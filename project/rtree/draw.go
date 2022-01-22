package rtree

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
)

const (
	R          = 50
	P          = 1
	RD float64 = R
)

var H = 0.0

func (t *Rtree) Draw() {
	fmt.Println("--------------------------------------------------")
	box := Rec{
		Low: Point{0, 0},
		Top: Point{18, 17},
	}
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, int(box.Top[x])*R, int(box.Top[y]*R)))
	H = box.Top[y] * R
	gc := draw2dimg.NewGraphicContext(dest)
	// Set some properties
	gc.SetFillColor(color.RGBA{255, 255, 255, 255})
	gc.SetStrokeColor(color.RGBA{0, 0, 0, 255})
	drawRectFill(gc, box)

	t.Root.Draw(gc)

	imaging.Save(dest, "tree.png")
}

func (n *Node) Draw(gc *draw2dimg.GraphicContext) {
	p := "nil"
	if n.parent != nil {
		p = n.parent.data
	}

	fmt.Printf("p: %v isLeaf: %t entries: %v \n", p, n.isLeaf, n.String())

	if n.data[0] == 'N' {
		gc.SetLineWidth(1)
		gc.SetLineDash([]float64{10, 20}, 10)
	} else {
		gc.SetLineWidth(4)
		gc.SetLineDash(nil, 0)
	}
	if n.parent != nil {
		drawRect(gc, n.bbox)
		drawText(gc, n.bbox, n.data)
	}

	for _, e := range n.entries {
		e.Draw(gc)
	}
}

func drawRect(gc *draw2dimg.GraphicContext, r Rec) {
	// Draw a closed shape
	gc.BeginPath()                        // Initialize a new path
	gc.MoveTo(r.Low[x]*RD, H-r.Low[y]*RD) // Move to a position to start the new path
	gc.LineTo(r.Low[x]*RD, H-r.Top[y]*RD)
	gc.LineTo(r.Top[x]*RD, H-r.Top[y]*RD)
	gc.LineTo(r.Top[x]*RD, H-r.Low[y]*RD)
	gc.Close()
	gc.Stroke()
}

func drawText(gc *draw2dimg.GraphicContext, r Rec, txt string) {
	gc.SetFillColor(color.RGBA{0, 0, 0, 255})
	gc.SetFontSize(10)
	gc.FillStringAt(txt, r.Low[x]*RD+10, H-r.Top[y]*RD+15)
}

func drawRectFill(gc *draw2dimg.GraphicContext, r Rec) {
	// Draw a closed shape
	gc.BeginPath()                      // Initialize a new path
	gc.MoveTo(r.Low[x]*RD, r.Low[y]*RD) // Move to a position to start the new path
	gc.LineTo(r.Low[x]*RD, r.Top[y]*RD)
	gc.LineTo(r.Top[x]*RD, r.Top[y]*RD)
	gc.LineTo(r.Top[x]*RD, r.Low[y]*RD)
	gc.Close()

	gc.FillStroke()
}

func (t *Rtree) bbox() Rec {
	var box Rec
	for _, v := range t.Root.entries {
		box = bbox(box, v.bbox)
	}
	return box
}

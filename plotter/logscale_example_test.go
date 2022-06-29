// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"image/color"
	"log"
	"math"

	"github.com/Emptywe/plot"
	"github.com/Emptywe/plot/plotter"
	"github.com/Emptywe/plot/vg"
)

// Example_logScale shows how to create a plot with a log-scale on the Y-axis.
func Example_logScale() {
	p := plot.New()
	p.Title.Text = "My Plot"
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{Prec: -1}
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	f := plotter.NewFunction(math.Exp)
	f.XMin = 0.2
	f.XMax = 10
	f.Color = color.RGBA{R: 255, A: 255}

	p.Add(f, plotter.NewGrid())
	p.Legend.Add("exp(x)", f)

	p.X.Min = f.XMin
	p.X.Max = f.XMax
	p.Y.Min = math.Exp(f.XMin)
	p.Y.Max = math.Exp(f.XMax)

	err := p.Save(10*vg.Centimeter, 10*vg.Centimeter, "testdata/logscale.png")
	if err != nil {
		log.Panic(err)
	}
}

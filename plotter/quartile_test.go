// Copyright ©2015 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plotter_test

import (
	"testing"

	"github.com/emptywe/plot/cmpimg"
)

func TestQuartPlot(t *testing.T) {
	cmpimg.CheckPlot(ExampleQuartPlot, t, "verticalQuartPlot.png",
		"horizontalQuartPlot.png",
		"groupedQuartPlot.png")
}

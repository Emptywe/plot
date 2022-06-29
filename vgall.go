// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !minimal
// +build !minimal

package plot // import "github.com/emptywe/plot"

import (
	_ "github.com/emptywe/plot/vg/vgeps"
	_ "github.com/emptywe/plot/vg/vgimg"
	_ "github.com/emptywe/plot/vg/vgpdf"
	_ "github.com/emptywe/plot/vg/vgsvg"
	_ "github.com/emptywe/plot/vg/vgtex"
)

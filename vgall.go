// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !minimal
// +build !minimal

package plot // import "gitlab.com/Emptywe/plot"

import (
	_ "gitlab.com/Emptywe/plot/vg/vgeps"
	_ "gitlab.com/Emptywe/plot/vg/vgimg"
	_ "gitlab.com/Emptywe/plot/vg/vgpdf"
	_ "gitlab.com/Emptywe/plot/vg/vgsvg"
	_ "gitlab.com/Emptywe/plot/vg/vgtex"
)

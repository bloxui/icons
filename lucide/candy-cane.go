package lucide

import x "github.com/plainkit/html"

// CandyCane creates a Candy Cane Lucide icon.
func CandyCane(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-candy-cane", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5.7 21a2 2 0 0 1-3.5-2l8.6-14a6 6 0 0 1 10.4 6 2 2 0 1 1-3.464-2 2 2 0 1 0-3.464-2Z"))),
		x.Child(x.Path(x.D("M17.75 7 15 2.1"))),
		x.Child(x.Path(x.D("M10.9 4.8 13 9"))),
		x.Child(x.Path(x.D("m7.9 9.7 2 4.4"))),
		x.Child(x.Path(x.D("M4.9 14.7 7 18.9"))),
	)
	return x.Svg(svgArgs...)
}

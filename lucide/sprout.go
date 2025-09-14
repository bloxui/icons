package lucide

import x "github.com/bloxui/blox"

// Sprout creates a Sprout Lucide icon.
func Sprout(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-sprout", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M14 9.536V7a4 4 0 0 1 4-4h1.5a.5.5 0 0 1 .5.5V5a4 4 0 0 1-4 4 4 4 0 0 0-4 4c0 2 1 3 1 5a5 5 0 0 1-1 3"))),
		x.Child(x.Path(x.D("M4 9a5 5 0 0 1 8 4 5 5 0 0 1-8-4"))),
		x.Child(x.Path(x.D("M5 21h14"))),
	)
	return x.Svg(svgArgs...)
}

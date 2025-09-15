package lucide

import x "github.com/plainkit/html"

// GlassWater creates a Glass Water Lucide icon.
func GlassWater(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-glass-water", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M5.116 4.104A1 1 0 0 1 6.11 3h11.78a1 1 0 0 1 .994 1.105L17.19 20.21A2 2 0 0 1 15.2 22H8.8a2 2 0 0 1-2-1.79z"))),
		x.Child(x.Path(x.D("M6 12a5 5 0 0 1 6 0 5 5 0 0 0 6 0"))),
	)
	return x.Svg(svgArgs...)
}

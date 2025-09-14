package lucide

import x "github.com/bloxui/blox"

// UmbrellaOff creates a Umbrella Off Lucide icon.
func UmbrellaOff(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-umbrella-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 13v7a2 2 0 0 0 4 0"))),
		x.Child(x.Path(x.D("M12 2v2"))),
		x.Child(x.Path(x.D("M18.656 13h2.336a1 1 0 0 0 .97-1.274 10.284 10.284 0 0 0-12.07-7.51"))),
		x.Child(x.Path(x.D("m2 2 20 20"))),
		x.Child(x.Path(x.D("M5.961 5.957a10.28 10.28 0 0 0-3.922 5.769A1 1 0 0 0 3 13h10"))),
	)
	return x.Svg(svgArgs...)
}

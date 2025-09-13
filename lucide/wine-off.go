package lucide

import x "github.com/bloxui/blox"

// WineOff creates a Wine Off Lucide icon.
func WineOff(args ...x.SvgArg) x.Component {
	svgArgs := buildLucideArgs("lucide lucide-wine-off", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M8 22h8"))),
		x.Child(x.Path(x.D("M7 10h3m7 0h-1.343"))),
		x.Child(x.Path(x.D("M12 15v7"))),
		x.Child(x.Path(x.D("M7.307 7.307A12.33 12.33 0 0 0 7 10a5 5 0 0 0 7.391 4.391M8.638 2.981C8.75 2.668 8.872 2.34 9 2h6c1.5 4 2 6 2 8 0 .407-.05.809-.145 1.198"))),
		x.Child(x.Line(x.X1("2"), x.X2("22"), x.Y1("2"), x.Y2("22"))),
	)
	return x.Svg(svgArgs...)
}

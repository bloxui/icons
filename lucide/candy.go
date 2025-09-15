package lucide

import x "github.com/plainkit/html"

// Candy creates a Candy Lucide icon.
func Candy(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-candy", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M10 7v10.9"))),
		x.Child(x.Path(x.D("M14 6.1V17"))),
		x.Child(x.Path(x.D("M16 7V3a1 1 0 0 1 1.707-.707 2.5 2.5 0 0 0 2.152.717 1 1 0 0 1 1.131 1.131 2.5 2.5 0 0 0 .717 2.152A1 1 0 0 1 21 8h-4"))),
		x.Child(x.Path(x.D("M16.536 7.465a5 5 0 0 0-7.072 0l-2 2a5 5 0 0 0 0 7.07 5 5 0 0 0 7.072 0l2-2a5 5 0 0 0 0-7.07"))),
		x.Child(x.Path(x.D("M8 17v4a1 1 0 0 1-1.707.707 2.5 2.5 0 0 0-2.152-.717 1 1 0 0 1-1.131-1.131 2.5 2.5 0 0 0-.717-2.152A1 1 0 0 1 3 16h4"))),
	)
	return x.Svg(svgArgs...)
}

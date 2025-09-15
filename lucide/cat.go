package lucide

import x "github.com/plainkit/html"

// Cat creates a Cat Lucide icon.
func Cat(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-cat", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 5c.67 0 1.35.09 2 .26 1.78-2 5.03-2.84 6.42-2.26 1.4.58-.42 7-.42 7 .57 1.07 1 2.24 1 3.44C21 17.9 16.97 21 12 21s-9-3-9-7.56c0-1.25.5-2.4 1-3.44 0 0-1.89-6.42-.5-7 1.39-.58 4.72.23 6.5 2.23A9.04 9.04 0 0 1 12 5Z"))),
		x.Child(x.Path(x.D("M8 14v.5"))),
		x.Child(x.Path(x.D("M16 14v.5"))),
		x.Child(x.Path(x.D("M11.25 16.25h1.5L12 17l-.75-.75Z"))),
	)
	return x.Svg(svgArgs...)
}

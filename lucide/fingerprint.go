package lucide

import x "github.com/plainkit/html"

// Fingerprint creates a Fingerprint Lucide icon.
func Fingerprint(args ...x.SvgArg) x.Node {
	svgArgs := buildLucideArgs("lucide lucide-fingerprint", args)
	svgArgs = append(svgArgs,
		x.Child(x.Path(x.D("M12 10a2 2 0 0 0-2 2c0 1.02-.1 2.51-.26 4"))),
		x.Child(x.Path(x.D("M14 13.12c0 2.38 0 6.38-1 8.88"))),
		x.Child(x.Path(x.D("M17.29 21.02c.12-.6.43-2.3.5-3.02"))),
		x.Child(x.Path(x.D("M2 12a10 10 0 0 1 18-6"))),
		x.Child(x.Path(x.D("M2 16h.01"))),
		x.Child(x.Path(x.D("M21.8 16c.2-2 .131-5.354 0-6"))),
		x.Child(x.Path(x.D("M5 19.5C5.5 18 6 15 6 12a6 6 0 0 1 .34-2"))),
		x.Child(x.Path(x.D("M8.65 22c.21-.66.45-1.32.57-2"))),
		x.Child(x.Path(x.D("M9 6.8a6 6 0 0 1 9 5.2v2"))),
	)
	return x.Svg(svgArgs...)
}

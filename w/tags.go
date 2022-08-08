package w

//go:noinline
func HPoint() *Ele {
	return Span().Style("all", "unset").Attr("data-point", "1")
	// return CreateTextNode("")
}

//go:noinline
func Div(text ...[]any) *Ele {
	return CreateElement("div")
}

//go:noinline
func Svg() *Ele {
	return CreateElement("svg")
}

//go:noinline
func Input() *Ele {
	return CreateElement("input")
}

//go:noinline
func Textarea() *Ele {
	return CreateElement("textarea")
}

//go:noinline
func Header() *Ele {
	return CreateElement("header")
}

//go:noinline
func Strong() *Ele {
	return CreateElement("strong")
}

//go:noinline
func Footer() *Ele {
	return CreateElement("footer")
}

//go:noinline
func Meta() *Ele {
	return CreateElement("meta")
}

//go:noinline
func Style() *Ele {
	return CreateElement("style")
}

//go:noinline
func Script() *Ele {
	return CreateElement("script")
}

//go:noinline
func Template() *Ele {
	return CreateElement("template")
}

//go:noinline
func Link() *Ele {
	return CreateElement("link")
}

//go:noinline
func Iframe() *Ele {
	return CreateElement("iframe")
}

//go:noinline
func P() *Ele {
	return CreateElement("p")
}

//go:noinline
func Span() *Ele {
	return CreateElement("span")
}

//go:noinline
func I() *Ele {
	return CreateElement("i")
}

//go:noinline
func Kbd() *Ele {
	return CreateElement("kbd")
}

//go:noinline
func Button() *Ele {
	return CreateElement("button")
}

//go:noinline
func Form() *Ele {
	return CreateElement("form")
}

//go:noinline
func Label() *Ele {
	return CreateElement("label")
}

//go:noinline
func Section() *Ele {
	return CreateElement("section")
}

//go:noinline
func Select() *Ele {
	return CreateElement("section")
}

//go:noinline
func Nav() *Ele {
	return CreateElement("nav")
}

//go:noinline
func Mark() *Ele {
	return CreateElement("mark")
}

//go:noinline
func Menu() *Ele {
	return CreateElement("menu")
}

//go:noinline
func Article() *Ele {
	return CreateElement("article")
}

//go:noinline
func Aside() *Ele {
	return CreateElement("aside")
}

//go:noinline
func Code() *Ele {
	return CreateElement("code")
}

//go:noinline
func Ol() *Ele {
	return CreateElement("ol")
}

//go:noinline
func Ul() *Ele {
	return CreateElement("ul")
}

//go:noinline
func Li() *Ele {
	return CreateElement("li")
}

//go:noinline
func Dt() *Ele {
	return CreateElement("dt")
}

//go:noinline
func Dd() *Ele {
	return CreateElement("dd")
}

//go:noinline
func H1() *Ele {
	return CreateElement("h1")
}

//go:noinline
func H2() *Ele {
	return CreateElement("h2")
}

//go:noinline
func H3() *Ele {
	return CreateElement("h3")
}

//go:noinline
func H4() *Ele {
	return CreateElement("h4")
}

//go:noinline
func H5() *Ele {
	return CreateElement("h5")
}

//go:noinline
func H6() *Ele {
	return CreateElement("h6")
}

//go:noinline
func Table() *Ele {
	return CreateElement("table")
}

//go:noinline
func Tr() *Ele {
	return CreateElement("tr")
}

//go:noinline
func Th() *Ele {
	return CreateElement("th")
}

//go:noinline
func Td() *Ele {
	return CreateElement("td")
}

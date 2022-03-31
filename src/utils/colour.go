package utils

const (
	colourReset  = "\033[0m"
	colourRed    = "\033[31m"
    colourGreen  = "\033[32m"
    colourYellow = "\033[33m"
    colourBlue   = "\033[34m"
    colourPurple = "\033[35m"
    colourCyan   = "\033[36m"
    colourWhite  = "\033[37m"
)

type textBuilder struct {
	text string
}

func BuildColourText() *textBuilder {
	builder := &textBuilder{}
	builder.Init()

	return builder
}

func (b *textBuilder) Init () {
	b.text = ""
}

func (b *textBuilder) Add (text string) *textBuilder {
	b.text += text
	return b
}

func (b *textBuilder) Reset () *textBuilder {
	b.text += colourReset
	return b
}


func (b *textBuilder) Red () *textBuilder {
	b.text += colourRed
	return b
}

func (b *textBuilder) Green () *textBuilder {
	b.text += colourGreen
	return b
}

func (b *textBuilder) Yellow () *textBuilder {
	b.text += colourYellow
	return b
}

func (b *textBuilder) Blue () *textBuilder {
	b.text += colourBlue
	return b
}

func (b *textBuilder) Purple () *textBuilder {
	b.text += colourPurple
	return b
}

func (b *textBuilder) Cyan () *textBuilder {
	b.text += colourCyan
	return b
}

func (b *textBuilder) White () *textBuilder {
	b.text += colourWhite
	return b
}

func (b *textBuilder) String () string {
	b.Reset()

	return b.text
}

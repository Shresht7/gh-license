package main

import (
	"bytes"
	"strings"
)

// wrap wraps the given text with the given string.
func wrap(str string, text ...string) string {
	for i, t := range text {
		text[i] = str + t + str
	}
	return strings.Join(text, "")
}

// wrapAndJoin wraps the given text with the given string and joins them with the given string.
func wrapAndJoin(str, join string, text ...string) string {
	for i, t := range text {
		text[i] = str + t + str
	}
	return strings.Join(text, join)
}

// Markdown is a markdown writer. It provides methods to write markdown elements to a buffer.
type Markdown struct {
	buf              *bytes.Buffer
	baseHeadingLevel int
}

// NewMarkdownWriter creates a new markdown writer.
func NewMarkdownWriter() *Markdown {
	return &Markdown{
		buf:              new(bytes.Buffer),
		baseHeadingLevel: 0,
	}
}

// Implement the Stringer interface.
func (m *Markdown) String() string {
	return m.buf.String()
}

// Reset the markdown buffer.
func (m *Markdown) Reset() {
	m.buf.Reset()
}

// Write a line element to the markdown buffer.
func (m *Markdown) WriteLine(s ...string) {
	str := strings.Join(s, "") + "\n"
	m.buf.WriteString(str)
}

// Write a paragraph element to the markdown buffer.
func (m *Markdown) WriteParagraph(lines ...string) {
	str := strings.Join(lines, "\n") + "\n\n"
	m.buf.WriteString(str)
}

// Write a heading element to the markdown buffer. The level is relative to the base heading level.
// For example, if the base heading level is 2, then a level of 1 will result in a heading of level 3.
func (m *Markdown) WriteHeading(level int, text string) {
	m.WriteParagraph(strings.Repeat("#", m.baseHeadingLevel+level) + " " + text)
}

// Write a heading element to the markdown buffer with a link. The level is relative to the base heading level.
// For example, if the base heading level is 2, then a level of 1 will result in a heading of level 3.
func (m *Markdown) WriteHeadingWithLink(level int, text, link string) {
	m.WriteParagraph(strings.Repeat("#", m.baseHeadingLevel+level) + " [" + text + "](" + link + ")")
}

// Write a inline code block element to the markdown buffer.
func (m *Markdown) WriteInlineCodeBlock(text string) {
	m.buf.WriteString(wrap("`", text))
}

// Write a code block element to the markdown buffer.
func (m *Markdown) WriteCodeBlock(code string, language ...string) {
	m.WriteLine("```" + strings.Join(language, " "))
	m.buf.WriteString(code)
	m.WriteLine("\n```\n")
}

// Write a list element to the markdown buffer.
func (m *Markdown) WriteList(items ...string) {
	for _, item := range items {
		m.WriteLine("- " + item)
	}
	m.WriteLine()
}

// Write a table element to the markdown buffer.
func (m *Markdown) WriteTable(headers []string, rows [][]string) {
	// Write the headers
	m.WriteLine("| " + strings.Join(headers, " | ") + " |")
	m.WriteLine("|" + strings.Repeat(" --- |", len(headers)))

	// Write the rows
	for _, row := range rows {
		m.WriteLine("| " + strings.Join(row, " | "))
	}
	m.WriteLine()
}

// Write a link element to the markdown buffer.
func (m *Markdown) WriteLink(text, link string) {
	m.buf.WriteString("[" + text + "](" + link + ")")
}

// Write an image element to the markdown buffer.
func (m *Markdown) WriteImage(alt, link string) {
	m.buf.WriteString("![alt](" + link + ")")
}

// Write a horizontal rule element to the markdown buffer.
func (m *Markdown) WriteHorizontalRule() {
	m.WriteParagraph("---")
}

// Write a block quote element to the markdown buffer.
func (m *Markdown) WriteBlockQuote(text string) {
	m.WriteParagraph("> " + text)
}

// Write a bold element to the markdown buffer.
func (m *Markdown) WriteBold(text string) {
	m.buf.WriteString("**" + text + "**")
}

// Write a italic element to the markdown buffer.
func (m *Markdown) WriteItalic(text string) {
	m.buf.WriteString("*" + text + "*")
}

// Write a strikethrough element to the markdown buffer.
func (m *Markdown) WriteStrikethrough(text string) {
	m.buf.WriteString("~~" + text + "~~")
}

// Write a superscript element to the markdown buffer.
func (m *Markdown) WriteSuperscript(text string) {
	m.buf.WriteString("^" + text + "^")
}

// Write a subscript element to the markdown buffer.
func (m *Markdown) WriteSubscript(text string) {
	m.buf.WriteString("~" + text + "~")
}

// Write a footnote reference element to the markdown buffer.
func (m *Markdown) WriteFootnoteReference(n string) {
	m.buf.WriteString("[^" + n + "]")
}

// Write a footnote definition element to the markdown buffer.
func (m *Markdown) WriteFootnoteDefinition(n, text string) {
	m.WriteLine("[^" + n + "]: " + text)
}

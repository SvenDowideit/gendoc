package mmark

import (
	"github.com/miekg/mmark"
)

// Render a Markdown fragment into a html fragment
func Render(md string) (html string, err error) {
	// not rendering entire page - that will happen using the tmpl html
	page := false
	css := ""
	head := ""

	// set up options
	extensions := 0
	extensions |= mmark.EXTENSION_TABLES
	extensions |= mmark.EXTENSION_FENCED_CODE
	extensions |= mmark.EXTENSION_AUTOLINK
	extensions |= mmark.EXTENSION_SPACE_HEADERS
	extensions |= mmark.EXTENSION_CITATION
	extensions |= mmark.EXTENSION_TITLEBLOCK_TOML
	extensions |= mmark.EXTENSION_HEADER_IDS
	extensions |= mmark.EXTENSION_AUTO_HEADER_IDS
	extensions |= mmark.EXTENSION_UNIQUE_HEADER_IDS
	extensions |= mmark.EXTENSION_FOOTNOTES
	extensions |= mmark.EXTENSION_SHORT_REF
	extensions |= mmark.EXTENSION_INCLUDE
	extensions |= mmark.EXTENSION_PARTS
	extensions |= mmark.EXTENSION_ABBREVIATIONS
	extensions |= mmark.EXTENSION_DEFINITION_LISTS

	var renderer mmark.Renderer
	htmlFlags := 0
	if page {
		htmlFlags |= mmark.HTML_COMPLETE_PAGE
	}
	renderer = mmark.HtmlRenderer(htmlFlags, css, head)

	output := mmark.Parse([]byte(md), renderer, extensions).Bytes()
	return string(output), nil
}

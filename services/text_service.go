package services

import "html"

// HtmlDecode Twice to handle sequences like: "&amp;mdash;"
func HtmlDecode(s string) string {
	return html.UnescapeString(html.UnescapeString(s))
}

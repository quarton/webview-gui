// +build windows

package gui

import "net/url"

// windows needs QueryEscape not PathEscape
func encode(html string) string {
	return "data:text/html," + url.QueryEscape(html)
}

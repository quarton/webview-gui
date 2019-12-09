// +build windows

package gui

import "net/url"

// windows needs query escape not PathEscape
func encode(html string) string {
	return "data:text/html," + url.QueryEscape(html)
}

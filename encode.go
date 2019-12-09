// +build !windows

package gui

import "net/url"

func encode(html string) string {
	return "data:text/html," + url.PathEscape(html)
}

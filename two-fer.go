// Package twofer is to determine what to say as you give away the extra cookie.
package twofer

// ShareWith function is to send the comment while giving the extra cookie
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}

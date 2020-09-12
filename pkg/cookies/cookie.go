package cookies

import (
	"strings"
)

// Cookies represent a common set of configuration settings
// applicable to the cookies issued by the plugin.
type Cookies struct {
	Domain string `json:"domain,omitempty"`
	Path   string `json:"path,omitempty"`
}

// GetAttributes returns cookie attributes.
func (c *Cookies) GetAttributes() string {
	var sb strings.Builder
	if c.Domain != "" {
		sb.WriteString(" Domain=" + c.Domain + ";")
	}
	if c.Path != "" {
		sb.WriteString(" Path=" + c.Path + ";")
	}
	sb.WriteString(" Secure; HttpOnly;")
	return sb.String()
}

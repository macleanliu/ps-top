// Package context stores some common information used in various places
package context

import (
	"strings"
	"time"

	"github.com/sjmudd/ps-top/global"
	"github.com/sjmudd/ps-top/lib"
	"github.com/sjmudd/ps-top/version"
)

// Context holds the common information
type Context struct {
	last      time.Time
	status    *global.Status
	uptime    int
	variables *global.Variables
	version   string
}

// NewContext returns the pointer to a new (empty) context
func NewContext(status *global.Status) *Context {
	c := new(Context)
	c.status = status

	return c
}

// SetVariables provides access to the global variables
func (c *Context) SetVariables(variables *global.Variables) {
	c.variables = variables
}

// Hostname returns the current short hostname
func (c Context) Hostname() string {
	hostname := c.variables.Get("hostname")
	if index := strings.Index(hostname, "."); index >= 0 {
		hostname = hostname[0:index]
	}
	return hostname
}

// MySQLVersion returns the current MySQL version
func (c Context) MySQLVersion() string {
	return c.variables.Get("version")
}

// Version returns the Application version
func (c Context) Version() string {
	return version.Version()
}

// MyName returns the program's name
func (c Context) MyName() string {
	return lib.MyName()
}

// Uptime returns the time that MySQL has been up
func (c Context) Uptime() int {
	return c.status.Get("Uptime")
}

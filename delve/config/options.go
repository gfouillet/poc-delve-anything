package config

import (
	"fmt"
	"maps"
)

// Options is a type that represents a map of configuration options where the key is a string and the value can be any type.
// They are passed to Delve as command line option.
type Options map[string]any

// Default returns an Options map with a default configuration settings for the application.
func Default() Options {
	return Options{
		"headless":           true,
		"continue":           true,
		"listen":             ":0",
		"api-version":        2,
		"accept-multiclient": true,
	}
}

// Args converts a variadic list of Options into a slice of command line argument strings.
func Args(opts ...Options) []string {
	if len(opts) > 0 {
		return opts[0].merge(opts[1:]...).args()
	}
	return nil
}

// args constructs a slice of command line argument strings from the Options map.
func (o Options) args() []string {
	args := make([]string, 0, len(o))
	for k, v := range o {
		switch val := v.(type) {
		case bool:
			if val {
				args = append(args, fmt.Sprintf("--%s", k))
			}
		default:
			if v != nil {
				args = append(args, fmt.Sprintf("--%s=%v", k, v))
			}
		}
	}
	return args
}

// merge combines multiple Options into one, overriding any duplicate keys with values from the latest Options in the list.
func (o Options) merge(opts ...Options) Options {
	for _, opt := range opts {
		maps.Insert(o, maps.All(opt))
	}
	return o
}

// WithOption creates an Options instance with a given key and one or more values.
func WithOption(key string, value ...any) Options {
	switch len(value) {
	case 0:
		return Options{key: true}
	case 1:
		return Options{key: value[0]}
	default:
		return Options{key: value}
	}
}

// WithPort configures the listening port for the application by setting the "listen" option to the specified port.
func WithPort(port int) Options {
	return WithOption("listen", fmt.Sprintf(":%d", port))
}

// WithApiVersion sets the API version in the options with the specified version number.
func WithApiVersion(version int) Options {
	return WithOption("api-version", version)
}

// Headless sets the "headless" option to true, which runs delve in server mode
func Headless() Options {
	return WithOption("headless", true)
}

// WaitDebugger returns an Options map with the "continue" option set to false,
// the debugged application will wait until a debugger is attached, after having generating a first
// log indicating on which endpoint it listens.
func WaitDebugger() Options {
	return Options{
		"continue": false,
	}
}

// NoWait returns an Options map with settings to continue execution immediately. Debugged application
// wouldn't wait for a debugger to be attached.
func NoWait() Options {
	return Options{
		"continue":           true,
		"accept-multiclient": true,
	}
}

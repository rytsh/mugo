package fstore

import (
	"github.com/rytsh/mugo/pkg/fstore/registry"
	"github.com/worldline-go/logz"
)

type options struct {
	disableFuncs   map[string]struct{}
	disableGroups  map[string]struct{}
	specificFunc   map[string]struct{}
	specificGroups map[string]struct{}

	trust           bool
	log             logz.Adapter
	workDir         string
	executeTemplate registry.ExecuteTemplate
}

type Option func(options *options)

// WithSpecificFuncs is a option for just enable specific functions.
func WithSpecificFuncs(specificFuncs ...string) Option {
	return func(options *options) {
		for _, f := range specificFuncs {
			options.specificFunc[f] = struct{}{}
		}
	}
}

// WithSpecificGroups is a option for just enable specific direct add groups.
//
//	WithSpecificGroups("sprig")
func WithSpecificGroups(specificGroups ...string) Option {
	return func(options *options) {
		for _, f := range specificGroups {
			options.specificGroups[f] = struct{}{}
		}
	}
}

// WithDisableGroups is a option for disable direct groups.
//
//	WithDisableGroups("sprig")
func WithDisableGroups(disableGroups ...string) Option {
	return func(options *options) {
		for _, g := range disableGroups {
			options.disableGroups[g] = struct{}{}
		}
	}
}

// WithDisableFuncs is a option for disableFuncs.
//
//	WithDisableFuncs("exec", "execTemplate")
func WithDisableFuncs(disableFuncs ...string) Option {
	return func(options *options) {
		for _, f := range disableFuncs {
			options.disableFuncs[f] = struct{}{}
		}
	}
}

// WithTrust is a option for trust.
// Some functions are not safe to use such as "exec".
func WithTrust(trust bool) Option {
	return func(options *options) {
		options.trust = trust
	}
}

// WithWorkDir is a option for workDir.
func WithWorkDir(workDir string) Option {
	return func(options *options) {
		options.workDir = workDir
	}
}

func WithExecuteTemplate(t registry.ExecuteTemplate) Option {
	return func(options *options) {
		options.executeTemplate = t
	}
}

func WithLog(log logz.Adapter) Option {
	return func(options *options) {
		options.log = log
	}
}

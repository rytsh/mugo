package fstore

type option struct {
	disableFuncs   map[string]struct{}
	disableGroups  map[string]struct{}
	specificFunc   map[string]struct{}
	specificGroups map[string]struct{}

	trust           bool
	log             Adapter
	workDir         string
	executeTemplate ExecuteTemplate
}

type OptionFunc func(option *option)

// WithSpecificFuncs is a option for just enable specific functions.
func WithSpecificFuncs(specificFuncs ...string) OptionFunc {
	return func(option *option) {
		for _, f := range specificFuncs {
			option.specificFunc[f] = struct{}{}
		}
	}
}

// WithSpecificGroups is a option for just enable specific direct add groups.
//
//	WithSpecificGroups("sprig")
func WithSpecificGroups(specificGroups ...string) OptionFunc {
	return func(option *option) {
		for _, f := range specificGroups {
			option.specificGroups[f] = struct{}{}
		}
	}
}

// WithDisableGroups is a option for disable direct groups.
//
//	WithDisableGroups("sprig")
func WithDisableGroups(disableGroups ...string) OptionFunc {
	return func(option *option) {
		for _, g := range disableGroups {
			option.disableGroups[g] = struct{}{}
		}
	}
}

// WithDisableFuncs is a option for disableFuncs.
//
//	WithDisableFuncs("exec", "execTemplate")
func WithDisableFuncs(disableFuncs ...string) OptionFunc {
	return func(option *option) {
		for _, f := range disableFuncs {
			option.disableFuncs[f] = struct{}{}
		}
	}
}

// WithTrust is a option for trust.
// Some functions are not safe to use such as "exec".
func WithTrust(trust bool) OptionFunc {
	return func(option *option) {
		option.trust = trust
	}
}

// WithWorkDir is a option for workDir.
func WithWorkDir(workDir string) OptionFunc {
	return func(option *option) {
		option.workDir = workDir
	}
}

func WithExecuteTemplate(t ExecuteTemplate) OptionFunc {
	return func(option *option) {
		option.executeTemplate = t
	}
}

func WithLog(log Adapter) OptionFunc {
	return func(option *option) {
		option.log = log
	}
}

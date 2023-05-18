package profile

import "errors"

var (
	ErrNoRouteProvided      = errors.New("no routes provided")
	ErrNeedTimeIntervalName = errors.New("time interval name must not be empty")
	ErrRootMissReceiver     = errors.New("root route must specify a default receiver")
	ErrRootMustNoMatcher    = errors.New("root route must not have any matchers")
	ErrRootMustNoMute       = errors.New("root route must not have any mute time intervals")
	ErrRootMustNoActive     = errors.New("root route must not have any active time intervals")
)

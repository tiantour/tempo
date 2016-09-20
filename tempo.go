package tempo

// tempo
var (
	Time     = &iTime{}
	Calendar = &iCalendar{}
	Now      = &iNow{}
	Unix     = &iUnix{}
	String   = &iString{}
)

type (
	iTime     struct{}
	iCalendar struct{}
	iNow      struct{}
	iUnix     struct{}
	iString   struct{}
)

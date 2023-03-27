package task

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

var stateTransitionMap = map[State]map[State]interface{}{
	Pending: {Scheduled: nil},
	Scheduled: {
		Scheduled: nil,
		Running:   nil,
		Failed:    nil,
	},
	Running: {
		Running:   nil,
		Completed: nil,
		Failed:    nil,
	},
	Completed: {},
	Failed:    {},
}

func ValidStateTransition(src State, dst State) bool {
	states, ok := stateTransitionMap[src]
	if !ok {
		return false
	}
	_, ok = states[dst]
	return ok
}

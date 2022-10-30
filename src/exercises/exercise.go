package exercises

type Exercise struct {
	Name  string
	Path  string
	Mode  string
	Hint  string
	State State
}

type State int

const (
	Pending State = iota + 1
	Done
)

func (s State) String() string {
	return [...]string{"Pending", "Done"}[s-1]
}

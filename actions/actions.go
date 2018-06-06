package actions

const ( // actions
	Top   = iota
	Bot
	Left
	Right
)

type Action struct {
	Name  string
	Value int
}

var None = Action{
	"None",
	-1,
}

var L = [4]Action{
	{
		"Top",
		Top,
	},
	{
		"Bot",
		Bot,
	},
	{
		"Left",
		Left,
	},
	{
		"Right",
		Right,
	},
}

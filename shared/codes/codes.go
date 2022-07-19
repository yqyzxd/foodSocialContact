package codes

type Code int

const (
	SUCCESS   Code = 0
	ERROR     Code = 1
	NOT_LOGIN Code = -100
)

func (c Code) Int() int {
	return int(c)
}

package sthttp

type Attempt func() error

type Attempter interface {
	Attempt(a Attempt) error
}

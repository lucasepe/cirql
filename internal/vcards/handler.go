package vcards

type CardHandler interface {
	Handle(c Card) error
}

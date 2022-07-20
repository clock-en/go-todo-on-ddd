package vo

type IHash interface {
	Value() string
	Verify() bool
}

type Hash struct {
	IHash
	value string
}

func NewHash(value string) (*Hash, error) {
	return &Hash{value: value}, nil
}

func (h Hash) Value() string {
	return h.value
}

func (h Hash) Verify(password Password) bool {
	comparison := password.Hash()
	return h.value == comparison
}

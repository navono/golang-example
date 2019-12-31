package option

type (
	Treasure struct {
		key    string
		option Option
	}

	Option struct {
		typ string
		num int
	}

	OptionFn func(op *Option)
)

func WithType(typ string) OptionFn {
	return func(op *Option) {
		op.typ = typ
	}
}

func WithNum(num int) OptionFn {
	return func(op *Option) {
		op.num = num
	}
}

func NewTreasure(key string, ops ...OptionFn) *Treasure {
	op := Option{
		typ: "bracelet",
		num: 0,
	}

	for _, fn := range ops {
		fn(&op)
	}

	return &Treasure{
		key:    key,
		option: op,
	}
}

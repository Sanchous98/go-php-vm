package vm

type Opline struct {
	Opcode           Opcode
	Op1, Op2, Result ZVal
	LineNo           int
}

func (o *Opline) Run(ctx Context) (err error) {
	o.Result, err = o.Opcode.Handler()(ctx, o.Op1, o.Op2)
	return
}

type Oparray []*Opline

func (o *Oparray) Run(ctx Context) error {
	for _, opline := range *o {
		if err := opline.Run(ctx); err != nil {
			return err
		}
	}

	return nil
}

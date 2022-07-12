package vm

import (
	"context"
	"fmt"
	"time"
)

type Context interface {
	context.Context
	Parent() Context
	SetErrorHandler(func(Context, error))
	ThrowError(error)
	GetVariable(string) *ZVariable
	NewVariable(ZVal) *ZVariable
	GetClass(string)
	GetFunction(string)
	Exec()
}

type GlobalContext struct {
	Functions []any
	Classes   []any
	Constants []any
	Variables []*ZVariable
}

func (g *GlobalContext) generateVarName() string {
	return fmt.Sprintf("!%d", len(g.Variables))
}

func (g *GlobalContext) NewVariable(val ZVal) (v *ZVariable) {
	v = AllocVariable(g.generateVarName(), val)
	g.Variables = append(g.Variables, v)
	return
}

func (g *GlobalContext) Deadline() (deadline time.Time, ok bool) {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) Done() <-chan struct{} {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) Err() error {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) Value(key any) any {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) Parent() Context {
	return nil
}

func (g *GlobalContext) SetErrorHandler(f func(Context, error)) {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) ThrowError(err error) {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) GetVariable(name string) *ZVariable {
	for _, variable := range g.Variables {
		if variable.Name == name {
			return variable
		}
	}

	return nil
}

func (g *GlobalContext) SetVariable(variable *ZVariable) {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) GetClass(s string) {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) GetFunction(s string) {
	//TODO implement me
	panic("implement me")
}

func (g *GlobalContext) Exec() {}

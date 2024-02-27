package main

type footware interface {
	do() error
}

type personStub struct {
	fw footware
}

func newPerson(fw footware) *personStub {
	return &personStub{fw: fw}
}

func (p personStub) do() error {
	if err := p.fw.do(); err != nil {
		return err
	}
	return nil
}

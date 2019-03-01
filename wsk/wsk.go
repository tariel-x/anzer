package wsk

import "io"

type Wsk struct {
}

func New() Wsk {
	return Wsk{}
}

func (w Wsk) Update(action io.Reader, name, runtime string) error {
	return nil
}

func (w Wsk) Create(action io.Reader, name, runtime string) error {
	return nil
}

func (w Wsk) Upsert(action io.Reader, name, runtime string) error {
	return nil
}

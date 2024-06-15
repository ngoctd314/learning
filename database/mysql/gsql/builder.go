package gsql

import (
	"strings"

	"gorm.io/gorm"
)

type Builder struct {
	cond        *Cond
	listPreload []*preload
}

type builderOpt func(*Builder)

func NewBuilder(opts ...builderOpt) *Builder {
	b := new(Builder)

	for _, opt := range opts {
		opt(b)
	}

	return b
}

func (b *Builder) Build(tx *gorm.DB) *gorm.DB {
	// build condition
	if b.cond != nil {
		tx = b.cond.applyTx(tx)
	}

	// build preload
	for _, v := range b.listPreload {
		tx = v.applyTx(tx)
	}

	// build join

	return tx
}

func WithCondition(cond *Cond) builderOpt {
	return func(b *Builder) {
		if cond != nil {
			b.cond = cond
		}
	}
}

func WithPreload(table string, cond *Cond) builderOpt {
	return func(b *Builder) {
		if notEmptyString(table) {
			b.listPreload = append(b.listPreload, &preload{
				table: table,
				cond:  cond,
			})
		}
	}
}

func notEmptyString(s string) bool {
	return strings.TrimSpace(s) != ""
}

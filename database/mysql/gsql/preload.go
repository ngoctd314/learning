package gsql

import "gorm.io/gorm"

type preload struct {
	cond  *Cond
	table string
}

func (p *preload) applyTx(tx *gorm.DB) *gorm.DB {
	if notEmptyString(p.table) {
		tx = tx.Preload(p.table, func(tmpTx *gorm.DB) *gorm.DB {
			if p.cond != nil {
				tmpTx = p.cond.applyTx(tmpTx)
			}
			return tmpTx
		})
	}

	return tx
}

package main

import (
	"log/slog"
	"sync"
)

type serializable struct {
	rows map[int]any
	rwm  map[int]*sync.RWMutex
}

func (i *serializable) execSelect(pk int) {
	i.rwm[pk].RLock()
	slog.Info("SELECT", "rows", i.rows[pk])
}

func (i *serializable) execUpdate(pk int) {
	i.rwm[pk].Lock()
	slog.Info("UPDATE", "rows", i.rows[pk])
}

func (i *serializable) commit(pk int) {
	if i.rwm[pk].TryRLock() {
		i.rwm[pk].RUnlock()
	}
	if i.rwm[pk].TryLock() {
		i.rwm[pk].Unlock()
	}
}

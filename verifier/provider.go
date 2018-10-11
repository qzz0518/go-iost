package verifier

import (
	"github.com/iost-official/go-iost/core/tx"
	"github.com/iost-official/go-iost/vm"
)

type ProviderImpl struct {
	cache    []*tx.Tx
	iter     vm.TxIter
	droplist map[*tx.Tx]error
}

func NewProvider(iter vm.TxIter) *ProviderImpl {
	return &ProviderImpl{
		cache:    make([]*tx.Tx, 0),
		droplist: make(map[*tx.Tx]error),
		iter:     iter,
	}
}

func (p *ProviderImpl) Tx() *tx.Tx {
	if len(p.cache) != 0 {
		t := p.cache[len(p.cache)-1]
		p.cache = p.cache[:len(p.cache)-1]
		return t
	}
	return p.iter.Next()
}
func (p *ProviderImpl) Return(t *tx.Tx) {
	p.cache = append(p.cache, t)
}
func (p *ProviderImpl) Drop(t *tx.Tx, err error) {
	p.droplist[t] = err
}
func (p *ProviderImpl) List() (a []*tx.Tx, b []error) {
	a = make([]*tx.Tx, 0)
	b = make([]error, 0)
	for k, v := range p.droplist {
		a = append(a, k)
		b = append(b, v)
	}
	return
}

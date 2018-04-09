package bills

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

type bill struct {
	tok        tokens.Token
	clients    []bills.Payer
	processors []bills.Payee
	leaders    []bills.Payee
	verifiers  []bills.Payee
	blocker    bills.Payee
}

func createBill(tok tokens.Token, clients []bills.Payer, processors []bills.Payee, leaders []bills.Payee, verifiers []bills.Payee, blocker bills.Payee) bills.Bill {
	out := bill{
		tok:        tok,
		clients:    clients,
		processors: processors,
		leaders:    leaders,
		verifiers:  verifiers,
		blocker:    blocker,
	}

	return &out
}

// returns the token
func (bi *bill) GetToken() tokens.Token {
	return bi.tok
}

// GetClient the clients
func (bi *bill) GetClients() []bills.Payer {
	return bi.clients
}

// GetProcessors the processors
func (bi *bill) GetProcessors() []bills.Payee {
	return bi.processors
}

// GetLeaders the leaders
func (bi *bill) GetLeaders() []bills.Payee {
	return bi.leaders
}

// GetVerifiers the verifiers
func (bi *bill) GetVerifiers() []bills.Payee {
	return bi.verifiers
}

// GetBlocker the blocker
func (bi *bill) GetBlocker() bills.Payee {
	return bi.blocker
}

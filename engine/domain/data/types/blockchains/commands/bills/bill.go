package bills

import (
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
)

// Bill represents a bill to pay block nodes
type Bill interface {
	GetToken() tokens.Token
	GetClients() []Payer
	GetProcessors() []Payee
	GetLeaders() []Payee
	GetVerifiers() []Payee
	GetBlocker() Payee
}

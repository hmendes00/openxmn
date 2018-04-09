package bills

import (
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	org_servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	usr_wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

// BillBuilderFactory represents a concrete BillBuilderFactory implementation
type BillBuilderFactory struct {
	payeeBuilderFactory bills.PayeeBuilderFactory
	payerBuilderFactory bills.PayerBuilderFactory
	servs               map[string]org_servers.Server
	wals                map[string]usr_wallets.Wallet
}

// CreateBillBuilderFactory creates a new BillBuilderFactory instance
func CreateBillBuilderFactory(payeeBuilderFactory bills.PayeeBuilderFactory, payerBuilderFactory bills.PayerBuilderFactory, servs map[string]org_servers.Server, wals map[string]usr_wallets.Wallet) bills.BillBuilderFactory {
	out := BillBuilderFactory{
		payeeBuilderFactory: payeeBuilderFactory,
		payerBuilderFactory: payerBuilderFactory,
		servs:               servs,
		wals:                wals,
	}
	return &out
}

// Create creates a new BillBuilder instance
func (fac *BillBuilderFactory) Create() bills.BillBuilder {
	out := createBillBuilder(fac.payeeBuilderFactory, fac.payerBuilderFactory, fac.servs, fac.wals)
	return out
}

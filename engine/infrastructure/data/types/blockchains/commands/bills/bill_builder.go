package bills

import (
	"encoding/json"
	"errors"

	validated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated"
	bills "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/commands/bills"
	signed_trs "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	org_servers "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/servers"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	usr_wallets "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users/wallets"
)

type billBuilder struct {
	payeeBuilderFactory bills.PayeeBuilderFactory
	payerBuilderFactory bills.PayerBuilderFactory
	servs               map[string]org_servers.Server
	wals                map[string]usr_wallets.Wallet
	tok                 tokens.Token
	signedBlk           validated.SignedBlock
}

func createBillBuilder(payeeBuilderFactory bills.PayeeBuilderFactory, payerBuilderFactory bills.PayerBuilderFactory, servs map[string]org_servers.Server, wals map[string]usr_wallets.Wallet) bills.BillBuilder {
	out := billBuilder{
		payeeBuilderFactory: payeeBuilderFactory,
		payerBuilderFactory: payerBuilderFactory,
		servs:               servs,
		wals:                wals,
		tok:                 nil,
		signedBlk:           nil,
	}

	return &out
}

// Create initializes the bill builder
func (build *billBuilder) Create() bills.BillBuilder {
	build.tok = nil
	build.signedBlk = nil
	return build
}

// WithToken adds a token to the bill builder
func (build *billBuilder) WithToken(tok tokens.Token) bills.BillBuilder {
	build.tok = tok
	return build
}

// WithBlock adds a signed block to the bill builder
func (build *billBuilder) WithBlock(signedBlk validated.SignedBlock) bills.BillBuilder {
	build.signedBlk = signedBlk
	return build
}

// Now builds a new bill instance
func (build *billBuilder) Now() (bills.Bill, error) {

	retrieveBlockerPayee := func() (bills.Payee, error) {
		signedBlkJS, signedBlkJSErr := json.Marshal(build.signedBlk)
		if signedBlkJSErr != nil {
			return nil, signedBlkJSErr
		}

		//incoming bandwidth:
		inBw := len(signedBlkJS)

		//create the payee builder:
		payeeBuilder := build.payeeBuilderFactory.Create().Create().WithIncomingBandwidthInBytes(inBw)

		//retrieve the server:
		usrIDAsString := build.signedBlk.GetSignature().GetUser().GetMetaData().GetID().String()
		if serv, ok := build.servs[usrIDAsString]; ok {
			payeeBuilder.WithServer(serv)
		}

		payee, payeeErr := payeeBuilder.Now()
		if payeeErr != nil {
			return nil, payeeErr
		}

		return payee, nil
	}

	retrieveVerifierPayees := func() ([]bills.Payee, error) {
		signedBlk := build.signedBlk.GetBlock().GetBlock()
		signedBlkJS, signedBlkJSErr := json.Marshal(signedBlk)
		if signedBlkJSErr != nil {
			return nil, signedBlkJSErr
		}

		//incoming bandwidth:
		inBw := len(signedBlkJS)

		//create the payee builder:
		payeeBuilder := build.payeeBuilderFactory.Create().Create().WithIncomingBandwidthInBytes(inBw)

		payees := []bills.Payee{}
		sigs := build.signedBlk.GetBlock().GetSignatures().GetSignatures()
		for _, oneSig := range sigs {
			sigJS, sigJSErr := json.Marshal(oneSig)
			if sigJSErr != nil {
				return nil, sigJSErr
			}

			//outgoing bandwidth:
			outBw := inBw + len(sigJS)

			//add the outgoing bandwidth:
			payeeBuilder.WithOutgoingBandwidthInBytes(outBw)

			//retrieve the server:
			usrIDAsString := build.signedBlk.GetSignature().GetUser().GetMetaData().GetID().String()
			if serv, ok := build.servs[usrIDAsString]; ok {
				payeeBuilder.WithServer(serv)
			}

			payee, payeeErr := payeeBuilder.Now()
			if payeeErr != nil {
				return nil, payeeErr
			}

			payees = append(payees, payee)
		}

		return payees, nil
	}

	retrieveLeaderPayees := func() ([]bills.Payee, error) {
		payees := []bills.Payee{}
		aggregatedSignedTrs := build.signedBlk.GetBlock().GetBlock().GetBlock().GetTransactions()
		for _, oneSignedAggrTrs := range aggregatedSignedTrs {
			trs := oneSignedAggrTrs.GetTransactions()
			trsJS, trsJSErr := json.Marshal(trs)
			if trsJSErr != nil {
				return nil, trsJSErr
			}

			signedAggrTrsJS, signedAggrTrsJSErr := json.Marshal(oneSignedAggrTrs)
			if signedAggrTrsJSErr != nil {
				return nil, signedAggrTrsJSErr
			}

			//bandwidth:
			inBw := len(trsJS)
			outBw := len(signedAggrTrsJS)

			//create the payee builder:
			payeeBuilder := build.payeeBuilderFactory.Create().Create().WithIncomingBandwidthInBytes(inBw).WithOutgoingBandwidthInBytes(outBw)

			//retrieve the server:
			usrIDAsString := oneSignedAggrTrs.GetSignature().GetUser().GetMetaData().GetID().String()
			if serv, ok := build.servs[usrIDAsString]; ok {
				payeeBuilder.WithServer(serv)
			}

			payee, payeeErr := payeeBuilder.Now()
			if payeeErr != nil {
				return nil, payeeErr
			}

			payees = append(payees, payee)
		}

		return payees, nil
	}

	retrieveProcessorPayees := func() ([]bills.Payee, error) {
		retrieveProcessorPayeesFromAtomicTransactions := func(signedAtomicTrs signed_trs.AtomicTransactions) ([]bills.Payee, error) {
			payees := []bills.Payee{}
			atomicTrs := signedAtomicTrs.GetTransactions()
			for _, oneAtomicTrs := range atomicTrs {

				atomicTrsJS, atomicTrsJSErr := json.Marshal(oneAtomicTrs)
				if atomicTrsJSErr != nil {
					return nil, atomicTrsJSErr
				}

				trs := oneAtomicTrs.GetTransactions()
				trsJS, trsJSErr := json.Marshal(trs)
				if trsJSErr != nil {
					return nil, trsJSErr
				}

				//bandwidth usage:
				inBw := len(trsJS)
				outBw := len(atomicTrsJS)

				//create the payee builder:
				payeeBuilder := build.payeeBuilderFactory.Create().Create().WithIncomingBandwidthInBytes(inBw).WithOutgoingBandwidthInBytes(outBw)

				//retrieve the server:
				usrIDAsString := oneAtomicTrs.GetSignature().GetUser().GetMetaData().GetID().String()
				if serv, ok := build.servs[usrIDAsString]; ok {
					payeeBuilder.WithServer(serv)
				}

				payee, payeeErr := payeeBuilder.Now()
				if payeeErr != nil {
					return nil, payeeErr
				}

				payees = append(payees, payee)
			}

			return payees, nil
		}

		retrieveProcessorPayeesFromTransactions := func(signedTrs signed_trs.Transactions) ([]bills.Payee, error) {
			payees := []bills.Payee{}
			signedTrans := signedTrs.GetTransactions()
			for _, oneSignedTrs := range signedTrans {
				signedTrsJS, signedTrsJSErr := json.Marshal(oneSignedTrs)
				if signedTrsJSErr != nil {
					return nil, signedTrsJSErr
				}

				trs := oneSignedTrs.GetTransaction()
				trsJS, trsJSErr := json.Marshal(trs)
				if trsJSErr != nil {
					return nil, trsJSErr
				}

				//bandwidth usage:
				inBw := len(trsJS)
				outBw := len(signedTrsJS)

				//create the payee builder:
				payeeBuilder := build.payeeBuilderFactory.Create().Create().WithIncomingBandwidthInBytes(inBw).WithOutgoingBandwidthInBytes(outBw)

				//retrieve the server:
				usrIDAsString := oneSignedTrs.GetSignature().GetUser().GetMetaData().GetID().String()
				if serv, ok := build.servs[usrIDAsString]; ok {
					payeeBuilder.WithServer(serv)
				}

				payee, payeeErr := payeeBuilder.Now()
				if payeeErr != nil {
					return nil, payeeErr
				}

				payees = append(payees, payee)
			}

			return payees, nil
		}

		payees := []bills.Payee{}
		aggregatedSignedTrs := build.signedBlk.GetBlock().GetBlock().GetBlock().GetTransactions()
		for _, oneAggregatedSignedTrs := range aggregatedSignedTrs {
			aggregatedTrs := oneAggregatedSignedTrs.GetTransactions()
			if aggregatedTrs.HasAtomicTransactions() {
				atomicTrs := aggregatedTrs.GetAtomicTransactions()
				atomicPayees, atomicPayeesErr := retrieveProcessorPayeesFromAtomicTransactions(atomicTrs)
				if atomicPayeesErr != nil {
					return nil, atomicPayeesErr
				}

				for _, oneAtomicPayee := range atomicPayees {
					payees = append(payees, oneAtomicPayee)
				}
			}

			if aggregatedTrs.HasTransactions() {
				trs := aggregatedTrs.GetTransactions()
				trsPayees, trsPayeesErr := retrieveProcessorPayeesFromTransactions(trs)
				if trsPayeesErr != nil {
					return nil, trsPayeesErr
				}

				for _, oneTrsPayee := range trsPayees {
					payees = append(payees, oneTrsPayee)
				}
			}
		}

		return payees, nil
	}

	retrieveClientPayers := func() ([]bills.Payer, error) {

		retrieveClientPayersFromAtomicTransactions := func(signedAtomicTrs signed_trs.AtomicTransactions) ([]bills.Payer, error) {
			payers := []bills.Payer{}
			atomicTrs := signedAtomicTrs.GetTransactions()
			for _, oneAtomicTrs := range atomicTrs {

				atomicTrsJS, atomicTrsJSErr := json.Marshal(oneAtomicTrs)
				if atomicTrsJSErr != nil {
					return nil, atomicTrsJSErr
				}

				//bandwidth usage:
				outBw := len(atomicTrsJS)

				//create the payee builder:
				payerBuilder := build.payerBuilderFactory.Create().Create().WithOutgoingBandwidthInBytes(outBw)

				//retrieve the server:
				usrIDAsString := oneAtomicTrs.GetSignature().GetUser().GetMetaData().GetID().String()
				if wal, ok := build.wals[usrIDAsString]; ok {
					payerBuilder.WithWallet(wal)
				}

				payer, payerErr := payerBuilder.Now()
				if payerErr != nil {
					return nil, payerErr
				}

				payers = append(payers, payer)
			}

			return payers, nil
		}

		retrieveClientPayersFromTransactions := func(signedTrs signed_trs.Transactions) ([]bills.Payer, error) {
			payers := []bills.Payer{}
			trans := signedTrs.GetTransactions()
			for _, oneTrs := range trans {

				trsJS, trsJSErr := json.Marshal(oneTrs)
				if trsJSErr != nil {
					return nil, trsJSErr
				}

				//bandwidth usage:
				outBw := len(trsJS)

				//create the payee builder:
				payerBuilder := build.payerBuilderFactory.Create().Create().WithOutgoingBandwidthInBytes(outBw)

				//retrieve the server:
				usrIDAsString := oneTrs.GetSignature().GetUser().GetMetaData().GetID().String()
				if wal, ok := build.wals[usrIDAsString]; ok {
					payerBuilder.WithWallet(wal)
				}

				payer, payerErr := payerBuilder.Now()
				if payerErr != nil {
					return nil, payerErr
				}

				payers = append(payers, payer)
			}

			return payers, nil
		}

		payers := []bills.Payer{}
		aggregatedSignedTrs := build.signedBlk.GetBlock().GetBlock().GetBlock().GetTransactions()
		for _, oneAggregatedSignedTrs := range aggregatedSignedTrs {
			aggregatedTrs := oneAggregatedSignedTrs.GetTransactions()
			if aggregatedTrs.HasAtomicTransactions() {
				atomicTrs := aggregatedTrs.GetAtomicTransactions()
				atomicPayers, atomicPayersErr := retrieveClientPayersFromAtomicTransactions(atomicTrs)
				if atomicPayersErr != nil {
					return nil, atomicPayersErr
				}

				for _, oneAtomicPayer := range atomicPayers {
					payers = append(payers, oneAtomicPayer)
				}
			}

			if aggregatedTrs.HasTransactions() {
				trs := aggregatedTrs.GetTransactions()
				trsPayers, trsPayersErr := retrieveClientPayersFromTransactions(trs)
				if trsPayersErr != nil {
					return nil, trsPayersErr
				}

				for _, oneTrsPayer := range trsPayers {
					payers = append(payers, oneTrsPayer)
				}
			}
		}

		return payers, nil
	}

	if build.tok == nil {
		return nil, errors.New("the token is mandatory in order to build a bill instance")
	}

	if build.signedBlk == nil {
		return nil, errors.New("the signed block is mandatory in order to build a bill instance")
	}

	blocker, blockerErr := retrieveBlockerPayee()
	if blockerErr != nil {
		return nil, blockerErr
	}

	verifiers, verifiersErr := retrieveVerifierPayees()
	if verifiersErr != nil {
		return nil, verifiersErr
	}

	leaders, leadersErr := retrieveLeaderPayees()
	if leadersErr != nil {
		return nil, leadersErr
	}

	processors, processorsErr := retrieveProcessorPayees()
	if processorsErr != nil {
		return nil, processorsErr
	}

	clients, clientsErr := retrieveClientPayers()
	if clientsErr != nil {
		return nil, clientsErr
	}

	out := createBill(build.tok, clients, processors, leaders, verifiers, blocker)
	return out, nil
}

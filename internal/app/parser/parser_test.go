package parser

import (
	"context"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"

	"github.com/iam047801/tonidx/internal/app"
	"github.com/iam047801/tonidx/internal/core/db"
)

var _testService *Service

var ctx = context.Background()

func testService(t *testing.T) *Service {
	if _testService != nil {
		return _testService
	}

	conn, err := db.Connect(ctx, "clickhouse://localhost:9000/idx?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	s, err := NewService(context.Background(), &app.ParserConfig{DB: conn})
	if err != nil {
		t.Fatal(err)
	}

	_testService = s
	return _testService
}

func getCurrentMaster(t *testing.T) *tlb.BlockInfo {
	s := testService(t)

	master, err := s.api.GetMasterchainInfo(ctx)
	if err != nil {
		t.Fatal(errors.Wrap(err, "cannot get masterchain info"))
	}
	master, err = s.api.LookupBlock(ctx, master.Workchain, master.Shard, master.SeqNo)
	if err != nil {
		t.Fatal(errors.Wrap(err, "lookup block"))
	}

	return master
}

func getTransactionOnce(t *testing.T, addr *address.Address, lt uint64, txHash []byte) *tlb.Transaction {
	transactions, err := testService(t).api.ListTransactions(ctx, addr, 1, lt, txHash)
	if err != nil {
		t.Fatal(err)
	}
	if len(transactions) == 0 {
		t.Fatal(fmt.Errorf("no transactions"))
	}
	return transactions[0]
}

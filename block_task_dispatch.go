package main

import (
	"sync"
	"time"

	"github.com/blocto/solana-go-sdk/rpc"
)

const (
	Commitment = rpc.CommitmentFinalized
)

type BlockTaskDispatch struct {
	cli rpc.RpcClient
}

func NewBlockTaskDispatch() *BlockTaskDispatch {
	cli := rpc.NewRpcClient(conf.Solana.RpcEndpoint)
	return &BlockTaskDispatch{cli: cli}
}

func (btd *BlockTaskDispatch) keepDispatchTaskMock(wg *sync.WaitGroup, startSlot uint64, count uint64, taskCh chan uint64) {
	defer wg.Done()

	start := startSlot
	end := startSlot + count
	for start < end {
		taskCh <- start
		start += 1
	}

	close(taskCh)
}

func (btd *BlockTaskDispatch) keepDispatchTask(wg *sync.WaitGroup, startSlot uint64, count uint64, taskCh chan uint64) {
	defer wg.Done()

	config, err := btd.cli.GetSlotWithConfig(nil, rpc.GetSlotConfig{Commitment: Commitment})
	if err != nil {
		time.Sleep(conf.Solana.RpcReqInterval)
	}

	close(taskCh)
}

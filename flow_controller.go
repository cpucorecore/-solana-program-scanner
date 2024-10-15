package main

import (
	"sync"
	"time"
)

type FlowController interface {
	onOk()
	onErr()
}

type flowController struct {
	okCnt int
	okMu  sync.Mutex

	okTimes   []time.Time
	okTimesMu sync.Mutex

	errCnt int
	errMu  sync.Mutex

	targetTps   int
	errWaitUnit time.Duration
}

func NewFlowController(waitUnit time.Duration) FlowController {
	return &flowController{
		errCnt:      0,
		okCnt:       0,
		errWaitUnit: waitUnit,
	}
}

func (fc *flowController) onOk() {
	now := time.Now()
	oneSecondAgo := now.Add(-time.Second)
	tps := 0
	fc.okTimesMu.Lock()
	fc.okTimes = append(fc.okTimes, now)
	for _, ot := range fc.okTimes {
		if ot.Before(oneSecondAgo) {
			tps++
			fc.okTimes = append(fc.okTimes[:1], fc.okTimes[2:]...)
		}
	}
	fc.okTimesMu.Unlock()

	fc.okMu.Lock()
	fc.okCnt++
	fc.okMu.Unlock()

	fc.errMu.Lock()
	fc.errCnt = 0
	fc.errMu.Unlock()
}

func (fc *flowController) onErr() {
	var errCnt int

	fc.errMu.Lock()
	fc.errCnt++
	errCnt = fc.errCnt
	fc.errMu.Unlock()

	time.Sleep(fc.errWaitUnit * time.Duration(errCnt))
}

var _ FlowController = &flowController{}

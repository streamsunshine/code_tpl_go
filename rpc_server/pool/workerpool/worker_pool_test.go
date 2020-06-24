package workerpool_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"git.code.oa.com/trpc-go/trpc-go/pool/workerpool"
)

func TestNew(t *testing.T) {
	var wp *workerpool.WorkerPool
	wp = workerpool.New()
	assert.NotNil(t, wp)

	wp = workerpool.New(workerpool.WithMaxWorkersCount(10))
	assert.NotNil(t, wp)
}

func TestWorkerPool_Run(t *testing.T) {
	var wp *workerpool.WorkerPool
	wp = workerpool.New()
	assert.NotNil(t, wp)

	err := wp.Run(func() {})
	assert.Nil(t, err)
}

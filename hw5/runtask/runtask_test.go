package runtask_test

import (
	"hw5/runtask"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	testCases := []struct {
		desc  string
		funcs []func() error
		N     int
		M     int
		wont  int
	}{
		{desc: "Not Error", funcs: runtask.GenTask(20, 0), N: 10, M: 0, wont: 20},
		{desc: "Not Error2", funcs: runtask.GenTask(5, 0), N: 10, M: 0, wont: 5},
		{desc: "LessOrEqual N+M", funcs: runtask.GenTask(20, 3), N: 10, M: 3, wont: 13},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := runtask.Run(tC.funcs, tC.N, tC.M)
			switch tC.desc {
			case "LessOrEqual N+M":
				assert.LessOrEqual(t, got, tC.wont)
			default:
				assert.Equal(t, tC.wont, got)
			}
		})
	}
}

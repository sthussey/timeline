package timeline

import (
	"golang.org/x/sys/unix"
	"syscall"
	"testing"
	"time"
)

func TestSignalRecv(t *testing.T) {
	t.Parallel()
	inputs := make(map[string]string)
	inputs["signal"] = "SIGUSR1"

	sm := getDefaultSyncMap()
	syncFunc, ok := sm["SignalRecv"]
	if !ok {
		t.Errorf("Default Sync Map does not contain 'SignalRecv'")
	}

	sigRecv := make(chan struct{})
	waitDur, _ := time.ParseDuration("2s")
	timeoutDur, _ := time.ParseDuration("10s")

	go func() {
		syncFunc(inputs, nil)
		close(sigRecv)
	}()
	// Avoid a race condition setting up the signal handler
	wait := time.NewTimer(waitDur)
	<-wait.C

	timeout := time.NewTimer(timeoutDur)
	unix.Kill(unix.Getpid(), syscall.SIGUSR1)

	select {
	case <-sigRecv:
		return
	case <-timeout.C:
		t.Errorf("Timed out waiting for SignalRecv sync")
	}
}

func TestTimerWait(t *testing.T) {
	t.Parallel()
	inputs := make(map[string]string)
	inputs["duration"] = "5s"

	sm := getDefaultSyncMap()
	syncFunc, ok := sm["TimerWait"]
	if !ok {
		t.Errorf("Default Sync Map does not contain 'TimerWait'")
	}

	timerExp := make(chan struct{})
	lowDur, _ := time.ParseDuration("3s")
	highDur, _ := time.ParseDuration("7s")
	lowFired := false

	go func() {
		syncFunc(inputs, nil)
		close(timerExp)
	}()

	low := time.NewTimer(lowDur)
	high := time.NewTimer(highDur)

	for {
		select {
		case <-low.C:
			lowFired = true
		case <-timerExp:
			if lowFired {
				return
			} else {
				t.Errorf("Timer expired too quickly, low watermark didn't fire.")
			}
		case <-high.C:
			t.Errorf("Timer didn't fire in time, timeout.")
		}
	}
}

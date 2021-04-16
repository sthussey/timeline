package timeline

import (
	"syscall"
	"testing"
	"time"
	"golang.org/x/sys/unix"
)

func TestSignalRecv(t *testing.T){
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

	go func(){	syncFunc(inputs)
				close(sigRecv)}()
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

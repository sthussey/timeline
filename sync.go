package timeline

import (
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"time"
)

func basicTimer(inputs interface{}, vars map[string]interface{}) error {
	inputMap, ok := inputs.(map[string]string)

	if !ok {
		return fmt.Errorf("Error: basicTimer input wrong type, should be map[string]string")
	}

	durationString, ok := inputMap["duration"]

	if !ok {
		return fmt.Errorf("Error: basicTimer input requires a 'duration' key")
	}

	duration, err := time.ParseDuration(durationString)

	if err != nil {
		return fmt.Errorf("Error: 'duration' not a valid duration string")
	}

	timer := time.NewTimer(duration)

	<-timer.C

	return nil
}

func receiveSignal(inputs interface{}, vars map[string]interface{}) error {
	inputMap, ok := inputs.(map[string]string)

	if !ok {
		return fmt.Errorf("Error: receiveSignal input wrong type, should be map[string]string")
	}

	signalName, ok := inputMap["signal"]

	if !ok {
		return fmt.Errorf("Error: receiveSignal input requires a 'signal' key")
	}

	sig := unix.SignalNum(signalName)

	if sig == 0 {
		return fmt.Errorf("Error: signal %s not defined on this platform.", signalName)
	}

	c := make(chan os.Signal, 2)

	signal.Notify(c, sig)
	<-c

	signal.Reset(sig)
	return nil
}

func getDefaultSyncMap() map[string]func(interface{}, map[string]interface{}) error {
	sm := make(map[string]func(interface{}, map[string]interface{}) error)
	sm["TimerWait"] = basicTimer
	sm["SignalRecv"] = receiveSignal
	return sm
}

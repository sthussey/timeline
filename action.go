package timeline

import (
	"fmt"
	"io"
	"os"
)

func logMessage(inputs interface{}, vars map[string]interface{}) error {
	inputMap, ok := inputs.(map[string]string)

	if !ok {
		return fmt.Errorf("Error: logMessage input wrong type, should be map[string]string")
	}

	msg, ok := inputMap["message"]
	output, ok := inputMap["output"]

	if !ok {
		return fmt.Errorf("Error: logMessage input requires a 'message' key")
	}

	var f io.Writer

	if output == "stdout" {
		f = os.Stdout
	} else if output == "stderr" {
		f = os.Stderr
	} else {
		v, ok := vars[output]
		if !ok {
			return fmt.Errorf("Error: logMessage does not recognize output %s", output)
		}
		f, ok = v.(io.Writer)
		if !ok {
			return fmt.Errorf("Error: logMessage requires output to be an io.Writer")
		}
	}

	fmt.Fprintf(f, "%s\n", msg)

	return nil
}

func getDefaultActionMap() map[string]func(interface{}, map[string]interface{}) error {
	am := make(map[string]func(interface{}, map[string]interface{}) error)
	am["LogMessage"] = logMessage
	return am
}

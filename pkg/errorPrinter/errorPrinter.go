package errorPrinter

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func PrintCallerFunctionName(err error) {
	pc, filePath, lineNumber, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("No information about caller function")
		return
	}
	caller := runtime.FuncForPC(pc)
	functionName := caller.Name()

	file, err_ := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 755)
	if err_ != nil {
		fmt.Println("Error opening file__:", err_)
		return
	}
	defer file.Close()

	if err != nil {
		currentTime := time.Now().Format("2006-01-02 15:04:05") // Получение текущего времени
		_, err_ := file.WriteString(fmt.Sprintf("Time: %s, Error: %s, Function name: %s File: %s, Line: %d\n", currentTime, err.Error(), functionName, filePath, lineNumber))
		if err_ != nil {
			fmt.Println("Error writing to file:", err_)
		}
	}
}

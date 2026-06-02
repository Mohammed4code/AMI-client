package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)
 
var (
	logFile *os.File
	logWriter *bufio.Writer
	logMu sync.Mutex
)

func initLogger(){
	var err error
	logFile , err = os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err != nil {
		log.Fatalf("Error openig log file: %v", err )
	}

	logWriter = bufio.NewWriter(logFile)

	logtoFile(strings.Repeat("=",80))
	logtoFile(fmt.Sprintf("[%s]Start client AMI", time.Now().Format("2006-01-02-15:04:05") ))
	logtoFile(strings.Repeat("=",80))


}

func logtoFile(msg string)  {
		logMu.Lock()
		defer logMu.Unlock()
		fmt.Fprintln(logWriter,msg)
		logWriter.Flush()
	}
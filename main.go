package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/staskobzar/goami2"
)

func main()  {
	initLogger()
	defer logFile.Close()


		sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n\n[!] Shutting down gracefully...")
		fmt.Println("[!] Flushing logs and closing file...")
		if logWriter != nil {
			logWriter.Flush()
		}
		if logFile != nil {
			logFile.Close()
		}
		fmt.Println("[✓] Done. Exiting.")
		os.Exit(0)
	}()
	

	conn, err := net.Dial("tcp",AMI_HOST+":"+AMI_PORT)
	if err != nil {
		logtoFile("Connuicantion error :" + err.Error())
		log.Fatalf("connuicantion error :%s\n",err)

	}
	defer conn.Close()


	client , err := goami2.NewClient(conn,AMI_USER,AMI_PASS)
	if err != nil{
		logtoFile("login error :"+ err.Error())
		log.Fatalf("login error :%s\n",err)
	}
	defer client.Close()


	    logtoFile("Successfully connected to Asterisk AMI")
	    fmt.Println("Successfully connected to Asterisk AMI")
	    fmt.Println("Events are being recorded at:", LOG_FILE)
	    fmt.Println("   (Press Ctrl+C to exit)")

	for {
		select{
		case msg :=<- client.AllMessages():
			handleMessage(msg)
			
		case err =<-client.Err():
			logtoFile("Connecion error:"+err.Error())
			log.Fatalf("connecion:%s\n",err)
		}
	}
}
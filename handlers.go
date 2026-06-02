package main

import (
	"encoding/json"
	"fmt"
	"time"

	
	"github.com/staskobzar/goami2"
)
func handleMessage(msg *goami2.Message) {
    timestamp := time.Now().Format("2006-01-02 15:04:05.000")
    jsonData := msg.JSON()
    
    logtoFile(fmt.Sprintf("\n[%s] New message:", timestamp))
    logtoFile(fmt.Sprintf(" %s", jsonData))

    var fields map[string]interface{}
    if err := json.Unmarshal([]byte(jsonData), &fields); err != nil {
        
        logtoFile(fmt.Sprintf("  ERROR parsing JSON: %v", err))
        return
    }
    
  
    for key, value := range fields {
        logtoFile(fmt.Sprintf("  %s: %v", key, value))
    }
    
    if eventType, ok := fields["Event"].(string); ok && eventType == "Cdr" {
        handleCDR(fields)
    }
    
    // تحتاج لتحويل map[string]interface{} إلى map[string]string قبل تمريرها
    printSummary(convertToStringMap(fields))
}

// دالة مساعدة للتحويل
func convertToStringMap(fields map[string]interface{}) map[string]string {
    result := make(map[string]string)
    for k, v := range fields {
        if v != nil {
            result[k] = fmt.Sprintf("%v", v)
        }
    }
    return result
}
func handleCDR(fields map[string]interface{}) {
    logtoFile("")
    logtoFile(" ===== details CDR =====")
    logtoFile(fmt.Sprintf("  source (Source): %s", fmt.Sprint(fields["Source"])))
    logtoFile(fmt.Sprintf("  destination (Destination): %s", fmt.Sprint(fields["Destination"])))
    logtoFile(fmt.Sprintf("  callerID (CallerID): %s", fmt.Sprint(fields["CallerID"])))
    logtoFile(fmt.Sprintf("  channel (Channel): %s", fmt.Sprint(fields["Channel"])))
    logtoFile(fmt.Sprintf("  starttime (StartTime): %s", fmt.Sprint(fields["StartTime"])))
    logtoFile(fmt.Sprintf("  answertime (AnswerTime): %s", fmt.Sprint(fields["AnswerTime"])))
    logtoFile(fmt.Sprintf("  endtime (EndTime): %s", fmt.Sprint(fields["EndTime"])))
    logtoFile(fmt.Sprintf("  duration (Duration): %s second", fmt.Sprint(fields["Duration"])))
    logtoFile(fmt.Sprintf("  billableseconds (BillableSeconds): %s", fmt.Sprint(fields["BillableSeconds"])))
    logtoFile(fmt.Sprintf("  disposition (Disposition): %s", fmt.Sprint(fields["Disposition"])))
    logtoFile("=========================")
    logtoFile("")
}


func printSummary(fields map[string]string) {
    eventName := fields["Event"]

    switch eventName {
    case "Newchannel":
        fmt.Printf("[%s] New channel: %s\n", eventName, fields["Channel"])
    case "Newstate":
        fmt.Printf("[%s] New channel status: %s -> %s\n", eventName, fields["Channel"], fields["ChannelStateDesc"])
    case "Hangup":
        fmt.Printf("[%s] The call ends: %s\n", eventName, fields["Channel"])
    case "Cdr":
        fmt.Printf("[%s] Record a call: %s -> %s | Duration: %s s | Condition: %s\n",
            eventName, fields["Source"], fields["Destination"], fields["Duration"], fields["Disposition"])
    case "BridgeEnter":
        fmt.Printf("[%s] Enter to the bridge: %s\n", eventName, fields["Channel"])
    }
}
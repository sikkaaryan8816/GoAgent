package index

/*

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "ndlb_encode.h"


int WrapHeader(char *s,int apiReqLen,int awsReqLen,int funcNameLen,int tagslength,short agentType,short messageType){
    int len = 0;
    wrapheader_t wrapHeader;
    memcpy(s, "^",1);
    len += 1;
    
    wrapHeader.wrapheadervar.apiReqLen = apiReqLen ;
    wrapHeader.wrapheadervar.awsReqLen = awsReqLen ;
    wrapHeader.wrapheadervar.funcNameLen = funcNameLen ;
    wrapHeader.wrapheadervar.tagslength = tagslength ;
    wrapHeader.wrapheadervar.agentType = agentType ;
    wrapHeader.wrapheadervar.messageType = messageType;
    wrapHeader.wrapheadervar.whLen = sizeof(wrapheadervar_t)+wrapHeader.wrapheadervar.awsReqLen+wrapHeader.wrapheadervar.apiReqLen+wrapHeader.wrapheadervar.funcNameLen+wrapHeader.wrapheadervar.tagslength+1;

    memcpy(s + len , (char *)&(wrapHeader.wrapheadervar), sizeof(wrapHeader.wrapheadervar));
    len += sizeof(wrapheadervar_t);

  return len;
}

int ValueStore(char *s,char *value,int len,int number)
{

      memcpy(s + len, value ,number);
      len += number;
      return len;
}

int StartTransaction(char *s,int fp_header,int url,int btHeaderValue,int ndCookieSet,int nvCookieSet,int correlationHeader,long long flowpathinstance,long qTimeMS,long long startTimeFP,int len)
{
    transactionStart_t node;
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;
    
    node.transactionStartVar.qTimeMS = qTimeMS ;
    node.transactionStartVar.startTimeFP = startTimeFP;
    node.transactionStartVar.fp_header = fp_header ;
    node.transactionStartVar.url = url ;
    node.transactionStartVar.btHeaderValue = btHeaderValue ;
    node.transactionStartVar.ndCookieSet = ndCookieSet ;
    node.transactionStartVar.nvCookieSet = nvCookieSet ;
    node.transactionStartVar.correlationHeader = correlationHeader ;
    node.transactionStartVar.flowpathinstance = flowpathinstance;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof (transactionStartVar_t) + msgHdr.header_len + node.transactionStartVar.fp_header +
    node.transactionStartVar.ndCookieSet + node.transactionStartVar.nvCookieSet +
    node.transactionStartVar.correlationHeader +
    node.transactionStartVar.btHeaderValue +
    node.transactionStartVar.url + 3;
    msgHdr.msg_type = 2;

    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);


    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node.transactionStartVar), sizeof(node.transactionStartVar));
    len += sizeof(node.transactionStartVar) ;

    return len;
}

int MethodEntry(char *s,int urlParameter,int methodName,int query_string,int mid,long long flowpathinstance,long threadId,long long startTime,int len)
{
    msgHdr_t msgHdr;
    MethodEntry_t node;
    memcpy(s+len, "^",1);
    len += 1;
    
    node.MethodEntryVar.methodName = methodName ;
    node.MethodEntryVar.threadId= threadId;
    node.MethodEntryVar.query_string = query_string ;
    node.MethodEntryVar.urlParameter = urlParameter ;
    node.MethodEntryVar.mid = mid ;
    node.MethodEntryVar.startTime = startTime ;
    node.MethodEntryVar.flowpathinstance = flowpathinstance;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof(MethodEntryVar_t) + msgHdr.header_len + node.MethodEntryVar.methodName +
    node.MethodEntryVar.query_string+ node.MethodEntryVar.urlParameter + 3;
    msgHdr.msg_type = 0;
    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);


    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node.MethodEntryVar), sizeof(MethodEntryVar_t));
    len += sizeof(MethodEntryVar_t);

    return len;
}

int MethodExit(char *s,int statusCode,int mid,int eventType,int isCallout,long duration,long threadId,long long cpuTime,long long flowpathinstance,long long tierCallOutSeqNum,long long endTime,int methodName,int backend_header,int requestNotificationPhase,int len)
{
    MethodExit_t node;
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;

    node.MethodExitVar.statusCode = statusCode ;
    node.MethodExitVar.mid = mid ;
    node.MethodExitVar.eventType = eventType;
    node.MethodExitVar.isCallout = isCallout;
    node.MethodExitVar.duration = duration;
    node.MethodExitVar.threadId = threadId;
    node.MethodExitVar.cpuTime = cpuTime;
    node.MethodExitVar.tierCallOutSeqNum = tierCallOutSeqNum ;
    node.MethodExitVar.endTime = endTime ;
    node.MethodExitVar.methodName = methodName ;
    node.MethodExitVar.backend_header = backend_header;
    node.MethodExitVar.requestNotificationPhase = requestNotificationPhase ;

    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = sizeof(MethodExitVar_t)+ msgHdr.header_len + node.MethodExitVar.methodName + node.MethodExitVar.backend_header +
    node.MethodExitVar.requestNotificationPhase + 3;
    msgHdr.msg_type = 1;
    memcpy(s + len , (char *)&(msgHdr), msgHdr.header_len);
    len += sizeof(msgHdr);


    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(node.MethodExitVar), sizeof(node.MethodExitVar));
    len += sizeof(MethodExitVar_t);

    return len;
}

int EndTransaction(char *s,int statuscode,long long endTime,long long flowpathinstance,long long cpuTime ,int len)
{
    transactionEnd_t transactionEnd;
    int N=sizeof (transactionEnd_t);
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;
    
    transactionEnd.statuscode = statuscode ;
    transactionEnd.endTime = endTime ;
    transactionEnd.flowpathinstance = flowpathinstance;
    transactionEnd.cpuTime = cpuTime;
   
    msgHdr.header_len = sizeof(msgHdr_t);
    msgHdr.total_len = msgHdr.header_len + N + 3;
    msgHdr.msg_type = 3;

    memcpy(s + len , (char *)&(msgHdr), sizeof(msgHdr_t));
    len += sizeof(msgHdr_t);


    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (void *)&transactionEnd, N);
    len += N;

    return len;

}
int last(char *s,int len)
{
    memcpy(s + len, "\n", 1);
}
*/
import "C"

import (
    "fmt"
    "log"
    "net"
    "time"
    "unsafe"
    
)



func Header(buf []byte) C.int {

    var apiReqId = "akkdnfjnflffk"
    var awsReqId = "lksjskwjdkdldl"
    var funcName = "lamdafunion_test1"
    //var agentType= 0
    //var messageType = 0
    var tags = "tierName=tier_test1;ndAppServerHost=server_test1;appName=lamdafunion_test1"

    var apiReqLen = C.int(len(apiReqId))
    var awsReqLen = C.int(len(awsReqId))
    var funcNameLen = C.int(len(funcName))
    var tagslength = C.int(len(tags))
    var agentType = C.short(0)
    var messageType = C.short(0)

    len := C.WrapHeader((*C.char)(unsafe.Pointer(&buf[0])), apiReqLen, awsReqLen, funcNameLen, tagslength, agentType, messageType)
    a := C.CString(apiReqId)
    b := C.CString(awsReqId)
    c := C.CString(funcName)
    d := C.CString(tags)
    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))
    defer C.free(unsafe.Pointer(d))

    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), a, len, apiReqLen)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), b, len, awsReqLen)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), c, len, funcNameLen)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), d, len, tagslength)

    return len
}


type aiRecord struct {
    conn net.Conn
}

var aiRecObj *aiRecord = nil
var err error

func NewAIRecord() *aiRecord {
    r := aiRecord{}
    r.conn, err = net.Dial("udp", "66.220.31.147:1224")
    if err != nil {
        fmt.Printf("Some error %v", err)
    }
    fmt.Println("conn value", r.conn)
    return &r
}

func UDPConnection() {

    fmt.Println("udp_call")

    aiRecObj = NewAIRecord()
    fmt.Println(aiRecObj)

    time.Sleep(1 * time.Second)

}

func condition(a string) C.int {

    if a == "" {
        return 0
    }
    return C.int(len(a))
}

/*func ReceiveMessageFromServer() {
    request := make([]byte, 1024)
    a, err := aiRecObj.conn.Read(request)
    //a :=len(request)
    fmt.Println("request=", string(request), a)
}*/

//func main() {
func create_start_transaction_message(bt_name, nil string) {
    bt_name := "defult"
    UDPConnection()

    StartTransactionMessage(bt_name, "")
    method_entry()
    method_exit()
    end_business_transaction()
    //fmt.Println(generate_bt())
    //bt := 1
    //return bt
    //closeUDP()
}

/*func generate_bt() {
    id := uuid.New().String()
    var i big.Int
    i.SetString(strings.Replace(id, "-", "", 4), 16)
    return i.String()
}*/

func StartTransactionMessage(bt_name string, correlationHeader string) {

    var buf = make([]byte, 1024)
    len := Header(buf)
    fp_header := "dummy_fp_header"
    url := bt_name
    btHeaderValue := "dummy_btHeaderValue"
    ndCookieSet := ""
    nvCookieSet := ""
    correlationHeader = "dummy_correlationHeader"

    var transaction transactionStart_t

    transaction.transactionStartVar.fp_header = (condition(fp_header))
    transaction.transactionStartVar.url = (condition(url))
    transaction.transactionStartVar.btHeaderValue = (condition(btHeaderValue))
    transaction.transactionStartVar.ndCookieSet = (condition(ndCookieSet))
    transaction.transactionStartVar.nvCookieSet = (condition(nvCookieSet))
    transaction.transactionStartVar.correlationHeader = (condition(correlationHeader))
    transaction.transactionStartVar.flowpathinstance = 0
    transaction.transactionStartVar.startTimeFP = 0
    transaction.transactionStartVar.qTimeMS = 0
    var msgHdr msgHdr_t
   
    msgHdr.header_len = 12
    msgHdr.total_len = 12 + 48 + transaction.transactionStartVar.fp_header + transaction.transactionStartVar.url + transaction.transactionStartVar.btHeaderValue + transaction.transactionStartVar.ndCookieSet + transaction.transactionStartVar.nvCookieSet + transaction.transactionStartVar.correlationHeader + 3

    msgHdr.msg_type = 2

    len = C.main3((*C.char)(unsafe.Pointer(&buf[0])), (*C.msgHdr_t)(unsafe.Pointer(&msgHdr)), (*C.transactionStart_t)(unsafe.Pointer(&transaction)), len)
    a := C.CString(fp_header)
    b := C.CString(url)
    c := C.CString(btHeaderValue)
    d := C.CString(ndCookieSet)
    e := C.CString(nvCookieSet)
    f := C.CString(correlationHeader)
    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))
    defer C.free(unsafe.Pointer(d))
    defer C.free(unsafe.Pointer(e))
    defer C.free(unsafe.Pointer(f))

    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), a, len, transaction.transactionStartVar.fp_header)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), b, len, transaction.transactionStartVar.url)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), c, len, transaction.transactionStartVar.btHeaderValue)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), d, len, transaction.transactionStartVar.ndCookieSet)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), e, len, transaction.transactionStartVar.nvCookieSet)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), f, len, transaction.transactionStartVar.correlationHeader)
    C.last((*C.char)(unsafe.Pointer(&buf[0])), len)
    _, err := aiRecObj.conn.Write(buf)

    fmt.Println("send data_start")
    if err != nil {
        log.Fatal(err)
        fmt.Println("err not null")

    }

    

}

func method_entry() {
    var buf = make([]byte, 1024)
    len := Header(buf)
    fmt.Println("len in Wraph", len)
    var node MethodEntry_t
    var msgHdr msgHdr_t

    methodName := "custom_method_name"
    query_string := "select * from countries"
    urlParameter := ""

    node.MethodEntryVar.mid = 23
    node.MethodEntryVar.flowpathinstance = 0
    node.MethodEntryVar.threadId = 0
    node.MethodEntryVar.startTime = 0
    node.MethodEntryVar.methodName = (condition(methodName))
    node.MethodEntryVar.query_string = (condition(query_string))
    node.MethodEntryVar.urlParameter = (condition(urlParameter))
    fmt.Println(node.MethodEntryVar.methodName, "hee")
    msgHdr.header_len = 12
    msgHdr.total_len = 12 + 40 + node.MethodEntryVar.methodName + node.MethodEntryVar.query_string + node.MethodEntryVar.urlParameter + 3
    msgHdr.msg_type = 0

    len = C.main4((*C.char)(unsafe.Pointer(&buf[0])), (*C.msgHdr_t)(unsafe.Pointer(&msgHdr)), (*C.MethodEntry_t)(unsafe.Pointer(&node)), len)
   
    
    a := C.CString(methodName)
    b := C.CString(query_string)
    c := C.CString(urlParameter)
    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))

    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), a, len, node.MethodEntryVar.methodName)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), b, len, node.MethodEntryVar.query_string)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), c, len, node.MethodEntryVar.urlParameter)

    C.last((*C.char)(unsafe.Pointer(&buf[0])), len)

    _, err := aiRecObj.conn.Write(buf)
    fmt.Println("send data_MEntry")
    if err != nil {
        log.Fatal(err)

    }
}

func method_exit() {

    var buf = make([]byte, 1024)
    len := Header(buf)
    var node MethodExit_t
    var msgHdr msgHdr_t
    methodName := "custom_method_name"
    backend_header := "NA|10.20.0.85|NA|NA|mydb|mysql|NA|NA|NA|root"
    requestNotificationPhase := ""

    node.MethodExitVar.statusCode = 200
    node.MethodExitVar.mid = 12
    node.MethodExitVar.eventType = 1
    node.MethodExitVar.isCallout = 1
    node.MethodExitVar.duration = 363
    node.MethodExitVar.threadId = 0
    node.MethodExitVar.cpuTime = 0
    node.MethodExitVar.flowpathinstance = 0
    node.MethodExitVar.tierCallOutSeqNum = 45
    node.MethodExitVar.endTime = 0
    node.MethodExitVar.methodName = condition(methodName)
    node.MethodExitVar.backend_header = condition(backend_header)
    node.MethodExitVar.requestNotificationPhase = condition(requestNotificationPhase)

    msgHdr.header_len = 12
    msgHdr.total_len = 12 + 76 + node.MethodExitVar.methodName + node.MethodExitVar.backend_header + node.MethodExitVar.requestNotificationPhase + 3
    msgHdr.msg_type = 1

    len = C.main5((*C.char)(unsafe.Pointer(&buf[0])), (*C.msgHdr_t)(unsafe.Pointer(&msgHdr)), (*C.MethodExit_t)(unsafe.Pointer(&node)), len)
    a := C.CString(methodName)
    b := C.CString(backend_header)
    c := C.CString(requestNotificationPhase)

    defer C.free(unsafe.Pointer(a))
    defer C.free(unsafe.Pointer(b))
    defer C.free(unsafe.Pointer(c))

    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), a, len, node.MethodExitVar.methodName)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), b, len, node.MethodExitVar.backend_header)
    len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), c, len, node.MethodExitVar.requestNotificationPhase)

    C.last((*C.char)(unsafe.Pointer(&buf[0])), len)
    _, err := aiRecObj.conn.Write(buf)
    fmt.Println("send data_MExit")
    if err != nil {
        log.Fatal(err)

    }

}

func end_business_transaction() {
    UDPConnection()
    var buf = make([]byte, 1024)
    len := Header(buf)

    var transactionEnd transactionEnd_t
    var msgHdr msgHdr_t
    msgHdr.header_len = 12
    msgHdr.total_len = 12 + 28 + 3
    msgHdr.msg_type = 3

    transactionEnd.statuscode = 200
    transactionEnd.endTime = 0
    transactionEnd.flowpathinstance = 0
    transactionEnd.cpuTime = 0

    len = C.main6((*C.char)(unsafe.Pointer(&buf[0])), (*C.msgHdr_t)(unsafe.Pointer(&msgHdr)), (*C.transactionEnd_t)(unsafe.Pointer(&transactionEnd)), len)
    C.last((*C.char)(unsafe.Pointer(&buf[0])), len)
    _, err := aiRecObj.conn.Write(buf)
    fmt.Println("send data_end")
    if err != nil {
        log.Fatal(err)

    }
    closeUDP()
}



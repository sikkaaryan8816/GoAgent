package index

//package main

/*
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "ndlb_encode.h"

void check(){
	
    	msgHdr_t msgHdr;
    	wrapheader_t wrapHeader;
    	int len = 0;
    	//int bytesSent = 0;
    	char *apiReqId="akkdnfjnflffk";
    	char *awsReqId="lksjskwjdkdldl";
    	char *funcName="lamdafunion_test1";
    	int agentType=0;
    	//int pid=getpid();;
    char *tags ="tierName=tier_test1;appServerName=server_test1;appName=lamdafunion_test1";
    wrapHeader.wrapheadervar.apiReqLen = apiReqId ? strlen(apiReqId) : 0;
    wrapHeader.wrapheadervar.awsReqLen = awsReqId ? strlen(awsReqId) : 0;
    wrapHeader.wrapheadervar.funcNameLen = funcName ? strlen(funcName) : 0;
    wrapHeader.wrapheadervar.tagslength = tags ? strlen(tags) : 0;
    wrapHeader.wrapheadervar.agentType = agentType > 0 ? agentType : 0;
//    wrapHeader.wrapheadervar.pid = pid > 0 ? pid : 0;
    wrapHeader.wrapheadervar.whLen = sizeof(wrapheadervar_t)+wrapHeader.wrapheadervar.awsReqLen+wrapHeader.wrapheadervar.apiReqLen+wrapHeader.wrapheadervar.funcNameLen+wrapHeader.wrapheadervar.tagslength+1;
    fprintf(stderr, "size of struct =%ld\n",sizeof(wrapHeader.wrapheadervar) );
    memcpy(sBuff, "^", 1);
    len += 1;
    memcpy(sBuff + len , (char *)&(wrapHeader.wrapheadervar), sizeof(wrapHeader.wrapheadervar));
    fprintf(stderr, "asdsfd = %p\n",(char *)&(wrapHeader.wrapheadervar));
    len += sizeof(wrapheadervar_t);
    if(wrapHeader.wrapheadervar.apiReqLen>0)
    {
      memcpy(sBuff + len, apiReqId, wrapHeader.wrapheadervar.apiReqLen);
      len += wrapHeader.wrapheadervar.apiReqLen;
    }

    if(wrapHeader.wrapheadervar.awsReqLen>0)
    {
      memcpy(sBuff + len, awsReqId, wrapHeader.wrapheadervar.awsReqLen);
      len += wrapHeader.wrapheadervar.awsReqLen;
    }

    if(wrapHeader.wrapheadervar.funcNameLen>0)
    {
      memcpy(sBuff + len, funcName, wrapHeader.wrapheadervar.funcNameLen);
      len += wrapHeader.wrapheadervar.funcNameLen;
    }
    if(wrapHeader.wrapheadervar.tagslength>0)
    {
      memcpy(sBuff + len, tags, wrapHeader.wrapheadervar.tagslength);
      len += wrapHeader.wrapheadervar.tagslength;
    }
      transactionStart_t node;
      long qTimeMS =0;
      long long startTimeFP=0;
      char *fp_header="header";
      char *url ="http://10.20.0.85:81/PDO/pdo_test1.php";
      char *btHeaderValue=NULL;
      char *ndCookieSet=NULL;
      char *nvCookieSet=NULL;
      char *correlationHeader=""; 
      node.transactionStartVar.qTimeMS = qTimeMS > 0 ? qTimeMS : 0;
      node.transactionStartVar.startTimeFP = startTimeFP > 0 ? startTimeFP : 0;
      node.transactionStartVar.fp_header = fp_header ? strlen(fp_header) : 0;
      node.transactionStartVar.url = url ? strlen(url) : 0;
      node.transactionStartVar.btHeaderValue = btHeaderValue ? strlen(btHeaderValue) : 0;
      node.transactionStartVar.ndCookieSet = ndCookieSet ? strlen(ndCookieSet) : 0;
      node.transactionStartVar.nvCookieSet = nvCookieSet ? strlen(nvCookieSet) : 0;
      node.transactionStartVar.correlationHeader = correlationHeader  ? strlen(correlationHeader) : 0;
      msgHdr.header_len = sizeof(msgHdr_t);
      msgHdr.total_len = sizeof (transactionStartVar_t) + msgHdr.header_len + node.transactionStartVar.fp_header +
	node.transactionStartVar.ndCookieSet + node.transactionStartVar.nvCookieSet +
	node.transactionStartVar.correlationHeader +
	node.transactionStartVar.btHeaderValue +
	node.transactionStartVar.url + 3;
      msgHdr.msg_type = 2;
      memcpy(sBuff+len, "^", 1);
      len += 1;
      memcpy(sBuff + len , (char *)&(msgHdr), sizeof(msgHdr_t));
      len += sizeof(msgHdr_t);
      memcpy(sBuff + len , "|",1);
      len += 1;
      memcpy(sBuff + len, (char *)&(node.transactionStartVar), sizeof(node.transactionStartVar));

      len += sizeof(node.transactionStartVar) ;
      if(node.transactionStartVar.fp_header)
      {
	memcpy(sBuff + len, fp_header, node.transactionStartVar.fp_header);
	len += node.transactionStartVar.fp_header;
      }

      if(node.transactionStartVar.url)
      {
	memcpy(sBuff + len, url, node.transactionStartVar.url);
	len += node.transactionStartVar.url;
      }

      if(node.transactionStartVar.btHeaderValue)
      {
	memcpy(sBuff + len, btHeaderValue, node.transactionStartVar.btHeaderValue);
	len += node.transactionStartVar.btHeaderValue;
      }

      if(node.transactionStartVar.ndCookieSet)
      {
	memcpy(sBuff + len, ndCookieSet, node.transactionStartVar.ndCookieSet);
	len += node.transactionStartVar.ndCookieSet;
      }

      if(node.transactionStartVar.nvCookieSet)
      {
	memcpy(sBuff + len, nvCookieSet, node.transactionStartVar.nvCookieSet);
	len += node.transactionStartVar.nvCookieSet;
      }

      if(node.transactionStartVar.correlationHeader)
      {
	memcpy(sBuff + len, correlationHeader, node.transactionStartVar.correlationHeader);
	len += node.transactionStartVar.correlationHeader;
      }

  }

int main1(char *s, wrapheader_t *in){

    int len = 0;
    wrapheader_t wrapHeader;
    memcpy(s, "^",1);
    len += 1;
    memcpy(s + len , (char *)&(in->wrapheadervar), sizeof(wrapHeader.wrapheadervar));
    len += sizeof(wrapHeader.wrapheadervar);
    fprintf(stderr,"len WH=%d\n",len);
    fprintf(stderr,"len WH=%ld\n",sizeof(wrapHeader.wrapheadervar));

  return len;
}
int main2(char *s,char *value,int len,int number)
{
      memcpy(s + len, value ,number);
      len += number;
    return len;
}
int main3(char *s,msgHdr_t *in,transactionStart_t *in1,int len)
{
    transactionStart_t node;
    msgHdr_t msgHdr;
    memcpy(s+len, "^",1);
    len += 1;
    memcpy(s + len , (char *)in, sizeof(msgHdr_t) );
    fprintf(stderr,"in->header_len=%d\n",in->header_len);
    fprintf(stderr,"in->total_len=%d\n",in->total_len);
    fprintf(stderr,"in->msg_type=%d\n",in->msg_type);
    len += sizeof(msgHdr_t);
    fprintf(stderr,"len MH=%d\n",len);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(in1->transactionStartVar), sizeof(node.transactionStartVar));
    len += sizeof(node.transactionStartVar);
    fprintf(stderr,"len ST=%d\n",len);
    fprintf(stderr,"len MH=%ld\n",sizeof(node.transactionStartVar));
    memcpy(s + len, "\n", 1);
    len += 1;
    return len;


}
int main4(char *s,msgHdr_t *in, MethodEntry_t *in1,int len)
{
    msgHdr_t msgHdr;
    MethodEntry_t node;
    memcpy(s+len, "^",1);
    len += 1;
    memcpy(s + len , (char *)in, sizeof(msgHdr_t) );
    fprintf(stderr,"in->header_len=%d\n",in->header_len);
    fprintf(stderr,"in->total_len=%d\n",in->total_len);
    fprintf(stderr,"in->msg_type=%d\n",in->msg_type);
    len += sizeof(msgHdr_t);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(in1->MethodEntryVar), sizeof(node.MethodEntryVar));
    len += sizeof(node.MethodEntryVar);
    fprintf(stderr,"len MEY=%d\n",len);
    fprintf(stderr,"len MH=%ld\n",sizeof(node.MethodEntryVar));
    memcpy(s + len, "\n", 1);
    len += 1;
    return len;
}
int main5(char *s,msgHdr_t *in, MethodExit_t *in1,int len)
{
    msgHdr_t msgHdr;
    MethodExit_t node;
    memcpy(s+len, "^",1);
    len += 1;
    memcpy(s + len , (char *)in, sizeof(msgHdr_t) );
    fprintf(stderr,"in->header_len=%d\n",in->header_len);
    fprintf(stderr,"in->total_len=%d\n",in->total_len);
    fprintf(stderr,"in->msg_type=%d\n",in->msg_type);
    len += sizeof(msgHdr);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(in1->MethodExitVar), sizeof(node.MethodExitVar));
    len += sizeof(node.MethodExitVar);
    fprintf(stderr,"len MET=%d\n",len);
    fprintf(stderr,"len MH=%ld\n",sizeof(node.MethodExitVar));
    memcpy(s + len, "\n", 1);
    len += 1;
    return len;

}
int main6(char *s,msgHdr_t *in, transactionEnd_t *in1,int len)
{
    msgHdr_t msgHdr;
    transactionEnd_t node;
    memcpy(s+len, "^",1);
    len += 1;
    memcpy(s + len , (char *)in, sizeof(msgHdr_t) );
    fprintf(stderr,"in->header_len=%d\n",in->header_len);
    fprintf(stderr,"in->total_len=%d\n",in->total_len);
    fprintf(stderr,"in->msg_type=%d\n",in->msg_type);
    len += sizeof(msgHdr_t);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)in1, sizeof(node));
    len += sizeof(node);
    fprintf(stderr,"len ET=%d\n",len);
    fprintf(stderr,"len MH=%ld\n",sizeof(node));
    memcpy(s + len, "\n", 1);
    len += 1;
}
*/
import "C"

import (
	"fmt"
	"log"
	"net"
	"time"
	"unsafe"
	//"sync"
	//"bufio"
	//"github.com/google/uuid"
	//"math/big"
	//"strings"
)

func closeUDP() {
	aiRecObj.conn.Close()
	fmt.Println("close")
}

func Header(buf []byte) C.int {

	var apiReqId = "akkdnfjnflffk"
	var awsReqId = "lksjskwjdkdldl"
	var funcName = "lamdafunion_test1"
	//var agentType= 0
	//var messageType = 0
	var tags = "tierName=tier_test1;ndAppServerHost=server_test1;appName=Abhimanyulambda_test1"
	var wrapHeader wrapheader_t
	wrapHeader.wrapheadervar.apiReqLen = (condition(apiReqId))
	wrapHeader.wrapheadervar.awsReqLen = (condition(awsReqId))
	wrapHeader.wrapheadervar.funcNameLen = (condition(funcName))
	wrapHeader.wrapheadervar.tagslength = (condition(tags))
	wrapHeader.wrapheadervar.agentType = 0
	wrapHeader.wrapheadervar.messageType = 0
	wrapHeader.wrapheadervar.whLen = C.int(unsafe.Sizeof(wrapHeader.wrapheadervar)) + wrapHeader.wrapheadervar.apiReqLen + wrapHeader.wrapheadervar.awsReqLen + wrapHeader.wrapheadervar.funcNameLen + wrapHeader.wrapheadervar.tagslength + 1
	len := C.main1((*C.char)(unsafe.Pointer(&buf[0])), (*C.wrapheader_t)(unsafe.Pointer(&wrapHeader)))
	a := C.CString(apiReqId)
	b := C.CString(awsReqId)
	c := C.CString(funcName)
	d := C.CString(tags)
	defer C.free(unsafe.Pointer(a))
	defer C.free(unsafe.Pointer(b))
	defer C.free(unsafe.Pointer(c))
	defer C.free(unsafe.Pointer(d))

	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), a, len, wrapHeader.wrapheadervar.apiReqLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), b, len, wrapHeader.wrapheadervar.awsReqLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), c, len, wrapHeader.wrapheadervar.funcNameLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), d, len, wrapHeader.wrapheadervar.tagslength)

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
	fmt.Println("conn value",r.conn)
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
	//bt_name := "defult"
	UDPConnection()

	StartTransactionMessage(bt_name, "")
	//method_entry()
	//method_exit()
	//end_business_transaction()
	//fmt.Println(generate_bt())
	//bt := 1
	//return bt
	//closeUDP()
}

func method_entry() {

	var buf = make([]byte, 1024)
	len := Header(buf)
	var node MethodEntry_t
	methodName := "custom_method_name"

	query_string := ""
	urlParameter := ""
	// mid := 0
	// startTime := 0
	var msgHdr msgHdr_t

	node.MethodEntryVar.methodName = (condition(methodName))
	node.MethodEntryVar.threadId = 0
	node.MethodEntryVar.query_string = (condition(query_string))
	node.MethodEntryVar.urlParameter = (condition(urlParameter))
	node.MethodEntryVar.mid = 0
	node.MethodEntryVar.startTime = 0

	msgHdr.header_len = C.int(unsafe.Sizeof(msgHdr))
	msgHdr.total_len = C.int(unsafe.Sizeof(node.MethodEntryVar)) + msgHdr.header_len + node.MethodEntryVar.methodName + node.MethodEntryVar.query_string + node.MethodEntryVar.urlParameter + 3
	msgHdr.msg_type = 0
	len = C.main4((*C.char)(unsafe.Pointer(&buf[0])), (*C.msgHdr_t)(unsafe.Pointer(&msgHdr)), (*C.MethodEntry_t)(unsafe.Pointer(&node)), len)

	x := C.CString(methodName)
	y := C.CString(query_string)
	z := C.CString(urlParameter)

	defer C.free(unsafe.Pointer(x))
	defer C.free(unsafe.Pointer(y))
	defer C.free(unsafe.Pointer(z))

	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), x, len, node.MethodEntryVar.methodName)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), y, len, node.MethodEntryVar.query_string)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), z, len, node.MethodEntryVar.urlParameter)

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

	msgHdr.header_len = C.int(unsafe.Sizeof(msgHdr))
	msgHdr.total_len = C.int(unsafe.Sizeof(node.MethodExitVar)) + msgHdr.header_len + node.MethodExitVar.methodName + node.MethodExitVar.backend_header + node.MethodExitVar.requestNotificationPhase + 3
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
	fmt.Println("conn value2",aiRecObj.conn)
	_, err := aiRecObj.conn.Write(buf)
	fmt.Println("send data_MExit")
	if err != nil {
		fmt.Println("conn value3",aiRecObj.conn)
		fmt.Println("err = nil")
		log.Fatal(err)

	}

}

/*func generate_bt() {
	id := uuid.New().String()
	var i big.Int
	i.SetString(strings.Replace(id, "-", "", 4), 16)
	return i.String()
}*/

func StartTransactionMessage(bt_name string, correlationHeader string) {
	main1()
	/*var buf = make([]byte, 1024)
	var apiReqId = "akkdnfjnflffk"
	var awsReqId = "lksjskwjdkdldl"
	var funcName = "lamdafunion_test1"
	//var  agentType=0
	var tags = "tierName=tier_test1;ndAppServerHost=server_test1;appName=Abhimanyulambda_test1"
	var wrapHeader wrapheader_t
	wrapHeader.wrapheadervar.apiReqLen = (condition(apiReqId))
	wrapHeader.wrapheadervar.awsReqLen = (condition(awsReqId))
	wrapHeader.wrapheadervar.funcNameLen = (condition(funcName))
	wrapHeader.wrapheadervar.tagslength = (condition(tags))
	wrapHeader.wrapheadervar.whLen = C.int(unsafe.Sizeof(wrapHeader.wrapheadervar)) + wrapHeader.wrapheadervar.apiReqLen + wrapHeader.wrapheadervar.awsReqLen + wrapHeader.wrapheadervar.funcNameLen + wrapHeader.wrapheadervar.tagslength + 1

	len := C.main1((*C.char)(unsafe.Pointer(&buf[0])), (*C.wrapheader_t)(unsafe.Pointer(&wrapHeader)))
	u := C.CString(apiReqId)
	v := C.CString(awsReqId)
	x := C.CString(funcName)
	y := C.CString(tags)
	defer C.free(unsafe.Pointer(u))
	defer C.free(unsafe.Pointer(v))
	defer C.free(unsafe.Pointer(x))
	defer C.free(unsafe.Pointer(y))

	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), u, len, wrapHeader.wrapheadervar.apiReqLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), v, len, wrapHeader.wrapheadervar.awsReqLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), x, len, wrapHeader.wrapheadervar.funcNameLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])), y, len, wrapHeader.wrapheadervar.tagslength)
	
	
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
	fmt.Println("size of transaction.transactionStartVar-",unsafe.Sizeof(transaction.transactionStartVar))
	msgHdr.header_len = C.int(unsafe.Sizeof(msgHdr))
	msgHdr.total_len = C.int(unsafe.Sizeof(transaction.transactionStartVar)) + msgHdr.header_len + transaction.transactionStartVar.fp_header + transaction.transactionStartVar.url + transaction.transactionStartVar.btHeaderValue + transaction.transactionStartVar.ndCookieSet + transaction.transactionStartVar.nvCookieSet + transaction.transactionStartVar.correlationHeader + 3
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

	_, err := aiRecObj.conn.Write(buf)
	
	fmt.Println("send data_start")
	if err != nil {
		log.Fatal(err)
		fmt.Println("err not null")

	}

	//closeUDP()*/

}

func end_business_transaction() {
	UDPConnection()
	var buf = make([]byte, 1024)
	len := Header(buf)

	var transactionEnd transactionEnd_t
	var msgHdr msgHdr_t
	msgHdr.header_len = C.int(unsafe.Sizeof(msgHdr))
	msgHdr.total_len = C.int(unsafe.Sizeof(transactionEnd)) + msgHdr.header_len + 3
	msgHdr.msg_type = 3

	transactionEnd.statuscode = 200
	transactionEnd.endTime = 0
	transactionEnd.flowpathinstance = 0
	transactionEnd.cpuTime = 0

	C.main6((*C.char)(unsafe.Pointer(&buf[0])), (*C.msgHdr_t)(unsafe.Pointer(&msgHdr)), (*C.transactionEnd_t)(unsafe.Pointer(&transactionEnd)), len)
	_, err := aiRecObj.conn.Write(buf)
	fmt.Println("send data_end")
	if err != nil {
		log.Fatal(err)

	}
	closeUDP()
}

type transactionStartVar_t struct {
	fp_header         C.int
	url               C.int
	btHeaderValue     C.int
	ndCookieSet       C.int
	nvCookieSet       C.int
	correlationHeader C.int
	flowpathinstance  C.longlong
	qTimeMS           C.long
	startTimeFP       C.longlong
}

type transactionStart_t struct {
	transactionStartVar transactionStartVar_t

	fp_header         *C.char
	url               *C.char
	btHeaderValue     *C.char
	ndCookieSet       *C.char
	nvCookieSet       *C.char
	correlationHeader *C.char
}

type msgHdr_t struct {
	header_len C.int
	total_len  C.int
	msg_type   C.int
}

type MethodEntryVar_t struct {
	mid C.int

	flowpathinstance C.longlong
	threadId         C.long
	startTime        C.longlong

	methodName   C.int
	query_string C.int
	urlParameter C.int
}

type MethodEntry_t struct {
	MethodEntryVar MethodEntryVar_t

	methodName   *C.char
	query_string *C.char
	urlParameter *C.char
}
type MethodExitVar_t struct {
	statusCode       C.int
	mid              C.int
	eventType        C.int
	isCallout        C.int
	threadId         C.long
	duration         C.long
	flowpathinstance C.longlong
	cpuTime          C.longlong

	methodName               C.int
	backend_header           C.int
	requestNotificationPhase C.int
	tierCallOutSeqNum        C.longlong
	endTime                  C.longlong
}

type MethodExit_t struct {
	MethodExitVar MethodExitVar_t

	methodName               *C.char
	backend_header           *C.char
	requestNotificationPhase *C.char
}
type transactionEnd_t struct {
	statuscode       C.int
	flowpathinstance C.longlong
	endTime          C.longlong
	cpuTime          C.longlong
}
type wrapheadervar_t struct {
	whLen       C.int
	apiReqLen   C.int
	awsReqLen   C.int
	funcNameLen C.int
	tagslength  C.int
	agentType   C.short
	messageType C.short
}

type wrapheader_t struct {
	wrapheadervar wrapheadervar_t
	apiReqId      *C.char
	awsReqId      *C.char
	funcName      *C.char
	tags          *C.char //; separated key=value pairs
}

package index

//package main

/*
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "ndlb_encode.h"


int main1(char *s, wrapheader_t *in){

    int len = 0;

    memcpy(s, "^",1);
    len += 1;
    memcpy(s + len , (char *)&(in->wrapheadervar), sizeof(in->wrapheadervar));
    fprintf(stderr, "asdsfd = %p\n",(char *)&(in->wrapheadervar));
    len += sizeof(in->wrapheadervar);

  return len;
}
int main2(char *s,char *value,int len,int number)
{
      memcpy(s + len, value ,number);
      len += number;
    return len;
}
int main3(char *s,msgHdr_t *in,transactionStart_t *in1,int len,int header_len)
{

    memcpy(s+len, "^",1);
    len += 1;
    memcpy(s + len , (char *)&(in), sizeof(in) );

    len += sizeof(header_len);
    memcpy(s + len , "|",1);
    len += 1;
    memcpy(s + len, (char *)&(in1), sizeof(in1));
    len += sizeof(in1);
    return len;


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
	//var  agentType=0
	var tags = "tierName=tier_test1;appServerName=server_test1;appName=lamdafunion_test1"
	var wrapHeader wrapheader_t
	wrapHeader.wrapheadervar.apiReqLen = (condition(apiReqId))
	wrapHeader.wrapheadervar.awsReqLen = (condition(awsReqId))
	wrapHeader.wrapheadervar.funcNameLen = (condition(funcName))
	wrapHeader.wrapheadervar.tagslength = (condition(tags))
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
	r.conn, err = net.Dial("udp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("Some error %v", err)
	}
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
var buf = make([]byte, 1024)

//func main() {
func create_start_transaction_message(bt_name, nil string) {
	bt_name := "defult"
	UDPConnection()

	StartTransactionMessage(bt_name, "")

	//fmt.Println(generate_bt())
	//return bt
}

/*func generate_bt() {
	id := uuid.New().String()
	var i big.Int
	i.SetString(strings.Replace(id, "-", "", 4), 16)
	return i.String()
}*/

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

	msgHdr.header_len = C.int(unsafe.Sizeof(msgHdr))
	msgHdr.total_len = C.int(unsafe.Sizeof(transaction.transactionStartVar)) + msgHdr.header_len + transaction.transactionStartVar.fp_header + transaction.transactionStartVar.url + transaction.transactionStartVar.btHeaderValue + transaction.transactionStartVar.ndCookieSet + transaction.transactionStartVar.nvCookieSet + transaction.transactionStartVar.correlationHeader + 3
	msgHdr.msg_type = 2

	len = C.main3((*C.char)(unsafe.Pointer(&buf[0])), (*C.msgHdr_t)(unsafe.Pointer(&msgHdr)), (*C.transactionStart_t)(unsafe.Pointer(&transaction)), len, msgHdr.header_len)
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
	fmt.Println("send data")
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

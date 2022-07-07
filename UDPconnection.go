package index
/*
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h> 
#include <string.h>
#include "ndlb_encode.h"


int main1(char *s, wrapheader_t *in)
{
  
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


*/
import "C"

import (
	"fmt"
	"net"
	"unsafe"
	"log"
	"time"
	//"sync"
	//"bufio"	
)





func closeUDP(){
	aiRecObj.conn.Close()
	fmt.Println("close")
}

func main1(msgType int, ii int) {
	
	 
   	buf := make([]byte, 1024)
  	var apiReqId ="akkdnfjnflffk";
    var awsReqId ="lksjskwjdkdldl";
    var funcName ="lamdafunion_test1";
    //var  agentType=0
    var tags ="tierName=tier_test1;appServerName=server_test1;appName=lamdafunion_test1";
    var wrapHeader wrapheader_t;
   	wrapHeader.wrapheadervar.apiReqLen = (condition(apiReqId))
    wrapHeader.wrapheadervar.awsReqLen = (condition(awsReqId))
    wrapHeader.wrapheadervar.funcNameLen = (condition(funcName))
    wrapHeader.wrapheadervar.tagslength = (condition(tags))
    wrapHeader.wrapheadervar.whLen = C.int(unsafe.Sizeof(wrapHeader.wrapheadervar))+wrapHeader.wrapheadervar.apiReqLen+wrapHeader.wrapheadervar.awsReqLen+wrapHeader.wrapheadervar.funcNameLen+wrapHeader.wrapheadervar.tagslength+1
   
 
	len := C.main1((*C.char)(unsafe.Pointer(&buf[0])),(*C.wrapheader_t)(unsafe.Pointer(&wrapHeader)))
	a :=C.CString(apiReqId)
	b :=C.CString(awsReqId)
	c :=C.CString(funcName)
	d :=C.CString(tags)
	defer C.free(unsafe.Pointer(a))
	defer C.free(unsafe.Pointer(b))
	defer C.free(unsafe.Pointer(c))
	defer C.free(unsafe.Pointer(d))
	
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])),a,len,wrapHeader.wrapheadervar.apiReqLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])),b,len,wrapHeader.wrapheadervar.awsReqLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])),c,len,wrapHeader.wrapheadervar.funcNameLen)
	len = C.main2((*C.char)(unsafe.Pointer(&buf[0])),d,len,wrapHeader.wrapheadervar.tagslength)

	
	 _, err := aiRecObj.conn.Write(buf)
	fmt.Println("udp_data_send")
        if err != nil {
            log.Fatal(err)
	    fmt.Println("udp_data_not_send")
           
        }
      
}
type aiRecord struct {
	
	conn net.Conn
}

var aiRecObj *aiRecord = nil
var err error

func NewAIRecord() *aiRecord {
	r := aiRecord{}
	r.conn, err = net.Dial("udp", "66.220.31.147:12345")
	if err != nil {
        fmt.Printf("Some error %v", err)
     }
	return &r
}

func UDPConnection(){
	
	fmt.Println("udp_call")	
	
    aiRecObj = NewAIRecord()
    fmt.Println(aiRecObj)
    
	var ii = 0
		
			main1(2,ii)
			//time.Sleep(1*time.Second)
			//main1(0,ii)
			//main1(1,ii)
			//main1(3,ii)
		
//	ReceiveMessageFromServer()
	time.Sleep(1*time.Second)

	go closeUDP()

}


func condition(a string) C.int{
  
  if a == ""{
    return 0
  }
  return C.int(len(a))
}

/*func  ReceiveMessageFromServer(){
	request := make([]byte, 1024) 
	a,err := aiRecObj.conn.Read(request)
	//a :=len(request)
	fmt.Println("request=",string(request),a)
	
}*/

type wrapheadervar_t struct {
 whLen C.int
 apiReqLen C.int
 awsReqLen C.int
 funcNameLen C.int
 tagslength C.int
 agentType C.short
 messageType C.short

}

type wrapheader_t struct  {
wrapheadervar wrapheadervar_t
apiReqId *C.char
awsReqId *C.char
funcName *C.char
tags *C.char//; separated key=value pairs
}

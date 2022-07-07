#pragma pack(1)

enum udp_message_type {
  UDP_MESSAGE_DEFAULT = 0, //startfp,endfp,methodentry etc
  UDP_MESSAGE_HAVOC,
  UDP_CLIENT_INIT //first time a client connects to proxy
};
typedef struct wrapheadervar_t {
int whLen;
int apiReqLen;
int awsReqLen;
int funcNameLen;
int tagslength;
short agentType;
short messageType;
//int threadId;
//int pid;
}wrapheadervar_t;

typedef struct wrapheader_t {
wrapheadervar_t wrapheadervar;
char *apiReqId;
char *awsReqId;
char *funcName;
char *tags;//; separated key=value pairs
}wrapheader_t;

typedef struct msgHdr_t
{
  int header_len;
  int total_len;
  int msg_type;
}msgHdr_t;

extern msgHdr_t msgHdr;

typedef struct transactionStartVar_t
{ 
  int fp_header;
  int url;
  int btHeaderValue;
  int ndCookieSet;
  int nvCookieSet;
  int correlationHeader;
  long long flowpathinstance;
  long qTimeMS;
  long long startTimeFP;
}transactionStartVar_t;

extern transactionStartVar_t transactionStartVar;

typedef struct transactionStart_t
{ 
  transactionStartVar_t transactionStartVar;
  
  char *fp_header;
  char *url;
  char *btHeaderValue;
  char *ndCookieSet;
  char *nvCookieSet;
  char *correlationHeader;
}transactionStart_t;

extern transactionStart_t transactionStart;

typedef struct transactionReset_t
{
  long long flowpathinstance;
}transactionReset_t;

extern transactionReset_t transactionReset;

typedef struct transactionEnd_t
{
  int statuscode;
  long long flowpathinstance;
  long long endTime;
  long long cpuTime;
}transactionEnd_t;

extern transactionEnd_t transactionEnd;

typedef struct MethodEntryVar_t
{
  int mid;  

  long long flowpathinstance;
  long threadId;
  long long startTime;

  int methodName;
  int query_string;
  int urlParameter;
}MethodEntryVar_t;

extern MethodEntryVar_t MethodEntryVar;

typedef struct MethodEntry_t
{
  struct MethodEntryVar_t MethodEntryVar;

  char *methodName;
  char *query_string;
  char *urlParameter;
}MethodEntry_t;

extern MethodEntry_t MethodEntry;


typedef struct MethodExitVar_t
{
  int statusCode;
  int mid;
  int eventType;
  int isCallout;
  long threadId;
  long duration;
  long long flowpathinstance;
  long long cpuTime;

  int methodName;
  int backend_header;
  int requestNotificationPhase;
  long long tierCallOutSeqNum;
  long long endTime;
}MethodExitVar_t;

extern MethodExitVar_t MethodExitVar;

typedef struct MethodExit_t
{
  MethodExitVar_t MethodExitVar;

  char *methodName;
  char *backend_header;
  char *requestNotificationPhase;

}MethodExit_t;

extern MethodExit_t MethodExit;

typedef struct transactionEncodeVarHttp_t
{
    int  statuscode;
    int buffer_len;
    int type_len;
    long long flowpathinstance;

}transactionEncodeVarHttp_t;
extern transactionEncodeVarHttp_t transactionEncodeVarHttp;

typedef struct transactionEncodeHttp_t
{
    struct transactionEncodeVarHttp_t transactionEncodeVarHttp;
    char* buffer;
    char* type;
} transactionEncodeHttp_t;
extern transactionEncodeHttp_t transactionEncodeHttp;


typedef struct ExceptionEncodeVar_t
{
	int lineNumber;
	int lenofExceptionClassName;
	int lenofExceptionThrowingClassName;
	int lenofExceptionThrowingMethodName;
	int lenofExceptionMessage;
	int lenofExceptionCause;
	int lenofExceptionStackTrace;
	long long flowpathInstance;
	long long startTime;

}ExceptionEncodeVar_t;
extern ExceptionEncodeVar_t ExceptionEncodeVar;

typedef struct ExceptionEncodestruct_t
{
	struct ExceptionEncodeVar_t ExceptionEncodeVar;
	char* ExceptionClassName;
	char* ExceptionThrowingClassName;
	char* ExceptionThrowingMethodName;
	char* ExceptionMessage;
	char* ExceptionCause;
	char* ExceptionStackTrace;
	

}ExceptionEncodestruct_t;
extern ExceptionEncodestruct_t ExceptionEncodestruct;

typedef struct ExceptionMethodExitEncodeVar_t
{
  int ExceptionExitType;
  int lenofExceptionMethodName;
  int lenofExceptionBackendName;
  long long flowpathInstance;
  long long threadId;

}ExceptionMethodExitEncodeVar_t;
extern ExceptionMethodExitEncodeVar_t ExceptionMethodExitEncodeVar;

typedef struct ExceptionMethodExitEncodestruct_t
{
  struct ExceptionMethodExitEncodeVar_t ExceptionMethodExitEncodeVar;
  char* ExceptionMethodName;
  char* ExceptionBackendName;
}ExceptionMethodExitEncodestruct_t;
extern ExceptionMethodExitEncodestruct_t ExceptionMethodExitEncodestruct;


msgHdr_t msgHdr;
transactionStartVar_t transactionStartVar;
transactionStart_t transactionStart;
transactionReset_t transactionReset;
transactionEnd_t transactionEnd;
MethodEntryVar_t MethodEntryVar;
MethodEntry_t MethodEntry;
MethodExitVar_t MethodExitVar;
MethodExit_t MethodExit;
transactionEncodeVarHttp_t transactionEncodeVarHttp;
transactionEncodeHttp_t transactionEncodeHttp;
ExceptionEncodeVar_t ExceptionEncodeVar;
ExceptionEncodestruct_t ExceptionEncodestruct;
ExceptionMethodExitEncodeVar_t ExceptionMethodExitEncodeVar;
ExceptionMethodExitEncodestruct_t ExceptionMethodExitEncodestruct;
extern void MethodEntryEncode(int mid, long long flowpathinstance, long int threadId, char* methodName, char* query_string, char* urlParameter, long long startTime);
extern void transactionEndEncode(int statuscode, long long flowpathinstance, long long endTime);
extern void transactionResetEncode(long long flowpathinstance);
extern void transactionStartEncode(long qTimeMS,long long flowpathinstance, char *fp_header, char *url, char *btHeaderValue, char *ndCookieSet, char *nvCookieSet, char *correlationHeader, long long startTimeFP);
extern void MethodExitEncode(int mid, int statusCode, int eventType, int isCallout, long threadId, long duration, long long flowpathinstance, long cpuTime, char *methodName, char *backend_header, long long tierCallOutSeqNum, long long endTime, char* requestNotificationPhase);
extern int decodeMessage(char *buff, int readBytes, int appIndex, int processIndex);
extern int decodeUDPMessage(char *buff, int readBytes, int appIndex, int processIndex);
extern void encodeHttpHeaderReq(long long flowpathinstance, char* buffer, char* type, int statusCode);
#ifdef __cplusplus
extern "C" void ExceptionDataEncode(long long flowpathInstance, long long startTime, char* ExceptionclassName, char* Exceptionmessage, char* throwingClassName, char* throwingmethodName, char* ExceptionCause, int lineNumber, char* exceptionStackTrace);
#else
extern void ExceptionDataEncode(long long flowpathInstance, long long startTime, char* ExceptionclassName, char* Exceptionmessage, char* throwingClassName, char* throwingmethodName, char* ExceptionCause, int lineNumber, char* exceptionStackTrace);
#endif


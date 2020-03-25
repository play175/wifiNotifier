// +build windows

package wifiNotifier

/*
#cgo LDFLAGS: -lwlanapi -lole32

#ifndef UNICODE
#define UNICODE
#endif

#include <windows.h>
#include <wlanapi.h>
#include <windot11.h>           // for DOT11_SSID struct
#include <objbase.h>
#include <wtypes.h>

//#include <wchar.h>
#include <stdio.h>
#include <stdlib.h>

static HANDLE hClient = NULL;

static inline void openHandle()
{
	if(hClient != NULL)return;

    DWORD dwMaxClient = 2;
    DWORD dwCurVersion = 0;
    DWORD dwResult = 0;

    dwResult = WlanOpenHandle(dwMaxClient, NULL, &dwCurVersion, &hClient);
    if (dwResult != ERROR_SUCCESS) {
        wprintf(L"WlanOpenHandle failed with error: %u\n", dwResult);
    }

}

static inline void wchar2char(const wchar_t* wchar , char * m_char)
{
    int len= WideCharToMultiByte( CP_ACP ,0,wchar ,wcslen( wchar ), NULL,0, NULL ,NULL );
    WideCharToMultiByte( CP_ACP ,0,wchar ,wcslen( wchar ),m_char,len, NULL ,NULL );
    m_char[len]= '\0';
}

static inline char * getCurrentSSID(void) {

	openHandle();

	char *ssid = malloc(256);
	memset(ssid,0,256);

	DWORD dwResult = 0;
    unsigned int i;

    PWLAN_INTERFACE_INFO_LIST pIfList = NULL;
    PWLAN_INTERFACE_INFO pIfInfo = NULL;

    PWLAN_CONNECTION_ATTRIBUTES pConnectInfo = NULL;
    DWORD connectInfoSize = sizeof(WLAN_CONNECTION_ATTRIBUTES);
	WLAN_OPCODE_VALUE_TYPE opCode = wlan_opcode_value_type_invalid;

    dwResult = WlanEnumInterfaces(hClient, NULL, &pIfList);
    if (dwResult != ERROR_SUCCESS) {
        wprintf(L"WlanEnumInterfaces failed with error: %u\n", dwResult);
    } else {
        for (i = 0; i < (int) pIfList->dwNumberOfItems; i++) {
			pIfInfo = (WLAN_INTERFACE_INFO *) & pIfList->InterfaceInfo[i];

            if (pIfInfo->isState == wlan_interface_state_connected) {
                dwResult = WlanQueryInterface(hClient,
                                              &pIfInfo->InterfaceGuid,
                                              wlan_intf_opcode_current_connection,
                                              NULL,
                                              &connectInfoSize,
                                              (PVOID *) &pConnectInfo,
                                              &opCode);

                if (dwResult != ERROR_SUCCESS) {
                    wprintf(L"WlanQueryInterface failed with error: %u\n", dwResult);
                } else {

					//wprintf(L"  Profile name used:\t %ws\n", pConnectInfo->strProfileName);
					wchar2char(pConnectInfo->strProfileName,ssid);
                }
            }
        }

	}

    if (pConnectInfo != NULL) {
        WlanFreeMemory(pConnectInfo);
        pConnectInfo = NULL;
    }

    if (pIfList != NULL) {
        WlanFreeMemory(pIfList);
        pIfList = NULL;
    }

	return ssid;
}

static inline void onWifiChanged(PWLAN_NOTIFICATION_DATA data,PVOID context)
{
	extern void __onWifiChanged(char *);
	__onWifiChanged(getCurrentSSID());
}

static inline void setWifiNotifier()
{
	openHandle();

	DWORD hResult = ERROR_SUCCESS;
	DWORD pdwPrevNotifSource = 0;
	hResult=WlanRegisterNotification(hClient,
									WLAN_NOTIFICATION_SOURCE_ACM,
									TRUE,
									(WLAN_NOTIFICATION_CALLBACK)onWifiChanged,
									NULL,
									NULL,
									&pdwPrevNotifSource);
	if(hResult!=ERROR_SUCCESS)
	{
		printf("failed WlanRegisterNotification=%d \n",hResult);
	}

	while(TRUE){
		Sleep(10);
	}

	WlanCloseHandle(hClient,NULL);
	printf("WlanCloseHandle success \n");

}

*/
import "C"
import "unsafe"

var internalOnWifiChangedCb func(string)
var internalOnGetSSIDCb func(string)

//export __onWifiChanged
func __onWifiChanged(ssid *C.char) {
	goSsid := C.GoString(ssid)
	C.free(unsafe.Pointer(ssid))

	if internalOnWifiChangedCb != nil {
		internalOnWifiChangedCb(goSsid)
	}
}

func GetCurrentSSID() string {
	ssid := C.getCurrentSSID()
	goSsid := C.GoString(ssid)
	C.free(unsafe.Pointer(ssid))
	return goSsid
}

func SetWifiNotifier(cb func(string)) {
	internalOnWifiChangedCb = cb
	go C.setWifiNotifier()
	// log.Println("setWifiNotifier complated")
}

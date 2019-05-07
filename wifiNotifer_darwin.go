// +build darwin

package wifiNotifier

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa -framework SystemConfiguration -framework CoreWLAN -framework Foundation

#import <Cocoa/Cocoa.h>
#import <AppKit/NSApplication.h>
#include <stdlib.h>
#import <Foundation/Foundation.h>
#import <CoreWLAN/CoreWLAN.h>
#import <SystemConfiguration/SystemConfiguration.h>

static inline  char* nsstring2cstring(NSString *s){
    if (s == NULL) { return NULL; }

    char *cstr = strdup([s UTF8String]);
    return cstr;
}

#define NOT_CONNECTED @""

#pragma clang diagnostic ignored "-Wdeprecated-declarations"
static inline void wifi_network_changed(SCDynamicStoreRef store, CFArrayRef changedKeys, void *ctx)
{
	//NSLog(@"wifi_network_changed");

	CWInterface *WiFiInterface = [CWInterface interface];
	NSString *currentSSID = [WiFiInterface ssid] ? [WiFiInterface ssid] : NOT_CONNECTED;
	extern void __onWifiChanged(char *);
	__onWifiChanged(nsstring2cstring(currentSSID));
}

static inline char * getCurrentSSID(void) {
    CWInterface *WiFiInterface = [CWInterface interface];
	NSString *ssid = [WiFiInterface ssid] ? [WiFiInterface ssid] : NOT_CONNECTED;
	return nsstring2cstring(ssid);
}

#pragma clang diagnostic ignored "-Wdeprecated-declarations"
static inline void setWifiNotifier(void) {
    CWInterface *WiFiInterface = [CWInterface interface];

    NSSet *wifiInterfaces = [CWInterface interfaceNames];
    NSMutableArray *scKeys = [[NSMutableArray alloc] init];
    [wifiInterfaces enumerateObjectsUsingBlock:^(NSString *ifName, BOOL *stop)
     {
         [scKeys addObject: [NSString stringWithFormat:@"State:/Network/Interface/%@/AirPort", ifName]];
     }];

    SCDynamicStoreContext ctx = { 0, NULL, NULL, NULL, NULL };
    SCDynamicStoreRef store = SCDynamicStoreCreate(kCFAllocatorDefault, CFSTR("myapp"), wifi_network_changed, &ctx);

    SCDynamicStoreSetNotificationKeys(store, (__bridge CFArrayRef)scKeys, NULL);

    CFRunLoopSourceRef src = SCDynamicStoreCreateRunLoopSource(kCFAllocatorDefault, store, 0);
	CFRunLoopAddSource([[NSRunLoop currentRunLoop] getCFRunLoop], src, kCFRunLoopCommonModes);

	CFRunLoopRun();

}

*/
import "C"
import (
	"unsafe"
)

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

wifiNotifier
===

Detect when wifi changed(connected or disconnected or ssid changed),current only support Mac OS

wifiNotifier 是一个可以监听Wi-Fi连接变化的库，当wifi连接或者断开连接或者ssid发生变化时都可以给出通知，当前只支持Mac OS

examples

```
package main

import (
	"log"

	"github.com/play175/wifiNotifier"
)

func main() {

	wifiNotifier.SetWifiNotifier(func(ssid string) {
		log.Println("onWifiChanged,current ssid:" + ssid)
	})

	log.Println("current ssid:" + wifiNotifier.GetCurrentSSID())

	for {

	}
}
```
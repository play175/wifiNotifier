wifiNotifier
===

Detect when wifi changed(connected or disconnected or ssid changed),current only support Windows & Mac OS

wifiNotifier 是一个可以监听 Wi-Fi 连接变化的库，当 wifi 连接或者断开连接或者 ssid 发生变化时都可以给出通知，当前支持 Windows 系统 和 Mac OS 系统

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
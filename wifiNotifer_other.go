// +build !darwin

package wifiNotifier

func GetCurrentSSID() string {
	return ""
}

func SetWifiNotifier(cb func(string)) {
}

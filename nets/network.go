package nets

type WifiAuthMode int

const (
	WIFI_AUTH_OPEN            WifiAuthMode = iota // authenticate mode : open
	WIFI_AUTH_WEP                                 // authenticate mode : WEP
	WIFI_AUTH_WPA_PSK                             // authenticate mode : WPA_PSK
	WIFI_AUTH_WPA2_PSK                            // authenticate mode : WPA2_PSK
	WIFI_AUTH_WPA_WPA2_PSK                        // authenticate mode : WPA_WPA2_PSK
	WIFI_AUTH_WPA2_ENTERPRISE                     // authenticate mode : WPA2_ENTERPRISE
	WIFI_AUTH_WPA3_PSK                            // authenticate mode : WPA3_PSK
	WIFI_AUTH_MAX                                 //
)

type Network struct {
	Ssid     string       `json:"ssid"`
	Password string       `json:"password,omitempty"`
	RSSI     int          `json:"rssi"`
	Secure   WifiAuthMode `json:"secure"`
	Stored   bool         `json:"stored"`
}

type NetworkConnectResponse struct {
	Ssid        string `json:"ssid"`
	RedirectURL string `json:"redirectUrl"`
}

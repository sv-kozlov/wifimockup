package nets

import (
	"math/rand"
)

var storedNames = []string{
	"linksys", "stream-guest", "kedoo", "pixel",
}
var names = []string{
	"Programatic", "Expression", "Savvy", "Illuminate", "Netry", "Spire", "Fission",
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func Generator() map[string]Network {
	list := make(map[string]Network)

	for _, name := range storedNames {
		list[name] = Network{
			Ssid:     name,
			Password: RandomString(6),
			RSSI:     30 + rand.Intn(70),
			Secure:   WifiAuthMode(rand.Intn(5)),
			Stored:   true,
		}
	}

	for _, name := range names {
		list[name] = Network{
			Ssid: name,
			// Password: RandomString(6),
			RSSI:   30 + rand.Intn(70),
			Secure: WifiAuthMode(rand.Intn(5)),
			Stored: false,
		}
	}

	return list
}

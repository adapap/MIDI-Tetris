package main

import (
	"time"
	keys "github.com/micmonay/keybd_event"
)

func main() {
	time.Sleep(10 * time.Second)
	kb, err := keys.NewKeyBonding()
	if err != nil {
		panic(err)
	}
	
	for i := 0; i < 10; i++ {
		kb.SetKeys(keys.VK_SPACE)
		time.Sleep(200 * time.Millisecond)
		err = kb.Launching()
		if err != nil {
			panic(err)
		}
	}
}
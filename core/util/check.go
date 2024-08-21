package util

func Check(e error) {
	if e != nil {
		Logger.Fatal("Opps!", "err", e.Error())
	}
}

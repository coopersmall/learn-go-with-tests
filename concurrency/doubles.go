package concurrency

import "time"

func MockWebsiteChecker(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func SlowStubWebsiteChecker(url string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

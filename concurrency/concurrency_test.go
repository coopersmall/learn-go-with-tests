package concurrency

import (
	"reflect"
	"testing"
)

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://www.google.com":      true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(MockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "hi"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(SlowStubWebsiteChecker, urls)
	}
}

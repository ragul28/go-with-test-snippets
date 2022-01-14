package racer

import (
	"testing"
	"time"
)

func TestRacerSelect(t *testing.T) {

	t.Run("return fasturl", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := RacerSelect(slowURL, fastURL)

		if err != nil {
			t.Fatalf("expected an error but didn't get one %v", err)
		}

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacerSelect(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})

}

package atdd_redis

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	t.Run("Given a Redis server", func(t *testing.T) {
		t.Run("When establishing a connection to it", func(t *testing.T) {
			t.Run("Then it returns 'HELLO' and closes the connection", func(t *testing.T) {
				conn, err := net.DialTimeout("tcp", "localhost:6379", 5*time.Second)
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				t.Cleanup(func() { _ = conn.Close() })

				rawBytes, err := io.ReadAll(conn)
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}

				result := string(rawBytes)
				if got, want := result, "HELLO"; got != want {
					t.Fatalf("Redis returned with %q but we wanted %q", got, want)
				}
			})
		})
	})
}

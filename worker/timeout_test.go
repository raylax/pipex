package worker

import (
	"log"
	"testing"
	"time"
)

func TestTimeout_ToDuration(t *testing.T) {
	t.Run("second", func(t *testing.T) {
		d := Timeout("3600s").ToDuration()
		if d != 1*time.Hour {
			log.Fatal(d)
		}
	})
	t.Run("minute", func(t *testing.T) {
		d := Timeout("60m").ToDuration()
		if d != 1*time.Hour {
			log.Fatal(d)
		}
	})
	t.Run("hour", func(t *testing.T) {
		d := Timeout("1h").ToDuration()
		if d != 1*time.Hour {
			log.Fatal(d)
		}
	})
	t.Run("second", func(t *testing.T) {
		d := Timeout("3600S").ToDuration()
		if d != 1*time.Hour {
			log.Fatal(d)
		}
	})
	t.Run("minute", func(t *testing.T) {
		d := Timeout("60M").ToDuration()
		if d != 1*time.Hour {
			log.Fatal(d)
		}
	})
	t.Run("hour", func(t *testing.T) {
		d := Timeout("1H").ToDuration()
		if d != 1*time.Hour {
			log.Fatal(d)
		}
	})
	t.Run("error default", func(t *testing.T) {
		d := Timeout("").ToDuration()
		if d != timeoutDefault.ToDuration() {
			log.Fatal(d)
		}
		d = Timeout("3600SS").ToDuration()
		if d != timeoutDefault.ToDuration() {
			log.Fatal(d)
		}
		d = Timeout("3600x").ToDuration()
		if d != timeoutDefault.ToDuration() {
			log.Fatal(d)
		}
	})
}

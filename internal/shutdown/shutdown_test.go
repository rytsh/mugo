package shutdown

import "testing"

func TestAddDeleteRun(t *testing.T) {
	shutdown := New()
	runs := 0

	shutdown.Add("kept", func() error {
		runs++
		return nil
	})
	deleted := shutdown.AddAnonymous(func() error {
		runs++
		return nil
	})
	shutdown.Delete(deleted)
	shutdown.Run()

	if runs != 1 {
		t.Fatalf("callbacks ran %d times, want 1", runs)
	}
}

func TestRunAllowsMutation(t *testing.T) {
	shutdown := New()
	shutdown.Add("mutating", func() error {
		shutdown.Delete("mutating")
		return nil
	})

	shutdown.Run()
}

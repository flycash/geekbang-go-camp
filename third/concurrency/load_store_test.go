package concurrency

import "testing"

func TestSafeMap_LoadOrStore(t *testing.T) {
	m := &SafeMap{
		m: map[string]interface{}{},
	}

	for i := 0; i < 10; i++ {
		go func() {
			con := &connection{}
			nc, loaded := m.LoadOrStore("hello", con)
			if loaded {
				_ = con.Close()
			}
			_ = nc.(*connection).send()
		}()
	}
}

func TestSafeMap_LoadOrStoreHeavy(t *testing.T) {
	m := &SafeMap{
		m: map[string]interface{}{},
	}

	for i := 0; i < 10; i++ {
		go func() {
			nc, _ := m.LoadOrStoreHeavy("hello", func() interface{} {
				return &connection{}
			})
			_ = nc.(*connection).send()
		}()
	}
}
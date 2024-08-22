// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncmap

import (
	"math/rand"
	"reflect"
	"runtime"
	"strconv"
	"sync"
	"testing/quick"

	"testing"
)

type mapOp string

const (
	opLoad             = mapOp("Load")
	opStore            = mapOp("Store")
	opLoadOrStore      = mapOp("LoadOrStore")
	opLoadAndDelete    = mapOp("LoadAndDelete")
	opDelete           = mapOp("Delete")
	opSwap             = mapOp("Swap")
	opCompareAndSwap   = mapOp("CompareAndSwap")
	opCompareAndDelete = mapOp("CompareAndDelete")
)

var mapOps = [...]mapOp{
	opLoad,
	opStore,
	opLoadOrStore,
	opLoadAndDelete,
	opDelete,
	opSwap,
	opCompareAndSwap,
	opCompareAndDelete,
}

type mapCall struct {
	op      mapOp
	k, v, n string
}
type mapInterfaceStrStr mapInterface[string, string]

func (c mapCall) apply(m mapInterfaceStrStr) (string, bool) {
	switch c.op {
	case opLoad:
		return m.Load(c.k)
	case opStore:
		m.Store(c.k, c.v)
		return "", false
	case opLoadOrStore:
		return m.LoadOrStore(c.k, c.v)
	case opLoadAndDelete:
		return m.LoadAndDelete(c.k)
	case opDelete:
		m.Delete(c.k)
		return "", false
	case opSwap:
		return m.Swap(c.k, c.v)
	case opCompareAndSwap:
		if m.CompareAndSwap(c.k, c.v, strconv.Itoa(rand.Int())) {
			m.Delete(c.k)
			return c.v, true
		}
		return "", false
	case opCompareAndDelete:
		if m.CompareAndDelete(c.k, c.v) {
			if _, ok := m.Load(c.k); !ok {
				return "", true
			}
		}
		return "", false
	default:
		panic("invalid mapOp")
	}
}

type mapResult struct {
	value any
	ok    bool
}

func (mapCall) Generate(r *rand.Rand, _ int) reflect.Value {
	c := mapCall{op: mapOps[rand.Intn(len(mapOps))], k: randValue(r)}
	switch c.op {
	//case opStore, opLoadOrStore, opSwap, opCompareAndSwap, opCompareAndDelete:
	case opStore, opLoadOrStore:
		c.v = randValue(r)
	}
	return reflect.ValueOf(c)
}

func applyCalls(m mapInterfaceStrStr, calls []mapCall) (results []mapResult, final map[any]any) {
	for _, c := range calls {
		v, ok := c.apply(m)
		results = append(results, mapResult{v, ok})
	}

	final = make(map[any]any)
	m.Range(func(k, v string) bool {
		final[k] = v
		return true
	})

	return results, final
}

func randValue(r *rand.Rand) string {
	b := make([]byte, r.Intn(4))
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}

func applyMap(calls []mapCall) ([]mapResult, map[any]any) {
	sm := &SyncMap[string, string]{}
	return applyCalls(sm, calls)
}

func applyRWMutexMap(calls []mapCall) ([]mapResult, map[any]any) {
	return applyCalls(new(RWMutexMap[string, string]), calls)
}

func applyMidMutexMap(calls []mapCall) ([]mapResult, map[any]any) {
	return applyCalls(new(SmartMutexMap[string, string]), calls)
}

func applyDeepCopyMap(calls []mapCall) ([]mapResult, map[any]any) {
	return applyCalls(new(DeepCopyMap[string, string]), calls)
}

func applyClassicSyncMap(calls []mapCall) ([]mapResult, map[any]any) {
	return applyCalls(new(SyncMapClassic[string, string]), calls)
}

func TestMapMatchesRWMutex(t *testing.T) {
	if err := quick.CheckEqual(applyMap, applyRWMutexMap, &quick.Config{MaxCount: 10000}); err != nil {
		t.Error(err)
	}
}

func TestMapMatchesMidMutex(t *testing.T) {
	if err := quick.CheckEqual(applyMidMutexMap, applyRWMutexMap, &quick.Config{MaxCount: 10000}); err != nil {
		t.Error(err)
	}
}

func TestMapMatchesDeepCopy(t *testing.T) {
	if err := quick.CheckEqual(applyMap, applyDeepCopyMap, &quick.Config{MaxCount: 10000}); err != nil {
		t.Error(err)
	}
}

func TestMapMatchesClassicSyncMap(t *testing.T) {
	if err := quick.CheckEqual(applyMap, applyClassicSyncMap, nil); err != nil {
		t.Error(err)
	}
}

func TestConcurrentRange(t *testing.T) {
	const mapSize = 1 << 10

	m := &SyncMap[int64, int64]{}
	for n := int64(1); n <= mapSize; n++ {
		m.Store(n, int64(n))
	}

	done := make(chan struct{})
	var wg sync.WaitGroup
	defer func() {
		close(done)
		wg.Wait()
	}()
	for g := int64(runtime.GOMAXPROCS(0)); g > 0; g-- {
		r := rand.New(rand.NewSource(g))
		wg.Add(1)
		go func(g int64) {
			defer wg.Done()
			for i := int64(0); ; i++ {
				select {
				case <-done:
					return
				default:
				}
				for n := int64(1); n < mapSize; n++ {
					if r.Int63n(mapSize) == 0 {
						m.Store(n, n*i*g)
					} else {
						m.Load(n)
					}
				}
			}
		}(g)
	}

	iters := 1 << 10
	if testing.Short() {
		iters = 16
	}
	for n := iters; n > 0; n-- {
		seen := make(map[int64]bool, mapSize)

		m.Range(func(k, v int64) bool {
			if v%k != 0 {
				t.Fatalf("while Storing multiples of %v, Range saw value %v", k, v)
			}
			if seen[k] {
				t.Fatalf("Range visited key %v twice", k)
			}
			seen[k] = true
			return true
		})

		if len(seen) != mapSize {
			t.Fatalf("Range visited %v elements of %v-element Map", len(seen), mapSize)
		}
	}
}

func TestMapRangeNestedCall(t *testing.T) { // Issue 46399
	var m = &SyncMap[int, string]{}
	for i, v := range [3]string{"hello", "world", "Go"} {
		m.Store(i, v)
	}
	m.Range(func(key int, value string) bool {
		m.Range(func(key int, value string) bool {
			// We should be able to load the key offered in the Range callback,
			// because there are no concurrent Delete involved in this tested map.
			if v, ok := m.Load(key); !ok || v != value {
				t.Fatalf("Nested Range loads unexpected value, got %+v want %+v", v, value)
			}

			// We didn't keep 42 and a value into the map before, if somehow we loaded
			// a value from such a key, meaning there must be an internal bug regarding
			// nested range in the Map.
			if _, loaded := m.LoadOrStore(42, "dummy"); loaded {
				t.Fatalf("Nested Range loads unexpected value, want store a new value")
			}

			// Try to Store then LoadAndDelete the corresponding value with the key
			// 42 to the Map. In this case, the key 42 and associated value should be
			// removed from the Map. Therefore any future range won't observe key 42
			// as we checked in above.
			val := "sync.Map"
			m.Store(42, val)
			if v, loaded := m.LoadAndDelete(42); !loaded || v != val {
				t.Fatalf("Nested Range loads unexpected value, got %v, want %v", v, val)
			}
			return true
		})

		// Remove key from Map on-the-fly.
		m.Delete(key)
		return true
	})

	// After a Range of Delete, all keys should be removed and any
	// further Range won't invoke the callback. Hence length remains 0.
	length := 0
	m.Range(func(key int, value string) bool {
		length++
		return true
	})

	if length != 0 {
		t.Fatalf("Unexpected sync.Map size, got %v want %v", length, 0)
	}
}

func TestMapRangeNoAllocations(t *testing.T) { // Issue 62404
	//testenv.SkipIfOptimizationOff(t)
	var m = &SyncMap[int, int]{}
	allocs := testing.AllocsPerRun(10, func() {
		m.Range(func(key, value int) bool {
			return true
		})
	})
	if allocs > 0 {
		t.Errorf("AllocsPerRun of m.Range = %v; want 0", allocs)
	}
}

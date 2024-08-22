// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syncmap

import (
	"maps"
	"sync"
	"sync/atomic"
)

// This file contains reference map implementations for unit-tests.

// mapInterfaceAny is the interface Map implements.
type mapInterfaceAny interface {
	Load(any) (any, bool)
	Store(key, value any)
	LoadOrStore(key, value any) (actual any, loaded bool)
	LoadAndDelete(key any) (value any, loaded bool)
	Delete(any)
	Swap(key, value any) (previous any, loaded bool)
	CompareAndSwap(key, old, new any) (swapped bool)
	CompareAndDelete(key, old any) (deleted bool)
	Range(func(key, value any) (shouldContinue bool))
}

// mapInterface is the generic interface Map implements.
type mapInterface[K comparable, V any] interface {
	Load(key K) (V, bool)
	Store(key K, value V)
	LoadOrStore(key K, value V) (actual V, loaded bool)
	LoadAndDelete(key K) (value V, loaded bool)
	Delete(K)
	Swap(key K, value V) (previous V, loaded bool)
	CompareAndSwap(key K, old, new V) (swapped bool)
	CompareAndDelete(key K, old V) (deleted bool)
	Range(func(key K, value V) (shouldContinue bool))
}

var (
	_ mapInterfaceAny              = &rwMutexMap{}
	_ mapInterfaceAny              = &smartMutexMap{}
	_ mapInterfaceAny              = &sync.Map{}
	_ mapInterfaceAny              = &deepCopyMap{}
	_ mapInterface[string, string] = &RWMutexMap[string, string]{}
	_ mapInterface[string, string] = &DeepCopyMap[string, string]{}
	_ mapInterface[string, string] = &SyncMapClassic[string, string]{}
	_ mapInterface[int, int]       = &RWMutexMap[int, int]{}
	_ mapInterface[int, int]       = &DeepCopyMap[int, int]{}
	_ mapInterface[int, int]       = &SyncMapClassic[int, int]{}
	_ mapInterface[int, int]       = &SmartMutexMap[int, int]{}
)

type (
	// SyncMapClassic is an generic implementation of mapInterface + mapInterfaceAny using a classic sync.Map
	SyncMapClassic[K comparable, V any] struct {
		base sync.Map
	}
)

func (m *SyncMapClassic[K, V]) Load(s K) (value V, ok bool) {
	vi, ok := m.base.Load(s)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *SyncMapClassic[K, V]) Store(key K, value V) {
	m.base.Store(key, value)
}

func (m *SyncMapClassic[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	vi, loaded := m.base.LoadOrStore(key, value)
	if vi != nil {
		actual = vi.(V)
	}
	return
}

func (m *SyncMapClassic[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	vi, loaded := m.base.LoadAndDelete(key)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *SyncMapClassic[K, V]) Delete(s K) {
	m.base.Delete(s)
}

func (m *SyncMapClassic[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	vi, loaded := m.base.Swap(key, value)
	if vi != nil {
		previous = vi.(V)
	}
	return
}

func (m *SyncMapClassic[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return m.base.CompareAndSwap(key, old, new)
}

func (m *SyncMapClassic[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.base.CompareAndDelete(key, old)
}

func (m *SyncMapClassic[K, V]) Range(f func(key K, value V) (shouldContinue bool)) {
	m.base.Range(func(k, v any) bool {
		key, _ := k.(K)
		value, _ := v.(V)
		return f(key, value)
	})
}

type (
	SmartRWLocker struct {
		sync.RWMutex
	}
	// smartMutexMap is an implementation of mapInterfaceAny using a sync.RWMutex.
	smartMutexMap struct {
		mu    SmartRWLocker
		dirty map[any]any
	}
	// SmartMutexMap is an implementation of mapInterface using a sync.RWMutex.
	SmartMutexMap[K comparable, V any] struct {
		base smartMutexMap
	}
)

func (l *SmartRWLocker) Proceed(checkF func() bool, changeF func()) {
	l.RLock()
	needChange := checkF()
	l.RUnlock()

	if needChange {
		l.Lock()
		if checkF() {
			changeF()
		}
		l.Unlock()
	}
}

func (m *SmartMutexMap[K, V]) Load(s K) (value V, ok bool) {
	vi, ok := m.base.Load(s)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *SmartMutexMap[K, V]) Store(key K, value V) {
	m.base.Store(key, value)
}

func (m *SmartMutexMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	vi, loaded := m.base.LoadOrStore(key, value)
	if vi != nil {
		actual = vi.(V)
	}
	return
}

func (m *SmartMutexMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	vi, loaded := m.base.LoadAndDelete(key)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *SmartMutexMap[K, V]) Delete(s K) {
	m.base.Delete(s)
}

func (m *SmartMutexMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	vi, loaded := m.base.Swap(key, value)
	if vi != nil {
		previous = vi.(V)
	}
	return
}

func (m *SmartMutexMap[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return m.base.CompareAndSwap(key, old, new)
}

func (m *SmartMutexMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.base.CompareAndDelete(key, old)
}

func (m *SmartMutexMap[K, V]) Range(f func(key K, value V) (shouldContinue bool)) {
	m.base.Range(func(k, v any) bool {
		key, _ := k.(K)
		value, _ := v.(V)
		return f(key, value)
	})
}

func (m *smartMutexMap) Load(key any) (value any, ok bool) {
	m.mu.RLock()
	value, ok = m.dirty[key]
	m.mu.RUnlock()
	return
}

func (m *smartMutexMap) Store(key, value any) {
	m.mu.Lock()
	if m.dirty == nil {
		m.dirty = make(map[any]any)
	}
	m.dirty[key] = value
	m.mu.Unlock()
}

func (m *smartMutexMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	m.mu.Proceed(func() bool {
		actual, loaded = m.dirty[key]
		return !loaded
	}, func() {
		actual = value
		if m.dirty == nil {
			m.dirty = make(map[any]any)
		}
		m.dirty[key] = value
	})
	return actual, loaded
}

func (m *smartMutexMap) Swap(key, value any) (previous any, loaded bool) {
	m.mu.Proceed(func() bool {
		previous, loaded = m.dirty[key]
		return true
	}, func() {
		if m.dirty == nil {
			m.dirty = make(map[any]any)
		}
		m.dirty[key] = value
	})
	return
}

func (m *smartMutexMap) LoadAndDelete(key any) (value any, loaded bool) {
	m.mu.Proceed(func() bool {
		value, loaded = m.dirty[key]
		return loaded
	}, func() {
		delete(m.dirty, key)
	})
	return
}

func (m *smartMutexMap) Delete(key any) {
	m.mu.Lock()
	delete(m.dirty, key)
	m.mu.Unlock()
}

func (m *smartMutexMap) CompareAndSwap(key, old, new any) (swapped bool) {
	m.mu.Proceed(func() bool {
		value, loaded := m.dirty[key]
		if loaded && value == old {
			swapped = true
		}
		return swapped
	}, func() {
		m.dirty[key] = new
	})
	return
}

func (m *smartMutexMap) CompareAndDelete(key, old any) (deleted bool) {
	m.mu.Proceed(func() bool {
		value, loaded := m.dirty[key]
		if loaded && value == old {
			deleted = true
		}
		return deleted
	}, func() {
		delete(m.dirty, key)
	})
	return
}

func (m *smartMutexMap) Range(f func(key, value any) (shouldContinue bool)) {
	m.mu.RLock()
	keys := make([]any, 0, len(m.dirty))
	for k := range m.dirty {
		keys = append(keys, k)
	}
	m.mu.RUnlock()

	for _, k := range keys {
		v, ok := m.Load(k)
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}

type (
	// rwMutexMap is an implementation of mapInterfaceAny using a sync.RWMutex.
	rwMutexMap struct {
		mu    sync.RWMutex
		dirty map[any]any
	}
	// RWMutexMap is an implementation of mapInterface using a sync.RWMutex.
	RWMutexMap[K comparable, V any] struct {
		base rwMutexMap
	}
)

func (m *RWMutexMap[K, V]) Load(s K) (value V, ok bool) {
	vi, ok := m.base.Load(s)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *RWMutexMap[K, V]) Store(key K, value V) {
	m.base.Store(key, value)
}

func (m *RWMutexMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	vi, loaded := m.base.LoadOrStore(key, value)
	if vi != nil {
		actual = vi.(V)
	}
	return
}

func (m *RWMutexMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	vi, loaded := m.base.LoadAndDelete(key)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *RWMutexMap[K, V]) Delete(s K) {
	m.base.Delete(s)
}

func (m *RWMutexMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	vi, loaded := m.base.Swap(key, value)
	if vi != nil {
		previous = vi.(V)
	}
	return
}

func (m *RWMutexMap[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return m.base.CompareAndSwap(key, old, new)
}

func (m *RWMutexMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.base.CompareAndDelete(key, old)
}

func (m *RWMutexMap[K, V]) Range(f func(key K, value V) (shouldContinue bool)) {
	m.base.Range(func(k, v any) bool {
		key, _ := k.(K)
		value, _ := v.(V)
		return f(key, value)
	})
}

func (m *rwMutexMap) Load(key any) (value any, ok bool) {
	m.mu.RLock()
	value, ok = m.dirty[key]
	m.mu.RUnlock()
	return
}

func (m *rwMutexMap) Store(key, value any) {
	m.mu.Lock()
	if m.dirty == nil {
		m.dirty = make(map[any]any)
	}
	m.dirty[key] = value
	m.mu.Unlock()
}

func (m *rwMutexMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	m.mu.Lock()
	actual, loaded = m.dirty[key]
	if !loaded {
		actual = value
		if m.dirty == nil {
			m.dirty = make(map[any]any)
		}
		m.dirty[key] = value
	}
	m.mu.Unlock()
	return actual, loaded
}

func (m *rwMutexMap) Swap(key, value any) (previous any, loaded bool) {
	m.mu.Lock()
	if m.dirty == nil {
		m.dirty = make(map[any]any)
	}

	previous, loaded = m.dirty[key]
	m.dirty[key] = value
	m.mu.Unlock()
	return
}

func (m *rwMutexMap) LoadAndDelete(key any) (value any, loaded bool) {
	m.mu.Lock()
	value, loaded = m.dirty[key]
	if !loaded {
		m.mu.Unlock()
		return nil, false
	}
	delete(m.dirty, key)
	m.mu.Unlock()
	return value, loaded
}

func (m *rwMutexMap) Delete(key any) {
	m.mu.Lock()
	delete(m.dirty, key)
	m.mu.Unlock()
}

func (m *rwMutexMap) CompareAndSwap(key, old, new any) (swapped bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.dirty == nil {
		return false
	}

	value, loaded := m.dirty[key]
	if loaded && value == old {
		m.dirty[key] = new
		return true
	}
	return false
}

func (m *rwMutexMap) CompareAndDelete(key, old any) (deleted bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.dirty == nil {
		return false
	}

	value, loaded := m.dirty[key]
	if loaded && value == old {
		delete(m.dirty, key)
		return true
	}
	return false
}

func (m *rwMutexMap) Range(f func(key, value any) (shouldContinue bool)) {
	m.mu.RLock()
	keys := make([]any, 0, len(m.dirty))
	for k := range m.dirty {
		keys = append(keys, k)
	}
	m.mu.RUnlock()

	for _, k := range keys {
		v, ok := m.Load(k)
		if !ok {
			continue
		}
		if !f(k, v) {
			break
		}
	}
}

type (
	// deepCopyMap is an implementation of mapInterfaceAny using a Mutex and atomic.Value.
	//It makes deep copies of the map on every write to avoid acquiring the Mutex in Load.
	deepCopyMap struct {
		mu    sync.Mutex
		clean atomic.Value
	}
	// deepCopyMap is an implementation of mapInterface using a Mutex and atomic.Value.
	//It makes deep copies of the map on every write to avoid acquiring the Mutex in Load.
	DeepCopyMap[K comparable, V any] struct {
		base deepCopyMap
	}
)

func (m *DeepCopyMap[K, V]) Load(s K) (value V, ok bool) {
	vi, ok := m.base.Load(s)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *DeepCopyMap[K, V]) Store(key K, value V) {
	m.base.Store(key, value)
}

func (m *DeepCopyMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	vi, loaded := m.base.LoadOrStore(key, value)
	if vi != nil {
		actual = vi.(V)
	}
	return
}

func (m *DeepCopyMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	vi, loaded := m.base.LoadAndDelete(key)
	if vi != nil {
		value = vi.(V)
	}
	return
}

func (m *DeepCopyMap[K, V]) Delete(s K) {
	m.base.Delete(s)
}

func (m *DeepCopyMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	vi, loaded := m.base.Swap(key, value)
	if vi != nil {
		previous = vi.(V)
	}
	return
}

func (m *DeepCopyMap[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return m.base.CompareAndSwap(key, old, new)
}

func (m *DeepCopyMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.base.CompareAndDelete(key, old)
}

func (m *DeepCopyMap[K, V]) Range(f func(key K, value V) (shouldContinue bool)) {
	m.base.Range(func(k, v any) bool {
		key, _ := k.(K)
		value, _ := v.(V)
		return f(key, value)
	})
}

func (m *deepCopyMap) Load(key any) (value any, ok bool) {
	clean, _ := m.clean.Load().(map[any]any)
	value, ok = clean[key]
	return value, ok
}

func (m *deepCopyMap) Store(key, value any) {
	m.mu.Lock()
	dirty := m.dirtyClone()
	dirty[key] = value
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *deepCopyMap) LoadOrStore(key, value any) (actual any, loaded bool) {
	clean, _ := m.clean.Load().(map[any]any)
	actual, loaded = clean[key]
	if loaded {
		return actual, loaded
	}

	m.mu.Lock()
	// Reload clean in case it changed while we were waiting on m.mu.
	clean, _ = m.clean.Load().(map[any]any)
	actual, loaded = clean[key]
	if !loaded {
		dirty := m.dirtyClone()
		dirty[key] = value
		actual = value
		m.clean.Store(dirty)
	}
	m.mu.Unlock()
	return actual, loaded
}

func (m *deepCopyMap) Swap(key, value any) (previous any, loaded bool) {
	m.mu.Lock()
	dirty := m.dirtyClone()
	previous, loaded = dirty[key]
	dirty[key] = value
	m.clean.Store(dirty)
	m.mu.Unlock()
	return
}

func (m *deepCopyMap) LoadAndDelete(key any) (value any, loaded bool) {
	m.mu.Lock()
	dirty := m.dirtyClone()
	value, loaded = dirty[key]
	delete(dirty, key)
	m.clean.Store(dirty)
	m.mu.Unlock()
	return
}

func (m *deepCopyMap) Delete(key any) {
	m.mu.Lock()
	dirty := m.dirtyClone()
	delete(dirty, key)
	m.clean.Store(dirty)
	m.mu.Unlock()
}

func (m *deepCopyMap) CompareAndSwap(key, old, new any) (swapped bool) {
	clean, _ := m.clean.Load().(map[any]any)
	if previous, ok := clean[key]; !ok || previous != old {
		return false
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	dirty := m.dirtyClone()
	value, loaded := dirty[key]
	if loaded && value == old {
		dirty[key] = new
		m.clean.Store(dirty)
		return true
	}
	return false
}

func (m *deepCopyMap) CompareAndDelete(key, old any) (deleted bool) {
	clean, _ := m.clean.Load().(map[any]any)
	if previous, ok := clean[key]; !ok || previous != old {
		return false
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	dirty := m.dirtyClone()
	value, loaded := dirty[key]
	if loaded && value == old {
		delete(dirty, key)
		m.clean.Store(dirty)
		return true
	}
	return false
}

func (m *deepCopyMap) Range(f func(key, value any) (shouldContinue bool)) {
	clean, _ := m.clean.Load().(map[any]any)
	for k, v := range clean {
		if !f(k, v) {
			break
		}
	}
}

func (m *deepCopyMap) dirtyClone() map[any]any {
	clean, _ := m.clean.Load().(map[any]any)
	//dirty := make(map[any]any, len(clean)+1)
	dirty := maps.Clone(clean)
	if dirty == nil {
		dirty = make(map[any]any)
	}
	return dirty
}

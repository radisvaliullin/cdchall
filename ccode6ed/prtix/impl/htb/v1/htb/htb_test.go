package htb

import (
	"strconv"
	"testing"
)

func TestHTB(t *testing.T) {

	// go map for validat our implementation
	m := map[string]string{}

	// hash table
	ht := NewHTB(0)

	// generate random key and value
	for i := 0; i < 1000_000; i++ {
		k := strconv.Itoa(i)
		v := k
		m[k] = v
		ht.Add(k, v)
	}

	// validate
	for k, v := range m {
		if tv, ok := ht.Get(k); !ok {
			t.Fatalf("key - %v, val - %v not found", k, v)
		} else {
			if tv != v {
				t.Fatalf("key - %v has diff values: v - %v, hv - %v", k, v, tv)
			}
		}
	}
	mlen := len(m)
	htlen := ht.Len()
	if mlen != htlen {
		t.Fatalf("len of m not equal len of ht: mlen - %v, htlen - %v", mlen, htlen)
	}
	t.Logf("len of m and ht: mlen - %v, htlen - %v", mlen, htlen)
	t.Logf("ht buckets size: %v", len(ht.buckets))
}

func TestHTBDelete(t *testing.T) {

	// go map for validat our implementation
	m := map[string]string{}

	// hash table
	ht := NewHTB(0)

	// generate random key and value
	for i := 0; i < 1000_000; i++ {
		k := strconv.Itoa(i)
		v := k
		m[k] = v
		ht.Add(k, v)
	}
	for i := 0; i < 1000_000; i++ {
		if i%2 == 0 {
			k := strconv.Itoa(i)
			delete(m, k)
			ht.Del(k)
		}
	}

	// validate
	for k, v := range m {
		if tv, ok := ht.Get(k); !ok {
			t.Fatalf("key - %v, val - %v not found", k, v)
		} else {
			if tv != v {
				t.Fatalf("key - %v has diff values: v - %v, hv - %v", k, v, tv)
			}
		}
	}
	mlen := len(m)
	htlen := ht.Len()
	if mlen != htlen {
		t.Fatalf("len of m not equal len of ht: mlen - %v, htlen - %v", mlen, htlen)
	}
	t.Logf("len of m and ht: mlen - %v, htlen - %v", mlen, htlen)
	t.Logf("ht buckets size: %v", len(ht.buckets))
}

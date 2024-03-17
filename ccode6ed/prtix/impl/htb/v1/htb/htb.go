package htb

import (
	"encoding/json"
	"hash/maphash"
)

const (
	// multiple of two
	MaxSlotSize = 8
	// better if bucket size multiple of 2
	DefaultBucketSize = MaxSlotSize
)

// Implement simple hash table (not thread safe).
// Inspired from Go Map type implem (very simplified).
// I guess it is common approach but I looked example in Go source code.
// hash table (where key/value a string type)
// size of buckets should be multiple of 8
// number slots in bucket 8
// when slot overflow double size of buckets
// bucket size never reduced (after delete key)
type HTB struct {
	buckets []bucket

	// for hash function
	seed maphash.Seed
}

func NewHTB(size int) *HTB {
	// if 0 use default size
	if size == 0 {
		size = DefaultBucketSize
	}
	// make size multiple of 8
	if size%MaxSlotSize != 0 {
		size = (size/MaxSlotSize)*MaxSlotSize + MaxSlotSize
	}

	t := &HTB{
		buckets: make([]bucket, size),
	}
	// allocate slots
	for i := range t.buckets {
		t.buckets[i].slots = make([]slot, 0, MaxSlotSize)
	}
	// set seed
	t.seed = maphash.MakeSeed()
	return t
}

// add item to next empty slot
// slots of same backet stor items with same bucket idx (collision)
// for slots we use preallocated capacity
// use very simpel approach for reallocation buckets
// if size of slots of bucket equal capacity we need double size of buckets
// for spread slots between different buckets
type bucket struct {
	slots []slot
}

type slot struct {
	key *string
	val *string
}

func (t *HTB) Add(k, v string) {
	// key hash (use Go default map hash function)
	h := maphash.String(t.seed, k)
	// bucket index
	bidx := h % uint64(len(t.buckets))
	bucket := t.buckets[bidx]

	// set key val
	// update if exist
	for sidx := 0; sidx < len(bucket.slots); sidx++ {
		if *bucket.slots[sidx].key == k {
			bucket.slots[sidx].val = &v
			// if update we can return
			return
		}
	}
	// if not exist add
	// check if need reallocate
	if len(bucket.slots) == MaxSlotSize {
		t.reallocate()
		// find new bucket index and bslots
		bidx = h % uint64(len(t.buckets))
	}
	t.buckets[bidx].slots = append(t.buckets[bidx].slots, slot{key: &k, val: &v})
}

// return value and find status
func (t *HTB) Get(k string) (string, bool) {
	h := maphash.String(t.seed, k)
	bidx := h % uint64(len(t.buckets))
	for i := 0; i < len(t.buckets[bidx].slots); i++ {
		slot := t.buckets[bidx].slots[i]
		if *slot.key == k {
			return *slot.val, true
		}
	}
	return "", false
}

func (t *HTB) Del(k string) {
	h := maphash.String(t.seed, k)
	bidx := h % uint64(len(t.buckets))
	for i := 0; i < len(t.buckets[bidx].slots); i++ {
		slot := t.buckets[bidx].slots[i]
		if *slot.key == k {
			t.buckets[bidx].slots = append(t.buckets[bidx].slots[:i], t.buckets[bidx].slots[i+1:]...)
			return
		}
	}
}

// TODO need finish
func (t *HTB) reallocate() {

}

// just for debug
func (t *HTB) String() string {
	m := map[string]string{}
	for _, b := range t.buckets {
		for _, s := range b.slots {
			m[*s.key] = *s.val
		}
	}
	jb, _ := json.Marshal(&m)
	return string(jb)
}

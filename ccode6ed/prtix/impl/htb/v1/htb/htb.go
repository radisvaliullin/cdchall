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
	// bucket idx
	bidx := getBIdx(t.buckets, k, t.seed)
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
		// find new bucket index
		bidx = getBIdx(t.buckets, k, t.seed)
	}
	// we never reallocate slice array of slots because track size of slice (see if len check above)
	t.buckets[bidx].slots = append(t.buckets[bidx].slots, slot{key: &k, val: &v})
}

// return value and find status
func (t *HTB) Get(k string) (string, bool) {
	bidx := getBIdx(t.buckets, k, t.seed)
	for i := 0; i < len(t.buckets[bidx].slots); i++ {
		slot := t.buckets[bidx].slots[i]
		if *slot.key == k {
			return *slot.val, true
		}
	}
	return "", false
}

func (t *HTB) Del(k string) {
	bidx := getBIdx(t.buckets, k, t.seed)
	for i := 0; i < len(t.buckets[bidx].slots); i++ {
		slot := t.buckets[bidx].slots[i]
		if *slot.key == k {
			// we can use generic slices.Delete method, but for learn purpose I used old manual approach
			oldLen := len(t.buckets[bidx].slots)
			t.buckets[bidx].slots = append(t.buckets[bidx].slots[:i], t.buckets[bidx].slots[i+1:]...)
			// clean deleted slot pointers
			newLen := len(t.buckets[bidx].slots)
			delSlots := t.buckets[bidx].slots[newLen : oldLen+7]
			for i := range delSlots {
				delSlots[i].key, delSlots[i].val = nil, nil
			}
			return
		}
	}
}

// allocates new buckets and copy slots to new buckets
// simple implementation without optimizations
func (t *HTB) reallocate() {
	bSize := len(t.buckets)

	// use loop
	// if we again get slot overflow we need again double size of buckets
	var newBuckets []bucket
newbuckets:
	for {
		bSize = bSize * 2
		newBuckets = make([]bucket, bSize)
		// allocate slots
		for i := range newBuckets {
			newBuckets[i].slots = make([]slot, 0, MaxSlotSize)
		}
		// copy slots to new buckets
		for bidx := 0; bidx < len(t.buckets); bidx++ {
			for sidx := 0; sidx < len(t.buckets[bidx].slots); sidx++ {
				kp, vp := t.buckets[bidx].slots[sidx].key, t.buckets[bidx].slots[sidx].val
				// new
				nbidx := getBIdx(newBuckets, *kp, t.seed)
				newBuckets[nbidx].slots = append(newBuckets[nbidx].slots, slot{key: kp, val: vp})
				// check slots overflow
				if len(newBuckets[nbidx].slots) == MaxSlotSize {
					continue newbuckets
				}
			}
		}
		// if all slots copied break
		break newbuckets
	}
	t.buckets = newBuckets
}

func (t *HTB) Len() int {
	ln := 0
	for bidx := range t.buckets {
		for range t.buckets[bidx].slots {
			ln++
		}
	}
	return ln
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

// get buckets idx
func getBIdx(buckets []bucket, k string, seed maphash.Seed) uint64 {
	// key hash (use Go default map hash function)
	h := maphash.String(seed, k)
	return h % uint64(len(buckets))
}

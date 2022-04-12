package algs4

type node struct {
	key, value interface{}
	next       *node
}

type SequentialSearchST struct {
	first *node
	n     int
}

func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{}
}

func (st *SequentialSearchST) Get(key interface{}) interface{} {
	for node := st.first; node != nil; node = node.next {
		if node.key == key {
			return node.value
		}
	}
	return nil
}

func (st *SequentialSearchST) Put(key interface{}, value interface{}) {
	for node := st.first; node != nil; node = node.next {
		if node.key == key {
			node.value = value
			return
		}
	}
	newNode := node{
		key:   key,
		value: value,
		next:  st.first,
	}
	st.first = &newNode
	st.n++
	return
}

func (st *SequentialSearchST) Delete(key interface{}) {
	st.Put(key, nil)
}

func (st *SequentialSearchST) Size() int {
	return st.n
}

func (st SequentialSearchST) IsEmpty() bool {
	return st.Size() == 0
}

func (st *SequentialSearchST) Contains(key interface{}) bool {
	return st.Get(key) != nil
}

func (st *SequentialSearchST) Keys() []interface{} {
	var keys []interface{}
	for node := st.first; node != nil; node = node.next {
		keys = append(keys, node.key)
	}
	return keys
}

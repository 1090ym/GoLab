package algs4

// Key is an interface of the keys in ST
type Key interface {
	compareTo(interface{}) int
}

// StringKey implements Key
type StringKey string

// CompareTo ...
func (k StringKey) compareTo(other interface{}) int {
	if k < other.(StringKey) {
		return -1
	} else if k > other.(StringKey) {
		return 1
	}
	return 0
}

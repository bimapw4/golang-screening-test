package queue

/*
	struct will be used in queue process
*/
type queue []interface{}

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func New(size int) Queue {
	return nil
}

/*
	function will be used for save value in struct
*/
func (self *queue) Push(x interface{}) {
	*self = append(*self, x)
}

/*
	function will be used for get key,

	====== Note ======
	this function already done
*/
func (self *queue) Keys() []interface{} {
	return *self
}

/*
	function will be used for get length from struct
*/
func (self *queue) Len() int {
	return self.Len()
}

/*
	function will be used for get first value from struct
*/
func (self *queue) Pop() interface{} {
	h := *self
	var el interface{}
	l := len(h)
	el, *self = h[0], h[1:l]
	// Or use this instead for a Stack
	// el, *self = h[l-1], h[0:l-1]
	return el
}

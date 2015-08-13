package xmrq

type Queue struct {
	c                 chan int
	task              []interface{}
	head, tail, count int
}

func NewQueue() *Queue {
	return &Queue{
		c:    make(chan int, 1),
		task: make([]interface{}, 12),
	}
}

func (q *Queue) Length() int {
	return q.count
}

func (q *Queue) resize() {
	newTask := make([]interface{}, q.count*2)
	if q.tail > q.head {
		copy(newTask, q.task[q.head:q.tail])
	} else {
		n := copy(newTask, q.task[q.head:])
		copy(newTask[n:], q.task[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.task = newTask
}

func (q *Queue) Add(e interface{}) {
	q.c <- 1
	if q.count == len(q.task) {
		q.resize()
	}
	q.task[q.tail] = e
	q.tail = (q.tail + 1) % len(q.task)
	q.count++

	<-q.c
}

func (q *Queue) Peek() interface{} {
	q.c <- 1
	if q.count == 0 {
		panic("queue is empty !")
	}
	<-q.c

	return q.task[q.head]
}

func (q *Queue) Get(i int) interface{} {
	if i < 0 || i >= q.count {
		panic("number is error !")
	}

	return q.task[(q.head+i)%len(q.task)]
}

func (q *Queue) Remove() {
	q.c <- 1
	if q.count <= 0 {
		panic("queue is empty !")
	}
	q.task[q.head] = nil
	q.head = (q.head + 1) % len(q.task)
	q.count--
	if len(q.task) > 12 && q.count*4 == len(q.task) {
		q.resize()
	}
	<-q.c
}

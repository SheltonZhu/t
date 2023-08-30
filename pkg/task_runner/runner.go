package task_runner

type Runner interface {
	run(*DocumentTask)
	setNext(Runner)
}

type DocInfo struct {
	Name  string
	Path  string
	Type  string
	DocID int
}

type DocumentTask struct {
	DocInfo   []DocInfo
	TaskId    int
	TaskType  string
	parsed    bool
	compared  bool
	extracted bool
	reviewed  bool
}

// Parser 解析器
type Parser struct {
	next       Runner
	HandleFunc func(docTask *DocumentTask)
}

func (p *Parser) run(docTask *DocumentTask) {
	if docTask.parsed && p.next != nil {
		p.next.run(docTask)
		return
	}

	p.HandleFunc(docTask)
	docTask.parsed = true
	if p.next != nil {
		p.next.run(docTask)
	}
}
func (p *Parser) setNext(next Runner) { p.next = next }

// Comparator 比较器
type Comparator struct {
	next       Runner
	HandleFunc func(docTask *DocumentTask)
}

func (c *Comparator) run(docTask *DocumentTask) {
	if docTask.compared && c.next != nil {
		c.next.run(docTask)
		return
	}

	c.HandleFunc(docTask)
	docTask.compared = true
	if c.next != nil {
		c.next.run(docTask)
	}
}
func (c *Comparator) setNext(next Runner) { c.next = next }

// Extractor 抽取器
type Extractor struct {
	next       Runner
	HandleFunc func(docTask *DocumentTask)
}

func (e *Extractor) run(docTask *DocumentTask) {
	if docTask.extracted && e.next != nil {
		e.next.run(docTask)
		return
	}

	e.HandleFunc(docTask)
	docTask.extracted = true
	if e.next != nil {
		e.next.run(docTask)
	}
}
func (e *Extractor) setNext(next Runner) { e.next = next }

// Reviewer 审核器
type Reviewer struct {
	next       Runner
	HandleFunc func(docTask *DocumentTask)
}

func (r *Reviewer) run(docTask *DocumentTask) {
	if docTask.reviewed && r.next != nil {
		r.next.run(docTask)
		return
	}

	r.HandleFunc(docTask)
	docTask.reviewed = true
	if r.next != nil {
		r.next.run(docTask)
	}
}
func (r *Reviewer) setNext(next Runner) { r.next = next }

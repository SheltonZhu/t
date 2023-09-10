package workflow

// Runner 任务运行
type Runner interface {
	run(*DocumentTask) error
	setNext(Runner)
}

// RunnerHandleFunc 具体处理函数
type RunnerHandleFunc func(docTask *DocumentTask) error

// Parser 解析器
type Parser struct {
	next       Runner
	HandleFunc RunnerHandleFunc
}

func (p *Parser) run(docTask *DocumentTask) error {
	if docTask.parsed && p.next != nil {
		return p.next.run(docTask)
	}

	if err := p.HandleFunc(docTask); err != nil {
		return err
	}
	docTask.parsed = true
	if p.next != nil {
		return p.next.run(docTask)
	}
	return nil
}

func (p *Parser) setNext(next Runner) { p.next = next }

// Comparator 比较器
type Comparator struct {
	next       Runner
	HandleFunc RunnerHandleFunc
}

func (c *Comparator) run(docTask *DocumentTask) error {
	if docTask.compared && c.next != nil {
		return c.next.run(docTask)
	}

	if err := c.HandleFunc(docTask); err != nil {
		return err
	}
	docTask.compared = true
	if c.next != nil {
		return c.next.run(docTask)
	}
	return nil
}
func (c *Comparator) setNext(next Runner) { c.next = next }

// Extractor 抽取器
type Extractor struct {
	next       Runner
	HandleFunc RunnerHandleFunc
}

func (e *Extractor) run(docTask *DocumentTask) error {
	if docTask.extracted && e.next != nil {
		return e.next.run(docTask)
	}

	if err := e.HandleFunc(docTask); err != nil {
		return err
	}
	docTask.extracted = true
	if e.next != nil {
		return e.next.run(docTask)
	}
	return nil
}
func (e *Extractor) setNext(next Runner) { e.next = next }

// Reviewer 审核器
type Reviewer struct {
	next       Runner
	HandleFunc RunnerHandleFunc
}

func (r *Reviewer) run(docTask *DocumentTask) error {
	if docTask.reviewed && r.next != nil {
		return r.next.run(docTask)
	}

	if err := r.HandleFunc(docTask); err != nil {
		return err
	}
	docTask.reviewed = true
	if r.next != nil {
		return r.next.run(docTask)
	}
	return nil
}
func (r *Reviewer) setNext(next Runner) { r.next = next }

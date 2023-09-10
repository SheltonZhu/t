package workflow

// DocumentTask 文档任务
type DocumentTask struct {
	TaskType  string      `json:"type"`
	Payload   TaskPayload `json:"payload"`
	parsed    bool
	compared  bool
	extracted bool
	reviewed  bool
}

// TaskPayload 任务主体数据
type TaskPayload struct {
	DocInfo []DocInfo `json:"-"`
	TaskID  int       `json:"task_id"`
	Step    int       `json:"task_step"`
}

// DocInfo 文档信息
type DocInfo struct {
	Name  string
	Path  string
	Type  string
	DocID int
}

// Parsed 修改解析状态
func (t *DocumentTask) Parsed() {
	t.parsed = true
}

// Compared 修改比对状态
func (t *DocumentTask) Compared() {
	t.compared = true
}

// Extracted 修改抽取状态
func (t *DocumentTask) Extracted() {
	t.extracted = true
}

// Reviewed 修改审核状态
func (t *DocumentTask) Reviewed() {
	t.reviewed = true
}

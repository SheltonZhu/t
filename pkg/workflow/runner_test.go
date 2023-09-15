package workflow

import (
	// "encoding/json"
	// "fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	t.Run("Test Comparison Task", func(t *testing.T) {
		task := DocumentTask{
			TaskType: "comparison",
			Payload: TaskPayload{
				TaskID: 1,
				DocInfo: []DocInfo{{
					DocID: 1,
					Name:  "test",
					Type:  "pdf",
					Path:  "/test1.pdf",
				}, {
					DocID: 2,
					Name:  "test2",
					Type:  "docx",
					Path:  "/test2.docx",
				}},
			},
		}

		c := &Comparator{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("比对文件: \n%+v\n", docTask)
			return nil
		}}
		p := &Parser{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("解析文件: \n%+v\n", docTask)
			return nil
		}}
		p.setNext(c)
		p.run(&task)
	})

	t.Run("Test Extraction task", func(t *testing.T) {
		task := DocumentTask{
			TaskType: "extraction",
			Payload: TaskPayload{
				DocInfo: []DocInfo{{
					DocID: 3,
					Name:  "test",
					Type:  "pdf",
					Path:  "/test1.pdf",
				}},
				TaskID: 2,
			},
		}

		c := &Extractor{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("抽取文件: \n%+v\n", docTask)
			return nil
		}}
		p := &Parser{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("解析文件: \n%+v\n", docTask)
			return nil
		}}
		p.setNext(c)
		p.run(&task)
	})

	t.Run("Test Restoration task", func(t *testing.T) {
		task := DocumentTask{
			TaskType: "restoration",
			Payload: TaskPayload{
				DocInfo: []DocInfo{{
					DocID: 4,
					Name:  "test",
					Type:  "pdf",
					Path:  "/test1.pdf",
				}},
				TaskID: 3,
			},
		}

		p := &Parser{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("解析文件: \n%+v\n", docTask)
			return nil
		}}
		p.run(&task)
	})

	t.Run("Test review task", func(t *testing.T) {
		task := DocumentTask{
			TaskType: "review",
			Payload: TaskPayload{
				DocInfo: []DocInfo{{
					DocID: 5,
					Name:  "test",
					Type:  "pdf",
					Path:  "/test1.pdf",
				}},
				TaskID: 4,
			},
		}

		r := &Reviewer{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("审核文件: \n%+v\n", docTask)
			return nil
		}}
		c := &Extractor{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("抽取文件: \n%+v\n", docTask)
			return nil
		}}
		p := &Parser{HandleFunc: func(docTask *DocumentTask) error {
			// fmt.Printf("解析文件: \n%+v\n", docTask)
			return nil
		}}
		c.setNext(r)
		p.setNext(c)
		p.run(&task)

		task.parsed = true
		task.extracted = false
		task.reviewed = false
		p.run(&task)
	})

	t.Run("Test not exist combo", func(t *testing.T) {
		task := &DocumentTask{
			TaskType: "review",
			Payload: TaskPayload{
				DocInfo: []DocInfo{{
					DocID: 6,
					Name:  "test",
					Type:  "pdf",
					Path:  "/test1.pdf",
				}},
				TaskID: 5,
			},
		}
		e := &Extractor{HandleFunc: func(docTask *DocumentTask) error { return nil }}
		ee := &Extractor{HandleFunc: func(docTask *DocumentTask) error { return nil }}
		c := &Comparator{HandleFunc: func(docTask *DocumentTask) error { return nil }}
		cc := &Comparator{HandleFunc: func(docTask *DocumentTask) error { return nil }}
		r := &Reviewer{HandleFunc: func(docTask *DocumentTask) error { return nil }}
		rr := &Reviewer{HandleFunc: func(docTask *DocumentTask) error { return nil }}

		rr.setNext(ee)
		cc.setNext(rr)
		r.setNext(cc)
		c.setNext(r)
		e.setNext(c)

		task.extracted = true
		e.run(task)
	})
}

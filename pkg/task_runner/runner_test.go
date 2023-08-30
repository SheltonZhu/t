package task_runner

import (
	"fmt"
	"testing"
)

func TestRunner(t *testing.T) {
	t.Parallel()
	t.Run("Test Comparison Task", func(t *testing.T) {
		task := DocumentTask{
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
			TaskId:   1,
			TaskType: "comparison",
		}

		c := &Comparator{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("比对文件: \n%+v\n", docTask)
		}}
		p := &Parser{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("解析文件: \n%+v\n", docTask)
		}}
		p.setNext(c)
		p.run(&task)
	})

	t.Run("Test Extraction task", func(t *testing.T) {
		task := DocumentTask{
			DocInfo: []DocInfo{{
				DocID: 3,
				Name:  "test",
				Type:  "pdf",
				Path:  "/test1.pdf",
			}},
			TaskId:   2,
			TaskType: "extraction",
		}

		c := &Extractor{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("抽取文件: \n%+v\n", docTask)
		}}
		p := &Parser{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("解析文件: \n%+v\n", docTask)
		}}
		p.setNext(c)
		p.run(&task)
	})

	t.Run("Test Restoration task", func(t *testing.T) {
		task := DocumentTask{
			DocInfo: []DocInfo{{
				DocID: 4,
				Name:  "test",
				Type:  "pdf",
				Path:  "/test1.pdf",
			}},
			TaskId:   3,
			TaskType: "restoration",
		}

		p := &Parser{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("解析文件: \n%+v\n", docTask)
		}}
		p.run(&task)
	})

	t.Run("Test review task", func(t *testing.T) {
		task := DocumentTask{
			DocInfo: []DocInfo{{
				DocID: 5,
				Name:  "test",
				Type:  "pdf",
				Path:  "/test1.pdf",
			}},
			TaskId:   4,
			TaskType: "review",
		}

		r := &Reviewer{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("审核文件: \n%+v\n", docTask)
		}}
		c := &Extractor{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("抽取文件: \n%+v\n", docTask)
		}}
		p := &Parser{HandleFunc: func(docTask *DocumentTask) {
			fmt.Printf("解析文件: \n%+v\n", docTask)
		}}
		c.setNext(r)
		p.setNext(c)
		p.run(&task)

		task.parsed = true
		task.extracted = false
		task.reviewed = false
		p.run(&task)
	})
}

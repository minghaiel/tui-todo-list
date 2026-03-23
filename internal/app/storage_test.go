package app

import (
	"path/filepath"
	"testing"
)

func TestSaveAndLoadTodos(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "todos.json")
	in := []todo{{Title: "write tests", Category: "work", Priority: "high"}}

	if err := saveTodos(path, in); err != nil {
		t.Fatalf("saveTodos error: %v", err)
	}

	out, err := loadTodos(path)
	if err != nil {
		t.Fatalf("loadTodos error: %v", err)
	}

	if len(out) != 1 || out[0].Title != "write tests" || out[0].Priority != "high" {
		t.Fatalf("unexpected todos: %+v", out)
	}
}

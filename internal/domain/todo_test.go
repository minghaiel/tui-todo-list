package domain

import (
	"testing"
	"time"
)

func TestFilterAndSortIndexes_PriorityThenDueDate(t *testing.T) {
	now := time.Date(2026, 3, 23, 10, 0, 0, 0, time.Local)
	todos := []Todo{
		{Title: "low later", Priority: "low", DueDate: "2026-03-30"},
		{Title: "high later", Priority: "high", DueDate: "2026-03-28"},
		{Title: "high sooner", Priority: "high", DueDate: "2026-03-24"},
		{Title: "urgent", Priority: "urgent", DueDate: "2026-03-29"},
	}

	got := FilterAndSortIndexes(todos, Query{Status: FilterAll, Category: "all", Sort: SortPriorityDue}, now)
	want := []int{3, 2, 1, 0}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("index %d: got %v want %v", i, got, want)
		}
	}
}

func TestFilterAndSortIndexes_SearchAndStatus(t *testing.T) {
	now := time.Date(2026, 3, 23, 10, 0, 0, 0, time.Local)
	todos := []Todo{
		{Title: "write report", Category: "work", Priority: "high"},
		{Title: "buy fruit", Category: "life", Priority: "medium", Completed: true},
		{Title: "workout", Category: "health", Priority: "low"},
	}

	got := FilterAndSortIndexes(todos, Query{
		Status:   FilterOpen,
		Category: "all",
		Search:   "work",
		Sort:     SortPriorityDue,
	}, now)

	want := []int{0, 2}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("index %d: got %v want %v", i, got, want)
		}
	}
}

func TestCategoryOptions(t *testing.T) {
	todos := []Todo{
		{Category: "work"},
		{Category: "life"},
		{Category: ""},
	}
	got := CategoryOptions(todos)
	if len(got) != 4 || got[0] != "all" {
		t.Fatalf("unexpected category options: %v", got)
	}
}

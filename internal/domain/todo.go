package domain

import (
	"errors"
	"sort"
	"strings"
	"time"
)

const DateLayout = "2006-01-02"

type Todo struct {
	Title     string `json:"title"`
	Category  string `json:"category,omitempty"`
	Priority  string `json:"priority,omitempty"`
	DueDate   string `json:"due_date,omitempty"`
	Completed bool   `json:"completed"`
}

type StatusFilter int

const (
	FilterAll StatusFilter = iota
	FilterOpen
	FilterDone
)

type SortMode int

const (
	SortPriorityDue SortMode = iota
)

type Query struct {
	Status   StatusFilter
	Category string
	Search   string
	Sort     SortMode
}

func NormalizeTodos(todos []Todo, now time.Time) []Todo {
	if len(todos) == 0 {
		return []Todo{
			{Title: "升级这个 TUI", Category: "work", Priority: "high", DueDate: now.AddDate(0, 0, 2).Format(DateLayout)},
			{Title: "买点水果", Category: "life", Priority: "medium", DueDate: now.AddDate(0, 0, 1).Format(DateLayout)},
		}
	}

	for i := range todos {
		todos[i].Category = NormalizeCategory(todos[i].Category)
		todos[i].Priority = NormalizePriorityValue(todos[i].Priority)
		todos[i].DueDate = strings.TrimSpace(todos[i].DueDate)
	}
	return todos
}

func NormalizeCategory(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	if value == "" {
		return "inbox"
	}
	return value
}

func NormalizePriority(input string) (string, error) {
	priority := NormalizePriorityValue(input)
	if priority == "" {
		return "", errors.New("优先级必须是 low、medium、high 或 urgent。")
	}
	return priority, nil
}

func NormalizePriorityValue(input string) string {
	value := strings.TrimSpace(strings.ToLower(input))
	switch value {
	case "", "med":
		return "medium"
	case "low", "medium", "high", "urgent":
		return value
	default:
		return ""
	}
}

func PriorityOrder(priority string) int {
	switch NormalizePriorityValue(priority) {
	case "urgent":
		return 4
	case "high":
		return 3
	case "medium":
		return 2
	case "low":
		return 1
	default:
		return 0
	}
}

func IsOverdue(item Todo, now time.Time) bool {
	if item.Completed || strings.TrimSpace(item.DueDate) == "" {
		return false
	}
	date, err := time.Parse(DateLayout, item.DueDate)
	if err != nil {
		return false
	}
	ny, nm, nd := now.Date()
	today := time.Date(ny, nm, nd, 0, 0, 0, 0, now.Location())
	return date.Before(today)
}

func SameDay(a, b time.Time) bool {
	ay, am, ad := a.Date()
	by, bm, bd := b.Date()
	return ay == by && am == bm && ad == bd
}

func CategoryOptions(todos []Todo) []string {
	set := map[string]struct{}{"all": {}}
	for _, item := range todos {
		set[NormalizeCategory(item.Category)] = struct{}{}
	}

	var options []string
	for category := range set {
		options = append(options, category)
	}
	sort.Strings(options)
	for i, value := range options {
		if value == "all" && i > 0 {
			options[0], options[i] = options[i], options[0]
			break
		}
	}
	return options
}

func FilterAndSortIndexes(todos []Todo, query Query, now time.Time) []int {
	var indexes []int
	search := strings.TrimSpace(strings.ToLower(query.Search))
	category := NormalizeCategory(query.Category)
	if query.Category == "" || query.Category == "all" {
		category = "all"
	}

	for i, item := range todos {
		if !matchesStatus(item, query.Status) {
			continue
		}
		if category != "all" && NormalizeCategory(item.Category) != category {
			continue
		}
		if search != "" && !matchesSearch(item, search) {
			continue
		}
		indexes = append(indexes, i)
	}

	sort.SliceStable(indexes, func(i, j int) bool {
		left := todos[indexes[i]]
		right := todos[indexes[j]]
		return lessByPriorityDue(left, right, now)
	})

	return indexes
}

func matchesStatus(item Todo, status StatusFilter) bool {
	switch status {
	case FilterOpen:
		return !item.Completed
	case FilterDone:
		return item.Completed
	default:
		return true
	}
}

func matchesSearch(item Todo, search string) bool {
	haystack := strings.ToLower(strings.Join([]string{
		item.Title,
		NormalizeCategory(item.Category),
		NormalizePriorityValue(item.Priority),
		item.DueDate,
	}, " "))
	return strings.Contains(haystack, search)
}

func lessByPriorityDue(left, right Todo, now time.Time) bool {
	lp := PriorityOrder(left.Priority)
	rp := PriorityOrder(right.Priority)
	if lp != rp {
		return lp > rp
	}

	ld, lok := parseDueDate(left.DueDate, now)
	rd, rok := parseDueDate(right.DueDate, now)
	if lok && rok && !ld.Equal(rd) {
		return ld.Before(rd)
	}
	if lok != rok {
		return lok
	}

	if left.Completed != right.Completed {
		return !left.Completed
	}

	return strings.ToLower(left.Title) < strings.ToLower(right.Title)
}

func parseDueDate(value string, now time.Time) (time.Time, bool) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, false
	}
	date, err := time.Parse(DateLayout, value)
	if err != nil {
		return time.Time{}, false
	}
	y, m, d := date.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Location()), true
}

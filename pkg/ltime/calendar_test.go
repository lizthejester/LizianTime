package ltime_test

import (
	"testing"

	"github.com/Lizthejester/LizianTime/pkg/ltime"
)

func TestGetDayMonth(t *testing.T) {
	day := 31
	year := 2025
	month := "January"

	gotMonth, gotDay := ltime.GetDayMonth(year, month, day)
	wantMonth := "Menotheen"
	wantDay := 31

	if wantDay != gotDay {
		t.Errorf("day: got %d, want %d", gotDay, wantDay)
	}
	if gotMonth != wantMonth {
		t.Errorf("month: got %s, want %s", gotMonth, wantMonth)
	}
}

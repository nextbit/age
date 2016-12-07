package age

import (
	"testing"
	"time"
)

var peopleList = []struct {
	TestDate  time.Time
	Birthdate time.Time
	ExpAge    int
}{
	{
		time.Date(2016, time.July, 01, 15, 0, 0, 0, time.UTC),
		time.Date(1986, time.January, 03, 10, 0, 0, 0, time.UTC),
		30,
	},
	{
		time.Date(2016, time.July, 01, 15, 0, 0, 0, time.UTC),
		time.Date(1986, time.August, 03, 10, 0, 0, 0, time.UTC),
		29,
	},
	{
		time.Date(2016, time.July, 01, 15, 0, 0, 0, time.UTC),
		time.Date(1980, time.March, 18, 10, 0, 0, 0, time.UTC),
		36,
	},
	{
		time.Date(2016, time.July, 01, 15, 0, 0, 0, time.UTC),
		time.Date(1980, time.October, 18, 10, 0, 0, 0, time.UTC),
		35,
	},
	{
		time.Date(2016, time.July, 01, 15, 0, 0, 0, time.UTC),
		time.Date(1986, time.July, 01, 15, 0, 0, 0, time.UTC),
		30,
	},
}

func TestAge(t *testing.T) {
	for _, p := range peopleList {
		age, err := ageAt(p.TestDate, p.Birthdate)

		if err != nil {
			t.Errorf("expect the error to be nil but got: %v", err)
		}

		if p.ExpAge != age {
			t.Errorf("expect the age to be %v but was %v", p.ExpAge, age)
		}
	}
}

func TestBirthdateAfterNow(t *testing.T) {
	testDate := time.Date(2016, time.July, 01, 15, 0, 0, 0, time.UTC)
	birthdate := time.Date(2016, time.October, 18, 10, 0, 0, 0, time.UTC)

	age, err := ageAt(testDate, birthdate)

	if err == nil {
		t.Error("expect the error to not be nil but got nil")
	}

	if age != 0 {
		t.Errorf("expect the age to be 0 but was %v", age)
	}
}

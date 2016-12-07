package age

import (
	"fmt"
	"time"
)

// Age calculates the age based on the birthdate
func Age(b time.Time) (int, error) {
	return ageAt(time.Now(), b)
}

func ageAt(base, birthdate time.Time) (int, error) {
	if base.Before(birthdate) {
		return 0, fmt.Errorf("Birthdate should be before %v", time.Now())
	}

	age := base.Year() - birthdate.Year()

	if (base.Month() < birthdate.Month()) && (base.Day() < birthdate.Day()) {
		age--
	}

	return age, nil
}

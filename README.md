# Age

## Age calculations in Go.

A simple package to calculate the age based on the birthdate.

### Examples:

```go
// supose current date is 2016-07-01

package main

import "github.com/nextbit/age"

func main() {
	birthdate := time.Date(1986, time.July, 01, 15, 0, 0, 0, time.UTC),


	a, err := age.Age(birthdate)
	if err != nil {
		fmt.Println("Invalid date")
	}
	
	fmt.Println(a) // => will print 30
}
```
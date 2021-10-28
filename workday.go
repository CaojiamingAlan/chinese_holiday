package chinese_holiday

import (
	"time"
)

var LegalHolidays = []string{
	"2021-01-01","2021-01-02","2021-01-03",
	"2021-02-11","2021-02-12","2021-02-13","2021-02-14","2021-02-15","2021-02-16","2021-02-17",
	"2021-04-03","2021-04-04","2021-04-05",
	"2021-05-01","2021-05-02","2021-05-03","2021-05-04","2021-05-05",
	"2021-06-12","2021-06-13","2021-06-14",
	"2021-09-19","2021-09-20","2021-09-21",
	"2021-10-01","2021-10-02","2021-10-03","2021-10-04","2021-10-05","2021-10-06","2021-10-07",

	"2022-01-01","2022-01-02","2022-01-03",
	"2022-01-31","2022-02-01","2022-02-02","2022-02-03","2022-02-04","2022-02-05","2022-02-06",
	"2022-04-03","2022-04-04","2022-04-05",
	"2022-04-30","2022-05-01","2022-05-02","2022-05-03","2022-05-04",
	"2022-06-03","2022-06-04","2022-06-05",
	"2022-09-10","2022-09-11","2022-09-12",
	"2022-10-01","2022-10-02","2022-10-03","2022-10-04","2022-10-05","2022-10-06","2022-10-07",
}

var LegalWorkdays = []string{
	"2021-02-07","2021-02-20",
	"2021-04-25","2021-05-08",
	"2021-09-18",
	"2021-09-26","2021-10-09",

	"2022-01-29","2022-01-30",
	"2022-04-02",
	"2022-04-24","2022-05-07",
	"2022-10-08","2022-10-09",
}

func IsWorkday(date string) bool {
	if contains(LegalHolidays, date) {
		return false
	}

	if contains(LegalWorkdays, date) {
		return true
	}

	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		return false
	}

	if t.Weekday() == time.Saturday || t.Weekday() == time.Sunday {
		return false
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
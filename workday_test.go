package chinese_holiday

import "testing"

func TestIsWorkday(t *testing.T) {
	//2021-09-18 is a workday
	if IsWorkday("2021-09-18") == false {
		t.Fatal()
	}

	//2021-10-29 is Friday
	if IsWorkday("2021-10-29") == false {
		t.Fatal()
	}

	//2021-10-30 is Saturday
	if IsWorkday("2021-10-30") == true {
		t.Fatal()
	}
}
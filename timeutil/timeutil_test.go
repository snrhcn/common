package timeutil

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {

	ts := MakeTimestamp()

	millis, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return
	}
	tsTime := time.Unix(0, millis*1000000)
	tsTime = tsTime.UTC()

	tss, err := TimestampString(ts, "UTC")

	if err != nil {
		t.Error("Unexpected error during timestamp printing:", err)
	} else if tsTime.String() != tss {
		t.Error("Unexpected timestamp printing result:", tss)
	}

	_, err = TimestampString("abc", "UTC")

	if err.Error() != "strconv.ParseInt: parsing \"abc\": invalid syntax" {
		t.Error("Unexpected error during timestamp printing:", err)
	}

	_, err = TimestampString(ts, "U_B_C")

	if !strings.HasPrefix(err.Error(), "cannot find U_B_C") {
		t.Error("Unexpected error during timestamp printing:", err)
	}
}

package sum
import "testing"

func TestSum(t *testing.T) {
    actual := sum(1, 3)
    expected := 4
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
	}
}

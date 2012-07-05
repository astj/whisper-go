package whisper

import (
	"testing"
)


func TestQuantizeArchive(t *testing.T) {
	points := Archive{Point{0,0}, Point{3,0}, Point{10,0}}
	pointsOut := Archive{Point{0,0}, Point{2,0}, Point{10,0}}
	quantizedPoints := quantizeArchive(points, 2)
	for i := range quantizedPoints {
		if quantizedPoints[i] != pointsOut[i] {
			t.Errorf("%v != %v", quantizedPoints[i], pointsOut[i])
		}
	}
}

func TestAggregate(t *testing.T) {
	points := Archive{Point{0,0}, Point{0,1}, Point{0,2}, Point{0,1}}
	expected := Point{0,1}
	if p, err := aggregate(AGGREGATION_AVERAGE, points); (p != expected) || (err != nil) {
		t.Errorf("Average failed to average to %v, got %v: %v", expected, p, err)
	}

	expected = Point{0,4}
	if p, err := aggregate(AGGREGATION_SUM, points); (p != expected) || (err != nil) {
		t.Errorf("Sum failed to aggregate to %v, got %v: %v", expected, p, err)
	}

	expected = Point{0,1}
	if p, err := aggregate(AGGREGATION_LAST, points); (p != expected) || (err != nil) {
		t.Errorf("Last failed to aggregate to %v, got %v: %v", expected, p, err)
	}

	expected = Point{0,2}
	if p, err := aggregate(AGGREGATION_MAX, points); (p != expected) || (err != nil) {
		t.Errorf("Max failed to aggregate to %v, got %v: %v", expected, p, err)
	}

	expected = Point{0,0}
	if p, err := aggregate(AGGREGATION_MIN, points); (p != expected) || (err != nil) {
		t.Errorf("Min failed to aggregate to %v, got %v: %v", expected, p, err)
	}

	if _, err := aggregate(1000, points); err == nil {
		t.Errorf("No error for invalid aggregation")
	}
}


func TestParseArchiveInfo(t *testing.T) {
	tests := map[string]ArchiveInfo{
		"60:1440": ArchiveInfo{0, 60, 1440},	// 60 seconds per datapoint, 1440 datapoints = 1 day of retention
		"15m:8": ArchiveInfo{0, 15 * 60, 8},	// 15 minutes per datapoint, 8 datapoints = 2 hours of retention
		"1h:7d": ArchiveInfo{0, 3600, 168}, // 1 hour per datapoint, 7 days of retention
		"12h:2y": ArchiveInfo{0, 43200, 1456}, 	// 12 hours per datapoint, 2 years of retention
	}

	for info, expected := range tests {
		if a, err := ParseArchiveInfo(info); (a != expected) || (err != nil) {
			t.Errorf("%s: %v != %v, %v", info, a, expected, err)
		}
	}

}


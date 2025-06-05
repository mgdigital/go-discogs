package discogs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var letterRegex = regexp.MustCompile(`^([a-zA-Z])$`)
var letterNumberRegex = regexp.MustCompile(`^([a-zA-Z])(\d+)$`)
var trackNumberRegex = regexp.MustCompile(`^(\d+)$`)
var discTrackNumberRegex = regexp.MustCompile(`^\D*(\d+)[-.](\d+)$`)

func ParseDiscTrackNumber(position string) (int, int, error) {
	matches := trackNumberRegex.FindStringSubmatch(position)
	if len(matches) > 0 {
		disc := 1
		track, _ := strconv.Atoi(matches[1])
		return disc, track, nil
	}
	matches = letterRegex.FindStringSubmatch(position)
	if len(matches) > 0 {
		disc := 1 + int(strings.ToUpper(matches[1])[0]-'A')
		track := 1
		return disc, track, nil
	}
	matches = letterNumberRegex.FindStringSubmatch(position)
	if len(matches) > 0 {
		disc := 1 + int(strings.ToUpper(matches[1])[0]-'A')
		track, _ := strconv.Atoi(matches[2])
		return disc, track, nil
	}
	matches = discTrackNumberRegex.FindStringSubmatch(position)
	if len(matches) > 0 {
		disc, _ := strconv.Atoi(matches[1])
		track, _ := strconv.Atoi(matches[2])
		return disc, track, nil
	}
	return 0, 0, fmt.Errorf("invalid position: %s", position)
}

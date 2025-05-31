package discogs

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Release struct {
	ID              int             `json:"id"`
	Status          string          `json:"status"`
	Year            int             `json:"year"`
	ResourceURL     string          `json:"resource_url"`
	URI             string          `json:"uri"`
	Artists         []ReleaseArtist `json:"artists"`
	ArtistsSort     string          `json:"artists_sort"`
	Labels          []Company       `json:"labels"`
	Series          []Series        `json:"series"`
	Companies       []Company       `json:"companies"`
	Formats         []Format        `json:"formats"`
	DataQuality     string          `json:"data_quality"`
	Community       Community       `json:"community"`
	MasterID        int             `json:"master_id"`
	MasterURL       string          `json:"master_url"`
	Title           string          `json:"title"`
	Country         string          `json:"country"`
	Released        string          `json:"released"`
	Notes           string          `json:"notes"`
	Identifiers     []Identifier    `json:"identifiers"`
	Videos          []Video         `json:"videos"`
	Genres          []string        `json:"genres"`
	Styles          []string        `json:"styles"`
	TrackList       []Track         `json:"tracklist"`
	ExtraArtists    []ReleaseArtist `json:"extraartists"`
	Images          []Image         `json:"images"`
	Thumb           string          `json:"thumb"`
	EstimatedWeight int             `json:"estimated_weight"`
	BlockedFromSale bool            `json:"blocked_from_sale"`
	IsOffensive     bool            `json:"is_offensive"`
	DateAdded       time.Time       `json:"date_added"`
	DateChanged     time.Time       `json:"date_changed"`
}

type Master struct {
	ID                   int             `json:"id"`
	MainRelease          int             `json:"main_release"`
	MostRecentRelease    int             `json:"most_recent_release"`
	ResourceURL          string          `json:"resource_url"`
	URI                  string          `json:"uri"`
	VersionsURL          string          `json:"versions_url"`
	MainReleaseURL       string          `json:"main_release_url"`
	MostRecentReleaseURL string          `json:"most_recent_release_url"`
	NumForSale           int             `json:"num_for_sale"`
	LowestPrice          float32         `json:"lowest_price"`
	Images               []Image         `json:"images"`
	Genres               []string        `json:"genres"`
	Styles               []string        `json:"styles"`
	Year                 int             `json:"year"`
	TrackList            []Track         `json:"tracklist"`
	Artists              []ReleaseArtist `json:"artists"`
	Title                string          `json:"title"`
	DataQuality          string          `json:"data_quality"`
	Videos               []Video         `json:"videos"`
	DateAdded            time.Time       `json:"date_added"`
	DateChanged          time.Time       `json:"date_changed"`
}

type Artist struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	ResourceURL    string   `json:"resource_url"`
	ThumbnailURL   string   `json:"thumbnail_url"`
	URI            string   `json:"uri"`
	ReleaseURL     string   `json:"release_url"`
	Images         []Image  `json:"images"`
	Profile        string   `json:"profile"`
	URLs           []string `json:"urls"`
	NameVariations []string `json:"name_variations"`
	Members        []Member `json:"members"`
	DataQuality    string   `json:"data_quality"`
}

type Member struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ResourceURL string `json:"resource_url"`
	Active      bool   `json:"active"`
}

type ReleaseArtist struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Anv          string `json:"anv"`
	Join         string `json:"join"`
	Role         string `json:"role"`
	Tracks       string `json:"tracks"`
	ResourceURL  string `json:"resource_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}

type Company struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CatNo          string `json:"catno"`
	EntityType     string `json:"entity_type"`
	EntityTypeName string `json:"entity_type_name"`
	ResourceURL    string `json:"resource_url"`
	ThumbnailURL   string `json:"thumbnail_url"`
}

type Format struct {
	Name         string   `json:"name"`
	Qty          string   `json:"qty"`
	Descriptions []string `json:"descriptions"`
	Text         string   `json:"text"`
}

type Community struct {
	Have   int    `json:"have"`
	Want   int    `json:"want"`
	Rating Rating `json:"rating"`
}

type Rating struct {
	Count   int     `json:"count"`
	Average float32 `json:"average"`
}

type Identifier struct {
	Type        string `json:"type"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type Video struct {
	URI         string `json:"uri"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Embed       bool   `json:"embed"`
}

type Track struct {
	Position     string          `json:"position"`
	Type         string          `json:"type_"`
	Title        string          `json:"title"`
	Duration     string          `json:"duration"`
	Artists      []Artist        `json:"artists"`
	ExtraArtists []ReleaseArtist `json:"extraartists"`
}

func (t Track) DiscTrackNumber() (int, int, error) {
	if t.Type != "track" {
		return 0, 0, fmt.Errorf("invalid track type: %s", t.Type)
	}
	return ParseDiscTrackNumber(t.Position)
}

var letterRegex = regexp.MustCompile(`^[a-zA-Z]$`)
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
		disc := 1
		track := 1 + int(strings.ToUpper(matches[1])[0]-'A')
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

type Image struct {
	Type        string `json:"type"`
	URI         string `json:"uri"`
	ResourceURL string `json:"resource_url"`
	URI150      string `json:"uri150"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type Series struct {
	Name           string `json:"name"`
	ID             int    `json:"id"`
	CatNo          string `json:"catno"`
	EntityTypeName string `json:"entity_type_name"`
	EntityType     string `json:"entity_type"`
	ResourceURL    string `json:"resource_url"`
}

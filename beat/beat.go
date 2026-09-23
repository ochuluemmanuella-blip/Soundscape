package beat

import (
	"math/rand"
)

type Beat struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	Mood        string `json:"mood"`
	BPM         int    `json:"bpm"`
	Description string `json:"description"`
}

var genres = []string{
	"Trap", "Afrobeats", "Hip-hop", "R&B", "Pop", "Phonk",
}

var moods = []string{
	"Dark", "Nostalgic", "Chill", "Emotional", "Hype",
	"Energetic", "Mysterious",
}

var prefixes = []string{
	"Shadow", "Velvet", "Ocean", "Crystal",
	"Starlight",
}

var suffixes = []string{
	"Chaos", "Soul", "Dreams", "Love", "Nights",
	"Echo", "Vibes",
}

func GenerateBeat() Beat {
	//rand.Seed(time.Now().UnixNano())
	genre := genres[rand.Intn(len(genres))]
	mood := moods[rand.Intn(len(moods))]
	prefix := prefixes[rand.Intn(len(prefixes))]
	suffix := suffixes[rand.Intn(len(suffixes))]

	bpm := 70 + rand.Intn(90)
	name := prefix + " " + suffix

	return Beat{
		Name:  name,
		Genre: genre,
		Mood:  mood,
		BPM:   bpm,
	}

}

package seasons

import (
	"aninext/lib/printer"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type jikanResponse struct {
	Data []Anime `json:"data"`
}

type Anime struct {
	MalID    int     `json:"mal_id"`
	Title    string  `json:"title"`
	Score    float64 `json:"score"`
	Episodes int     `json:"episodes"`
	Status   string  `json:"status"`
	Synopsis string  `json:"synopsis"`
}

const url = "https://api.jikan.moe/v4/seasons/now"

func GetSeasonsNow() []string {
	green := printer.PrintGreen()
	red := printer.PrintRed()
	cyan := printer.PrintCyan()

	fmt.Printf("%s\n", green("[*] Fetching current popular animes..."))

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s\n", red("[!] Failed to MAKE REQUEST to fetch current popular animes!"))
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errorText := fmt.Sprintf("[!] API request /seasons/now failed with status code %d", res.StatusCode)
		fmt.Printf("%s\n", red(errorText))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("%s\n", red("[!] Failed to read response body for the API /seasons/now"))
	}

	var jikanRes jikanResponse

	if err := json.Unmarshal(body, &jikanRes); err != nil {
		fmt.Printf("%s\n", red("[!] Failed to deserialize response from the API /seasons/now"))
	}

	var aniList []string

	fmt.Println("\n--- Airing Anime This Season (Score > 7.0) ---")

	foundCount := 0
	for _, anime := range jikanRes.Data {
		if anime.Score > 7 {
			aniList = append(aniList, anime.Title)
			foundCount++
			goodAnime := fmt.Sprintf("#%d: %s (Score: %.2f)", foundCount, anime.Title, anime.Score)
			fmt.Printf("%s\n", cyan(goodAnime))
		}
	}

	if foundCount == 0 {
		fmt.Println("No good anime currently being aired with score > 7.0")
		return nil
	}

	return aniList
}

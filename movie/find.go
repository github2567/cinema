package movie

func FindMovieName(imdbID string) string {
	switch imdbID {
	case "tt4154796":
		return "Find!!! Avengers: Endgame"
	case "tt1825683":
		return "Find!!! Black Panther"
	}

	return "Find!!! not found."
}

package enums

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Medium Difficulty = "medium"
	Hard   Difficulty = "hard"
)

func ValidateDifficulty(dif string) bool {
	difficulty := map[string]Difficulty{
		string(Easy):   Easy,
		string(Medium): Medium,
		string(Hard):   Hard,
	}

	if _, exists := difficulty[dif]; exists {
		return exists
	}

	return false
}

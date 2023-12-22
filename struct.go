package hangman

type Party struct {
	Try string
	UserName string
	TimeStart  int
	Life          int
	WordsChoice   string
	WordUser      []string
	TabOfLetter   []string
	TabOfWords    []string
	Win           bool
}



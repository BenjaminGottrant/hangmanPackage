package hangman

func IfAlreadyTested(wordsOrLettre string, tab []string) bool {
	if len(tab) >= 1 {
		for i := 0; i < len(tab); i++ {
			if tab[i] == wordsOrLettre {
				return true
			}
		}
	}
	return false
}

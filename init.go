package hangman

func Init ()  *Party {
	wordsRandom:= TackeRandomWord("./hangman/data/words/words.txt")
	wordUser , tabOfLetter := EncryptWord(wordsRandom)
	newparty := Party{
		Life:          0,
		WordsChoice:   wordsRandom,
		WordUser:      wordUser,
		TabOfLetter:   tabOfLetter,
		Win:           false,
	}
	return &newparty
}
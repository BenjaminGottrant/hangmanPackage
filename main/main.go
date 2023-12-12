package main

import (
	"hangman"
	"os"
)

func main() {
	fileOfJose := "data/hangman.txt"
	ascii := true
	if len(os.Args) <= 1 {
		hangman.Help()
		return
	}
	args := os.Args[1:]
	if args[0] == "--startWith" {
		oldPart := hangman.TakeBackParty(args[1])
		hangman.Form(&oldPart)
	} else if args[0] == "--help" {
		hangman.Help()
		return
	} else {
		if len(args) == 1 {
			ascii = false
			args = append(args, " ")
		}
		wordschoice := hangman.TackeRandomWord(args[0])
		wordUser, tabOfLetter := hangman.EncryptWord(wordschoice)
		tabOfPosition := hangman.FetchFile(fileOfJose)
		newparty := hangman.Party{
			Life:          0,
			WordsChoice:   wordschoice,
			WordUser:      wordUser,
			TabOfLetter:   tabOfLetter,
			Win:           false,
		}
		hangman.Form(&newparty)
	}
}

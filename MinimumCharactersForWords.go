package main

func MinimumCharactersForWords(words []string) []string {
	characters := []string{}
	removedCharacter := []string{}
	appendedCharacters := []string{}
	for i := range words {
		for j := range words[i] {
			removed := false
			for x := range characters {
				if string(words[i][j]) == characters[x] {
					removedCharacter = append(removedCharacter, characters[x])
					characters = remove(characters, x)
					removed = true
					break
				}
			}
			if !removed {
				appendedCharacters = append(appendedCharacters, string(words[i][j]))
			}
		}
		for _, t := range removedCharacter {
			characters = append(characters, t)
		}
		for _, t := range appendedCharacters {
			characters = append(characters, t)
		}
		removedCharacter = nil
		appendedCharacters = nil

	}

	return characters
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

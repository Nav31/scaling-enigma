package codewars_probs

func DNAStrand(dna string) string {
  retStr := ""
  dnaMap := map[string]string{ "A": "T", "T": "A", "G": "C", "C": "G", }
  for _, letter := range dna {
  	retStr += dnaMap[string(letter)]
  }
  return retStr
}


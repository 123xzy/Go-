package intset

import(
	"fmt"
	"bytes"
)

type Intset struct{
	words []uint64
}

func (s *Intset) Has(x int) bool{
	word,bit := x/64,uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *Intset) Add(x int){
	word,bit := x/64,uint(x%64)
	for word >= len(s.words){
		s.words = append(s.words,0)
	}
	s.words[word] |= 1 << bit
}

func (s *Intset) UnionWith(t *Intset){
	for i,tword := range t.words{
		if i < len(s.words){
			s.words[i] |= tword
		}else{
			s.words = append(s.words,tword)
		}
	}
}

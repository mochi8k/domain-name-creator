package thesaurus

type Thesaurus interface {
  Synonsyms(term string) ([]string, error)
}

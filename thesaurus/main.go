package main

import (
  "encoding/json"
  "errors"
  "net/http"
)

type BigHuge struct {
  APIKey string
}

type synonyms struct {
  Noun *words `json:"noun"`
  Verb *words `json:"verb"`
}

type words struct {
  Syn []string `json:"syn"`
}

package main

import (
  "fmt"
  "log"
  "net/http"
  //"strings"
  "regexp"
  "sort"
)

func get_words_from(text string) []string{
  words:= regexp.MustCompile("\\w+")
  return words.FindAllString(text, -1)
}

func count_words (words []string) map[string]int{
  word_counts := make(map[string]int)
  for _, word :=range words{
    word_counts[word]++
  }
  return word_counts;
}

func console_out (word_counts map[string]int){
  for word, word_count :=range word_counts{
    fmt.Printf("%v %v\n",word, word_count)
  }
}

func hello(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.Error(w, "404 not found.", http.StatusNotFound)
    return
  }

  switch r.Method {
  case "GET":
    http.ServeFile(w, r, "form.html")
  case "POST":
    // Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
    if err := r.ParseForm(); err != nil {
      fmt.Fprintf(w, "ParseForm() err: %v", err)
      return
    }
    //fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
    name := r.FormValue("name")
    fmt.Fprintf(w, "Input Text are :  %s\n", name)


    // To create an empty map, use the builtin `make`:
    // `make(map[key-type]val-type)`.
    m := make(map[string]int)

    for word, word_count :=range count_words(get_words_from(name)){
      m[word] = word_count
    }

    n := map[int][]string{}
    var a []int
    for k, v := range m {
      n[v] = append(n[v], k)
    }
    for k := range n {
      a = append(a, k)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(a)))
    for _, k := range a {
      for _, s := range n[k] {
        fmt.Printf("%s, %d\n", s, k)
        //fmt.Fprintf(w, "%s, %d\n", s, k)
        fmt.Fprintf(w, "%s, %d\n", s, k)
      }
    }

  default:
    fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
  }
}

func main() {
  http.HandleFunc("/", hello)

  fmt.Printf("Starting server for testing HTTP POST...\n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }
}

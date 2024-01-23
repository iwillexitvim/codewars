package kata
import (
  "strings"
  "strconv"
)

func WordPattern(word string) string {
  var counter = 0
  var chMap = make(map[string]int, len(word))
  var pattern = make([]string, len(word))
  
  for i := range word { 
    
    var chKey = strings.ToLower(string(word[i]))
    number, ok := chMap[chKey]
    
    if (!ok) {
      chMap[chKey] = counter
      number = counter
      counter++
    }
    
    pattern[i] = strconv.Itoa(number)
  }
  
  
  return strings.Join(pattern, ".")
}


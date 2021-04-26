package main
import "fmt"

func tsort(graph map[string][]string, front string) []string {
   todo := []string{front}
   done := make(map[string]bool)
   for {
      back := len(todo) - 1
      done[todo[back]] = true
      for _, do := range graph[todo[back]] {
         if done[do] { continue }
         todo = append(todo, do)
      }
      if len(todo) - 1 > back { continue }
      todo = append(todo[back:], todo[:back]...)
      if todo[0] == front { break }
   }
   return todo
}

func main() {
   m := map[string][]string{
      "A": {"B", "C"},
      "B": {"D", "E"},
      "C": {"main"},
      "D": {"main"},
      "E": {"main"},
   }
   fmt.Println(tsort(m, "A"))
}

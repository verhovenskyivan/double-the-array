package kata
package kata

func Maps(x []int) []int {
  result := make([]int, len(x))
  for i, v := range x {
    result[i] = v + v
  }
  return result
}
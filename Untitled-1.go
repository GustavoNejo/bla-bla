package main
import "fmt"
func main() {

  var n1 float64 = 20
  var n2 float64 = 10
  const R1 int = 10
     fmt.Println( "Qual o valor da multiplicaçao dos valores", n1, "e", n2, "?")
    
    if n1 / n2 == 10{
        fmt.Println (n1/n2)
          fmt.Println ( "esta correto!! A resposta é", R1 ) 

  } else if n1*n2 != 10 {
    fmt.Println( "Esta incsorreto!! \n o valor correto seria", R1 )
  }
   i := 1 
    for i <= R1 {
      fmt.Println(i)
      i = i + 1
    }
  for {
      fmt.Print("loop")
      break
      } 
      for i1 := -1; i1 <= R1 ; i1++ {
        if i1 % 2 == 0 {
          continue
    }
    fmt.Println(i1)
  }
}

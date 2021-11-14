package instructions
import (
	"fmt"
	"time" )

var (
  greetingString   = []byte("Welcome/Bienvenue/歡迎")
  additionalSheets = false
  numQstns, numPgs = 3, 10
  openBook = true
)
func PrintInstructions() {
 fmt.Println(string(greetingString))
 fmt.Print("ExtraSheets:", additionalSheets)
 fmt.Print("The paper is openbook:", openBook)
 fmt.Printf("Pages:%d/Questions:%d\n", numQstns, numPgs)
}
func StartTimer() {
  timr := time.NewTimer(time.Minute * 120) //channel to read
  <-timr.C //Read channel,but ignore the value
  fmt.Println("Stop writing.. Adios/Ciao/再见")
}

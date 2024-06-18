package array
import "testing"
import "fmt"

func TestGenerateMatrix(t *testing.T){
	a := GenerateMatrix(2)
	for i, _ := range a {
		b := a[i]
		fmt.Printf("[")
		for j, _ := range b {
			fmt.Printf("%d", b[j])
			if j + 1 < len(b) {
				fmt.Printf(",")
			}
		}
		fmt.Printf("]\n")
	}
}


func TestSpiralOrder(t *testing.T){
	a := SpiralOrder([][]int{[]int{1,2,3,4}, []int{5,6,7,8}, []int{9, 10, 11, 12}})
	fmt.Printf("%v", a)
}
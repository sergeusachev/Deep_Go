package main
import(
	"fmt"
	"unsafe"
	bitz "deepgo/pkg/lesson_6"
)

func main(){
	fmt.Println("size: ", unsafe.Sizeof(bitz.PlayerMy{}))
}


package main

import (
	"fmt"

	"github.com/vibridi/cacooclip"
)

func main() {
	s, err := cacooclip.Read()
	if err != nil {
		fmt.Println("clipboard err:", err)
		return
	}
	fmt.Println(s)
}

// func main() {
// 	str := `{"target":"shapes","sheetId":"6F6B3","shapes":[{"uid":"2i1FXan7OYW-fkSkLE6zkce","type":3,"bounds":{"top":-381,"bottom":-281,"left":-410,"right":-310},"shapes":[{"type":1,"uid":"b5iwatPq0w8-dv8BrGPnBCQ","bounds":{"top":0,"bottom":100.00000447034836,"left":0,"right":100.00000447034836},"lineInfo":{"enabled":true,"thickness":2,"color":"4A4A4A"},"drawInfo":{"enabled":true,"fillRule":"nonzero","gradientColors":[{"color":"FFFFFF","offset":0},{"color":"000000","offset":100}]},"paths":[{"closed":true,"points":[[100.00000447034836,0],[0,0],[0,100.00000447034836],[100.00000447034836,100.00000447034836],[100.00000447034836,0]]}]},{"text":"CACOOCLIP TEST ABC","leading":2,"halign":1,"valign":1,"uid":"dbwxRmL4pqs-6mo2ZxP8PYI","bounds":{"top":1,"left":1,"right":-2,"bottom":-2,"topFixed":0,"leftFixed":2,"rightFixed":3,"bottomFixed":1},"type":5,"styles":[{"index":0,"font":"Open Sans","size":15,"color":"000000","bold":false,"italic":false,"underline":false,"strikeThrough":false,"backgroundColor":""}],"links":[]}],"connectionPoints":[[50.00000223517418,0],[100.00000447034836,50.00000223517418],[50.00000223517418,100.00000447034836],[0,50.00000223517418],[0,0],[100.00000447034836,0],[100.00000447034836,100.00000447034836],[0,100.00000447034836],[25.00000111758709,0],[75.00000335276127,0],[100.00000447034836,25.00000111758709],[100.00000447034836,75.00000335276127],[25.00000111758709,100.00000447034836],[75.00000335276127,100.00000447034836],[0,25.00000111758709],[0,75.00000335276127]],"attr":[{"name":"stencil-id","value":"250F0"}],"categoryName":"basic"}]}`
// 	err := cacooclip.Write(str)
// 	if err != nil {
// 		fmt.Println("clipboard err:", err)
// 		return
// 	}
// 	fmt.Println("content written successfully")
// }

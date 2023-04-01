package main

import (
	"fmt"
	"strings"
)

var (
	 coin = 50
	 users map[string]int
	 left int
	 names string = "eE, iI ,Oo ,Uu "
)


func main()  {
	fmt.Println(coin,users,left,names)
	left = func() int {
		namesArr := strings.Split(names,",")
		users = make(map[string]int)
		var count int
		for _,name := range namesArr {
			for _,v := range name {
				switch v {
				case 'i','I':
					count = 1
				case 'e','E':
					count = 2
				case 'o','O':
					count = 3
				case 'u','U':
					count = 4
				}
				users[name] += count
				coin -= count
				count = 0
			}
			//if strings.Contains(name,"iI") {
			//	users[name] += 1
			//	coin -= 1
			//}
			//if strings.Contains(name,"eE") {
			//	users[name] += 2
			//	coin -= 2
			//}
			//if strings.Contains(name,"Oo") {
			//	users[name] += 3
			//	coin -= 3
			//}
			//if strings.Contains(name,"Uu") {
			//	users[name] += 4
			//	coin -= 4
			//}
		}
		return coin
	}()

	fmt.Println(users,coin,left)
}

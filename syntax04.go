package main

import (
	"fmt"
	"math/rand"
	"time" //seed 생성용 패키지
)

//소수는 1과 자기자신 이외에는 나누어 떨어지지 않는수
func main() {
	소수 판정 프로그램 ver. 0.2
	 seed := time.Now().Unix()
	 rand.Seed(seed)

	 count := 0
	 number := rand.Intn(150) + 2 //2~151사이의 수
	 fmt.Println("임의로 추출된 수 : ", number)

	 for i := 2; i < number; i++ { //1과 number일때 반복문이 돌지않음
	 	if number%i == 0 {
	 		//count++
	 		count = count + 1
	 	}
	 }
	 if count == 0 {
	 	fmt.Println(number, "는 소수입니다")
	 } else {
	 	fmt.Println(number, "는 소수가 아닙니다")
	 }

	//소수 판정 프로그램 ver. 0.1
	//seed := time.Now().Unix()
	//rand.Seed(seed)

	//count := 0
	//number := rand.Intn(150) + 2 //2~151사이의 수
	//fmt.Println("임의로 추출된 수 : ", number)

	//for i := 1; i <= number; i++ {
	//	if number%i == 0 {
	//		count++
	//	}
	//}
	//if count == 2 {
	//	fmt.Println(number, "는 소수입니다")
	//} else {
	//	fmt.Println(number, "는 소수가 아닙니다")
	//}
}

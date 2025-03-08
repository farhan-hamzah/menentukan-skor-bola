package main

import "fmt"

func hitungSkor(i int, menang *int, kalah *int, draw *int, jumlahGol *int, jumlahBobol *int, selisih *int, skor *int, skorFinal *int) {
	var g, k, c int
	c = i
	*kalah = 0
	*menang = 0 
	*draw = 0
	*jumlahGol = 0
	*jumlahBobol = 0
	*skor =1
	*skorFinal = 0
	for i > 0 {
		fmt.Scan(&g, &k)
		*jumlahGol = *jumlahGol + g
		*jumlahBobol = *jumlahBobol + k
		if g > k {
			*menang++
			*skor *= 3
			*skorFinal += *skor
		}else if g < k{
			*kalah++
		}else{
			*draw++
			*skor *= 1
			*skorFinal += *skor
		}
		i--
		*skor = 1
	}
	*selisih = *jumlahGol - *jumlahBobol
	fmt.Print(c, *menang, *kalah, *draw, *jumlahGol, *jumlahBobol,*selisih, *skorFinal)
}

func main() {
	var permainan int
	fmt.Scan(&permainan)
	var menang, draw, kalah, jumlahGol, jumlahBobol, selisih,skor, skorFinal int
	hitungSkor(permainan, &menang, &kalah, &draw, &jumlahGol, &jumlahBobol, &selisih,&skor, &skorFinal)
}

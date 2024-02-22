package main
import "fmt"

func main() {
	var hp, attack, defense int
	var IV float64

	fmt.Scan(&hp, &attack, &defense)
	IV = ((float64(hp) + float64(attack) + float64(defense)) / 45) * 100
	fmt.Println(IV)

	if (hp >= 0 && hp <= 15) && (attack >= 0 && attack <= 15) && (defense >= 0 && defense <= 15) {
		if IV >= 0 && IV <= 100 {
			if IV >= 0 && IV <= 48.90 {
				fmt.Println("Overall, your pokémon has room for improvement as far as battling goes.", "0 stars")
			} else if IV >= 51.10 && IV <= 65.45 {
				fmt.Println("Overall, your pokémon is pretty decent!", "1 stars")
			} else if IV >= 66.60 && IV <= 80 {
				fmt.Println("Overall your pokémon is really strong!", "2 stars")
			} else if IV >= 82.2 && IV <= 97.8 {
				fmt.Println("Overall, your pokémon looks like it can really battle with the best of them!", "3 stars")
			} else if IV == 100 {
				fmt.Println("Overall, your pokémon looks like it can really battle with the best of them!", "4 stars")
			}
		} else {
			fmt.Println("Not a pokémon")
		}
	} else {
		fmt.Println("Not a pokémon")
	}
}

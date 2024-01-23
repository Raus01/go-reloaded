package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	data, err := ioutil.ReadFile(os.Args[1]) // loen kindlat kohta ja saan data byte
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(string(data), " ") // teen selle data byte stringiks ja lahutan iga stringid tühikuga

	for num, taht := range words { // lähen läbi kõikide stringi sõnade ja kui leian vastava stringi, muudan seda ja võtan algse stringi ära
		if taht == "(up)" {
			words[num-1] = strings.ToUpper(words[num-1])
			words = RemoveIndex(words, num)
		}
	}
	for num, taht := range words {
		if taht == "(cap)" {
			words[num-1] = strings.Title(words[num-1])
			words = RemoveIndex(words, num)
		}
		for num, taht := range words {
			if taht == "(low)" {
				words[num-1] = strings.ToLower(words[num-1])
				words = RemoveIndex(words, num)
			}
		}
		for num, taht := range words {
			if taht == "(hex)" {
				output, err := strconv.ParseInt(words[num-1], 16, 64) // muudan eelmine stringi kümnendarvuks
				words[num-1] = strconv.Itoa(int(output))              // siis mul vaja in64 muuta stringiks tagasi
				words = RemoveIndex(words, num)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		for num, taht := range words {
			if taht == "(bin)" {
				output, err := strconv.ParseInt(words[num-1], 2, 64)
				words[num-1] = strconv.Itoa(int(output))
				words = RemoveIndex(words, num)
				if err != nil {
					fmt.Println(err)
				}

			}
		}
		for num, taht := range words { // siin on mul vaja numbrit mis järgneb sellele
			if taht == "(cap," {
				edasi := strings.Trim(words[num+1], ")") // slice of stringi muudan stringiks ja võtan sealt järgneva numbri ja sulu
				arv, _ := strconv.Atoi(edasi)            // stringi muudan numbriks
				for i := 1; i <= arv; i++ {              // käin arvud läbi
					words[num-i] = strings.Title(words[num-i]) // muudan selle suureks
				}
				words = RemoveIndex(words, num) // võtan ära algse stringi
				words = RemoveIndex(words, num) // võtan ära järgneva stingi
			}
		}
		for num, taht := range words {
			if taht == "(up," {
				edasi := strings.Trim(words[num+1], ")")
				arv, _ := strconv.Atoi(edasi)
				for i := 1; i <= arv; i++ {
					words[num-i] = strings.ToUpper(words[num-i])
				}
				words = RemoveIndex(words, num)
				words = RemoveIndex(words, num)

			}
		}
		for num, taht := range words {
			if taht == "(low," {
				edasi := strings.Trim(words[num+1], ")")
				arv, _ := strconv.Atoi(edasi)
				for i := 1; i <= arv; i++ {
					words[num-i] = strings.ToLower(words[num-i])
				}
				words = RemoveIndex(words, num)
				words = RemoveIndex(words, num)

			}
		}
		for num, taht := range words {
			if taht == "a" || taht == "A" { // otsin üles need tähed
				switch words[num+1][0] { // vahetan nad korral kui järgmise sõna esimene täht on järgnev:
				case 'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H':
					words[num] = words[num] + "n" // sellel korral liidan lihtsalt n-i juurde
				}
			}
		}
		for num, taht := range words {
			if taht == "," { // leian vastava märgi
				words[num-1] += words[num]      // pannen selle märgi eelmise sõna külge
				words = RemoveIndex(words, num) // ja võtan selle märgi ss ära algsest kohast

			}
		}
		for num, taht := range words {
			if taht == ";" {
				words[num-1] += words[num]
				words = RemoveIndex(words, num)

			}
		}
		for num, taht := range words {
			if taht == ":" {
				words[num-1] += words[num]
				words = RemoveIndex(words, num)

			}
		}
		for num, taht := range words {
			if taht == "!" {
				words[num-1] += words[num]
				words = RemoveIndex(words, num)

			}
		}
		for num, taht := range words {
			if taht == "?" {
				words[num-1] += words[num]
				words = RemoveIndex(words, num)

			}
		}
		for num, taht := range words {
			if taht == "!?" {
				words[num-1] += words[num]
				words = RemoveIndex(words, num)

			}
		}
		for num, taht := range words {
			if taht == "." {
				words[num-1] += words[num]
				words = RemoveIndex(words, num)

			}
		}

		for num, taht := range words {
			if strings.HasPrefix(taht, ",") { // leain hasprefix-iga sõna mis algab kindla märgiga
				words[num] = strings.TrimPrefix(taht, ",") // trimprefix-iga võtan selle märgi ära
				words[num-1] = words[num-1] + ","          // liidan eelnevale sõnale selle märgi juurde
			}
		}
		for num, taht := range words {
			if strings.HasPrefix(taht, ".") {
				words[num] = strings.TrimPrefix(taht, ".")
				words[num-1] = words[num-1] + "."
			}
		}
		for num, taht := range words {
			if strings.HasPrefix(taht, ";") {
				words[num] = strings.TrimPrefix(taht, ";")
				words[num-1] = words[num-1] + ";"
			}
		}
		for num, taht := range words {
			if strings.HasPrefix(taht, ":") {
				words[num] = strings.TrimPrefix(taht, ":")
				words[num-1] = words[num-1] + ":"
			}
		}
		for num, taht := range words {
			if strings.HasPrefix(taht, "!") {
				words[num] = strings.TrimPrefix(taht, "!")
				words[num-1] = words[num-1] + "!"
			}
		}
		for num, taht := range words {
			if strings.HasPrefix(taht, "?") {
				words[num] = strings.TrimPrefix(taht, "?")
				words[num-1] = words[num-1] + "?"
			}
		}
		for num, taht := range words {
			if strings.HasPrefix(taht, "!?") {
				words[num] = strings.TrimPrefix(taht, "!?")
				words[num-1] = words[num-1] + "!?"
			}
		}

		kokku := strings.Join(words, " ") // võtan kõik kokku tühikuga, et saaks kirjutada faili

		re := regexp.MustCompile("' .+ '")                 // on vaja selleks stringi aga mul oli string of slice aga join stringiga sain selle stringiks
		for _, taht := range re.FindAllString(kokku, -1) { // otsin kõik stringid üles ja valin palju ma teen seda, -1 võtab kõik
			vahetus1 := strings.Replace(taht, "' ", "'", -1) // võtan tühiku vahelt ära
			vahetus2 := strings.Replace(vahetus1, " '", "'", -1)
			kokku = strings.Replace(kokku, taht, vahetus2, -1) // tagastas vahetades algse stringi uuega

		}
		re = regexp.MustCompile("‘ .+ ‘")
		for _, taht := range re.FindAllString(kokku, -1) {
			vahetus1 := strings.Replace(taht, "‘ ", "‘", -1)
			vahetus2 := strings.Replace(vahetus1, " ‘", "‘", -1)
			kokku = strings.Replace(kokku, taht, vahetus2, -1)

		}

		err2 := ioutil.WriteFile(os.Args[2], []byte(kokku), 0644) // kirjutan vaili, kus mul vaja see sing võtta slice of byte-ina
		if err2 != nil {
			log.Fatal(err)
		}

	}
}

func RemoveIndex(num []string, taht int) []string {
	return append(num[:taht], num[taht+1:]...)
}

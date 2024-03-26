package main

// func genData() {
// 	f, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
//
// 	// f.Write([]byte(fmt.Sprintf("%d %d\n", lines, lenOnLine)))
// 	m := 0
// 	for i := 0; i < lines; i++ {
// 		s := make([]string, 0, lenOnLine)
// 		prev := rand.Intn(1e2)
// 		for i := 0; i < lenOnLine; i++ {
// 			// rn := rand.Intn(1e3)
// 			// s = append(s, strconv.Itoa(rn))
// 			s = append(s, strconv.Itoa(prev+i))
// 			prev = prev + rand.Intn(8e3)
// 			// prev = prev + rand.Intn(4e4)
// 			if prev > m {
// 				m = prev
// 			}
// 		}
// 		_, err := f.Write([]byte(strings.Join(s, " ") + "\n"))
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	fmt.Println("max", m)
// }

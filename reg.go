package main

import (
	"fmt" //For work with registry we need get this package
	//go get golang.org/x/sys/windows/registry - daownload this packege
)

func main() {
	gtk.Init(nil)
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)

	fmt.Println("Создать расширенеи .zzz? (y - yes, n - no)")
	var answer string
	fmt.Scan(&answer)
	if answer == "y" {
		registr()
	} else {
		if answer == "n" {
			fmt.Print("Досвидание")
			fmt.Scan(&answer)
		} else {
			fmt.Println("Неверно ввод, попробуй снова")
			main()
		}
	}
}

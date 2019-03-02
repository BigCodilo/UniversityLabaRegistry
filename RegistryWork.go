package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"
)

func registr() {
	key, _, err := registry.CreateKey(registry.CLASSES_ROOT, ".zzz", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("CLASSES_ROOT/.zzz - Wasn't ctreated")
	} else {
		fmt.Println("CLASSES_ROOT/.zzz - Was created")
	}
	err = key.SetStringValue("", "ZZZ-File")
	if err != nil {
		fmt.Println("CLASSES_ROOT/.zzz Default wasn't set as ZZZ-File")
	} else {
		fmt.Println("CLASSES_ROOT/.zzz Default set as ZZZ-File")
	}
	key.Close()

	key, _, err = registry.CreateKey(registry.CLASSES_ROOT, "ZZZ-File", registry.ALL_ACCESS)
	key.SetStringValue("", "My expansion")
	if err != nil {
		fmt.Println("CLASSES_ROOT/ZZZ-File - Wasn't ctreated")
	} else {
		fmt.Println("CLASSES_ROOT/ZZZ-File - Was created")
	}
	key.Close()

	_, _, err = registry.CreateKey(registry.CLASSES_ROOT, "ZZZ-File\\shell", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell - Wasn't ctreated")
	} else {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell - Was created")
	}

	_, _, err = registry.CreateKey(registry.CLASSES_ROOT, "ZZZ-File\\shell\\Open", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell/Open - Wasn't ctreated")
	} else {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell/Open - Was created")
	}

	key, _, err = registry.CreateKey(registry.CLASSES_ROOT, "ZZZ-File\\shell\\Open\\command", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell/Open/command - Wasn't ctreated")
	} else {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell/Open/command - Was created")
	}
	err = key.SetStringValue("", "C:\\windows\\notepad.exe")
	if err != nil {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell/Open/command Default wasn't set as notepad")
	} else {
		fmt.Println("CLASSES_ROOT/ZZZ-File/shell/Open/command Default set as notepad")
	}
	key.Close()

	DoyouWantToCreateTheFile()
}

func DoyouWantToCreateTheFile() {
	var answer string
	fmt.Println("Создать файл расшиерния .zzz? (y - yes, n - no)")
	fmt.Scan(&answer)
	if answer == "y" {
		CreateFileZzz()
	} else {
		if answer == "n" {
			fmt.Print("Досвиданья")
		} else {
			fmt.Print("Неверный ввод, попробюуйт заного")
			DoyouWantToCreateTheFile()
		}
	}
}

func CreateFileZzz() {
	var stringForWait string
	_, err := os.Create("MyFile.zzz")
	if err != nil {
		fmt.Println("MyFile.zzz wasn't created")
	} else {
		fmt.Println("MyFile.zzz was created")
	}
	fmt.Scan(&stringForWait)
}

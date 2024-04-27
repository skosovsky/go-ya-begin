package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// получаем временную директорию
	tempDir := os.TempDir()
	dir := filepath.Join(tempDir, "margo")
	// создаём директорию
	// не забываем про ведущий ноль для прав доступа
	if err := os.Mkdir(dir, 0755); err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	filename := filepath.Join(dir, "readme.txt")
	// создаём файл
	// остальные пользователи не имеют никаких прав доступа
	if err := os.WriteFile(filename, []byte("привет"), 0660); err != nil {
		log.Fatal(err)
	}

	// Выводим идентификаторы текущего пользователя и группы
	fmt.Printf("UID: %d, GID: %d\n", os.Getuid(), os.Getgid())
	fmt.Println("Файл:", filename)

	// получаем информацию о файле
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	// выводим права доступа в числовом виде
	fmt.Printf("Права доступа %s: 0%o\n", fi.Name(), fi.Mode().Perm())
	// выводим значения расширенных прав
	fmt.Printf("SUID: %d, SGID: %d, StickyBit: %d\n", fi.Mode()&os.ModeSetuid,
		fi.Mode()&os.ModeSetgid, fi.Mode()&os.ModeSticky)
}

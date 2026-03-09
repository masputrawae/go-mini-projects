package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
)

var num = rand.Intn(10)

func main() {
	count := 1

	fmt.Println("Anda diberi kesempatan 5 kali")
	confirm("Tekan enter untuk memulai")
	for range 5 {
		cleanScreen()
		fmt.Println("Kesempatan ke:", count)

		var choice int
		fmt.Print("Pilih dari angka 0 - 10: ")
		fmt.Scan(&choice)

		if choice == num {
			fmt.Println("Selamat tebakan benar")
			break
		} else if choice < num {
			fmt.Println("Terlalu rendah")
			count++
		} else if choice > num {
			fmt.Println("Terlalu tinggi")
			count++
		}

		confirm("Tekan enter untuk melankutkan")
	}

	if count == 6 {
		fmt.Println("Jawaban yang benar adalah: ", num)
	}
}

func cleanScreen() {
	var cmd *exec.Cmd

	// Cek sistem operasi
	if runtime.GOOS == "windows" {
		// Untuk Windows
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		// Untuk Linux/macOS
		cmd = exec.Command("clear")
	}

	// Mengarahkan output dari perintah ke terminal saat ini.
	cmd.Stdout = os.Stdout

	// Jalankan cmd
	cmd.Run()
}

func confirm(message string) {
	// Buat opsi enter agar terminal tidak langsung di hapus
	fmt.Println(message)
	render := bufio.NewReader(os.Stdin)
	_, err := render.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
}

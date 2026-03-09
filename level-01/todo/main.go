package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Data todos
var todos []string

// Main function
func main() {
	// Opsi
	opts := `
========================
[1]. Tambah Tugas Baru
[2]. Edit Tugas
[3]. Hapus Tugas
[4]. Lihat Tugas

[0]. Keluar
========================
Pilih Opsi: `

	for {
		// Panggil fungsi untuk membersihkan terminal
		cleanScreen()

		// Ambil opsi yang dipilih
		var opt int
		fmt.Print(opts)
		fmt.Scan(&opt)

		// Operasi
		switch opt {

		// Hentikan perulangan jika memilih opsi 0 (keluar)
		case 0:
			fmt.Println("Goodbye...")
			return

		// Tugas baru
		case 1:
			task, err := inputTask()
			if err != nil {
				fmt.Println(err)
				return
			}

			todos = append(todos, task)

			// Tampilkan tugas
			showTodos(todos)

		// Edit Tugas
		case 2:
			// Ambil id
			var id int
			showTodos(todos)
			fmt.Print("Pilih ID: ")
			fmt.Scan(&id)
			index := id - 1

			// Cek apakah id tersedia
			if checkID(index, todos) {
				task, err := inputTask()
				if err != nil {
					fmt.Println(err)
					return
				}

				todos[index] = task
				fmt.Println("Berhasil diubah menjadi: ", task)
			} else {
				fmt.Printf("ID %d tidak ditemukan\n", id)
			}

		// Hapus Tugas
		case 3:
			// Ambil id
			var id int
			showTodos(todos)
			fmt.Print("Pilih ID: ")
			fmt.Scan(&id)

			index := id - 1

			// Cek apakah id tersedia
			if checkID(index, todos) {
				todos = append(todos[:index], todos[index+1:]...)
				fmt.Printf("ID %d berhasil di hapus\n", id)
			} else {
				fmt.Printf("ID %d tidak ditemukan\n", id)
			}

		// Default / 4 tampilkan semua tugas
		default:
			showTodos(todos)
		}

		// Buat opsi enter agar terminal tidak langsung di hapus
		fmt.Println("Tekan enter untuk melanjutkan...")
		render := bufio.NewReader(os.Stdin)
		_, err := render.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// Fungsi untuk menampilkan tugas
func showTodos(data []string) {
	if len(data) == 0 {
		fmt.Println("Tugas kosong atau tidak ditemukan")
		return
	}
	fmt.Println("==============================")
	for i, v := range data {
		fmt.Printf("[%d]: %s\n", i+1, v)
	}
	fmt.Println("==============================")
}

// Pengecekan id
func checkID(id int, data []string) bool {
	return id >= 0 && id < len(data)
}

// Fungsi untuk ambil input tugas
func inputTask() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var task string

	fmt.Print("Task: ")
	if scanner.Scan() {
		task = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return task, err
	}
	return task, nil
}

// Bersihkan terminal
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

package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"perang-kode/entity"
	"regexp"
	"runtime"
	"syscall"

	"golang.org/x/term"
)

var scanner = bufio.NewScanner(os.Stdin)

func DisplayMainMenu() {
	fmt.Printf("\nSelamat datang di Perang Kode CLI\n")
	fmt.Println("Menu:")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Print("Masukkan pilihan (1/2/3): ")
}

func MenuRegister(db *sql.DB) {
	emailRegex, _ := regexp.Compile(`^[\w-\.]+@(?:[\w-]+\.)+[\w-]{2,4}$`)
	birthRegex, _ := regexp.Compile(`^\d{4}\-(?:0[1-9]|1[012])\-(?:0[1-9]|[12][0-9]|3[01])$`)

	fmt.Printf("\nREGISTER\n")
	var name string
	for {
		fmt.Print("Masukkan nama\t\t: ")
		scanner.Scan()
		name = scanner.Text()

		if len(name) == 0 {
			fmt.Println("Input tidak boleh kosong")
			continue
		}
		break
	}

	var email string
	for {
		fmt.Print("Masukkan email\t\t: ")
		scanner.Scan()
		email = scanner.Text()

		if !emailRegex.MatchString(email) {
			fmt.Println("Format email tidak sesuai!")
			continue
		}
		break
	}

	var bytePassword []byte
	for {
		fmt.Print("Masukkan password\t: 🔒")
		bytePassword, _ = term.ReadPassword(int(syscall.Stdin))
		fmt.Println()

		if len(bytePassword) == 0 {
			fmt.Println("Input tidak boleh kosong")
			continue
		}
		break
	}

	var birth string
	for {
		fmt.Print("Tanggal lahir (YYYY-MM-DD): ")
		scanner.Scan()
		birth = scanner.Text()

		if !birthRegex.MatchString(birth) {
			fmt.Println("Format tanggal lahir tidak sesuai!")
			continue
		}
		break
	}

	var admin bool
	for {
		fmt.Print("Admin privilege? (y/n)\t: ")
		scanner.Scan()
		adminInput := scanner.Text()

		if adminInput == "y" {
			admin = true
			break
		} else if adminInput == "n" {
			admin = false
			break
		}
		fmt.Println("Input tidak valid!")
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Birth:    birth,
		Password: bytePassword,
		Admin:    admin,
	}

	if err := Register(user, db); err != nil {
		log.Fatal(err)
	} else {
		ClearTerminal()
		fmt.Printf("\nRegistrasi berhasil!\n")
	}
}

func MenuLogin() (string, []byte) {
	emailRegex, _ := regexp.Compile(`^[\w-\.]+@(?:[\w-]+\.)+[\w-]{2,4}$`)
	fmt.Printf("\nLOGIN\n")

	var email string
	for {
		fmt.Print("Email\t\t: ")
		scanner.Scan()
		email = scanner.Text()

		if !emailRegex.MatchString(email) {
			fmt.Println("Input email tidak valid!")
			continue
		}
		break
	}

	var bytePassword []byte
	for {
		fmt.Print("Password\t: 🔒")
		bytePassword, _ = term.ReadPassword(int(syscall.Stdin))
		fmt.Println()

		if len(bytePassword) == 0 {
			fmt.Println("Input tidak boleh kosong")
			continue
		}
		break
	}

	return email, bytePassword
}

func UserMenu(user entity.User, db *sql.DB) {
	for {
		var input int
		fmt.Printf("\nSelamat datang %v!\n", user.Name)
		fmt.Println("1. Tampilkan semua game")
		fmt.Println("2. Tampilkan cart")
		fmt.Println("3. Tambah game ke cart")
		fmt.Println("4. Hapus game dari cart")
		fmt.Println("5. Get Voucher")
		fmt.Println("6. Log Out")
		fmt.Print("Masukkan pilihan sub-menu (1/2/3/4/5/6): ")
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &input); err != nil {
			ClearTerminal()
			fmt.Printf("\nInput harus berupa angka!\n")
			continue
		}
		ClearTerminal()

		switch input {
		case 1:
			ShowAllGames(db)

		case 2:
			ShowCart(user, db)

		case 3:
			ShowAllGames(db)
			gameID := GameIdInput()
			if err := AddGameToCart(user, gameID, db); err != nil {
				fmt.Printf("\n%v\n", err)
			}

		case 4:
			ShowCart(user, db)
			gameID := GameIdInput()
			if err := RemoveGameFromCart(user, gameID, db); err != nil {
				fmt.Printf("\n%v\n", err)
			}

		case 5:
			if err := GetVoucher(&user, db); err != nil {
				log.Fatal(err)
			}

		case 6:
			fmt.Printf("\nLogging out...\n")
			return

		default:
			fmt.Printf("\nInput di luar range 1-6\n")
		}
	}
}

func GameIdInput() int {
	var gameID int
	for {
		fmt.Printf("\nMasukkan Game ID: ")
		_, err := fmt.Scanln(&gameID)
		if err != nil {
			fmt.Println("Terjadi kesalahan saat memproses masukan:", err)
			continue
		}
		break
	}
	return gameID
}

func AdminMenu(user entity.User, db *sql.DB) {
	for {
		var input int
		fmt.Printf("\nSelamat datang %v!\n", user.Name)
		fmt.Println("1. Update stock game")
		fmt.Println("2. Tampilkan report user")
		fmt.Println("3. Tampilkan report order")
		fmt.Println("4. Tampilkan report stock")
		fmt.Println("5. Log Out")
		fmt.Print("Masukkan pilihan sub-menu admin (1/2/3/4/5): ")
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &input); err != nil {
			fmt.Printf("\nInput harus berupa angka!\n")
			continue
		}
		ClearTerminal()

		switch input {
		case 1:
			for {
				DisplayStock(db)
				var input entity.Stock
				for {
					fmt.Printf("\nMasukkan ID game: ")
					scanner.Scan()
					if _, err := fmt.Sscanf(scanner.Text(), "%d", &input.Id); err != nil {
						fmt.Println("Input harus berupa angka")
						continue
					}
					break
				}
				for {
					fmt.Printf("Masukkan stock game: ")
					scanner.Scan()
					if _, err := fmt.Sscanf(scanner.Text(), "%d", &input.Stock); err != nil {
						fmt.Println("Input harus berupa angka")
						continue
					}
					break
				}
	
				if err := UpdateStock(db, input); err != nil {
					ClearTerminal()
					fmt.Println(err)
					continue
				}
				break
			}


		case 2:
			if err := UserReport(db); err != nil {
				log.Fatal(err)
			}

		case 3:
			if err := OrderReport(db); err != nil {
				log.Fatal(err)
			}

		case 4:
			if err := DisplayStock(db); err != nil {
				log.Fatal(err)
			}

		case 5:
			fmt.Printf("\nLogging out...\n")
			return

		default:
			fmt.Printf("\nInput di luar range 1-5\n")
		}
	}
}

func RunCmd(name string, arg ...string) {
    cmd := exec.Command(name, arg...)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func ClearTerminal() {
    switch runtime.GOOS {
    case "darwin":
        RunCmd("clear")
    case "linux":
        RunCmd("clear")
    case "windows":
        RunCmd("cmd", "/c", "cls")
    default:
        RunCmd("clear")
    }
}
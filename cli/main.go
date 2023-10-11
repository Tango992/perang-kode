package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"perang-kode/config"
	"perang-kode/entity"
	"perang-kode/handler"
	"regexp"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/term"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		var option int
		displayMainMenu()

		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d", &option)
		if err != nil {
			fmt.Printf("\nInput anda bukan merupakan integer\n\n")
			continue
		}

		switch option {
		case 1:
			menuRegister(db)

		case 2:
			email, password := menuLogin()
			user, authenticated, _ := handler.Login(email, password, db)

			if authenticated {
				if user.Admin {
					adminMenu(user)
				} else {
					userMenu(user)
				}
			} else {
				log.Fatal("Username / Password tidak sesuai")
			}

		case 3:
			fmt.Printf("\nSampai jumpa!\n")
			os.Exit(0)

		default:
			fmt.Printf("\nInput harus merupakan angka 1-3!\n\n")
		}
	}
}


func displayMainMenu() {
	fmt.Printf("\nSelamat datang di Perang Kode CLI\n")
	fmt.Println("Menu:")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Print("Masukkan pilihan (1/2/3): ")
}

func menuRegister(db *sql.DB) {
	emailRegex, _ := regexp.Compile(`^[\w-\.]+@(?:[\w-]+\.)+[\w-]{2,4}$`)
	birthRegex, _ := regexp.Compile(`^\d{4}\-(?:0[1-9]|1[012])\-(?:0[1-9]|[12][0-9]|3[01])$`)

	fmt.Printf("\nREGISTER\n")
	fmt.Print("Masukkan nama: ")
	scanner.Scan()
	name := scanner.Text()

	var email string
	for {
		fmt.Print("Masukkan email: ")
		scanner.Scan()
		email = scanner.Text()

		if !emailRegex.MatchString(email) {
			fmt.Println("Format email tidak sesuai!")
			continue
		}
		break
	}

	fmt.Print("Masukkan password: ")
	bytePassword, _ := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	
	var birth string
	for {
		fmt.Print("Masukkan tanggal lahir (YYYY-MM-DD): ")
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
		fmt.Print("Register sebagai admin? (y/n): ")
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
		Name: name,
		Email: email,
		Birth: birth,
		Password: bytePassword,
		Admin: admin,
	}

	if err := handler.Register(user, db); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("\nRegistrasi berhasil!\n")
	}
}


func menuLogin() (string, []byte) {
	fmt.Printf("\nLOGIN\n")
	fmt.Print("Email: ")
	scanner.Scan()
	email := scanner.Text()

	fmt.Print("Password: ")
	bytePassword, _ := term.ReadPassword(syscall.Stdin)
	fmt.Println()
	return email, bytePassword
}

func userMenu(user entity.User) {
	for {
		var input int
		fmt.Printf("\nSelamat datang %v!\n", user.Name)
		fmt.Println("1. Tampilkan semua game")
		fmt.Println("2. Tampilkan cart")
		fmt.Println("3. Tambah game ke cart")
		fmt.Println("4. Hapus game dari cart")
		fmt.Println("5. Log Out")
		fmt.Print("Masukkan pilihan sub-menu (1/2/3/4/5): ")
		scanner.Scan()
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &input); err != nil {
			fmt.Printf("\nInput harus berupa angka!\n")
			continue
		}

		switch input {
		case 1:
			fmt.Println("menu 1")
		case 2:
			fmt.Println("menu 2")
		case 3:
			fmt.Println("menu 3")
		case 4:
			fmt.Println("menu 4")
		case 5:
			fmt.Printf("\nLogging out...\n")
			return
		default:
			fmt.Printf("\nInput di luar range 1-5\n")
		}
	}
}

func adminMenu(user entity.User) {
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

		switch input {
		case 1:
			fmt.Println("menu 1")
		case 2:
			fmt.Println("menu 2")
		case 3:
			fmt.Println("menu 3")
		case 4:
			fmt.Println("menu 4")
		case 5:
			fmt.Printf("\nLogging out...\n")
			return
		default:
			fmt.Printf("\nInput di luar range 1-5\n")
		}
	}
}
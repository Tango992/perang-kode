package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/Tango992/perang-kode/config"
	"github.com/Tango992/perang-kode/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		var option int
		handler.DisplayMainMenu()

		scanner.Scan()
		_, err := fmt.Sscanf(scanner.Text(), "%d", &option)
		if err != nil {
			handler.ClearTerminal()
			fmt.Printf("\nInput anda bukan merupakan integer\n")
			continue
		}
		handler.ClearTerminal()

		switch option {
		case 1:
			handler.MenuRegister(db)

		case 2:
			for {
				email, password := handler.MenuLogin()
				user, authenticated, _ := handler.Login(email, password, db)
				
				if authenticated {
					if user.Admin {
						handler.ClearTerminal()
						handler.AdminMenu(user, db)
					} else {
						handler.ClearTerminal()
						handler.UserMenu(user, db)
					}
				} else {
					handler.ClearTerminal()
					fmt.Printf("\nEmail / password tidak sesuai\n")
					continue
				}
				break
			}

		case 3:
			fmt.Printf("\nSampai jumpa!\n")
			os.Exit(0)

		default:
			fmt.Printf("\nInput harus merupakan angka 1-3!\n\n")
		}
	}
}

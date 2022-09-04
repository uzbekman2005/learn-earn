package frontend

import (
	"fmt"
	"learn/config"
)

func MainMune() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 			 Choose one                 -+")
	fmt.Println("+-		1. Your Profile                     -+")
	fmt.Println("+-		2. HighScore                        -+")
	fmt.Println("+-		3. Update Profile                   -+")
	fmt.Println("+-		4. Play Game                        -+")
	fmt.Println("+-		5. Exit                             -+")
	fmt.Println("+----------------------------------------------------+")
}

func SignUpMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- To sign up succesfully enter all info asked      -+")
	fmt.Println("+----------------------------------------------------+")
}

func LogInMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- To log in succesfully enter all info asked       -+")
	fmt.Println("+----------------------------------------------------+")
}

func RegistrationMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	            Registration Menu               -+")
	fmt.Println("+- 	         1. Sign up                         -+")
	fmt.Println("+- 	         2. Log in                          -+")
	fmt.Println("+- 	         3. Exit                            -+")
	fmt.Println("+----------------------------------------------------+")
}

func GamesMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	                 Games list                  -+")
	fmt.Println("+- 	          1. Polychudis                      -+")
	fmt.Println("+- 	          2. Tic Tai Toi                     -+")
	fmt.Println("+- 	          3. Exit                            -+")
	fmt.Println("+----------------------------------------------------+")
}

func SuccesLogin() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Printf("+- %-48s -+\n", "You have successfully loged in")
	fmt.Println("+----------------------------------------------------+")
}

func SuccesSignUp() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Printf("+-  %-46s  -+\n", "You have successfully signed up")
	fmt.Println("+----------------------------------------------------+")
}

func UserProfile() {
	fmt.Println("+-------------+--------------------------------------+")
	fmt.Printf("|  Full Name  |  %-35s |\n", fmt.Sprintf("%s %s", config.CurrentUser.Last_name, config.CurrentUser.First_name))
	fmt.Println("+-------------+--------------------------------------+")
	fmt.Printf("|  Username   |  %-35s |\n", config.CurrentUser.Username)
	fmt.Println("+-------------+--------------------------------------+")
	fmt.Printf("|  Country    |  %-35s |\n", config.CurrentUser.Country)
	fmt.Println("+-------------+--------------------------------------+")
	fmt.Printf("|  Password   |  %-35s |\n", config.CurrentUser.Password)
	fmt.Println("+-------------+--------------------------------------+")
	fmt.Printf("|  Birth date |  %-35s |\n", config.CurrentUser.DateOfBirth)
	fmt.Println("+-------------+--------------------------------------+")
	fmt.Printf("|  Score      |  %-35d |\n", config.CurrentUser.Score)
	fmt.Println("+-------------+--------------------------------------+")
	fmt.Printf("|  High Score |  %-35d |\n", config.CurrentUser.HighScore)
	fmt.Println("+--------------+-------------------------------------+")
	fmt.Println()
}

func UpdateMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	                 Update page                 -+")
	fmt.Println("+- 	          1. First name                      -+")
	fmt.Println("+- 	          2. Last name                       -+")
	fmt.Println("+- 	          3. Country                         -+")
	fmt.Println("+- 	          4. Date of birth                   -+")
	fmt.Println("+- 	          5. Password                        -+")
	fmt.Println("+- 	          6. Username                        -+")
	fmt.Println("+- 	          7. Exit                            -+")
	fmt.Println("+----------------------------------------------------+")
}

func UsernameRequirement() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	        Username requirements                -+")
	fmt.Println("+- 1. Uniqueness                                    -+")
	fmt.Println("+- 2.	        Allowed characters                   -+")
	fmt.Println("+-   Letters, arabic numbers, underscore('_')       -+")
	fmt.Println("+----------------------------------------------------+")
}

func WasRegistered() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	 Did you really register before?             -+")
	fmt.Println("+- 1. Yes                                           -+")
	fmt.Println("+- 2. No I will register now                        -+")
	fmt.Println("+----------------------------------------------------+")
}

func DateFormat() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- Date format is yyyy-mm-dd                        -+")
	fmt.Println("+----------------------------------------------------+")
}

func PasswordRequirement() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	        Password requirements               -+")
	fmt.Println("+- 1. at least 5 in length                          -+")
	fmt.Println("+- 2. comma (',') and spaces are not allowed        -+")
	fmt.Println("+----------------------------------------------------+")
}

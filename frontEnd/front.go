package frontend

import "fmt"

func MainMune() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 				 Choose one                      -+")
	fmt.Println("+- 		1. You Profile                           -+")
	fmt.Println("+- 		2. HighScore                             -+")
	fmt.Println("+- 		3. Play Game                             -+")
	fmt.Println("+- 		4. Exit                                  -+")
	fmt.Println("+----------------------------------------------------+")
}

func SignUpMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- To sign up succesfully enter all info asked      -+")
	fmt.Println("+----------------------------------------------------+")
}

func LogInMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- To login up succesfully enter all info asked     -+")
	fmt.Println("+----------------------------------------------------+")
}

func RegistrationMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	             Registration Menu               -+")
	fmt.Println("+- 	          1. Sign up                         -+")
	fmt.Println("+- 	          2. Log in                          -+")
	fmt.Println("+- 	          3. Exit                            -+")
	fmt.Println("+----------------------------------------------------+")
}

func GamesMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	                 Games list                  -+")
	fmt.Println("+- 	          1. Polychudis                      -+")
	fmt.Println("+- 	          2. Exit                            -+")
	fmt.Println("+----------------------------------------------------+")
}

func SuccesLogin() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Printf("+- %-48s -+\n","You have successfully loged in")
	fmt.Println("+----------------------------------------------------+")
}

func SuccesSignUp() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Printf("+-  %-46s  -+\n", "You have successfully signed up")
	fmt.Println("+----------------------------------------------------+")
}

func UserProfile() {
	// This part will be wrttien after user struct created
}

func UpdateMenu() {
	fmt.Println("+----------------------------------------------------+")
	fmt.Println("+- 	                 Update page                 -+")
	fmt.Println("+- 	          1. First name                      -+")
	fmt.Println("+- 	          2. Last name                       -+")
	fmt.Println("+- 	          4. Country                         -+")
	fmt.Println("+- 	          5. Date of birth                   -+")
	fmt.Println("+- 	          6. Phone number                    -+")
	fmt.Println("+- 	          7. Exit                            -+")
	fmt.Println("+----------------------------------------------------+")

}

func UsernameRequirement(){
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
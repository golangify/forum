package usercontroller

import (
	"errors"
	"fmt"
	"forum/utils/declension"
)

func (c *userController) validateRegisterForm(registerRequest registerRequest) error {
	symbolDeclension := func(n int) string {
		return declension.Declension(n, "символа", "символов", "символов")
	}

	usernameRune := []rune(registerRequest.Username)
	if len(usernameRune) == 0 {
		return errors.New("имя не может быть пустым")
	}
	if len(usernameRune) < int(c.config.User.MinUsernameLength) {
		return fmt.Errorf("имя должно быть длиннее %d %s", c.config.User.MinUsernameLength, symbolDeclension(int(c.config.User.MinUsernameLength)))
	}
	if len(usernameRune) > int(c.config.User.MaxUsernameLength) {
		return fmt.Errorf("имя не может быть длиннее %d %s", c.config.User.MaxUsernameLength, symbolDeclension(int(c.config.User.MaxUsernameLength)))
	}

	if len(registerRequest.Password) == 0 {
		return errors.New("пароль не может быть пустым")
	}
	if len(registerRequest.Password) < int(c.config.User.MinPasswordLength) {
		return errors.New("слишком короткий пароль")
	}
	if len(registerRequest.Password) > int(c.config.User.MaxPasswordLength) {
		return errors.New("слишком длинный пароль")
	}

	if registerRequest.Username == registerRequest.Password {
		return errors.New("пароль не может быть таким-же как и имя")
	}

	return nil
}

func (c *userController) validateLoginForm(loginRequest loginRequest) error {
	if len(loginRequest.Password) == 0 {
		return errors.New("пароль не может быть пустым")
	}
	if len(loginRequest.Password) > int(c.config.User.MaxPasswordLength) {
		return errors.New("слишком длинный пароль")
	}

	return nil
}

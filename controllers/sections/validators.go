package sectioncontroller

import (
	"errors"
	"fmt"
	"forum/utils/declension"
)

func (c *sectionController) validateNewSectionForm(newSectionRequest newSectionRequest) error {
	symbolDeclension := func(n int) string {
		return declension.Declension(n, "символа", "символов", "символов")
	}

	titleRune := []rune(newSectionRequest.Title)
	if len(titleRune) == 0 {
		return errors.New("название раздела не может быть пустым")
	}
	if len(titleRune) < int(c.config.Section.MinTitleLength) {
		return fmt.Errorf("название раздела должно быть длиннее %d %s", c.config.Section.MinTitleLength, symbolDeclension(int(c.config.Section.MinTitleLength)))
	}
	if len(titleRune) > int(c.config.Section.MaxTitleLength) {
		return fmt.Errorf("название раздела должно быть короче %d %s", c.config.Section.MaxTitleLength, symbolDeclension(int(c.config.Section.MaxTitleLength)))
	}

	bodyRune := []rune(newSectionRequest.Body)
	if len(bodyRune) < int(c.config.Section.MinBodyLength) {
		return fmt.Errorf("содержимое раздела должно быть длиннее %d %s", c.config.Section.MinBodyLength, symbolDeclension(int(c.config.Section.MinBodyLength)))
	}
	if len(bodyRune) > int(c.config.Section.MaxBodyLength) {
		return fmt.Errorf("содержимое раздела не может быть длиннее %d %s", c.config.Section.MaxBodyLength, symbolDeclension(int(c.config.Section.MaxBodyLength)))
	}
	return nil
}

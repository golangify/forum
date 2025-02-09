package topiccontroller

import (
	"errors"
	"fmt"
	"forum/utils/declension"
)

func (c *topicController) validateNewTopicForm(newTopicRequest newTopicRequest) error {
	symbolDeclension := func(n int) string {
		return declension.Declension(n, "символа", "символов", "символов")
	}

	titleRune := []rune(newTopicRequest.Title)
	if len(titleRune) == 0 {
		return errors.New("название темы не может быть пустым")
	}
	if len(titleRune) < int(c.config.Topic.MinTitleLength) {
		return fmt.Errorf("название темы должно быть длиннее %d %s", c.config.Topic.MinTitleLength, symbolDeclension(int(c.config.Topic.MinTitleLength)))
	}
	if len(titleRune) > int(c.config.Topic.MaxTitleLength) {
		return fmt.Errorf("название темы должно быть короче %d %s", c.config.Topic.MaxTitleLength, symbolDeclension(int(c.config.Topic.MaxTitleLength)))
	}

	bodyRune := []rune(newTopicRequest.Body)
	if len(bodyRune) < int(c.config.Topic.MinBodyLength) {
		return fmt.Errorf("содержимое темы должно быть длиннее %d %s", c.config.Topic.MinBodyLength, symbolDeclension(int(c.config.Topic.MinBodyLength)))
	}
	if len(bodyRune) > int(c.config.Topic.MaxBodyLength) {
		return fmt.Errorf("содержимое темы не может быть длиннее %d %s", c.config.Topic.MaxBodyLength, symbolDeclension(int(c.config.Topic.MaxBodyLength)))
	}
	return nil
}

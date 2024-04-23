package command

import (
	"errors"
	"github.com/manifoldco/promptui"
	"time"
)

func PromptString(label string) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			if len(s) == 0 {
				return errors.New("参数不能为空")
			}
			return nil
		},
	}
	return prompt.Run()
}

func PromptDate(label string) (string, error) {
	layout := "2006-01-02"
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(s string) error {
			_, err := time.Parse(layout, s)
			return err
		},
	}
	return prompt.Run()
}

func PromptDateSToE() (string, string, error) {
	s, err := PromptDate("请输入开始时间")
	if err != nil {
		return "", "", err
	}
	e, err2 := PromptDate("请输入结束时间")
	if err2 != nil {
		return "", "", err2
	}
	return s, e, nil
}

func GenerateDateList(startTime, endTime time.Time) []time.Time {
	var list []time.Time

	for t := startTime; t.Before(endTime); t = t.AddDate(0, 0, 1) {
		list = append(list, t)
	}

	return list
}

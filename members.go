package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

var AllChatIDs []int64

func LoadMembersChatIDs() ([]int64, error) {
	var chatIDs []int64
	dat, err := os.ReadFile(MEMBERS_FILE_PATH)
	if err != nil {
		return nil, err
	}

	ids := strings.Split(string(dat), "\n")
	for _, id := range ids {
		if id != "" {
			numberChatID, err := strconv.Atoi(id)
			if err != nil {
				return nil, errors.New("encountered non numeric id in members file")
			}
			chatIDs = append(chatIDs, int64(numberChatID))
		}
	}

	return chatIDs, nil
}

func AddMemberChatID(chatID int64) error {
	for _, ID := range AllChatIDs {
		if chatID == ID {
			return nil
		}
	}

	AllChatIDs = append(AllChatIDs, chatID)
	err := SaveIDsToFile(AllChatIDs)
	if err != nil {
		return err
	}

	return nil
}

func SaveIDsToFile(chatIDs []int64) error {
	var finalChatIDs []string
	for _, ID := range chatIDs {
		finalChatIDs = append(finalChatIDs, strconv.Itoa(int(ID)))
	}

	file, err := os.OpenFile(MEMBERS_FILE_PATH, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return nil
	}

	file.Write([]byte(strings.Join(finalChatIDs, "\n")))

	if err := file.Close(); err != nil {
		return err
	}

	return nil
}

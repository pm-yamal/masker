package main

import (
	"bytes"
	"fmt"
)

func maskLinks(message string) string {
	// Преобразуем строку в байты
	messageToBytes := []byte(message)
	// Создаем новый байтовый срез для буфера
	buffer := make([]byte, 0, len(messageToBytes))
	// Ищем ссылки и маскируем их
	for i := 0; i < len(messageToBytes); i++ {
		// Проверяем, начинается ли с "http://"
		if i+7 <= len(messageToBytes) && bytes.Equal(messageToBytes[i:i+7], []byte("http://")) {
			// Добавляем "http://" в буфер
			buffer = append(buffer, messageToBytes[i:i+7]...)
			i += 6 // Пропускаем "http://"
			// Находим конец ссылки (пробел или конец строки)
			for i < len(messageToBytes) && messageToBytes[i] != ' ' {
				buffer = append(buffer, '*') // Заменяем на звёздочку
				i++
			}
			// В этом месте цикл продолжится, но i уже указывает на следующее значение
			continue
		}
		// Если это не ссылка, то просто добавляем байт
		buffer = append(buffer, messageToBytes[i])
	}
	// Преобразуем буфер обратно в строку и возвращаем
	return string(buffer)
}

func main() {
	message := "Hello, its my page: http://localhost123.com See you"
	maskedMessage := maskLinks(message)
	fmt.Println(maskedMessage)
}
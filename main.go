package main

import (
    "fmt"
)

func maskLinks(message string) string {
    // Создаем новый байтовый срез на основе входного сообщения
    buffer := []byte(message)
    // Ищем начало ссылки
    for i := 0; i < len(buffer); i++ {
        // Проверяем, начинается ли ссылка с "http://"
        if i+7 < len(buffer) && string(buffer[i:i+7]) == "http://" {
            // Найдем конец ссылки
            j := i + 7
            for j < len(buffer) && (buffer[j] != ' ' && buffer[j] != '>') {
                j++
            }
            // Маскируем ссылку символами '*'
            for k := i + 7; k < j; k++ {
                buffer[k] = '*'
            }
            // Перемещаем индекс i в конец найденной ссылки
            i = j - 1
        }
    }
    // Возвращаем результирующее сообщение
    return string(buffer)
}

func main() {
    input := "Hello, its my page: http://localhost123.com See you"
    output := maskLinks(input)
    fmt.Println(output)
}
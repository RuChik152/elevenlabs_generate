package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func runGenerate(text string, nameFile string) {
	url := "https://api.elevenlabs.io/v1/text-to-speech/WuLq5z7nEcrhppO0ZQJw?enable_logging=true&output_format=mp3_22050_32"

	dataRequest := fmt.Sprintf("{\n  \"text\": \"%s\",\n  \"model_id\": \"eleven_multilingual_v2\",\n  \"voice_settings\": {\n    \"similarity_boost\": 0.75,\n    \"style\": 0.25,\n    \"use_speaker_boost\": true,\n    \"stability\": 0.5\n  }\n}", sanitizeInput(text))

	payload := strings.NewReader(dataRequest)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("xi-api-key", "ef1b08266ff0440ea363bad990a0e599")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		fileName := fmt.Sprintf(".\\output\\%s.json", nameFile)
		err = os.WriteFile(fileName, body, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Аудио успешно сохранено в файл:", fileName)
	} else {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		// Сохранение аудио в файл
		//fileName := "output.mp3"
		fileName := fmt.Sprintf(".\\output\\%s.mp3", nameFile)
		err = os.WriteFile(fileName, body, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Аудио успешно сохранено в файл:", fileName)
	}

}

func sanitizeInput(text string) string {
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	re := regexp.MustCompile(`<color[^>]*>|</color>`)
	text = re.ReplaceAllString(text, "")
	fmt.Println("Выходная строка: ", text)
	return text
}
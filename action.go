package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

func runGenerate(text string, nameFile string, l string) {

	var url string

	if lang, err := getLangID(l); err != nil {
		log.Fatal(err)
	} else {
		url = fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s?enable_logging=true&output_format=mp3_22050_32", lang)
	}

	dataRequest := fmt.Sprintf("{\n  \"text\": \"%s\",\n  \"model_id\": \"eleven_multilingual_v2\",\n  \"voice_settings\": {\n    \"similarity_boost\": 0.75,\n    \"style\": 0.25,\n    \"use_speaker_boost\": true,\n    \"stability\": 0.5\n  }\n}", sanitizeInput(text))

	payload := strings.NewReader(dataRequest)

	req, _ := http.NewRequest("POST", url, payload)

	xi_api_key := os.Getenv("XI_API_KEY")

	req.Header.Add("xi-api-key", xi_api_key)
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

		fileName := path.Join(".", "error", l, fmt.Sprintf("%s.json", nameFile))

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

		fileName := path.Join(".", "output", l, fmt.Sprintf("%s.mp3", nameFile))

		err = os.WriteFile(fileName, body, 0644)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Audio successfully saved to file: ", fileName)
	}

}

func sanitizeInput(text string) string {
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	re := regexp.MustCompile(`<color[^>]*>|</color>`)
	text = re.ReplaceAllString(text, "")
	fmt.Println("Output line:", text)
	return text
}

type langID string

func getLangID(lang string) (langID, error) {
	switch lang {
	case "CHN":
		return langID(os.Getenv("SPEAKER_CHN_ID")), nil
	case "ENG":
		return langID(os.Getenv("SPEAKER_ENG_ID")), nil
	case "FR":
		return langID(os.Getenv("SPEAKER_FR_ID")), nil
	case "ESP":
		return langID(os.Getenv("SPEAKER_ESP_ID")), nil
	case "ARAB":
		return langID(os.Getenv("SPEAKER_ARAB_ID")), nil
	default:
		return "", fmt.Errorf("lang not found")
	}
}

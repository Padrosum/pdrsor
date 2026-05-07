package main

import (
        "bytes"
        "encoding/json"
        "fmt"
        "io"
        "net/http"
        "os"
        "os/exec"
        "path/filepath"
        "strings"
)

const (
        clrPrp = "\033[1;35m"
        clrGrn = "\033[1;32m"
        clrYlw = "\033[1;33m"
        clrRst = "\033[0m"
)

type OllamaRequest struct {
        Model  string `json:"model"`
        Prompt string `json:"prompt"`
        System string `json:"system"`
        Stream bool   `json:"stream"`
}

type OllamaResponse struct {
        Response string `json:"response"`
}

func printHeader() {
        fmt.Printf("%s┌──────────────────────────────────────────┐\n", clrPrp)
        fmt.Printf("│ %spdrsor%s - Local AI Terminal Assistant    │\n", clrGrn, clrPrp)
        fmt.Printf("└──────────────────────────────────────────┘\n%s", clrRst)
}

func getModel() string {
        home, _ := os.UserHomeDir()
        configPath := filepath.Join(home, ".config", "pdrsor_rc")

        // Config dosyasını oku
        data, err := os.ReadFile(configPath)
        if err == nil {
                return strings.TrimSpace(string(data))
        }

        // Dosya yoksa sor
        fmt.Printf("%s(!) %sİlk kullanım. Ollama model adını girin (örn: llama3): ", clrYlw, clrRst)
        var modelName string
        fmt.Scanln(&modelName)

        // Kaydet
        os.MkdirAll(filepath.Dir(configPath), 0755)
        os.WriteFile(configPath, []byte(modelName), 0644)
        return modelName
}

func copyToClipboard(text string) {
        var cmd *exec.Cmd
        // X11 veya Wayland kontrolü
        if os.Getenv("WAYLAND_DISPLAY") != "" {
                cmd = exec.Command("wl-copy")
        } else {
                cmd = exec.Command("xclip", "-selection", "clipboard")
        }

        cmd.Stdin = strings.NewReader(text)
        if err := cmd.Run(); err == nil {
                fmt.Printf("%s  [ℹ] Komut panoya kopyalandı.%s\n", clrYlw, clrRst)
        }
}

func main() {
        if len(os.Args) < 2 {
                printHeader()
                fmt.Printf("Kullanım: %spdrsor%s <istek>\n", clrGrn, clrRst)
                return
        }

        model := getModel()
        query := strings.Join(os.Args[1:], " ")

        reqBody := OllamaRequest{
                Model:  model,
                Prompt: query,
                System: "Sen bir Linux adminisin. Sadece ham bash komutu dön. Açıklama yapma.",
                Stream: false,
        }

        jsonData, _ := json.Marshal(reqBody)
        resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
        if err != nil {
                fmt.Printf("Hata: Ollama'ya bağlanılamadı. Servis çalışıyor mu?\n")
                return
        }
        defer resp.Body.Close()

        body, _ := io.ReadAll(resp.Body)
        var ollamaResp OllamaResponse
        json.Unmarshal(body, &ollamaResp)

        fmt.Printf("\n%s➜ Önerilen Komut:%s\n", clrYlw, clrRst)
        fmt.Printf("%s   %s%s\n\n", clrGrn, ollamaResp.Response, clrRst)
        copyToClipboard(ollamaResp.Response)
}

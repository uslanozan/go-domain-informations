# **Go Domain Informations** 🌐

Bu proje, belirli bir alan adıyla ilişkili bilgileri toplamak için **Go** programlama dili ile geliştirilmiştir. Proje, verilen bir site ile alakalı Google taramalarını yapar ve site içerisindeki bağlantılı linkleri toplar.

## ✨ Özellikler

- 🔍 Verilen site ile alakalı Google taraması yapar
- 📂 Google taraması sonucu bulunan Excel dosyalarının linklerini toplar
- 🔗 Verilen site içerisindeki bağlantılı linkleri toplar

## 🚀 Kurulum

### Gereksinimler

- Go 1.18+ sürümü
- Google Search API erişimi

### Adımlar

1. Repoyu klonlayın:
   ```bash
   git clone https://github.com/uslanozan/go-domain-informations.git
   cd go-domain-informations

2. Gerekli bağımlılıkları yükleyin:
   ```bash
   go mod tidy

3. Projeyi çalıştırın:
   ```bash
   go run main.go

4. Site URL'sini ve Google API anahtarınızı girin:
   ```bash
   Enter Site URL: YOUR_SITE_URL
   Enter Google API Key: YOUR_GOOGLE_API_KEY
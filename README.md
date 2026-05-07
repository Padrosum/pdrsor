## pdrsor 🔮

pdrsor, doğal dildeki isteklerinizi yerel bir yapay zeka (Ollama) kullanarak Bash komutlarına dönüştüren, Go ile yazılmış minimalist bir terminal asistanıdır. Verileriniz bilgisayarınızdan çıkmaz, gizlilik önceliklidir.
Özellikler

  %100 Yerel: Ollama API'sini kullanarak tüm işlemleri kendi makinenizde yapar.

  Otomatik Kopyalama: Üretilen komutları anında panoya (clipboard) kopyalar (xclip veya wl-copy desteği).  
    Hızlı Yapılandırma: İlk çalıştırmada modelinizi seçer, sonrasında ayarları hatırlar.  
    Görsel Arayüz: Terminal ruhuna uygun renkli ve tablolu çıktı düzeni.  
    Hafif: Go ile derlenmiş, bağımlılıksız tek bir binary.  
    
### Gereksinimler
    Ollama (Arka planda çalışıyor olmalı)  
    Go 1.20+ (Derlemek için)  
    xclip (X11) veya wl-clipboard (Wayland)  

### Kurulum

Projeyi klonlayıp Makefile üzerinden hızlıca kurabilirsiniz:
Bash
```
git clone https://github.com/kullaniciadin/pdrsor.git
cd pdrsor
go build -o pdrsor main.go

```
### Kullanım

Terminalde pdrsor komutundan sonra ne yapmak istediğinizi yazmanız yeterli:  

```
pdrsor "8080 portunu kullanan işlemi bul ve sonlandır"
```

Çıktı:
```
┌──────────────────────────────────────────┐
│ pdrsor - Local AI Terminal Assistant    │
└──────────────────────────────────────────┘

➜ Önerilen Komut:
   lsof -ti:8080 | xargs kill -9

  [ℹ] Komut panoya kopyalandı.
```

### Yapılandırma

Araç, ayarlarını **~/.config/pdrsor_rc** dosyasında tutar. Kullanılan modeli değiştirmek isterseniz bu dosyayı manuel olarak düzenleyebilir veya silebilirsiniz.
Neden pdrsor?

Bulut tabanlı AI asistanlarının aksine pdrsor, terminal geçmişinizi veya sistem yapınızı dış sunuculara göndermez. Sadece ihtiyacınız olan komutu üretir, panonuza kopyalar ve aradan çekilir.

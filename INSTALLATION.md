# 🚀 Инструкции по установке и запуску

## 📋 Предварительные требования

### Windows
1. **Go 1.22+** - [Скачать с golang.org](https://golang.org/dl/)
2. **Node.js 16+** - [Скачать с nodejs.org](https://nodejs.org/)
3. **Git** - [Скачать с git-scm.com](https://git-scm.com/download/win)

### macOS
1. **Go 1.22+** - [Скачать с golang.org](https://golang.org/dl/)
2. **Node.js 16+** - [Скачать с nodejs.org](https://nodejs.org/)
3. **Git** - обычно уже установлен, или через [Homebrew](https://brew.sh/)

### Linux (Ubuntu/Debian)
```bash
# Установка Go
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc

# Установка Node.js
curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash -
sudo apt-get install -y nodejs

# Установка Git
sudo apt-get install git
```

# Установка Node.js
curl -fsSL https://rpm.nodesource.com/setup_18.x | sudo bash -
sudo yum install -y nodejs

### Шаг 2: Установка Wails
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Шаг 3: Установка зависимостей

#### Backend (Go модули)
```bash
go mod tidy
```

#### Frontend (Node.js пакеты)
```bash
cd frontend
npm install
cd ..
```

### Шаг 4: Сборка фронтенда
```bash
cd frontend
npm run build
cd ..
```

### Шаг 5: Запуск приложения

#### Вариант A: Режим разработки (рекомендуется)
```bash
wails dev
```

#### Вариант B: Сборка готового приложения
```bash
# Сборка
wails build

# Запуск (Windows)
./build/bin/todo-list-wails.exe

# Запуск (macOS/Linux)
./build/bin/todo-list-wails
```

## 🔧 Устранение проблем

### Проблема: "wails: command not found"
**Решение:**
```bash
# Проверьте что Go установлен
go version

# Переустановите Wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# Добавьте Go bin в PATH (если нужно)
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

### Проблема: "npm: command not found"
**Решение:**
- Windows: Переустановите Node.js с официального сайта
- macOS: `brew install node`
- Linux: Следуйте инструкциям выше

### Проблема: "go: command not found"
**Решение:**
- Windows: Переустановите Go с официального сайта
- macOS: `brew install go`
- Linux: Следуйте инструкциям выше

### Проблема: "wails dev" не работает
**Решение:**
```bash
# Убедитесь что фронтенд собран
cd frontend
npm run build
cd ..

# Проверьте версию Wails
wails version

# Попробуйте пересобрать
wails build
```

### Проблема: Приложение не запускается
**Решение:**
```bash
# Проверьте версии
go version
node --version
npm --version

# Переустановите зависимости
go mod tidy
cd frontend
rm -rf node_modules package-lock.json
npm install
npm run build
cd ..
```

Если у вас возникли проблемы:
1. Проверьте версии: `go version`, `node --version`
2. Убедитесь что все зависимости установлены
3. Попробуйте пересобрать проект с нуля
4. Проверьте логи в терминале на наличие ошибок


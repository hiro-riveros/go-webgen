# go-webgen

CLI para generar proyectos web en Go de forma rápida y modular.

---

## 🚀 ¿Qué es go-webgen?

`go-webgen` es una herramienta de línea de comandos que te permite scaffoldear proyectos web en Go con las librerías y arquitecturas más comunes, como Gin, GORM y configuración modular. Facilita iniciar nuevos proyectos con estructura limpia y código base listo para producción.

---

## ⚙️ Instalación

### Desde código fuente

```bash
git clone https://github.com/hiro-riveros/go-webgen.git
cd go-webgen
go build -o go-webgen
```

## Usando Homebrew

```bash
brew tap hiro-riveros/tap
brew install go-webgen
```

## 🧪 Uso básico
Generar un proyecto Gin:
```bash
go-webgen generate myapp --gin
cd myapp
```

### Esto crea un proyecto con:
- Estructura modular (config, models, internal/api)
- Middlewares comunes (CORS, rate limit)
- Conexión lista para PostgreSQL con GORM
- Endpoint de ejemplo /ping


## 📦 Características
Generación rápida y confiable de boilerplates
Uso de plantillas embebidas con embed
Archivos Docker y Makefile para desarrollo y despliegue
Soporte para futuras arquitecturas y frameworks

## 🤝 Contribuir
Contribuciones son bienvenidas!
Abre un issue o pull request en el repositorio.

## 📄 Licencia
MIT License
# URL Shortener

Un servicio de acortamiento de URLs implementado en Go, con interfaz web y almacenamiento en Redis.

## Características

- Interfaz web simple y moderna
- Generación de URLs cortas aleatorias
- Almacenamiento persistente en Redis
- Redirección automática de URLs cortas
- Diseño modular y mantenible

## Requisitos

- Go 1.16 o superior
- Redis Server

## Instalación

1. Clona el repositorio:
```bash
git clone https://github.com/joaquinvulcano/url-shortener.git
cd url-shortener
```

2. Instala las dependencias:
```bash
go mod tidy
```

3. Asegúrate de tener Redis ejecutándose en `localhost:6379`

## Ejecución

1. Desde el directorio raíz del proyecto:
```bash
go run cmd/main.go
```

2. Abre tu navegador y visita `http://localhost:8080`

## Estructura del Proyecto

```
url-shortener/
│── cmd/
│   └── main.go          # Punto de entrada de la aplicación
│── internal/
│   ├── handlers/        # Manejadores HTTP
│   ├── models/          # Modelos de datos
│   ├── storage/         # Lógica de almacenamiento en Redis
│── templates/           # Archivos HTML para el frontend
│── static/             # Archivos estáticos (CSS, JS)
│── go.mod              # Módulo de Go
│── README.md           # Este archivo
```

## Uso

1. Accede a la página principal en `http://localhost:8080`
2. Ingresa la URL larga que deseas acortar
3. Haz clic en "Acortar URL"
4. Copia la URL corta generada o haz clic en ella para probar la redirección

## Contribuir

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Licencia

Este proyecto está licenciado bajo la Licencia MIT - ver el archivo `LICENSE` para más detalles.

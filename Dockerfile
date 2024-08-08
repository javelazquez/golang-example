# Utiliza la imagen oficial de Go como base
FROM golang:1.22

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo go.mod y go.sum para instalar las dependencias primero
COPY go.mod go.sum ./

# Descarga las dependencias del proyecto
RUN go mod download

# Copia el código fuente del proyecto al contenedor
COPY . .

# Compila la aplicación Go
RUN go build -o /myapp

# Define el comando por defecto para ejecutar la aplicación
CMD ["/myapp"]

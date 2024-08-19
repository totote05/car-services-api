#/bin/bash

# Nombre del archivo: run_tests.sh

# Definir variables
COVERAGE_DIR="coverage"
COVERAGE_FILE="$COVERAGE_DIR/coverage.out"
COVERAGE_FILE_FILTERED="$COVERAGE_DIR/coverage.filtered.out"
COVERAGE_REPORT="$COVERAGE_DIR/index.html"
COVERAGE_IGNORE=".covignore"

# Crear la carpeta coverage si no existe
if [ ! -d "$COVERAGE_DIR" ]; then
  mkdir "$COVERAGE_DIR"
fi

# Vaciar la carpeta coverage si existe
if [ -d "$COVERAGE_DIR" ]; then
  rm -rf "$COVERAGE_DIR"/*
fi

# Lista de paquetes por defecto
PKGS="./..."

# Construye un patrón de grep para excluir líneas basadas en el archivo de ignorados
IGNORE_PATTERN=$(grep -v '^[[:space:]]*$\|^\s*#' "$COVERAGE_IGNORE")

# Si el patrón de ignorados no está vacío, se genera la lista de paquetes para el coverage
if [ -n "$IGNORE_PATTERN" ]; then
  # Se crea una lista separada por "," de los paquetes deseados filtrando los ignorados
  PKGS=$(go list ./... | grep -vE "$IGNORE_PATTERN" | tr '\n' ',' | sed 's/,$//')
fi

# Ejecuta los tests con coverage
go test \
  -race \
  -shuffle=on \
  -coverprofile="$COVERAGE_FILE" \
  -coverpkg="$PKGS" \
  ./...


# Verificar si el patrón de ignorados está vacío
if [ -n "$IGNORE_PATTERN" ]; then
  # Filtra los resultados de cobertura excluyendo los patrones de ignorados
  grep -vE "$IGNORE_PATTERN" "$COVERAGE_FILE" > "$COVERAGE_FILE_FILTERED"
else
  # Si el patrón está vacío, simplemente copia el archivo de cobertura
  cp "$COVERAGE_FILE" "$COVERAGE_FILE_FILTERED"
fi

# Genera el resultado para la consola
go tool cover -func=$COVERAGE_FILE_FILTERED

# Genera el resultado en formato html si no se corre en el ci
if [ -z "$CI" ] || [ "$CI" != "true" ]; then
  gocov convert $COVERAGE_FILE_FILTERED | gocov-html -t kit > $COVERAGE_REPORT
fi

# Obtiene el porcentaje total de la cobertura
COVERAGE_PERCENTAGE=$(go tool cover -func=$COVERAGE_FILE_FILTERED | grep total | awk '{print int($3)}')

# Muestra un mensajes según el porcentaje de cobertura
if [ "$COVERAGE_PERCENTAGE" -eq 100 ]; then
  echo "Wow! 100% coverage, it's awesome! 😍"
elif [ "$COVERAGE_PERCENTAGE" -ge 80 ]; then
  echo "$COVERAGE_PERCENTAGE% - Coverage is looking good! 😇 Keep it up! 😎"
elif [ "$COVERAGE_PERCENTAGE" -ge 70 ]; then
  echo "$COVERAGE_PERCENTAGE% - Hmm, didn't quite reach the 80% coverage goal, but close! 😅 Keep pushing! 💪"
  exit 1
else
  echo "$COVERAGE_PERCENTAGE% - Uh oh, coverage seems a bit low... 💩"
  exit 1
fi
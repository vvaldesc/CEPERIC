# Scripts de utilidad para el proyecto

## Desarrollo Local
```bash
npm install
npm start
```

## Docker - Desarrollo
```bash
# Construir y ejecutar
docker-compose up dev

# Reconstruir imagen
docker-compose up --build dev

# Ejecutar en segundo plano
docker-compose up -d dev

# Ver logs
docker-compose logs -f dev

# Detener
docker-compose down
```

## Docker - Producción
```bash
# Construir y ejecutar
docker-compose up prod

# Reconstruir imagen
docker-compose up --build prod

# Ejecutar en segundo plano
docker-compose up -d prod

# Ver logs
docker-compose logs -f prod

# Detener
docker-compose down
```

## Comandos útiles
```bash
# Generar componente
ng generate component nombre-componente

# Generar servicio
ng generate service nombre-servicio

# Generar módulo
ng generate module nombre-modulo

# Ejecutar tests
npm test

# Build de producción local
npm run build

# Limpiar node_modules y reinstalar
rm -rf node_modules package-lock.json
npm install
```

## Limpiar Docker
```bash
# Eliminar contenedores
docker-compose down

# Eliminar contenedores y volúmenes
docker-compose down -v

# Eliminar todo (contenedores, volúmenes e imágenes)
docker-compose down -v --rmi all
```

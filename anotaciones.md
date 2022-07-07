# Comando donde vienen las credenciales del docker que se ejecuto
docker run --name some-postgres -e POSTGRES_USER=fazt -e POSTGRES_PASSWORD=mysecretpassword -p 5433:5432 -d postgres

# Accedemos al docker de manera interactiva
docker exec -it some-postgres bash

# acedemos al postgres
psql -U fazt --password

# despues solo creamos la base de datos
CREATE DATABASE gorm;
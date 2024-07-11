# Como usar

## Download dos módulos para cache local

```shell
make install
```

## Inicialização

```shell
make start
``` 

## Execução de testes

```shell
make run-test
``` 

# Curls

## Caminhões

### Cadastro de caminhão

```shell
curl --location 'http://localhost:8080/api/truck' \
--header 'Content-Type: application/json' \
--data '{
    "plateNumber": "GLI2462"
}'
```

#### Resposta

```json
{
    "id": "bfdc7e25-7fdb-4b09-a098-9003f995d8fd",
    "plateNumber": "GLI2462",
    "createdAt": "2024-07-11T10:24:39.441774023-03:00"
}
```

### Listagem de caminhões

```shell
curl --location 'http://localhost:8080/api/truck' \
--header 'Content-Type: application/json' \
--data ''
```

```shell
curl --location 'http://localhost:8080/api/truck?id=5e337efe-e7ad-46f6-9ac9-71e8c2cb9634' \
--header 'Content-Type: application/json' \
--data ''
```

```shell
curl --location 'http://localhost:8080/api/truck?plateNumber=6C1C45' \
--header 'Content-Type: application/json' \
--data ''
```

#### Resposta

```json
[
    {
        "id": "fc7c95cc-3176-46bd-88f2-073e24a7f948",
        "plateNumber": "6C1C45",
        "createdAt": "2024-07-11T10:31:42.880187Z"
    }
]
```

### Atualização de caminhão

```shell
curl --location --request PUT 'http://localhost:8080/api/truck' \
--header 'Content-Type: application/json' \
--data '{
    "id": "bfdc7e25-7fdb-4b09-a098-9003f995d8fd",
    "plateNumber": "GLI2464",
    "createdAt": "2024-07-10T10:24:39.441774023-03:00"
}'
```

#### Resposta

```json
{
    "id": "bfdc7e25-7fdb-4b09-a098-9003f995d8fd",
    "plateNumber": "GLI2464",
    "createdAt": "2024-07-10T10:24:39.441774023-03:00"
}
```

### Remoção de caminhão

```shell
curl --location --request DELETE 'http://localhost:8080/api/truck?id=4b9279e8-bb5d-4be0-96e9-55b96897399a' \
--header 'Content-Type: application/json' \
--data ''
```

#### Resposta status 200

## Motoristas

### Cadastro de motorista

```shell
curl --location 'http://localhost:8080/api/driver' \
--header 'Content-Type: application/json' \
--data '{
    "document": "54935848082"
}'
```

#### Resposta

```json
{
    "id": "e7dc3e9c-c714-4d2e-9a1c-77048f50b0ee",
    "document": "54935848082",
    "createdAt": "2024-07-11T12:00:30.68343841-03:00"
}
```

### Listagem de motoristas

```shell
curl --location 'http://localhost:8080/api/driver' \
--data ''
```

```shell
curl --location 'http://localhost:8080/api/driver?document=30534152082' \
--data ''
```

#### Resposta

```json
[
    {
        "id": "64596e48-fa08-4550-bd79-9251eaa54b02",
        "document": "30534152082",
        "createdAt": "2024-07-10T21:59:29.31226Z"
    }
]
```

### Atualização de motorista

```shell
curl --location --request PUT 'http://localhost:8080/api/driver' \
--header 'Content-Type: application/json' \
--data '{
    "id": "64596e48-fa08-4550-bd79-9251eaa54b02",
    "document": "28727572013",
    "createdAt": "2024-07-05T22:43:48.487Z"
}'
```

#### Resposta

```json
{
    "id": "64596e48-fa08-4550-bd79-9251eaa54b02",
    "document": "28727572013",
    "createdAt": "2024-07-05T22:43:48.487Z"
}
```

### Remoção de motorista

```shell
curl --location --request DELETE 'http://localhost:8080/api/driver?id=461616e6-b01f-4940-8ce3-f5ef3cef8428' \
--data ''
```

#### Resposta status 200

## Relação Caminhão x Motorista

### Criação de vínculo

```shell
curl --location 'http://localhost:8080/api/truck/relation' \
--header 'Content-Type: application/json' \
--data '{
    "driverId": "dcb03fd6-fb66-44b7-857f-9f598f56f570",
    "truckId": "1f79d093-62f0-4e9b-81e7-d635d0357ed0"
}'
```

#### Resposta

```json
{
    "driverId": "dcb03fd6-fb66-44b7-857f-9f598f56f570",
    "truckId": "1f79d093-62f0-4e9b-81e7-d635d0357ed0",
    "id": "aa46547a-9dd4-4328-89bb-ca1f5e84f18c",
    "createdAt": "2024-07-11T12:06:13.153168091-03:00"
}
```

### Listagem de vínculos

```shell
curl --location 'http://localhost:8080/api/truck/relation' \
--data ''
```

#### Resposta

```json
[
    {
        "driverId": "7ba22e91-d3c3-4423-8902-fd7c284c079a",
        "truckId": "4873700e-956f-4f28-bf52-06fd16622779",
        "id": "267a60f2-69a2-4136-8c79-f47a2c2cbc1a",
        "createdAt": "2024-07-11T11:50:09.049207Z"
    },
    {
        "driverId": "dcb03fd6-fb66-44b7-857f-9f598f56f570",
        "truckId": "1f79d093-62f0-4e9b-81e7-d635d0357ed0",
        "id": "aa46547a-9dd4-4328-89bb-ca1f5e84f18c",
        "createdAt": "2024-07-11T12:06:13.153168Z"
    }
]
```

### Atualização de vínculos

```shell
curl --location --request PUT 'http://localhost:8080/api/truck/relation' \
--header 'Content-Type: application/json' \
--data '{
    "id": "461616e6-b01f-4940-8ce3-f5ef3cef8428",
    "driverId": "64596e48-fa08-4550-bd79-9251eaa54b02",
    "createdAt": "2024-07-05T22:43:48.487Z"
}'
```

#### Resposta

```json
{
    "driverId": "64596e48-fa08-4550-bd79-9251eaa54b02",
    "id": "461616e6-b01f-4940-8ce3-f5ef3cef8428",
    "createdAt": "2024-07-05T22:43:48.487Z"
}
```

### Remoção de vínculos

```shell
curl --location --request DELETE 'http://localhost:8080/api/truck/relation?id=dcd44d29-40e1-43ac-9ee4-e715923a5cdb' \
--data ''
```

#### Resposta status 200
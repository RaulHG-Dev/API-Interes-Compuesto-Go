# API

## Client

### `/v1.0.0/compound-interest/year`

GET - Interés compuesto por año.

#### Parameters

- `capitalInicial` - Capital inicial de inversión.
- `anios` - Número de años a calcular.
- `interes` - Interés anual de inversión.

#### Request body

```json
{
    "capitalInicial": 100,
    "anios": 2,
    "interes": 15
}
```

#### Response

```json
{
  "success": true,
  "code": 200,
  "message": "OK",
  "data": [
    {
      "year": 1,
      "capital": 115
    },
    {
      "year": 2,
      "capital": 132.25
    }
  ]
}
```

### `/v1.0.0/compound-interest/month`

GET - Interés compuesto por mes.

#### Parameters

- `capitalInicial` - Capital inicial de inversión.
- `anios` - Número de años a calcular.
- `interes` - Interés anual de inversión.

#### Request body

```json
{
    "capitalInicial": 100,
    "anios": 2,
    "interes": 15
}
```

#### Response

```json
{
    "success": true,
    "code": 200,
    "message": "OK",
    "data": [
        {
            "year": 1,
            "month": 1,
            "montfOfTheYear": 1,
            "monthName": "Enero",
            "capital": 101.171491691985
        },
        {
            "year": 1,
            "month": 2,
            "montfOfTheYear": 2,
            "monthName": "Febrero",
            "capital": 102.356707311815
        },
        {
            "year": 1,
            "month": 3,
            "montfOfTheYear": 3,
            "monthName": "Marzo",
            "capital": 103.555807634162
        },
        ...
    ]
}
```

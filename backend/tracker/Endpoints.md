# Tracker API V1.0
## Endpoints and JSON Data

### Communication with Product-API

GET http://localhost:8082/tracker/products

Get Array of Products from Product-API
#### Response Data
```json
[
  {
    "id": 1,
    "name": "Skyr",
    "brand": "K-Classic",
    "caloriesPer100Grams": 64,
    "proteinsInGrams": 11,
    "fatsInGrams": 0.2,
    "carbsInGrams": 4
  },
  {
    "id": 2,
    "name": "Protein Riegel",
    "brand": "ESN",
    "caloriesPer100Grams": 395,
    "proteinsInGrams": 28,
    "fatsInGrams": 18,
    "carbsInGrams": 35
  },
  {
    "id": 3,
    "name": "Kartoffeln",
    "brand": "K-Classic",
    "caloriesPer100Grams": 76,
    "proteinsInGrams": 2,
    "fatsInGrams": 0.5,
    "carbsInGrams": 16
  }
]
```
---
GET http://localhost:8082/tracker/products/4

Get the product with the ID 4 from Product-API
#### Response Data
```json
{
  "id": 4,
  "name": "Skyr",
  "brand": "Milbona",
  "caloriesPer100Grams": 62,
  "proteinsInGrams": 11,
  "fatsInGrams": 0.2,
  "carbsInGrams": 4
}
```
---
GET http://localhost:8082/tracker/products/name/Skyr

Get one or more products with the name from the parameter.
If more than one product, result comes in array.
#### Response Data
```json
[
  {
    "id": 1,
    "name": "Skyr",
    "brand": "K-Classic",
    "caloriesPer100Grams": 64,
    "proteinsInGrams": 11,
    "fatsInGrams": 0.2,
    "carbsInGrams": 4
  },
  {
    "id": 4,
    "name": "Skyr",
    "brand": "Milbona",
    "caloriesPer100Grams": 62,
    "proteinsInGrams": 11,
    "fatsInGrams": 0.2,
    "carbsInGrams": 4
  }
]
```

### Nutrition-Settings

GET http://localhost:8082/tracker/settings

Get the settings (needed as "Soll" for Dashboard)
#### Expected Input Data
```json
{
  "PlannedCalories": 2500,
  "ProteinsInGrams": 201,
  "FatsInGrams": 54,
  "CarbsInGrams": 287
}
```
-----------
PATCH http://localhost:8082/tracker/settings

Update the Nutrition-Settings. Calculate the percentage with the PlannedCalories 
and store the calculated Grams in the db
#### Needed Data
```json
{
  "PlannedCalories": 2000,
  "ProteinsPercentage": 20,
  "FatsPercentage": 40,
  "CarbsPercentage": 40
}
```

### Consume Product
These endpoint simulates the addition of a product to the diary (Lebensmittel hinzufügen)

POST http://localhost:8082/tracker/consume
#### Expected Input Data
(Der Endpunkt wird benötigt wenn du ein Product ausgewählt hast von der Product-API.
Du musst die Produktdaten des ausgewählten Products an den POST Request übergeben. 
Zusätzlich kommen Datum, Gewicht in Gramm und Kategorie dazu)

"Category" kann folgende Strings annehmen: Frühstück, Mittagessen, Abendesen, Snack
```json
{
  "Product": {
    "id": 6,
    "name": "Hähnchenbrustfilet",
    "brand": "Landjunker",
    "caloriesPer100Grams": 107,
    "proteinsInGrams": 23,
    "fatsInGrams": 2.6,
    "carbsInGrams": 0.5
  },
  "Date": "2025-02-10",
  "Weight": 200,
  "Category": "Frühstück"
}
```

### Diary
These endpoints are needed to get all consumed food for a specific date.
There's also one endpoint to delete consumed Products

GET http://localhost:8082/tracker/diary/date/2025-02-10

Get the products that are consumed on a specific date.
Is the standard view when open the diary(tagebuch) tab.
```json
{
  "date": "2025-02-10",
  "products": [
    {
      "ID": 1,
      "DailyProductsConsumedID": 1,
      "ProductID": 6,
      "ProductName": "Hähnchenbrustfilet",
      "Category": "Snack",
      "WeightInGrams": 500,
      "Calories": 535,
      "ProteinsInGrams": 115,
      "FatsInGrams": 13,
      "CarbsInGrams": 2.5
    },
    {
      "ID": 2,
      "DailyProductsConsumedID": 1,
      "ProductID": 6,
      "ProductName": "Hähnchenbrustfilet",
      "Category": "Abendessen",
      "WeightInGrams": 500,
      "Calories": 535,
      "ProteinsInGrams": 115,
      "FatsInGrams": 13,
      "CarbsInGrams": 2.5
    }
  ]
}
```
---
GET http://localhost:8082/tracker/diary/5

Get a product by the ID (ID is from ConsumedProduct Table).

```json
{
  "ID": 5,
  "DailyProductsConsumedID": 2,
  "ProductID": 6,
  "Category": "Abendessen",
  "WeightInGrams": 500,
  "Calories": 535,
  "ProteinsInGrams": 115,
  "FatsInGrams": 13,
  "CarbsInGrams": 2.5
}
```
---
DELETE http://localhost:8082/tracker/diary/2

Delete a product by the ID (ID is from ConsumedProduct Table).
We can get the id when picking a product from the diary list (the first get call from diary)
#### expected message after deletion
```json
{
  "message": "Product deleted successfully"
}
```

### Nutrition-Statistics

GET http://localhost:8082/tracker/nutrition/date/2025-02-09

Get the Nutrition-Statistics for a specific Date (IST-Werte fürs Dashboard)

```json
{
  "ID": 2,
  "Date": "2025-02-09",
  "ConsumedCalories": 0,
  "ConsumedProteins": 0,
  "ConsumedFats": 0,
  "ConsumedCarbs": 0
}
```

Funktion zum Updaten der Werte in der DB folgt noch

### Shopping-List

Hier fehlen noch die genauen Daten welche Pascal durch seine API bereitstellt

Endpoints sind nur vorläufig bis genaue Implementierung der API bereit steht

GET http://localhost:8082/tracker/shopping-list (Liste aller Shopping-Listen)

GET http://localhost:8082/tracker/shopping-List/:id (eine ausgewählte Liste anzeigen)

OPTIONAL: GET http://localhost:8082/tracker/shopping-List/name/:name (eine ausgewählte Liste per Namenssuche)

DELETE http://localhost:8082/tracker/shopping-List/:id (eine ausgewählte Liste löschen)

POST http://localhost:8082/tracker/shopping-list (Liste erstellen)

#### Vorschlag der JSON Daten

```json
{
  "ID": 1,
  "Date": "2025-02-10",
  "Name": "LIDL-Einkauf",
  "Products": [
    {
      "ID": 1,
      "ShoppingListID": 1,
      "Name": "Test",
      "Quantity": 500,
      "Unit": "grams",
      "shopped": false
    },
    {
      "ID": 2,
      "ShoppingListID": 1,
      "Name": "Test",
      "Quantity": 500,
      "Unit": "kg",
      "shopped": true
    }
  ]
}
```

#### Produkte hinzufügen

GET http://localhost:8082/tracker/shopping-list/:id/products (Alle Products in der Liste anzeigen)

POST http://localhost:8082/tracker/shopping-list/:id (Produkt anlegen)

OPTIONAL: PATCH http://localhost:8082/tracker/shopping-list/:id/products/:id (spezifisches Produkt bearbeiten)

DELETE http://localhost:8082/tracker/shopping-list/:id/products/:id (spezifisches Produkt löschen)

#### Vorschlag der JSON Daten
```json
{
  "ID": 1,
  "ShoppingListID": 1,
  "Name": "Test",
  "Quantity": 500,
  "Unit": "grams",
  "shopped": false
}
```

### Calender

Hier fehlen noch die genauen Daten welche Pascal durch seine API bereitstellt

Endpoints sind nur vorläufig bis genaue Implementierung der API bereit steht

#### Vorschlag der JSON Daten

GET http://localhost:8082/tracker/calender/date/:date (Alle einträge für spezifisches datum)

GET http://localhost:8082/tracker/calender/date/:date/:id (Eintrag an spezifischem Datum)

DELETE http://localhost:8082/tracker/calender/date/:date/:id (Eintrag sn spezifischem Datum löschen)

```json
{
  "ID": 1,
  "Date": "2025-02-10",
  "Name": "Test",
  "Category": "Shopping, Gym, ...",
}
```
Optional noch Uhrzeit

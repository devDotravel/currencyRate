# currencyRate

Creates json with currency rates for British Pounds (**GBP**) and US dollars (**USD**) compared to euros (**EUR**). Uses fixer.io API.

i.e.
*"GBP: 1.15"* means 1 pound equals 1.15 euros

***You'll need a valid API key for fixer.io*** in a plain file called *apiKey* in config folder

*config.template* must be changed to *config.go* with your own variables.

In order to have different currencies to GBP and USD, you'll have to alter the fields of the structs *rates* and *currencyReverseRates*, as well as the string *urlRequest*, in *currencyRate.go*

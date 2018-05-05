# Polycon CLI
CLI For simple Polygon.io functions.

Required args:

	--apiKey=YOUR_API_KEY

## Streams
You can stream channels: 
- Trades
- Quotes
- Currencies

For example, to stream all Trade ticks:
	
	polygon-cli stream --channel T.*

## Metadata 

Get Symbol Financials
	
	polygon-cli financials --symbol MFST

Get Symbol Earnings

	polygon-cli earnings --symbol MFST

Get Exchanges

	polygon-cli exchanges

Get SymbolType Maps

	polygon-cli symboltypes


### TODO:
There is still lots to do:
- BONDS, COMMODITIES, METALS streams
- All RESTful Stock endpoints
- All RESTful Currencies endpoints


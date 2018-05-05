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
	
	polygon-cli --apiKey=YOUR_API_KEY stream --channel=T.*


### TODO:
There is still lots to do:
- BONDS, COMMODITIES, METALS streams
- All RESTful Stock endpoints
- All RESTful Currencies endpoints


# Polygon CLI
CLI For simple Polygon.io functions.

Required args:

	--apiKey=YOUR_API_KEY

## Streams
You can stream channels: 
- Trades `T.*`
- Quotes `Q.*`
- Currencies `C.*`

For example, to stream all Trade ticks:
	
	polygon-cli stream --channel T.*

## Metadata 

Get Symbol Financials
	
	polygon-cli financials --symbol MFST

Get Symbol Earnings

	polygon-cli earnings --symbol MFST

Get Exchanges

	polygon-cli exchanges

	  ID |                    NAME                     | TAPE |   TYPE   |   MARKET   | MIC   
	+----+---------------------------------------------+------+----------+------------+------+
	   0 | Multiple                                    |    - | TRF      | equities   | TFF   
	   1 | NYSE American (AMEX)                        | A    | exchange | equities   | XASE  
	   2 | NASDAQ OMX BX                               | B    | exchange | equities   | XBOS  
	   3 | National Stock Exchange                     | C    | TRF      | equities   | XCIS  
	   4 | FINRA                                       | D    | TRF      | equities   | FINR  
	   5 | Consolidated Quote System                   | E    | TRF      | equities   | CQS   
	   6 | International Securities Exchange           | I    | TRF      | equities   | XISX  
	   7 | Cboe EDGA                                   | J    | exchange | equities   | EDGA  
	   8 | Cboe EDGX                                   | K    | exchange | equities   | EDGX  
	   9 | Chicago Stock Exchange, Inc                 | M    | exchange | equities   | XCHI  
	  10 | New York Stock Exchange                     | N    | exchange | equities   | XNYS  
	  11 | NYSE Arca                                   | P    | exchange | equities   | ARCX  
	  12 | Nasdaq                                      | Q    | exchange | equities   | XNGS  
	  13 | Consolidated Tape System                    | S    | TRF      | equities   | CTS   
	  14 | OTC Bulletin Board                          | U    | TRF      | equities   | XOTC  
	  15 | IEX                                         | V    | exchange | equities   | IEXG  
	  16 | Chicago Board Options Exchange              | W    | TRF      | equities   | XCBO  
	  17 | Nasdaq PSX                                  | X    | exchange | equities   | XPHL  
	  18 | Cboe BYX                                    | Y    | exchange | equities   | BATY  
	  19 | Cboe BZX                                    | Z    | exchange | equities   | BATS  
	  33 | NASDAQ BX Options/ETF                       | B    | exchange | equities   | XBOS  
	  37 | Chicago Board Options Exchange ( Indecies ) | W    | TRF      | indecies   | XCBO  
	  20 | Currency Banks 1                            |      | banking  | currencies |       
	  44 | Currency Banks 2                            |      | banking  | currencies |       
	  60 | Currency Banks 3                            |      | banking  | currencies |       





Get SymbolType Maps

	polygon-cli symboltypes


### TODO:
There is still lots to do:
- BONDS, COMMODITIES, METALS streams
- All RESTful Stock endpoints
- All RESTful Currencies endpoints


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

#### Get Symbol Earnings

	polygon-cli earnings --symbol MFST

	  REPORT DATE |            TYPE             |          VALUE            
	+-------------+-----------------------------+--------------------------+
	  2018-04-26  | Fiscal Period               |                  Q3 2018  
	              | Fiscal End Date             | 2018-03-31T00:00:00.000Z  
	              | Actual EPS                  |                 0.950000  
	              | Consensus EPS               |                 0.850000  
	              | Estimated EPS               |                 0.850000  
	              | Announce Time               |                      AMC  
	              | Number Of Estimates         |                       14  
	              | EPS Surprise Dollar         |                 0.100000  
	              | Year Ago                    |                 0.730000  
	              | Year Ago Change Percent     |                       30  
	              | Estimated Change Percent    |                       16  
	  ---------   | --------------------------- |       ------------------  
	  2018-01-31  | Fiscal Period               |                  Q2 2018  
	              | Fiscal End Date             | 2017-12-31T00:00:00.000Z  
	              | Actual EPS                  |                 0.960000  
	              | Consensus EPS               |                 0.860000  
	              | Estimated EPS               |                 0.860000  
	              | Announce Time               |                      AMC  
	              | Number Of Estimates         |                       14  
	              | EPS Surprise Dollar         |                 0.100000  
	              | Year Ago                    |                 0.830000  
	              | Year Ago Change Percent     |                       16  
	              | Estimated Change Percent    |                        4  
	  ---------   | --------------------------- |       ------------------  



#### Get Exchanges

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





#### Get SymbolType Maps

	polygon-cli symboltypes

	  TYPE  |           DESCRIPTION             
	+-------+----------------------------------+
	  fdr   | Foreign Ordinary Shares           
	  ost   | Other Security Type               
	  adr   | American Depository Receipt       
	  cef   | Closed-End Fund                   
	  reit  | Real Estate Investment Trust      
	  trak  | Tracking stock or targeted stock  
	  rylt  | Royalty Trust                     
	  pfd   | Preferred Stock                   
	  fund  | Fund                              
	  cs    | Common Stock                      
	  etp   | Exchange Traded Product/Fund      
	  ltdp  | Limited Partnership               
	  mlp   | Master Limited Partnership        
	  mf    | Mutual Fund                       
	  sp    | Structured Product                
	  wrt   | Equity WRT                        
	  pub   | Public                            
	  nyrs  | New York Registry Shares          
	  unit  | Unit                              
	  right | Right                             



### TODO:
There is still lots to do:
- BONDS, COMMODITIES, METALS streams
- All RESTful Stock endpoints
- All RESTful Currencies endpoints


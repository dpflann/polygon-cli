# Polygon CLI
CLI For simple Polygon.io functions.

Required args:

	--apiKey=YOUR_API_KEY

# Streams
You can stream channels: 
- Trades `T.*`
- Quotes `Q.*`
- Currencies `C.*`

#### Streaming Quotes
	
	polygon-cli stream --channel Q.*
	Symbol:  NTNX 	 Bid: 54.39 	 Ask:  54.62 	 BidExch:  12 	 AskExch:  12 	 BidSize:  2 	 AskSize:  2 	 Timestamp:  1525477920173
	Symbol:  NTNX 	 Bid: 54.62 	 Ask:  54.63 	 BidExch:  12 	 AskExch:  11 	 BidSize:  8 	 AskSize:  3 	 Timestamp:  1525477920173
	Symbol:  TVIX 	 Bid: 6.77 	 Ask:  6.79 	 BidExch:  8 	 AskExch:  8 	 BidSize:  10 	 AskSize:  48 	 Timestamp:  1525477921622
	Symbol:  KOSS 	 Bid: 1.8 	 Ask:  2.14 	 BidExch:  8 	 AskExch:  11 	 BidSize:  14 	 AskSize:  2 	 Timestamp:  1525477923597
	Symbol:  NTNX 	 Bid: 54.62 	 Ask:  54.63 	 BidExch:  12 	 AskExch:  11 	 BidSize:  3 	 AskSize:  3 	 Timestamp:  1525477924533
	Symbol:  NTNX 	 Bid: 54.39 	 Ask:  54.63 	 BidExch:  12 	 AskExch:  11 	 BidSize:  2 	 AskSize:  3 	 Timestamp:  1525477924533
	Symbol:  NVDA 	 Bid: 239.85 	 Ask:  239.99 	 BidExch:  11 	 AskExch:  12 	 BidSize:  5 	 AskSize:  7 	 Timestamp:  1525477924706
	Symbol:  NTNX 	 Bid: 54.39 	 Ask:  54.62 	 BidExch:  12 	 AskExch:  12 	 BidSize:  2 	 AskSize:  2 	 Timestamp:  1525477924541
	Symbol:  NTNX 	 Bid: 54.39 	 Ask:  54.63 	 BidExch:  12 	 AskExch:  11 	 BidSize:  2 	 AskSize:  3 	 Timestamp:  1525477925750
	Symbol:  DRRX 	 Bid: 1.78 	 Ask:  2.49 	 BidExch:  8 	 AskExch:  11 	 BidSize:  1 	 AskSize:  68 	 Timestamp:  1525477925550
	Symbol:  GUSH 	 Bid: 32.9 	 Ask:  33.24 	 BidExch:  8 	 AskExch:  8 	 BidSize:  7 	 AskSize:  1 	 Timestamp:  1525477925972
	Symbol:  GLD 	 Bid: 124.56 	 Ask:  124.6 	 BidExch:  11 	 AskExch:  8 	 BidSize:  1 	 AskSize:  1 	 Timestamp:  1525477925986
	Symbol:  DIS 	 Bid: 101.25 	 Ask:  101.45 	 BidExch:  8 	 AskExch:  11 	 BidSize:  2 	 AskSize:  2 	 Timestamp:  1525477926146
	Symbol:  DIS 	 Bid: 101.37 	 Ask:  101.45 	 BidExch:  8 	 AskExch:  11 	 BidSize:  7 	 AskSize:  2 	 Timestamp:  1525477926147
	Symbol:  SQQQ 	 Bid: 16.32 	 Ask:  16.34 	 BidExch:  11 	 AskExch:  11 	 BidSize:  3 	 AskSize:  64 	 Timestamp:  1525477927478
	Symbol:  NVDA 	 Bid: 239.85 	 Ask:  239.99 	 BidExch:  11 	 AskExch:  12 	 BidSize:  5 	 AskSize:  6 	 Timestamp:  1525477927637
	Symbol:  DBX 	 Bid: 29.22 	 Ask:  29.25 	 BidExch:  11 	 AskExch:  8 	 BidSize:  3 	 AskSize:  8 	 Timestamp:  1525477927897
	Symbol:  MU 	 Bid: 47.63 	 Ask:  47.65 	 BidExch:  11 	 AskExch:  8 	 BidSize:  10 	 AskSize:  19 	 Timestamp:  1525477928608

#### Streaming Trades:

	polygon-cli stream --channel T.*
	Symbol:  LB 	 Price: 34.6 	 Size:  50 	 Exchange:  8 	 Timestamp:  1525478065420 	 Conditions:  [14 12 37]
	Symbol:  CHK 	 Price: 3.02 	 Size:  527 	 Exchange:  11 	 Timestamp:  1525478066537 	 Conditions:  [14 12]
	Symbol:  CHK 	 Price: 3.02 	 Size:  173 	 Exchange:  11 	 Timestamp:  1525478066537 	 Conditions:  [12]
	Symbol:  TVIX 	 Price: 6.79 	 Size:  250 	 Exchange:  4 	 Timestamp:  1525478067693 	 Conditions:  [12]
	Symbol:  BAC 	 Price: 29.3 	 Size:  50 	 Exchange:  12 	 Timestamp:  1525478068963 	 Conditions:  [12 37]
	Symbol:  KOPN 	 Price: 3.45 	 Size:  2000 	 Exchange:  4 	 Timestamp:  1525478070234 	 Conditions:  [12]
	Symbol:  SZC 	 Price: 17.55 	 Size:  1500 	 Exchange:  11 	 Timestamp:  1525478070592 	 Conditions:  [14 12]
	Symbol:  AMD 	 Price: 11.3 	 Size:  100 	 Exchange:  4 	 Timestamp:  1525478071810 	 Conditions:  [12]
	Symbol:  BRK.B 	 Price: 195.7 	 Size:  15 	 Exchange:  4 	 Timestamp:  1525478072398 	 Conditions:  [12 37]
	Symbol:  FB 	 Price: 176.7 	 Size:  35 	 Exchange:  11 	 Timestamp:  1525478073429 	 Conditions:  [12 37]
	Symbol:  AAPL 	 Price: 184.05 	 Size:  747 	 Exchange:  12 	 Timestamp:  1525478073487 	 Conditions:  [14 12]
	Symbol:  KHC 	 Price: 58.17 	 Size:  60 	 Exchange:  4 	 Timestamp:  1525478074297 	 Conditions:  [12 37]
	Symbol:  GPRO 	 Price: 5.45 	 Size:  1000 	 Exchange:  11 	 Timestamp:  1525478075956 	 Conditions:  [14 12]



# Metadata 

#### Get Symbol Financials
	
	polygon-cli financials --symbol MFST

	  REPORT DATE |            TYPE             |       VALUE         
	+-------------+-----------------------------+--------------------+
	  2018-03-31  | Gross Profit                | $17,550,000,000     
	              | Cost of Revenue             | $9,269,000,000      
	              | Operating Revenue           | $26,819,000,000     
	              | Total Revenue               | $26,819,000,000     
	              | Operating Income            | $8,292,000,000      
	              | Net Income                  | $7,424,000,000      
	              | R & D                       | $3,715,000,000      
	              | Operating Expenses          | $9,258,000,000      
	              | Current Assets              | $156,659,000,000    
	              | Total Assets                | $245,497,000,000    
	              | Total Liabilities           | $0                  
	              | Current Cash                | $9,221,000,000      
	              | Current Debt                | $3,677,000,000      
	              | Total Cash                  | $132,270,000,000    
	              | Shareholder Equity          | $79,239,000,000     
	              | Cash Change                 | -$3,663,000,000     
	              | Cash Flow                   | $12,151,000,000     
	              | Operating Gains Losses      | -$347,000,000       
	    --------- | --------------------------- | ------------------  
	  2017-12-31  | Gross Profit                | $17,854,000,000     
	              | Cost of Revenue             | $11,064,000,000     
	              | Operating Revenue           | $28,918,000,000     
	              | Total Revenue               | $28,918,000,000     
	              | Operating Income            | $8,679,000,000      
	              | Net Income                  | -$6,302,000,000     
	              | R & D                       | $3,504,000,000      
	              | Operating Expenses          | $9,175,000,000      
	              | Current Assets              | $167,633,000,000    
	              | Total Assets                | $256,003,000,000    
	              | Total Liabilities           | $177,643,000,000    
	              | Current Cash                | $12,859,000,000     
	              | Current Debt                | $15,912,000,000     
	              | Total Cash                  | $142,780,000,000    
	              | Shareholder Equity          | $78,360,000,000     
	              | Cash Change                 | $5,992,000,000      
	              | Cash Flow                   | $7,875,000,000      
	              | Operating Gains Losses      | -$1,749,000,000     
	    --------- | --------------------------- | ------------------  


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



#### Get Symbol Dividends

	polygon-cli dividends --symbol MFST
	        REPORT DATE        |             KEY             |          VALUE            
	+--------------------------+-----------------------------+--------------------------+
	  2018-02-14T00:00:00.000Z | Execution Date              | 2018-02-14T00:00:00.000Z  
	                           | Type                        |          Dividend income  
	                           | Payment Date                | 2018-03-08T00:00:00.000Z  
	                           | Record Date                 | 2018-02-15T00:00:00.000Z  
	                           | Declared Date               | 2017-11-29T00:00:00.000Z  
	                           | Amount                      |                 0.420000  
	                           | Qualified                   |                        Q  
	                           | Flag                        |                           
	  ---------                | --------------------------- |       ------------------  
	  2017-11-15T00:00:00.000Z | Execution Date              | 2017-11-15T00:00:00.000Z  
	                           | Type                        |          Dividend income  
	                           | Payment Date                | 2017-12-14T00:00:00.000Z  
	                           | Record Date                 | 2017-11-16T00:00:00.000Z  
	                           | Declared Date               | 2017-09-19T00:00:00.000Z  
	                           | Amount                      |                 0.420000  
	                           | Qualified                   |                        Q  
	                           | Flag                        |                           



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


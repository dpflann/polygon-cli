
package cmd

import (
	"os"
	"log"
	"fmt"
	"bytes"
	"strconv"
	"encoding/json"
	cobra "github.com/spf13/cobra"
	schemas "github.com/polygon-io/go-schemas"
	gorequest "github.com/parnurzeal/gorequest"
	tablewriter "github.com/olekukonko/tablewriter"
	accounting "github.com/leekchan/accounting"
)


var metaCmdOps = struct {
	Symbol     string
}{}


func getUrl( url string, obj interface{} ){
	resp, body, err := gorequest.New().Get( url ).Query("apiKey="+apiKey).End(); if err != nil {
		log.Fatal( err )
	}
	if resp.Status != "200 OK" {
		fmt.Println("WARNING: Got a non 200 response. Got:", resp.Status)
	}
	jerr := json.Unmarshal([]byte(body), obj); if jerr != nil {
		fmt.Println("JSON Unmarshal Error:", jerr)
	}
}


// Exchanges:
var exchangesCmd = &cobra.Command{
	Use: "exchanges",
	Short: "Get the list of Exchanges",
	Run: func(cmd *cobra.Command, args []string) {
		var exchanges []schemas.Exchange
		getUrl( "https://api.polygon.io/v1/meta/exchanges", &exchanges )
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth( 60 )
		table.SetBorder(false)
		table.SetHeader([]string{"ID", "Name", "Tape", "Type", "Market", "mic"})
		for _, exch := range exchanges {
			table.Append([]string{ strconv.Itoa( exch.ID ), exch.Name, exch.Tape, exch.Type, exch.Market, exch.MIC })
		}
		println("")
		table.Render()
		println("")
	},
}



// Exchanges:
var cryptoExchangesCmd = &cobra.Command{
	Use: "cryptoexchanges",
	Short: "Get the list of Crypto Exchanges",
	Run: func(cmd *cobra.Command, args []string) {
		var exchanges []schemas.CryptoExchange
		getUrl( "https://api.polygon.io/v1/meta/crypto-exchanges", &exchanges )
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth( 60 )
		table.SetBorder(false)
		table.SetHeader([]string{"ID", "Name", "Type", "Market", "url"})
		for _, exch := range exchanges {
			table.Append([]string{ strconv.Itoa( exch.ID ), exch.Name, exch.Type, exch.Market, exch.URL })
		}
		println("")
		table.Render()
		println("")
	},
}


var symbolTypesCmd = &cobra.Command{
	Use: "symboltypes",
	Short: "Get the map for symbol types",
	Run: func(cmd *cobra.Command, args []string) {
		var symbolTypes map[string]string
		getUrl( "https://api.polygon.io/v1/meta/symbol-types", &symbolTypes )
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth( 60 )
		table.SetBorder(false)
		table.SetHeader([]string{"Type", "Description"})
		for k, v := range symbolTypes {
			table.Append([]string{ k, v })
		}
		println("")
		table.Render()
		println("")
	},
}


var symbolFinancialsCmd = &cobra.Command{
	Use: "financials",
	Short: "Get the financials for a symbol",
	Run: func(cmd *cobra.Command, args []string) {
		if len(metaCmdOps.Symbol) == 0 {
			log.Fatal("ERROR: Please enter a symbol. eg: --symbol=MSFT")
		}
		ac := accounting.Accounting{Symbol: "$", Precision: 0}
		var financials []schemas.Financial
		getUrl( "https://api.polygon.io/v1/meta/symbols/"+metaCmdOps.Symbol+"/financials", &financials )
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth( 60 )
		table.SetBorder(false)
		table.SetHeader([]string{"Report Date", "Key", "Value"})
		for _, fin := range financials {
			table.Append([]string{ fin.ReportDate, "Gross Profit", ac.FormatMoney( fin.GrossProfit ) })
			table.Append([]string{ "", "Cost of Revenue", ac.FormatMoney( fin.CostOfRevenue ) })
			table.Append([]string{ "", "Operating Revenue", ac.FormatMoney( fin.OperatingRevenue ) })
			table.Append([]string{ "", "Total Revenue", ac.FormatMoney( fin.TotalRevenue ) })
			table.Append([]string{ "", "Operating Income", ac.FormatMoney( fin.OperatingIncome ) })
			table.Append([]string{ "", "Net Income", ac.FormatMoney( fin.NetIncome ) })
			table.Append([]string{ "", "R & D", ac.FormatMoney( fin.ResearchAndDevelopment ) })
			table.Append([]string{ "", "Operating Expenses", ac.FormatMoney( fin.OperatingExpense ) })
			table.Append([]string{ "", "Current Assets", ac.FormatMoney( fin.CurrentAssets ) })
			table.Append([]string{ "", "Total Assets", ac.FormatMoney( fin.TotalAssets ) })
			table.Append([]string{ "", "Total Liabilities", ac.FormatMoney( fin.TotalLiabilities ) })
			table.Append([]string{ "", "Current Cash", ac.FormatMoney( fin.CurrentCash ) })
			table.Append([]string{ "", "Current Debt", ac.FormatMoney( fin.CurrentDebt ) })
			table.Append([]string{ "", "Total Cash", ac.FormatMoney( fin.TotalCash ) })
			table.Append([]string{ "", "Shareholder Equity", ac.FormatMoney( fin.ShareholderEquity ) })
			table.Append([]string{ "", "Cash Change", ac.FormatMoney( fin.CashChange ) })
			table.Append([]string{ "", "Cash Flow", ac.FormatMoney( fin.CashFlow ) })
			table.Append([]string{ "", "Operating Gains Losses", ac.FormatMoney( fin.OperatingGainsLosses ) })
			table.Append([]string{ "---------", "---------------------------", "------------------" })
		}
		println("")
		table.Render()
		println("")
	},
}


var symbolEarningsCmd = &cobra.Command{
	Use: "earnings",
	Short: "Get the earnings for a symbol",
	Run: func(cmd *cobra.Command, args []string) {
		if len(metaCmdOps.Symbol) == 0 {
			log.Fatal("ERROR: Please enter a symbol. eg: --symbol=MSFT")
		}
		var earnings []schemas.Earning
		getUrl( "https://api.polygon.io/v1/meta/symbols/"+metaCmdOps.Symbol+"/earnings", &earnings )
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth( 60 )
		table.SetBorder(false)
		table.SetColumnAlignment([]int{ tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT })
		table.SetHeader([]string{"Report Date", "Key", "Value"})
		for _, fin := range earnings {
			table.Append([]string{ fin.EPSReportDateStr, "Fiscal Period", fin.FiscalPeriod })
			table.Append([]string{ "", "Fiscal End Date", fin.FiscalEndDate })
			table.Append([]string{ "", "Actual EPS", fmt.Sprintf("%.6f", fin.ActualEPS ) })
			table.Append([]string{ "", "Consensus EPS", fmt.Sprintf("%.6f", fin.ConsensusEPS ) })
			table.Append([]string{ "", "Estimated EPS", fmt.Sprintf("%.6f", fin.EstimatedEPS ) })
			table.Append([]string{ "", "Announce Time", fin.AnnounceTime })
			table.Append([]string{ "", "Number Of Estimates", strconv.Itoa( fin.NumberOfEstimates ) })
			table.Append([]string{ "", "EPS Surprise Dollar", fmt.Sprintf("%.6f", fin.EPSSurpriseDollar ) })
			table.Append([]string{ "", "Year Ago", fmt.Sprintf("%.6f", fin.YearAgo ) })
			table.Append([]string{ "", "Year Ago Change Percent", strconv.Itoa( fin.YearAgoChangePercent ) })
			table.Append([]string{ "", "Estimated Change Percent", strconv.Itoa( fin.EstimatedChangePercent ) })

			table.Append([]string{ "---------", "---------------------------", "------------------" })
		}
		println("")
		table.Render()
		println("")
	},
}



var symbolDividendsCmd = &cobra.Command{
	Use: "dividends",
	Short: "Get the dividends for a symbol",
	Run: func(cmd *cobra.Command, args []string) {
		if len(metaCmdOps.Symbol) == 0 {
			log.Fatal("ERROR: Please enter a symbol. eg: --symbol=MSFT")
		}
		var dividends []schemas.Dividend
		getUrl( "https://api.polygon.io/v1/meta/symbols/"+metaCmdOps.Symbol+"/dividends", &dividends )
		table := tablewriter.NewWriter(os.Stdout)
		table.SetColWidth( 60 )
		table.SetBorder(false)
		table.SetColumnAlignment([]int{ tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT })
		table.SetHeader([]string{"Execution Date", "Key", "Value"})
		for _, fin := range dividends {
			table.Append([]string{ fin.ExDate, "Execution Date", fin.ExDate })
			table.Append([]string{ "", "Type", fin.Type })
			table.Append([]string{ "", "Payment Date", fin.PaymentDate })
			table.Append([]string{ "", "Record Date", fin.RecordDate })
			table.Append([]string{ "", "Declared Date", fin.DeclaredDate })
			table.Append([]string{ "", "Amount", fmt.Sprintf("%.6f", fin.Amount ) })
			table.Append([]string{ "", "Qualified", fin.Qualified })
			table.Append([]string{ "", "Flag", fin.Flag })
			table.Append([]string{ "---------", "---------------------------", "------------------" })
		}
		println("")
		table.Render()
		println("")
	},
}


func init() {
	RootCmd.AddCommand(exchangesCmd)
	RootCmd.AddCommand(cryptoExchangesCmd)
	RootCmd.AddCommand(symbolTypesCmd)
	RootCmd.AddCommand(symbolFinancialsCmd)
	RootCmd.AddCommand(symbolEarningsCmd)
	RootCmd.AddCommand(symbolDividendsCmd)
	symbolFinancialsCmd.Flags().StringVar(&metaCmdOps.Symbol, "symbol", "", "Symbol eg: MSFT")
	symbolEarningsCmd.Flags().StringVar(&metaCmdOps.Symbol, "symbol", "", "Symbol eg: MSFT")
	symbolDividendsCmd.Flags().StringVar(&metaCmdOps.Symbol, "symbol", "", "Symbol eg: MSFT")
}

func prettyprint(b []byte) ([]byte, error) {
    var out bytes.Buffer
    err := json.Indent(&out, b, "", "  ")
    return out.Bytes(), err
}

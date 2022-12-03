package main
import (
    "context"
    "fmt"
    "log"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)
var spreadsheetID = "1ythUvqWfra0Ihaw0D_PS7R5feJE8bCMPRSfMdlq-uts"
func main() {
    credential := option.WithCredentialsFile("credentials/secret.json")
    srv, err := sheets.NewService(context.TODO(), credential)
    if err != nil {
        log.Fatal(err)
    }

	// 読み込め
    readRange := "シート1!A1:B3"
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
    if err != nil {
        log.Fatalln(err)
    }
    if len(resp.Values) == 0 {
        log.Fatalln("data not found")
    }
    for _, row := range resp.Values {
        fmt.Printf("%s, %s\n", row[0], row[1])
    }

	// 書き込み
	ctx := context.Background()

	// 更新範囲の指定
    valueRange := "A1:B2"

	 // 更新値の指定
	 rb := &sheets.ValueRange{
        MajorDimension: "ROWS",
        Values: [][]interface{}{
            []interface{}{"123", "hoge"},
            []interface{}{"1.23", "=B1&B1"},
        },
    }

    srv.Spreadsheets.Values.Update(spreadsheetID, valueRange, rb).ValueInputOption("USER_ENTERED").Context(ctx).Do()

}
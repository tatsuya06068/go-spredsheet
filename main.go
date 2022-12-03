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
}
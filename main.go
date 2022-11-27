package main
import (
    "context"
    "fmt"
    "log"
    "google.golang.org/api/option"
    "google.golang.org/api/sheets/v4"
)
var spreadsheetID = "https://docs.google.com/spreadsheets/d/1ythUvqWfra0Ihaw0D_PS7R5feJE8bCMPRSfMdlq-uts/edit"
func main() {
    credential := option.WithCredentialsFile("credentials/secret.json")
    srv, err := sheets.NewService(context.TODO(), credential)
    if err != nil {
		log.Println("1")
        log.Fatal(err)
    }
    // readRange := "シート1!A1:B3"
    resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, "A1").Do()
    if err != nil {
		log.Println("2")
        log.Fatalln(err)
    }
    if len(resp.Values) == 0 {
        log.Fatalln("data not found")
    }
    for _, row := range resp.Values {
        fmt.Printf("%s, %s\n", row[0], row[1])
    }
}
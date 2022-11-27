# Create SpredSheets
- URLから`Google Sheets API`を有効にする
  - https://console.cloud.google.com/apis/api/sheets.googleapis.com/metrics?project=eastern-lattice-256111
- client作成
  - https://pkg.go.dev/google.golang.org/api/sheets/v4













<br>
  - クエリ調査
  <br>
  
  ```sq;
    SELECT 
        count(*)
    FROM 
        resumes
    WHERE 
        NOT EXISTS
        (
            SELECT 
                * 
            FROM 
                TableA AS a
            WHERE 
                job_id = ''

            UNION 
  
            SELECT 
                * 
            FROM 
                TableB ASas b
            WHERE 
                job_id = ''  
        );
  ```

```
    SELECT 
        count(*)
    FROM 
        resumes
    WHERE 
        NOT EXISTS
    (
    SELECT 
        * 
    FROM 
        TableA AS a
    WHERE 
        job_id = ''
    )
    OR
    NOT EXISTS
    (
    SELECT 
        * 
    FROM 
    TableB as b
    WHERE 
        job_id = ''  
    );
```
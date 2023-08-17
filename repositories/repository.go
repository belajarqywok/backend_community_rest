package repositories


import (
    "time"
    "context"

    "google.golang.org/api/sheets/v4"

    model "backend_community_rest/models"
    exception "backend_community_rest/exceptions"
    config "backend_community_rest/configurations"
)


func AddJoinRequestRepo(RequestData model.JoinRequest) (bool) {
    // Get Environment Variable
    env := config.EnvironmentVariables()

    /* 
     * Spreadsheet ID and Key
     * Example :
     *   - https://docs.google.com/spreadsheets/d/<SPREADSHEETID>/edit#gid=<SHEETID>
     *
     */
    sheetId := env.SpreadsheetId
    spreadsheetKey := env.SpreadsheetKey

    // Create API Context
    ctx := context.Background()

    // Database Configuration
    srv := config.SpreadsheetDatabase()

    // Convert Sheet ID To Sheet Name.
    sheetid_conv, err := srv.Spreadsheets.Get(
        spreadsheetKey,

    ).Fields(
        "sheets(properties(sheetId,title))", 
    ).Do()

    exception.TryCatchError(err)

    if err != nil || sheetid_conv.HTTPStatusCode != 200 {
        exception.TryCatchError(err)
        return false
    }

    sheetName := ""
    for _, v := range sheetid_conv.Sheets {
        prop := v.Properties

        if prop.SheetId == int64(sheetId) {
            sheetName = prop.Title
            break
        }

    }

    // Append Value To The Sheet.
    row := &sheets.ValueRange {
        Values: [][]interface{}{
            {
                time.Now(),
                RequestData.UsernameGithub,
                "not accept",
            },
        },
    }

    // Append Data To Sheet
    append, err := srv.Spreadsheets.Values.Append(
        spreadsheetKey,
        sheetName,
        row,
    ).ValueInputOption(
        "USER_ENTERED",
    ).InsertDataOption(
        "INSERT_ROWS",
    ).Context(ctx).Do()

    exception.TryCatchError(err)

    if append.HTTPStatusCode != 200 {
        exception.TryCatchError(err)
        return false
    }


    return true
}

package configurations

import (
        "context"
        b64 "encoding/base64"

        "golang.org/x/oauth2/google"
        "google.golang.org/api/option"
        "google.golang.org/api/sheets/v4"

        exception "backend_community_rest/exceptions"
)


func SpreadsheetDatabase() (*sheets.Service) {
    // Environment Variables
    env := EnvironmentVariables()

    // Create API Context
    ctx := context.Background()

    // Get Bytes From Base64 Encoded Google Service Accounts Key
    credBytes, err := b64.StdEncoding.DecodeString(env.KeyJsonBase64)
    exception.TryCatchError(err)

    // Authenticate and Get Configuration
    config, err := google.JWTConfigFromJSON(credBytes, "https://www.googleapis.com/auth/spreadsheets")
    exception.TryCatchError(err)

    // Create Client With Config and Context
    client := config.Client(ctx)

    // Create New Service Using Client
    srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
    exception.TryCatchError(err)

    return srv
}
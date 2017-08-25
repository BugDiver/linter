package upload

import (
	"fmt"
	"os"

	"github.com/bugdiver/linter/version"
	"github.com/stacktic/dropbox"
)

// Upload to dropbox
func Upload() error {
	var err error
	var db *dropbox.Dropbox

	var clientID, clientSecret string
	var token string

	clientID = os.Getenv("DB_KEY")
	clientSecret = os.Getenv("DB_SECRET")
	token = os.Getenv("DB_TOKEN")

	db = dropbox.NewDropbox()

	db.SetAppInfo(clientID, clientSecret)
	db.SetAccessToken(token)

	var e *dropbox.Entry
	pkg := fmt.Sprintf("spider%s.pkg", version.GetVersion())
	if e, err = db.UploadFile("./deploy/"+pkg, "macOS/"+pkg, true, ""); err != nil {
		return err
	}
	fmt.Printf("Folder %v successfully created\n", e)
	return nil
}

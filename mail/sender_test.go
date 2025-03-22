package mail

import (
	"testing"

	"github.com/osisupermoses/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGamil(t *testing.T) {
	// this checks for the "-short" flag in the command and if present skips this test
	if testing.Short() {
		t.Skip()
	}
	
	config, err := util.LoadConfig("..") // ".." means the parent folder (where the app.env file is located)
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
	<h1>Hello world<h1>
	<p>This is a test message from <a href="https://kwenchoandco.com">Tech School</a></p>
	`

	to := []string{"daraprovidence@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}

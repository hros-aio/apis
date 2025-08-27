package keys

import "fmt"

func AuthSessionId(sessionId string) string {
	return fmt.Sprintf("session_%s", sessionId)
}

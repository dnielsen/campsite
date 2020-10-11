package service

import "encoding/base64"

func encodeCursor(cursor []byte) string {
	return base64.StdEncoding.EncodeToString(cursor)
}

func decodeCursor(encodedCursor string) ([]byte, error) {
	decodedCursor, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return nil, err
	}
	return decodedCursor, nil
}
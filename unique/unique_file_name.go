package unique

import "github.com/google/uuid"

// generate a unique file name
func GenerateUniqueFileName() string {
	id := uuid.New()
	return id.String()
}

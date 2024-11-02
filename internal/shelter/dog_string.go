package shelter

import "fmt"

func (d Dog) String() string {
	return fmt.Sprintf(`
ID=%v,
NAME=%v,
BREED=%v,
CREATED_AT=%v,
UPDATED_AT=%v
`,
		d.ID,
		d.Name,
		d.Breed,
		d.CreatedAt.Time.String(),
		d.UpdatedAt.Time.String(),
	)
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ValueChange struct {
	// Represents the previous value of the element.
	From string `json:"From"`
	// Represents the current value of the element.
	To string `json:"To"`
}

func (o *ValueChange) GetFrom() string {
	if o == nil {
		return ""
	}
	return o.From
}

func (o *ValueChange) GetTo() string {
	if o == nil {
		return ""
	}
	return o.To
}

// A SchemaDiff represents a diff of two Schemas.
type SchemaDiff struct {
	// Holds every addition change in the diff.
	Additions []string `json:"additions"`
	// Holds every deletion change in the diff.
	Deletions []string `json:"deletions"`
	// Holds every modification change in the diff.
	Modifications map[string]ValueChange `json:"modifications"`
}

func (o *SchemaDiff) GetAdditions() []string {
	if o == nil {
		return []string{}
	}
	return o.Additions
}

func (o *SchemaDiff) GetDeletions() []string {
	if o == nil {
		return []string{}
	}
	return o.Deletions
}

func (o *SchemaDiff) GetModifications() map[string]ValueChange {
	if o == nil {
		return map[string]ValueChange{}
	}
	return o.Modifications
}

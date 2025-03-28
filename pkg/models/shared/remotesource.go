// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// RemoteSource - Remote source configuration
type RemoteSource struct {
	Inputs []RemoteDocument `json:"inputs"`
	// A document hosted in the registry
	Output   RemoteDocument   `json:"output"`
	Overlays []RemoteDocument `json:"overlays,omitempty"`
}

func (o *RemoteSource) GetInputs() []RemoteDocument {
	if o == nil {
		return []RemoteDocument{}
	}
	return o.Inputs
}

func (o *RemoteSource) GetOutput() RemoteDocument {
	if o == nil {
		return RemoteDocument{}
	}
	return o.Output
}

func (o *RemoteSource) GetOverlays() []RemoteDocument {
	if o == nil {
		return nil
	}
	return o.Overlays
}

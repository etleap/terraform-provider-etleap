// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RetractEventRequest struct {
	// Unique identifier of this batch.
	ExternalBatchID *string `json:"externalBatchId,omitempty"`
}

func (o *RetractEventRequest) GetExternalBatchID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalBatchID
}

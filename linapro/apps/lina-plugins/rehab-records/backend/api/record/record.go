// This file defines the record API contract implemented by the rehab-records plugin.

package record

import (
	"context"

	"lina-plugin-rehab-records/backend/api/record/v1"
)

// IRecordV1 declares training record endpoints.
type IRecordV1 interface {
	List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
}

// This file defines the summary API contract implemented by the rehab-records plugin.

package summary

import (
	"context"

	"lina-plugin-rehab-records/backend/api/summary/v1"
)

// ISummaryV1 declares training summary endpoints.
type ISummaryV1 interface {
	Daily(ctx context.Context, req *v1.DailyReq) (res *v1.DailyRes, err error)
}

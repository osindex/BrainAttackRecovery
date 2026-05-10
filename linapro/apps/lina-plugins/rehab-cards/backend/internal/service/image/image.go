// Package image implements a safe Wikimedia Commons image reverse proxy used by
// the rehab-cards plugin. The proxy keeps the patient H5 inside the same origin
// even when commons.wikimedia.org is unreachable from the patient device.

package image

import (
	"context"
	"errors"
	"net/url"
	"regexp"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
)

// Service defines the Wikimedia image proxy contract.
type Service interface {
	// FetchWikimedia downloads the requested file from Wikimedia Commons
	// and returns its content type and bytes for streaming through the host response.
	FetchWikimedia(ctx context.Context, fileName string) (contentType string, body []byte, err error)
}

// Ensure serviceImpl implements Service.
var _ Service = (*serviceImpl)(nil)

// serviceImpl implements Service using the GoFrame HTTP client with a hardened timeout.
type serviceImpl struct {
	httpClient *gclient.Client
}

// New creates and returns an image proxy service instance.
func New() Service {
	client := g.Client()
	client.SetTimeout(15 * time.Second)
	client.SetHeader("User-Agent", "BrainAttackRecovery/0.1 (rehab cards proxy; contact local admin)")
	return &serviceImpl{httpClient: client}
}

// safeFileNamePattern restricts proxied file names to a small, safe character set.
var safeFileNamePattern = regexp.MustCompile(`^[A-Za-z0-9._\-()\s]+\.(jpg|jpeg|png|gif|webp)$`)

// errFileNameNotAllowed is returned when the requested file name fails the proxy allowlist.
var errFileNameNotAllowed = errors.New("file name not allowed by image proxy")

// FetchWikimedia downloads the requested file from Wikimedia Commons.
func (s *serviceImpl) FetchWikimedia(ctx context.Context, fileName string) (string, []byte, error) {
	if !safeFileNamePattern.MatchString(fileName) {
		return "", nil, errFileNameNotAllowed
	}
	target := "https://commons.wikimedia.org/wiki/Special:FilePath/" + url.PathEscape(fileName)
	response, err := s.httpClient.Get(ctx, target)
	if err != nil {
		return "", nil, err
	}
	defer func() { _ = response.Close() }()
	if response.StatusCode >= 400 {
		return "", nil, errors.New("wiki image upstream not ok")
	}
	body := response.ReadAll()
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}
	return contentType, body, nil
}

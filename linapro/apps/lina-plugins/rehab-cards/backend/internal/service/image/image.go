// Package image implements a safe Wikimedia Commons image reverse proxy used by
// the rehab-cards plugin. The proxy keeps the patient H5 inside the same origin
// even when commons.wikimedia.org is unreachable from the patient device.

package image

import (
	"context"
	"crypto/md5"
	"encoding/hex"
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
	client.SetTimeout(20 * time.Second)
	client.SetHeader("User-Agent", "BrainAttackRecovery/0.1 (rehab cards proxy; contact local admin)")
	client.SetHeader("Accept", "image/*,*/*;q=0.8")
	return &serviceImpl{httpClient: client}
}

// safeFileNamePattern restricts proxied file names to a small, safe character set.
var safeFileNamePattern = regexp.MustCompile(`^[A-Za-z0-9._\-()\s]+\.(jpg|jpeg|png|gif|webp|JPG|JPEG|PNG|GIF|WEBP)$`)

// errFileNameNotAllowed is returned when the requested file name fails the proxy allowlist.
var errFileNameNotAllowed = errors.New("file name not allowed by image proxy")

// FetchWikimedia downloads the requested file directly from upload.wikimedia.org
// using the canonical static asset path derived from the file name's MD5 hash.
// This avoids the Special:FilePath redirect which is unreachable on some networks.
func (s *serviceImpl) FetchWikimedia(ctx context.Context, fileName string) (string, []byte, error) {
	if !safeFileNamePattern.MatchString(fileName) {
		return "", nil, errFileNameNotAllowed
	}
	target := uploadURLForCommonsFile(fileName)
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

// uploadURLForCommonsFile returns the canonical upload.wikimedia.org URL for a
// Wikimedia Commons file. Wikimedia Commons hashes the file name with MD5 and
// uses the first one and two hex characters as directory prefixes.
func uploadURLForCommonsFile(fileName string) string {
	normalized := normalizeCommonsFileName(fileName)
	hash := md5.Sum([]byte(normalized))
	hexHash := hex.EncodeToString(hash[:])
	return "https://upload.wikimedia.org/wikipedia/commons/" + hexHash[:1] + "/" + hexHash[:2] + "/" + url.PathEscape(normalized)
}

// normalizeCommonsFileName converts spaces to underscores so the URL hash matches
// the canonical Wikimedia storage layout.
func normalizeCommonsFileName(fileName string) string {
	out := make([]byte, 0, len(fileName))
	for _, r := range fileName {
		if r == ' ' {
			out = append(out, '_')
			continue
		}
		out = append(out, string(r)...)
	}
	return string(out)
}

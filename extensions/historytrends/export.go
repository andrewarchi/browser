// Copyright (c) 2021 Andrew Archibald
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package historytrends parses Chromium browsing history exports from
// the History Trends Unlimited extension.
package historytrends

import (
	"fmt"
	"time"

	"github.com/andrewarchi/browser/chrome"
)

// Export contains browsing history exported from History Trends
// Unlimited. The export time is in the local timezone at the time the
// export was created. When the local timezone cannot be determined, the
// location is incorrectly set to UTC.
type Export struct {
	Filename   string // filename of tsv within zip or as given
	Type       ExportType
	ExportTime time.Time
	Visits     []Visit
}

// Visit is a page visit in browsing history. URL and visit time
// combined are unique; no two visits have the same URL and visit time.
// The visit time is in UTC, not local time.
type Visit struct {
	URL        string
	VisitTime  time.Time // UTC
	Transition chrome.PageTransition
	PageTitle  string
}

// ExportType is the format of export.
type ExportType uint8

// Values for ExportType:
const (
	_ ExportType = iota
	AnalysisExport
	ArchivedExport
)

func (typ ExportType) String() string {
	switch typ {
	case AnalysisExport:
		return "analysis"
	case ArchivedExport:
		return "archived"
	default:
		return fmt.Sprintf("export(%d)", typ)
	}
}

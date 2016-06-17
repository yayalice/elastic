// Copyright 2012-2015 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

import (
	"net/url"
	"strings"

	"gopkg.in/olivere/elastic.v2/uritemplates"
)

// AnalyzeService analyzes text.
type AnalyzeService struct {
	client        *Client
	indices       []string
	pretty        bool
	timeout       string
	masterTimeout string
	analyzer      string
	tokenizer     string
	tokenFilters  []string
	charFilters   []string
	body          string
}

// NewAnalyzeService returns a new AnalyzeService.
func NewAnalyzeService(client *Client) *AnalyzeService {
	return &AnalyzeService{client: client}
}

// Index sets the name of the index to use for analyzing.
func (s *AnalyzeService) Index(index string) *AnalyzeService {
	if s.indices == nil {
		s.indices = make([]string, 0)
	}
	s.indices = append(s.indices, index)
	return s
}

// Indices sets the names of the indices to use for analyzing.
func (s *AnalyzeService) Indices(indices ...string) *AnalyzeService {
	if s.indices == nil {
		s.indices = make([]string, 0)
	}
	s.indices = append(s.indices, indices...)
	return s
}

// Timeout the explicit operation timeout, e.g. "5s".
func (s *AnalyzeService) Timeout(timeout string) *AnalyzeService {
	s.timeout = timeout
	return s
}

// MasterTimeout specifies the timeout for connection to master.
func (s *AnalyzeService) MasterTimeout(masterTimeout string) *AnalyzeService {
	s.masterTimeout = masterTimeout
	return s
}

// Analyzer specifies the analyzer to use.
func (s *AnalyzeService) Analyzer(analyzer string) *AnalyzeService {
	s.analyzer = analyzer
	return s
}

// Tokenizer sets the tokenizer to use for analyzing.
func (s *AnalyzeService) Tokenizer(tokenizer string) *AnalyzeService {
	s.tokenizer = tokenizer
	return s
}

// TokenFilters sets the names of the token_filters to use for analyzing.
func (s *AnalyzeService) TokenFilters(filters ...string) *AnalyzeService {
	if s.tokenFilters == nil {
		s.tokenFilters = make([]string, 0)
	}
	s.tokenFilters = append(s.tokenFilters, filters...)
	return s
}

// CharFilters sets the names of the char_filters to use for analyzing.
func (s *AnalyzeService) CharFilters(filters ...string) *AnalyzeService {
	if s.charFilters == nil {
		s.charFilters = make([]string, 0)
	}
	s.charFilters = append(s.charFilters, filters...)
	return s
}

// Body specifies the text to analyze.
func (s *AnalyzeService) Body(body string) *AnalyzeService {
	s.body = body
	return s
}

// Pretty indicates that the JSON response be indented and human readable.
func (s *AnalyzeService) Pretty(pretty bool) *AnalyzeService {
	s.pretty = pretty
	return s
}

// Do executes the operation.
func (s *AnalyzeService) Do() (*AnalyzeResult, error) {
	// Build url
	path := "/"

	// Indices part
	var indexPart []string
	for _, index := range s.indices {
		index, err := uritemplates.Expand("{index}", map[string]string{
			"index": index,
		})
		if err != nil {
			return nil, err
		}
		indexPart = append(indexPart, index)
	}
	path += strings.Join(indexPart, ",")

	// Analyze
	path += "/_analyze"

	params := make(url.Values)
	if s.pretty {
		params.Set("pretty", "1")
	}
	if s.masterTimeout != "" {
		params.Set("master_timeout", s.masterTimeout)
	}
	if s.timeout != "" {
		params.Set("timeout", s.timeout)
	}
	if s.analyzer != "" {
		params.Set("analyzer", s.analyzer)
	}

	// Setup HTTP request body
	var body = s.body

	// Get response
	res, err := s.client.PerformRequest("GET", path, params, body)
	if err != nil {
		return nil, err
	}

	ret := new(AnalyzeResult)
	if err := s.client.decoder.Decode(res.Body, ret); err != nil {
		return nil, err
	}
	return ret, nil
}

// -- Result of a analyze request.

// AnalyzeResult is the outcome of analyzing text.
type AnalyzeResult struct {
	Tokens []Token `json:"tokens"`
}

// Token is a single token for a broken down text.
type Token struct {
	Token       string `json:"token"`
	StartOffset int64  `json:"start_offset"`
	EndOffset   int64  `json:"end_offset"`
	Position    int64  `json:"position"`
	Type        string `json:"type"`
}

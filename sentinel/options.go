/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sentinel

import (
	"context"

	"github.com/cloudwego-contrib/cwgo-pkg/opensergo/sentinel/kxadapter"
)

type Option = kxadapter.Option

func DefaultBlockFallback(ctx context.Context, req, resp interface{}, blockErr error) error {
	return blockErr
}

func DefaultResourceExtract(ctx context.Context, req, resp interface{}) string {
	return kxadapter.DefaultResourceExtract(ctx, req, resp)
}

// WithResourceExtract sets the resource extractor
func WithResourceExtract(f func(ctx context.Context, req, resp interface{}) string) Option {
	return kxadapter.WithResourceExtract(f)
}

// WithBlockFallback sets the fallback handler
func WithBlockFallback(f func(ctx context.Context, req, resp interface{}, blockErr error) error) Option {
	return kxadapter.WithBlockFallback(f)
}

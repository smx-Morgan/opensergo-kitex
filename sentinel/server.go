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
	"github.com/cloudwego-contrib/cwgo-pkg/opensergo/sentinel/kxadapter"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

// SentinelServerMiddleware returns new server.Middleware
// Default resource name is {service's name}:{method}
// Default block fallback is returning blockError
// Define your own behavior by setting serverOptions
func SentinelServerMiddleware(opts ...Option) func(endpoint.Endpoint) endpoint.Endpoint {
	return kxadapter.SentinelServerMiddleware(opts...)
}

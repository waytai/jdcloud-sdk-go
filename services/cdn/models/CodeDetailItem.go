// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type CodeDetailItem struct {

    /*  (Optional) */
    TimeStamp int64 `json:"timeStamp"`

    /*  (Optional) */
    Ok float64 `json:"ok"`

    /*  (Optional) */
    BadGateway float64 `json:"badGateway"`

    /*  (Optional) */
    BadRequest float64 `json:"badRequest"`

    /*  (Optional) */
    Forbidden float64 `json:"forbidden"`

    /*  (Optional) */
    Found float64 `json:"found"`

    /*  (Optional) */
    GatewayTimeout float64 `json:"gatewayTimeout"`

    /*  (Optional) */
    InternalServerError float64 `json:"internalServerError"`

    /*  (Optional) */
    MovedPermanently float64 `json:"movedPermanently"`

    /*  (Optional) */
    NotFound float64 `json:"notFound"`

    /*  (Optional) */
    NotModified float64 `json:"notModified"`

    /*  (Optional) */
    PartialContent float64 `json:"partialContent"`

    /*  (Optional) */
    RequestedRangeNotSuitable float64 `json:"requestedRangeNotSuitable"`

    /*  (Optional) */
    Other float64 `json:"other"`

    /*  (Optional) */
    ServiceUnavailable float64 `json:"serviceUnavailable"`
}

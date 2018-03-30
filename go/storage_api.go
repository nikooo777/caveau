/*
 * Swagger Petstore
 *
 * Backup solution for binaries
 *
 * API version: 1.0.0
 * Contact: niko@lbry.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"net/http"
)

func StoreWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

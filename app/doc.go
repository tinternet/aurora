// Package app is the controller layer holding market APIs and data processors.
// It fetches raw trade data from the market APIs, subscribes for updates and
// is responsible to push the raw data to the registered processors.
// Another basic function is to pull metrics from the market APIs.
package app

/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package client contains the implementation of the client side communication with the
Kubernetes master. The Client class provides methods for reading, creating, updating,
and deleting pods, replication controllers, services, and minions.

Most consumers should use the Config object to create a Client:

    config := &client.Config{
      Host:     "http://localhost:8080",
      Username: "test",
      Password: "password",
    }
    client, err := client.New(&config)
    if err != nil {
      // handle error
    }
    client.Pods(ns).List()

More advanced consumers may wish to provide their own transport via a http.RoundTripper:

    config := &client.Config{
      Host:      "https://localhost:8080",
      Transport: oauthclient.Transport(),
    }
    client, err := client.New(&config)

The RESTClient type implements the Kubernetes API conventions (see `docs/api-conventions.md`)
for a given API path and is intended for use by consumers implementing their own Kubernetes
compatible APIs.
*/
package client

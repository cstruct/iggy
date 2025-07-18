// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package tcp_test

import (
	. "github.com/apache/iggy/foreign/go/contracts"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LOGIN FEATURE:", func() {
	When("user is already logged in", func() {
		Context("and tries to log with correct data", func() {
			client := createAuthorizedConnection()
			user, err := client.LoginUser("iggy", "iggy")

			itShouldNotReturnError(err)
			itShouldReturnUserId(user, 1)
		})

		Context("and tries to log with invalid credentials", func() {
			client := createAuthorizedConnection()
			user, err := client.LoginUser("incorrect", "random")

			itShouldReturnError(err)
			itShouldNotReturnUser(user)
		})
	})

	When("user is not logged in", func() {
		Context("and tries to log with correct data", func() {
			client := createClient()
			user, err := client.LoginUser("iggy", "iggy")

			itShouldNotReturnError(err)
			itShouldReturnUserId(user, 1)
		})

		Context("and tries to log with invalid credentials", func() {
			client := createClient()
			user, err := client.LoginUser("incorrect", "random")

			itShouldReturnError(err)
			itShouldNotReturnUser(user)
		})
	})
})

func itShouldReturnUserId(user *IdentityInfo, id uint32) {
	It("should return user id", func() {
		Expect(user.UserId).To(Equal(id))
	})
}

func itShouldNotReturnUser(user *IdentityInfo) {
	It("should return user id", func() {
		Expect(user).To(BeNil())
	})
}

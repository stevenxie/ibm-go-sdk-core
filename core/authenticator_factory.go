package core

/**
 * Copyright 2019 IBM All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"fmt"
)

// GetAuthenticatorFromEnvironment: instantiates an Authenticator using service properties
// retrieved from external config sources.
func GetAuthenticatorFromEnvironment(credentialKey string) (authenticator Authenticator, err error) {
	properties, err := GetServiceProperties(credentialKey)
	if properties == nil || len(properties) == 0 {
		return
	}

	// Default the authentication type to IAM if not specified.
	authType := properties[PROPNAME_AUTH_TYPE]
	if authType == "" {
		authType = AUTHTYPE_IAM
	}

	// Construct the appropriate authenticator according to the authentication type.
	switch authType {
	case AUTHTYPE_BASIC:
		authenticator, err = NewBasicAuthenticatorFromMap(properties)
	case AUTHTYPE_BEARER_TOKEN:
		authenticator, err = NewBearerTokenAuthenticatorFromMap(properties)
	case AUTHTYPE_IAM:
		authenticator, err = NewIamAuthenticatorFromMap(properties)
	case AUTHTYPE_CP4D:
		authenticator, err = NewCloudPakForDataAuthenticatorFromMap(properties)
	case AUTHTYPE_NOAUTH:
		authenticator, err = NewNoAuthAuthenticator()
	default:
		err = fmt.Errorf(ERRORMSG_AUTHTYPE_UNKNOWN, authType)
	}

	return
}

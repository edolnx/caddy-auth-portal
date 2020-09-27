
## OAuth 2.0 and OpenID Connect (OIDC) Authentication Backend

### OAuth 2.0 Flow

Generally, a developer must create an "Application" with and identity provider.
e.g. Google, Okta, Azure, etc., and obtain OAuth 2.0 "Client ID" and
"Client Secret".

The authentication flow begins with obtaining "Authorization Code" from an
identity provider.

* What does the portal send to Okta?
  - Client ID
  - Redirect URI
  - Response Type
  - Scope

* What does the portal receive from Okta?
  - Authorization Code

Once, the portal has the "Authorization Code", it could get "Access Token"
to access the user's data at the identity provider.

* What does the portal send to Okta?
  - Authorization Code
  - Client ID
  - Client Secret

* What does Okta respond with?
  - Access Token

* What could the portal use "Access Token" for?
  -  Make API calls to obtain user information

The OpenID Connect (OIDC) adds login and profile information about the person
who is logged in. The differences between standard OAuth2.0 flow are:

1. In the initial request, a specific scope of `openid` is used
2. In the final exchange the Client receives both an "Access Token" and an "ID Token" (JWT Token).

References:
* [Mozilla - OIDC in a nutshell](https://infosec.mozilla.org/guidelines/iam/openid_connect.html#oidc-in-a-nutshell)

[:arrow_up: Back to Top](#table-of-contents)

### OAuth 2.0 Authorization Servers and Identity Providers

#### Okta

Create an "Application," by browsing to "Applications" and clicking
"Add Application" button.

For a website, the choice is "Web".

![Okta Developer - New App Choice](./assets/docs/images/oauth2_okta_new_app_choice.png)

Provided your application is running on localhost port 8443, provide the following
information:

Base URI entries:
* `https://myapp.localdomain.local:8443/`
* `https://localhost:8443/`

Login redirect URIs:
* `https://myapp.localdomain.local:8443/auth/oauth2/okta/authorization-code-callback`
* `https://localhost:8443/auth/oauth2/okta/authorization-code-callback`

Logout redirect URIs:
* `https://myapp.localdomain.local:8443/auth/logout`
* `https://localhost:8443/auth/logout`

Group Assignments:
* Everyone
* Viewer
* Editor
* Administrator

Grant type allowed:
* Client acting on behalf of itself
  - Client Credentials: No
* Client acting on behalf of a user
  - Authorization Code: Yes
  - Refresh Token: No
  - Implicit (Hybrid): No


![Okta Developer - New App Setup](./assets/docs/images/oauth2_okta_new_app.png)

Review the newly created application.

![Okta Developer - Settings - General](./assets/docs/images/oauth2_okta_app_settings_01.png)

Store the credentials securely.

![Okta Developer - Settings - Client Credentials](./assets/docs/images/oauth2_okta_app_settings_02.png)

Review default Sign On Policy.

![Okta Developer - Settings - Sign On](./assets/docs/images/oauth2_okta_app_settings_03.png)

By default, the default Authorization Server has no `groups` scope.

Therefore, browse to "API", "Authorization Servers" and select "default".

![Okta Developer - API](./assets/docs/images/okta_configure_scope_01.png)

![Okta Developer - API](./assets/docs/images/okta_configure_scope_02.png)

Next, browse to "Scopes" and click "Add Scope".

Fill out the "Add Scope" form:
* Name: `groups`
* Description: `This allows the app to view your group memberships.`
* Check "Set as a default scope"
* Check "Include in public metadata"

![Okta Developer - Add Scope](./assets/docs/images/okta_configure_scope_03.png)

Next, browse to "Claims" and click "Add Claim".

Fill out the "Add Claim" form:
* Name: `groups`
* Value type: Groups
* Filter: Matches `.*` regex
* Include in: The "groups" scope

![Okta Developer - Add Scope](./assets/docs/images/okta_configure_scope_04.png)

Next, review [Okta OpenID Connect and OAuth 2.0 API - Get Started](https://developer.okta.com/docs/reference/api/oidc/#get-started).

The Caddyfile snipper for Okta OAuth 2.0 backend is as follows:

```
127.0.0.1, localhost {
  route /auth* {
    auth_portal {
      backends {
        okta_oauth2_backend {
          method oauth2
          realm okta
          provider okta
          domain dev-680653.okta.com
          client_id 0oa121qw81PJW0Tj34x7
          client_secret b3aJC5E59hU18YKC7Yca3994F4qFhWiAo_ZojanF
          server_id default
        }
```

Provided the Okta domain is `dev-680653.okta.com`, the authorization server is
`default`, and Client ID is `0oa121qw81PJW0Tj34x7`, check OpenID configuration:

```bash
curl -X GET "https://dev-680653.okta.com/oauth2/default/.well-known/openid-configuration?client_id=0oa121qw81PJW0Tj34x7" | jq
```

The expected output is:

```
{
  "issuer": "https://dev-680653.okta.com/oauth2/default",
  "authorization_endpoint": "https://dev-680653.okta.com/oauth2/default/v1/authorize",
  "token_endpoint": "https://dev-680653.okta.com/oauth2/default/v1/token",
  "userinfo_endpoint": "https://dev-680653.okta.com/oauth2/default/v1/userinfo",
  "registration_endpoint": "https://dev-680653.okta.com/oauth2/v1/clients/0oa121qw81PJW0Tj34x7",
  "jwks_uri": "https://dev-680653.okta.com/oauth2/default/v1/keys",
  "response_types_supported": [
    "code"
  ],
  "response_modes_supported": [
    "query",
    "fragment",
    "form_post",
    "okta_post_message"
  ],
  "grant_types_supported": [
    "authorization_code"
  ],
  "subject_types_supported": [
    "public"
  ],
  "id_token_signing_alg_values_supported": [
    "RS256"
  ],
  "scopes_supported": [
    "openid",
    "profile",
    "email",
    "address",
    "phone"
  ],
  "token_endpoint_auth_methods_supported": [
    "client_secret_basic"
  ],
  "claims_supported": [
    "iss",
    "ver",
    "sub",
    "aud",
    "iat",
    "exp",
    "jti",
    "auth_time",
    "amr",
    "idp",
    "nonce",
    "name",
    "nickname",
    "preferred_username",
    "given_name",
    "middle_name",
    "family_name",
    "email",
    "email_verified",
    "profile",
    "zoneinfo",
    "locale",
    "address",
    "phone_number",
    "picture",
    "website",
    "gender",
    "birthdate",
    "updated_at",
    "at_hash",
    "c_hash"
  ],
  "code_challenge_methods_supported": [
    "S256"
  ],
  "introspection_endpoint": "https://dev-680653.okta.com/oauth2/default/v1/introspect",
  "introspection_endpoint_auth_methods_supported": [
    "client_secret_basic"
  ],
  "revocation_endpoint": "https://dev-680653.okta.com/oauth2/default/v1/revoke",
  "revocation_endpoint_auth_methods_supported": [
    "client_secret_basic"
  ],
  "end_session_endpoint": "https://dev-680653.okta.com/oauth2/default/v1/logout",
  "request_parameter_supported": true,
  "request_object_signing_alg_values_supported": [
    "HS256",
    "HS384",
    "HS512"
  ]
}
```

The following returns a JSON Web Key Set (JWKS) that contains the public keys that
can be used to verify the signatures of tokens received from Okta's authorization server.

```bash
curl https://dev-680653.okta.com/oauth2/default/v1/keys | jq
```

The expected output is:

```bash
{
  "keys": [
    {
      "kty": "RSA",
      "alg": "RS256",
      "kid": "7NJ-w2OhocVRbyIT1HUY3pKaFTa8apJxSDe6hMpIRNQ",
      "use": "sig",
      "e": "AQAB",
      "n": "mRg7MIf9I6Mp2SRdPyyHNCuf7ca695bTOGEUXrqQq51YO4wqs_LJX6T2JqJTreR9vIV856shOFlRvb0ZjCLSxQu6phdXFj7V3Qycm2wyZkKkQ8v1s52cec5KyNe3IRchtMfBVAcK0PXFwUAw6fu7xPW-2B7yD2_Xop-03GkSi7RmH7y4xnCz8eetyNU11cQZu9_nXex3Eo9_poiyDHHHO24menmjA6rMLraR2JlBMeI7VTFSaTyqEQ_qtdo1wRD0umGqN6jfsL6XUyvmKEsulZxj7KWbqtl6VyMxrtT3KDnwHVJ7I5V68NsjAZCF7iwgVpNQeXbRKi7SpoTet48bvQ"
    },
    {
      "kty": "RSA",
      "alg": "RS256",
      "kid": "u4louF5gsEhI28_Oj6fQHcd6fViOT_P9GTedCQAiooY",
      "use": "sig",
      "e": "AQAB",
      "n": "phxaY-ZGuuCiHbprkz3MeQJwgIBD5TqEb1IhnjhRlNrhYjtnj-0g5YqaOc_qPJGc8uFaPgtnk_UXMFAmNEqSYnrItBW9uh4hR_JHSEyq4TMiW6-1aqvBPUzVjYUYFKzkRjwUSqMafJyMlQD1DqOQf1MjZLW2RnI5IjsC6YHb0LOtqa890CeuDhllr5iBr2Q9LcU9e8YnQk-vzjLoN8F1ooXp0MFJ52eZObdJq0Gl6rnDhmtmJykmz6niTcEPM5UdqsRyh2L2Mv7oFA5y_qG1WDnqLugWynh8Pc-q6BW3DMRcxN3HNFOAJrnfy5CiqHY1Cleql9f2agQCCtypdCqNow"
    }
  ]
}
```

[:arrow_up: Back to Top](#table-of-contents)

#### Google Identity Platform

References:
* [Google Identity Platform - Identity Platform - OpenID Connect](https://developers.google.com/identity/protocols/oauth2/openid-connect)

First, create new application, e.g. "My Gatekeeper".

![Google Identity Platform - Identity Platform - New Application](./assets/docs/images/oauth2_google_new_app.png)

After the creation of the app, you will land on Credentials page.

![Google Identity Platform - Identity Platform - Credentials](./assets/docs/images/oauth2_google_credentials.png)

Click "Configure Consent Screen" and select an appropriate option, e.g. "External".

![Google Identity Platform - Identity Platform - Consent Screen](./assets/docs/images/oauth2_google_consent_screen.png)

Next, provide the name for the application, e.g. "My Gatekeeper" and select appropriate support email.

![Google Identity Platform - Identity Platform - Consent Screen Configuration](./assets/docs/images/oauth2_google_consent_screen_config.png)

After configuring the consent screen you will see the following.

![Google Identity Platform - Identity Platform - Consent Screen Verification](./assets/docs/images/oauth2_google_consent_screen_verification.png)

Next, browse to "Credentials" and click "Create Credentials". Then, choose "OAuth client ID":

![Google Identity Platform - Identity Platform - New Credentials](./assets/docs/images/oauth2_google_new_credentials.png)

First, choose the type of credentials:

![Google Identity Platform - Identity Platform - Consent Screen](./assets/docs/images/oauth2_google_new_credentials_type_choice.png)

Next, provide Redirect URL:

![Google Identity Platform - Identity Platform - Consent Screen](./assets/docs/images/oauth2_google_new_credentials_uri_choice.png)

Finally, you will get a confirmation. Store the Client ID and Client Secret securely.

![Google Identity Platform - Identity Platform - Consent Screen](./assets/docs/images/oauth2_google_new_credentials_confirm.png)

[:arrow_up: Back to Top](#table-of-contents)

#### Auth0

TBD.

[:arrow_up: Back to Top](#table-of-contents)

#### OneLogin

TBD.

[:arrow_up: Back to Top](#table-of-contents)


#### Microsoft

TBD.

[:arrow_up: Back to Top](#table-of-contents)

<!--- end of section -->
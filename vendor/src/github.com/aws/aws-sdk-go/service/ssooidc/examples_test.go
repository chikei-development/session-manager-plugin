// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssooidc_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssooidc"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// Call OAuth/OIDC /token endpoint for Device Code grant with Secret authentication
//

func ExampleSSOOIDC_CreateToken_shared00() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.CreateTokenInput{
		ClientId:     aws.String("_yzkThXVzLWVhc3QtMQEXAMPLECLIENTID"),
		ClientSecret: aws.String("VERYLONGSECRETeyJraWQiOiJrZXktMTU2NDAyODA5OSIsImFsZyI6IkhTMzg0In0"),
		DeviceCode:   aws.String("yJraWQiOiJrZXktMTU2Njk2ODA4OCIsImFsZyI6IkhTMzIn0EXAMPLEDEVICECODE"),
		GrantType:    aws.String("urn:ietf:params:oauth:grant-type:device-code"),
	}

	result, err := svc.CreateToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientException:
				fmt.Println(ssooidc.ErrCodeInvalidClientException, aerr.Error())
			case ssooidc.ErrCodeInvalidGrantException:
				fmt.Println(ssooidc.ErrCodeInvalidGrantException, aerr.Error())
			case ssooidc.ErrCodeUnauthorizedClientException:
				fmt.Println(ssooidc.ErrCodeUnauthorizedClientException, aerr.Error())
			case ssooidc.ErrCodeUnsupportedGrantTypeException:
				fmt.Println(ssooidc.ErrCodeUnsupportedGrantTypeException, aerr.Error())
			case ssooidc.ErrCodeInvalidScopeException:
				fmt.Println(ssooidc.ErrCodeInvalidScopeException, aerr.Error())
			case ssooidc.ErrCodeAuthorizationPendingException:
				fmt.Println(ssooidc.ErrCodeAuthorizationPendingException, aerr.Error())
			case ssooidc.ErrCodeSlowDownException:
				fmt.Println(ssooidc.ErrCodeSlowDownException, aerr.Error())
			case ssooidc.ErrCodeAccessDeniedException:
				fmt.Println(ssooidc.ErrCodeAccessDeniedException, aerr.Error())
			case ssooidc.ErrCodeExpiredTokenException:
				fmt.Println(ssooidc.ErrCodeExpiredTokenException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Call OAuth/OIDC /token endpoint for Refresh Token grant with Secret authentication
//

func ExampleSSOOIDC_CreateToken_shared01() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.CreateTokenInput{
		ClientId:     aws.String("_yzkThXVzLWVhc3QtMQEXAMPLECLIENTID"),
		ClientSecret: aws.String("VERYLONGSECRETeyJraWQiOiJrZXktMTU2NDAyODA5OSIsImFsZyI6IkhTMzg0In0"),
		GrantType:    aws.String("refresh_token"),
		RefreshToken: aws.String("aorvJYubGpU6i91YnH7Mfo-AT2fIVa1zCfA_Rvq9yjVKIP3onFmmykuQ7E93y2I-9Nyj-A_sVvMufaLNL0bqnDRtgAkc0:MGUCMFrRsktMRVlWaOR70XGMFGLL0SlcCw4DiYveIiOVx1uK9BbD0gvAddsW3UTLozXKMgIxAJ3qxUvjpnlLIOaaKOoa/FuNgqJVvr9GMwDtnAtlh9iZzAkEXAMPLEREFRESHTOKEN"),
		Scope: []*string{
			aws.String("codewhisperer:completions"),
		},
	}

	result, err := svc.CreateToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientException:
				fmt.Println(ssooidc.ErrCodeInvalidClientException, aerr.Error())
			case ssooidc.ErrCodeInvalidGrantException:
				fmt.Println(ssooidc.ErrCodeInvalidGrantException, aerr.Error())
			case ssooidc.ErrCodeUnauthorizedClientException:
				fmt.Println(ssooidc.ErrCodeUnauthorizedClientException, aerr.Error())
			case ssooidc.ErrCodeUnsupportedGrantTypeException:
				fmt.Println(ssooidc.ErrCodeUnsupportedGrantTypeException, aerr.Error())
			case ssooidc.ErrCodeInvalidScopeException:
				fmt.Println(ssooidc.ErrCodeInvalidScopeException, aerr.Error())
			case ssooidc.ErrCodeAuthorizationPendingException:
				fmt.Println(ssooidc.ErrCodeAuthorizationPendingException, aerr.Error())
			case ssooidc.ErrCodeSlowDownException:
				fmt.Println(ssooidc.ErrCodeSlowDownException, aerr.Error())
			case ssooidc.ErrCodeAccessDeniedException:
				fmt.Println(ssooidc.ErrCodeAccessDeniedException, aerr.Error())
			case ssooidc.ErrCodeExpiredTokenException:
				fmt.Println(ssooidc.ErrCodeExpiredTokenException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Call OAuth/OIDC /token endpoint for Authorization Code grant with IAM authentication
//

func ExampleSSOOIDC_CreateTokenWithIAM_shared00() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.CreateTokenWithIAMInput{
		ClientId:    aws.String("arn:aws:sso::123456789012:application/ssoins-111111111111/apl-222222222222"),
		Code:        aws.String("yJraWQiOiJrZXktMTU2Njk2ODA4OCIsImFsZyI6IkhTMzg0In0EXAMPLEAUTHCODE"),
		GrantType:   aws.String("authorization_code"),
		RedirectUri: aws.String("https://mywebapp.example/redirect"),
		Scope: []*string{
			aws.String("openid"),
			aws.String("aws"),
			aws.String("sts:identity_context"),
		},
	}

	result, err := svc.CreateTokenWithIAM(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientException:
				fmt.Println(ssooidc.ErrCodeInvalidClientException, aerr.Error())
			case ssooidc.ErrCodeInvalidGrantException:
				fmt.Println(ssooidc.ErrCodeInvalidGrantException, aerr.Error())
			case ssooidc.ErrCodeUnauthorizedClientException:
				fmt.Println(ssooidc.ErrCodeUnauthorizedClientException, aerr.Error())
			case ssooidc.ErrCodeUnsupportedGrantTypeException:
				fmt.Println(ssooidc.ErrCodeUnsupportedGrantTypeException, aerr.Error())
			case ssooidc.ErrCodeInvalidScopeException:
				fmt.Println(ssooidc.ErrCodeInvalidScopeException, aerr.Error())
			case ssooidc.ErrCodeAuthorizationPendingException:
				fmt.Println(ssooidc.ErrCodeAuthorizationPendingException, aerr.Error())
			case ssooidc.ErrCodeSlowDownException:
				fmt.Println(ssooidc.ErrCodeSlowDownException, aerr.Error())
			case ssooidc.ErrCodeAccessDeniedException:
				fmt.Println(ssooidc.ErrCodeAccessDeniedException, aerr.Error())
			case ssooidc.ErrCodeExpiredTokenException:
				fmt.Println(ssooidc.ErrCodeExpiredTokenException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			case ssooidc.ErrCodeInvalidRequestRegionException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestRegionException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Call OAuth/OIDC /token endpoint for Refresh Token grant with IAM authentication
//

func ExampleSSOOIDC_CreateTokenWithIAM_shared01() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.CreateTokenWithIAMInput{
		ClientId:     aws.String("arn:aws:sso::123456789012:application/ssoins-111111111111/apl-222222222222"),
		GrantType:    aws.String("refresh_token"),
		RefreshToken: aws.String("aorvJYubGpU6i91YnH7Mfo-AT2fIVa1zCfA_Rvq9yjVKIP3onFmmykuQ7E93y2I-9Nyj-A_sVvMufaLNL0bqnDRtgAkc0:MGUCMFrRsktMRVlWaOR70XGMFGLL0SlcCw4DiYveIiOVx1uK9BbD0gvAddsW3UTLozXKMgIxAJ3qxUvjpnlLIOaaKOoa/FuNgqJVvr9GMwDtnAtlh9iZzAkEXAMPLEREFRESHTOKEN"),
	}

	result, err := svc.CreateTokenWithIAM(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientException:
				fmt.Println(ssooidc.ErrCodeInvalidClientException, aerr.Error())
			case ssooidc.ErrCodeInvalidGrantException:
				fmt.Println(ssooidc.ErrCodeInvalidGrantException, aerr.Error())
			case ssooidc.ErrCodeUnauthorizedClientException:
				fmt.Println(ssooidc.ErrCodeUnauthorizedClientException, aerr.Error())
			case ssooidc.ErrCodeUnsupportedGrantTypeException:
				fmt.Println(ssooidc.ErrCodeUnsupportedGrantTypeException, aerr.Error())
			case ssooidc.ErrCodeInvalidScopeException:
				fmt.Println(ssooidc.ErrCodeInvalidScopeException, aerr.Error())
			case ssooidc.ErrCodeAuthorizationPendingException:
				fmt.Println(ssooidc.ErrCodeAuthorizationPendingException, aerr.Error())
			case ssooidc.ErrCodeSlowDownException:
				fmt.Println(ssooidc.ErrCodeSlowDownException, aerr.Error())
			case ssooidc.ErrCodeAccessDeniedException:
				fmt.Println(ssooidc.ErrCodeAccessDeniedException, aerr.Error())
			case ssooidc.ErrCodeExpiredTokenException:
				fmt.Println(ssooidc.ErrCodeExpiredTokenException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			case ssooidc.ErrCodeInvalidRequestRegionException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestRegionException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Call OAuth/OIDC /token endpoint for JWT Bearer grant with IAM authentication
//

func ExampleSSOOIDC_CreateTokenWithIAM_shared02() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.CreateTokenWithIAMInput{
		Assertion: aws.String("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImtpZCI6IjFMVE16YWtpaGlSbGFfOHoyQkVKVlhlV01xbyJ9.eyJ2ZXIiOiIyLjAiLCJpc3MiOiJodHRwczovL2xvZ2luLm1pY3Jvc29mdG9ubGluZS5jb20vOTEyMjA0MGQtNmM2Ny00YzViLWIxMTItMzZhMzA0YjY2ZGFkL3YyLjAiLCJzdWIiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFJa3pxRlZyU2FTYUZIeTc4MmJidGFRIiwiYXVkIjoiNmNiMDQwMTgtYTNmNS00NmE3LWI5OTUtOTQwYzc4ZjVhZWYzIiwiZXhwIjoxNTM2MzYxNDExLCJpYXQiOjE1MzYyNzQ3MTEsIm5iZiI6MTUzNjI3NDcxMSwibmFtZSI6IkFiZSBMaW5jb2xuIiwicHJlZmVycmVkX3VzZXJuYW1lIjoiQWJlTGlAbWljcm9zb2Z0LmNvbSIsIm9pZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC02NmYzLTMzMzJlY2E3ZWE4MSIsInRpZCI6IjkxMjIwNDBkLTZjNjctNGM1Yi1iMTEyLTM2YTMwNGI2NmRhZCIsIm5vbmNlIjoiMTIzNTIzIiwiYWlvIjoiRGYyVVZYTDFpeCFsTUNXTVNPSkJjRmF0emNHZnZGR2hqS3Y4cTVnMHg3MzJkUjVNQjVCaXN2R1FPN1lXQnlqZDhpUURMcSFlR2JJRGFreXA1bW5PcmNkcUhlWVNubHRlcFFtUnA2QUlaOGpZIn0.1AFWW-Ck5nROwSlltm7GzZvDwUkqvhSQpm55TQsmVo9Y59cLhRXpvB8n-55HCr9Z6G_31_UbeUkoz612I2j_Sm9FFShSDDjoaLQr54CreGIJvjtmS3EkK9a7SJBbcpL1MpUtlfygow39tFjY7EVNW9plWUvRrTgVk7lYLprvfzw-CIqw3gHC-T7IK_m_xkr08INERBtaecwhTeN4chPC4W3jdmw_lIxzC48YoQ0dB1L9-ImX98Egypfrlbm0IBL5spFzL6JDZIRRJOu8vecJvj1mq-IUhGt0MacxX8jdxYLP-KUu2d9MbNKpCKJuZ7p8gwTL5B7NlUdh_dmSviPWrw"),
		ClientId:  aws.String("arn:aws:sso::123456789012:application/ssoins-111111111111/apl-222222222222"),
		GrantType: aws.String("urn:ietf:params:oauth:grant-type:jwt-bearer"),
	}

	result, err := svc.CreateTokenWithIAM(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientException:
				fmt.Println(ssooidc.ErrCodeInvalidClientException, aerr.Error())
			case ssooidc.ErrCodeInvalidGrantException:
				fmt.Println(ssooidc.ErrCodeInvalidGrantException, aerr.Error())
			case ssooidc.ErrCodeUnauthorizedClientException:
				fmt.Println(ssooidc.ErrCodeUnauthorizedClientException, aerr.Error())
			case ssooidc.ErrCodeUnsupportedGrantTypeException:
				fmt.Println(ssooidc.ErrCodeUnsupportedGrantTypeException, aerr.Error())
			case ssooidc.ErrCodeInvalidScopeException:
				fmt.Println(ssooidc.ErrCodeInvalidScopeException, aerr.Error())
			case ssooidc.ErrCodeAuthorizationPendingException:
				fmt.Println(ssooidc.ErrCodeAuthorizationPendingException, aerr.Error())
			case ssooidc.ErrCodeSlowDownException:
				fmt.Println(ssooidc.ErrCodeSlowDownException, aerr.Error())
			case ssooidc.ErrCodeAccessDeniedException:
				fmt.Println(ssooidc.ErrCodeAccessDeniedException, aerr.Error())
			case ssooidc.ErrCodeExpiredTokenException:
				fmt.Println(ssooidc.ErrCodeExpiredTokenException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			case ssooidc.ErrCodeInvalidRequestRegionException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestRegionException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Call OAuth/OIDC /token endpoint for Token Exchange grant with IAM authentication
//

func ExampleSSOOIDC_CreateTokenWithIAM_shared03() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.CreateTokenWithIAMInput{
		ClientId:           aws.String("arn:aws:sso::123456789012:application/ssoins-111111111111/apl-222222222222"),
		GrantType:          aws.String("urn:ietf:params:oauth:grant-type:token-exchange"),
		RequestedTokenType: aws.String("urn:ietf:params:oauth:token-type:access_token"),
		SubjectToken:       aws.String("aoak-Hig8TUDPNX1xZwOMXM5MxOWDL0E0jg9P6_C_jKQPxS_SKCP6f0kh1Up4g7TtvQqkMnD-GJiU_S1gvug6SrggAkc0:MGYCMQD3IatVjV7jAJU91kK3PkS/SfA2wtgWzOgZWDOR7sDGN9t0phCZz5It/aes/3C1Zj0CMQCKWOgRaiz6AIhza3DSXQNMLjRKXC8F8ceCsHlgYLMZ7hZDIFFERENTACCESSTOKEN"),
		SubjectTokenType:   aws.String("urn:ietf:params:oauth:token-type:access_token"),
	}

	result, err := svc.CreateTokenWithIAM(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientException:
				fmt.Println(ssooidc.ErrCodeInvalidClientException, aerr.Error())
			case ssooidc.ErrCodeInvalidGrantException:
				fmt.Println(ssooidc.ErrCodeInvalidGrantException, aerr.Error())
			case ssooidc.ErrCodeUnauthorizedClientException:
				fmt.Println(ssooidc.ErrCodeUnauthorizedClientException, aerr.Error())
			case ssooidc.ErrCodeUnsupportedGrantTypeException:
				fmt.Println(ssooidc.ErrCodeUnsupportedGrantTypeException, aerr.Error())
			case ssooidc.ErrCodeInvalidScopeException:
				fmt.Println(ssooidc.ErrCodeInvalidScopeException, aerr.Error())
			case ssooidc.ErrCodeAuthorizationPendingException:
				fmt.Println(ssooidc.ErrCodeAuthorizationPendingException, aerr.Error())
			case ssooidc.ErrCodeSlowDownException:
				fmt.Println(ssooidc.ErrCodeSlowDownException, aerr.Error())
			case ssooidc.ErrCodeAccessDeniedException:
				fmt.Println(ssooidc.ErrCodeAccessDeniedException, aerr.Error())
			case ssooidc.ErrCodeExpiredTokenException:
				fmt.Println(ssooidc.ErrCodeExpiredTokenException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			case ssooidc.ErrCodeInvalidRequestRegionException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestRegionException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Call OAuth/OIDC /register-client endpoint
//

func ExampleSSOOIDC_RegisterClient_shared00() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.RegisterClientInput{
		ClientName:             aws.String("My IDE Plugin"),
		ClientType:             aws.String("public"),
		EntitledApplicationArn: aws.String("arn:aws:sso::ACCOUNTID:application/ssoins-1111111111111111/apl-1111111111111111"),
		GrantTypes: []*string{
			aws.String("authorization_code"),
			aws.String("refresh_token"),
		},
		IssuerUrl: aws.String("https://identitycenter.amazonaws.com/ssoins-1111111111111111"),
		RedirectUris: []*string{
			aws.String("127.0.0.1:PORT/oauth/callback"),
		},
		Scopes: []*string{
			aws.String("sso:account:access"),
			aws.String("codewhisperer:completions"),
		},
	}

	result, err := svc.RegisterClient(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidScopeException:
				fmt.Println(ssooidc.ErrCodeInvalidScopeException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientMetadataException:
				fmt.Println(ssooidc.ErrCodeInvalidClientMetadataException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			case ssooidc.ErrCodeInvalidRedirectUriException:
				fmt.Println(ssooidc.ErrCodeInvalidRedirectUriException, aerr.Error())
			case ssooidc.ErrCodeUnsupportedGrantTypeException:
				fmt.Println(ssooidc.ErrCodeUnsupportedGrantTypeException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// Call OAuth/OIDC /start-device-authorization endpoint
//

func ExampleSSOOIDC_StartDeviceAuthorization_shared00() {
	svc := ssooidc.New(session.New())
	input := &ssooidc.StartDeviceAuthorizationInput{
		ClientId:     aws.String("_yzkThXVzLWVhc3QtMQEXAMPLECLIENTID"),
		ClientSecret: aws.String("VERYLONGSECRETeyJraWQiOiJrZXktMTU2NDAyODA5OSIsImFsZyI6IkhTMzg0In0"),
		StartUrl:     aws.String("https://identitycenter.amazonaws.com/ssoins-111111111111"),
	}

	result, err := svc.StartDeviceAuthorization(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ssooidc.ErrCodeInvalidRequestException:
				fmt.Println(ssooidc.ErrCodeInvalidRequestException, aerr.Error())
			case ssooidc.ErrCodeInvalidClientException:
				fmt.Println(ssooidc.ErrCodeInvalidClientException, aerr.Error())
			case ssooidc.ErrCodeUnauthorizedClientException:
				fmt.Println(ssooidc.ErrCodeUnauthorizedClientException, aerr.Error())
			case ssooidc.ErrCodeSlowDownException:
				fmt.Println(ssooidc.ErrCodeSlowDownException, aerr.Error())
			case ssooidc.ErrCodeInternalServerException:
				fmt.Println(ssooidc.ErrCodeInternalServerException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

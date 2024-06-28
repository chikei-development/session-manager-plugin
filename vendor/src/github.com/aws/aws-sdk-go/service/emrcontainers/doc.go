// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package emrcontainers provides the client and types for making API
// requests to Amazon EMR Containers.
//
// Amazon EMR on EKS provides a deployment option for Amazon EMR that allows
// you to run open-source big data frameworks on Amazon Elastic Kubernetes Service
// (Amazon EKS). With this deployment option, you can focus on running analytics
// workloads while Amazon EMR on EKS builds, configures, and manages containers
// for open-source applications. For more information about Amazon EMR on EKS
// concepts and tasks, see What is Amazon EMR on EKS (https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/emr-eks.html).
//
// Amazon EMR containers is the API name for Amazon EMR on EKS. The emr-containers
// prefix is used in the following scenarios:
//
//   - It is the prefix in the CLI commands for Amazon EMR on EKS. For example,
//     aws emr-containers start-job-run.
//
//   - It is the prefix before IAM policy actions for Amazon EMR on EKS. For
//     example, "Action": [ "emr-containers:StartJobRun"]. For more information,
//     see Policy actions for Amazon EMR on EKS (https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/security_iam_service-with-iam.html#security_iam_service-with-iam-id-based-policies-actions).
//
//   - It is the prefix used in Amazon EMR on EKS service endpoints. For example,
//     emr-containers.us-east-2.amazonaws.com. For more information, see Amazon
//     EMR on EKSService Endpoints (https://docs.aws.amazon.com/emr/latest/EMR-on-EKS-DevelopmentGuide/service-quotas.html#service-endpoints).
//
// See https://docs.aws.amazon.com/goto/WebAPI/emr-containers-2020-10-01 for more information on this service.
//
// See emrcontainers package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/emrcontainers/
//
// # Using the Client
//
// To contact Amazon EMR Containers with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon EMR Containers client EMRContainers for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/emrcontainers/#New
package emrcontainers

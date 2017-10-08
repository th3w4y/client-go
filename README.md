# Go API client for swagger

An API to insert and retrieve annotations on cloud artifacts.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 0.1
- Package version: 1.0.0
- Build date: 2017-10-08T17:36:22.846-04:00
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*GrafeasApi* | [**CreateNote**](docs/GrafeasApi.md#createnote) | **Post** /v1alpha1/projects/{projectsId}/notes | 
*GrafeasApi* | [**CreateOccurrence**](docs/GrafeasApi.md#createoccurrence) | **Post** /v1alpha1/projects/{projectsId}/occurrences | 
*GrafeasApi* | [**DeleteNote**](docs/GrafeasApi.md#deletenote) | **Delete** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
*GrafeasApi* | [**DeleteOccurrence**](docs/GrafeasApi.md#deleteoccurrence) | **Delete** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
*GrafeasApi* | [**GetNote**](docs/GrafeasApi.md#getnote) | **Get** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
*GrafeasApi* | [**GetOccurrence**](docs/GrafeasApi.md#getoccurrence) | **Get** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
*GrafeasApi* | [**GetOccurrenceNote**](docs/GrafeasApi.md#getoccurrencenote) | **Get** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId}/notes | 
*GrafeasApi* | [**GetOperation**](docs/GrafeasApi.md#getoperation) | **Get** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 
*GrafeasApi* | [**ListNoteOccurrences**](docs/GrafeasApi.md#listnoteoccurrences) | **Get** /v1alpha1/projects/{projectsId}/notes/{notesId}/occurrences | 
*GrafeasApi* | [**ListNotes**](docs/GrafeasApi.md#listnotes) | **Get** /v1alpha1/projects/{projectsId}/notes | 
*GrafeasApi* | [**ListOccurrences**](docs/GrafeasApi.md#listoccurrences) | **Get** /v1alpha1/projects/{projectsId}/occurrences | 
*GrafeasApi* | [**ListOperations**](docs/GrafeasApi.md#listoperations) | **Get** /v1alpha1/projects/{projectsId}/operations | 
*GrafeasApi* | [**UpdateNote**](docs/GrafeasApi.md#updatenote) | **Put** /v1alpha1/projects/{projectsId}/notes/{notesId} | 
*GrafeasApi* | [**UpdateOccurrence**](docs/GrafeasApi.md#updateoccurrence) | **Put** /v1alpha1/projects/{projectsId}/occurrences/{occurrencesId} | 
*GrafeasApi* | [**UpdateOperation**](docs/GrafeasApi.md#updateoperation) | **Put** /v1alpha1/projects/{projectsId}/operations/{operationsId} | 


## Documentation For Models

 - [AliasContext](docs/AliasContext.md)
 - [Artifact](docs/Artifact.md)
 - [Basis](docs/Basis.md)
 - [BuildDetails](docs/BuildDetails.md)
 - [BuildProvenance](docs/BuildProvenance.md)
 - [BuildSignature](docs/BuildSignature.md)
 - [BuildType](docs/BuildType.md)
 - [CloudRepoSourceContext](docs/CloudRepoSourceContext.md)
 - [CloudWorkspaceId](docs/CloudWorkspaceId.md)
 - [CloudWorkspaceSourceContext](docs/CloudWorkspaceSourceContext.md)
 - [Command](docs/Command.md)
 - [CreateOperationRequest](docs/CreateOperationRequest.md)
 - [CustomDetails](docs/CustomDetails.md)
 - [Deployable](docs/Deployable.md)
 - [Deployment](docs/Deployment.md)
 - [Derived](docs/Derived.md)
 - [Detail](docs/Detail.md)
 - [Discovered](docs/Discovered.md)
 - [Discovery](docs/Discovery.md)
 - [Distribution](docs/Distribution.md)
 - [Empty](docs/Empty.md)
 - [ExtendedSourceContext](docs/ExtendedSourceContext.md)
 - [FileHashes](docs/FileHashes.md)
 - [Fingerprint](docs/Fingerprint.md)
 - [GerritSourceContext](docs/GerritSourceContext.md)
 - [GitSourceContext](docs/GitSourceContext.md)
 - [Hash](docs/Hash.md)
 - [Installation](docs/Installation.md)
 - [Layer](docs/Layer.md)
 - [ListNoteOccurrencesResponse](docs/ListNoteOccurrencesResponse.md)
 - [ListNotesResponse](docs/ListNotesResponse.md)
 - [ListOccurrencesResponse](docs/ListOccurrencesResponse.md)
 - [ListOperationsResponse](docs/ListOperationsResponse.md)
 - [Location](docs/Location.md)
 - [ModelPackage](docs/ModelPackage.md)
 - [Note](docs/Note.md)
 - [Occurrence](docs/Occurrence.md)
 - [Operation](docs/Operation.md)
 - [PackageIssue](docs/PackageIssue.md)
 - [ProjectRepoId](docs/ProjectRepoId.md)
 - [RelatedUrl](docs/RelatedUrl.md)
 - [RepoId](docs/RepoId.md)
 - [RepoSource](docs/RepoSource.md)
 - [Source](docs/Source.md)
 - [SourceContext](docs/SourceContext.md)
 - [Status](docs/Status.md)
 - [StorageSource](docs/StorageSource.md)
 - [UpdateOperationRequest](docs/UpdateOperationRequest.md)
 - [Version](docs/Version.md)
 - [VulnerabilityDetails](docs/VulnerabilityDetails.md)
 - [VulnerabilityLocation](docs/VulnerabilityLocation.md)
 - [VulnerabilityType](docs/VulnerabilityType.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author




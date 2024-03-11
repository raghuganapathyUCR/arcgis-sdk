<p align="center">
  <img src="/LOGO.png" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">arcgis-sdk</h1>
</p>
<p align="center">
    <h6 align="center">v0.0.2-alpha</h6>
</p>
<p align="center">
    <em><code>A lightweight Golang SDK for Arcgis REST Services. Inspired by
    <a>arcgis-rest-js</></code></em>
</p>
<p align="center">
   <img src="https://img.shields.io/github/actions/workflow/status/raghuganapathyUCR/arcgis-sdk/go.yml" alt="repo-languages">
	<img src="https://img.shields.io/codecov/c/github/raghuganapathyUCR/arcgis-sdk" alt="license">
	<img src="https://img.shields.io/github/languages/top/raghuganapathyUCR/arcgis-sdk?style=default&color=0080ff" alt="repo-top-language">
   <img src="https://img.shields.io/github/last-commit/raghuganapathyUCR/arcgis-sdk?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">


<p>


<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Tests](#-tests)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)
</details>
<hr>

##  Overview

<code><i>WORK IN PROGRESS:</i> arcgis-sdk is a lightweight, modular golang wrapper for Esri's ArcGis rest services. Currently it supports API Key based Authentication and Geocoding/Reverse Geocoding services. The SDK is designed to be modular and extensible, allowing for easy integration with other ArcGis services.
</code>

---

##  Features

<code>► Geocode and Reverse geocode your data into well structured lightweight structs!</code>

---

##  Repository Structure

```sh
└── arcgis-sdk/
    ├── .github
    │   └── workflows
    ├── LICSENSE.md
    ├── auth
    │   ├── auth.go
    │   ├── auth_errors.go
    │   ├── auth_manager.go
    │   └── auth_test.go
    ├── geocode
    │   ├── Address.go
    │   ├── errors.go
    │   ├── geocode.go
    │   ├── geocode_test.go
    │   └── reverse.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── requests
    │   ├── request.go
    │   ├── request_options.go
    │   └── requests_test.go
    └── utils
        └── utils.go
```

---

##  Modules

<details closed><summary>auth</summary>

| File                                                                                                | Summary                         |
| ---                                                                                                 | ---                             |
| [auth_test.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/auth/auth_test.go)       | <code>Authentication Manager tests</code> |
| [auth.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/auth/auth.go)                 | <code>Main auth package, exports the APIKeymanager</code> |
| [auth_errors.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/auth/auth_errors.go)   | <code>Auth Specific Error Logic and Types</code> |
| [auth_manager.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/auth/auth_manager.go) | <code>Auth Manager and definitions</code> |

</details>

<details closed><summary>utils</summary>

| File                                                                                   | Summary                         |
| ---                                                                                    | ---                             |
| [utils.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/utils/utils.go) | <code>► All utility functions</code> |

</details>

<details closed><summary>requests</summary>

| File                                                                                                          | Summary                         |
| ---                                                                                                           | ---                             |
| [requests_test.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/requests/requests_test.go)     | <code>Tests for Requests module</code> |
| [request.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/requests/request.go)                 | <code>Main Requests wrapper, can be used to ingest custom request options for different rest services</code> |
| [request_options.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/requests/request_options.go) | <code>Type Definition for Request Options</code> |

</details>

<details closed><summary>.github.workflows</summary>

| File                                                                                           | Summary                         |
| ---                                                                                            | ---                             |
| [go.yml](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/.github/workflows/go.yml) | <code>Release Pipeline</code> |

</details>

<details closed><summary>geocode</summary>

| File                                                                                                   | Summary                         |
| ---                                                                                                    | ---                             |
| [Address.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/geocode/Address.go)           | <code>Address type definitions</code> |
| [geocode.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/geocode/geocode.go)           | <code>Main Geocoder and related code</code> |
| [reverse.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/geocode/reverse.go)           | <code>Main reverse Geocoder and related code</code> |
| [geocode_test.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/geocode/geocode_test.go) | <code>geocode and reversegeocode test</code> |
| [errors.go](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/master/geocode/errors.go)             | <code>API Error Definitions</code> |

</details>

---

##  Getting Started

**System Requirements:**

* **Go**: `version 1.21`

###  Installation
<h4>From <code>pkg.go.dev</code></h4>

> 1. Install the package using the command below:
> ```console
> $ go get github.com/raghuganapathyUCR/arcgis-sdk
> ```
> We recommend downloading the modules you need using the `go get` command.
> Example: ``` go get github.com/raghuganapathyUCR/arcgis-sdk/geocode``` to get the geocode module.

<h4>From <code>source</code></h4>

> 1. Clone the arcgis-sdk repository:
>
> ```console
> $ git clone https://github.com/raghuganapathyUCR/arcgis-sdk
> ```
>
> 2. Change to the project directory:
> ```console
> $ cd arcgis-sdk
> ```
>
> 3. Install the dependencies:
> ```console
> $ go build -o myapp
> ```

###  Usage

<h4>From <code>source</code></h4>

> Run arcgis-sdk using the command below:
> ```console
> $ ./myapp
> ```

###  Tests

> Run the test suite using the command below:
> ```console
> $ go test ./...  
> ```

---

##  Project Roadmap

- [X] `► Complete the Basic SDK Structure`
- [X] `► Complete the Geocoder and Reverse Geocoder`
- [ ] `► Implement the Bulk and Batch Geocoding`
- [ ] `► Implement Enterprise Support`
- [ ] `► Implement the Routing and Navigation`
- [ ] `► Implement the Spatial Analysis`
- [ ] `► Implement Places and Location Services`


---

##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/raghuganapathyUCR/arcgis-sdk/issues)**: Submit bugs found or log feature requests for the `arcgis-sdk` project.
- **[Submit Pull Requests](https://github.com/raghuganapathyUCR/arcgis-sdk/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/raghuganapathyUCR/arcgis-sdk/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/raghuganapathyUCR/arcgis-sdk
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="center">
   <a href="https://github.com{/raghuganapathyUCR/arcgis-sdk/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=raghuganapathyUCR/arcgis-sdk">
   </a>
</p>
</details>

---

##  License

This project is protected under the [MIT](LICSENSE.md) License. For more details, refer to the [LICENSE](LICSENSE.md) file.

---

##  Acknowledgments

- Many thanks to [ESRI](https://www.esri.com) for their [arcgis-rest-js](https://github.com/esri/arcgis-rest-js/) project, which inspired this SDK.

---

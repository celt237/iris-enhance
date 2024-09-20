[![Release](https://github.com/celt237/iris-enhance/actions/workflows/go.yml/badge.svg)](https://github.com/celt237/iris-enhance/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/celt237/iris-enhance/badge.svg?branch=master)](https://coveralls.io/repos/github/celt237/iris-enhance/badge.svg?branch=master)

# iris-enhance

[中文版 README](README_CN.md)

## Iris Enhancement (including Swagger documentation generation, custom annotations, etc.)

### Project Description

- Enhances development with annotations, enabling a Java-like controller development approach
- Reduces code duplication and improves development efficiency
- One-click generation of Swagger documentation
- Supports custom annotations for implementing features like logging and access control

### Example

1. Create a service file in the service folder
2. Run the iris-enhance command to generate corresponding handler and router code
3. Register routes
4. Bind Swagger documentation (UI from knife4j)

For detailed examples, please refer to the Chinese README.

### Supported Annotation Tags

- @zService
- @zResult
- @zSummary
- @zDescription
- @zTags
- @zParam
- @zResultData
- @zAccept
- @zProduce
- @zRouter

For detailed explanations of these tags, please refer to the Chinese README.

### Command Usage

1. Installation:

   ```shell
   go get github.com/celt237/iris-enhance/cmd/iris-enhance@latest
   ```

2. Add github.com/celt237/iris-enhance dependency to your project

3. Run the iris-enhance command:

   ```shell
   iris-enhance --servicePath=xxx --handlePath=xxx --result=xxx --errorCode=xxx
   ```

For detailed instructions and troubleshooting, please refer to the Chinese README.

### MIT LICENSE

[LICENSE](./LICENSE)

### Links

- knife4j: <https://github.com/xiaomin/knife4j>
- go-annotation: <https://github.com/celt237/go-annotation>

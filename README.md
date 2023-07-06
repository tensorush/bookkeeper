## :hamster: :books: **bookkeeper**

[![CI][ci-shield]][ci-url]
[![Schema][schema-shield]][schema-url]
[![Codecov][codecov-shield]][codecov-url]
[![License][license-shield]][license-url]

### Server backend for a bookkeeping service.

### :sparkles: Features

- Server API support:
    - REST JSON using Gin framework,
    - gRPC using protobuf compiler,
    - REST JSON using gRPC gateway.

- Storage support: PostgreSQL and Redis.

- Access token support: PASETO and JWT.

- Email notifications for new users.

### :arrow_down: Dependencies

- [`just` - command runner](https://just.systems/)

- [`protoc` - Protocol Buffer compiler](https://grpc.io/schema/protoc-installation/)

- [`sqlc` - SQL query interface generator](https://sqlc.dev/)

- [`migrate` - Database migrations framework](https://github.com/golang-migrate/migrate)

- [`mockgen` - Mock testing framework for Go](https://github.com/golang/mock)

- [`statik` - Static file embedding tool for Go](https://github.com/rakyll/statik)

- [`dbdocs` - Database documentation tool](https://dbdocs.io/)

- [`dbml2sql` - DBML-to-SQL converter](https://dbml.dbdiagram.io/cli/)

<!-- MARKDOWN LINKS -->

[ci-shield]: https://img.shields.io/github/actions/workflow/status/tensorush/bookkeeper/ci.yaml?branch=main&style=for-the-badge&logo=github&label=CI&labelColor=black
[ci-url]: https://github.com/tensorush/bookkeeper/blob/main/.github/workflows/ci.yaml
[schema-shield]: https://img.shields.io/badge/click-2596BE?style=for-the-badge&logo=go&logoColor=2596BE&label=schema&labelColor=black
[schema-url]: https://dbdocs.io/tensorush/bookkeeper
[codecov-shield]: https://img.shields.io/codecov/c/github/tensorush/bookkeeper?style=for-the-badge&labelColor=black
[codecov-url]: https://app.codecov.io/gh/tensorush/bookkeeper
[license-shield]: https://img.shields.io/github/license/tensorush/bookkeeper.svg?style=for-the-badge&labelColor=black
[license-url]: https://github.com/tensorush/bookkeeper/blob/main/LICENSE.md

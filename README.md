## :hamster: :books: **bookkeeper**

[![CI][ci-shield]][ci-url]
[![Schema][schema-shield]][schema-url]
[![Codecov][codecov-shield]][codecov-url]
[![License][license-shield]][license-url]

### Backend for a bookkeeping service based on gRPC.

### :rocket: Usage

#### See [`justfile`](justfile) for all development and deployment commands.

### :sparkles: Features

- Server API support:
    - gRPC using protobuf compiler,
    - REST JSON using gRPC gateway,
    - REST JSON using Gin framework.

- Storage support: PostgreSQL and Redis.

- Access token support: PASETO and JWT.

- Email notifications for new users via SMTP.

### :arrow_down: Dependencies

- [`just`](https://just.systems/) - command runner.

- [`protoc`](https://github.com/golang/protobuf) - Protocol Buffer compiler.

- [`dbml2sql`](https://dbml.dbdiagram.io/cli/) - DBML-to-SQL converter.

- [`sqlc`](https://sqlc.dev/) - SQL query interface generator.

- [`dbdocs`](https://dbdocs.io/) - Database documentation tool.

- [`mockgen`](https://github.com/golang/mock) - Mock testing framework for Go.

- [`statik`](https://github.com/rakyll/statik) - Static file embedding tool for Go.

- [`migrate`](https://github.com/golang-migrate/migrate) - Database migrations framework.

<!-- MARKDOWN LINKS -->

[ci-shield]: https://img.shields.io/github/actions/workflow/status/tensorush/bookkeeper/ci.yaml?branch=main&style=for-the-badge&logo=github&label=CI&labelColor=black
[ci-url]: https://github.com/tensorush/bookkeeper/blob/main/.github/workflows/ci.yaml
[schema-shield]: https://img.shields.io/badge/click-2596BE?style=for-the-badge&logo=go&logoColor=2596BE&label=schema&labelColor=black
[schema-url]: https://dbdocs.io/tensorush/bookkeeper
[codecov-shield]: https://img.shields.io/codecov/c/github/tensorush/bookkeeper?style=for-the-badge&labelColor=black
[codecov-url]: https://app.codecov.io/gh/tensorush/bookkeeper
[license-shield]: https://img.shields.io/github/license/tensorush/bookkeeper.svg?style=for-the-badge&labelColor=black
[license-url]: https://github.com/tensorush/bookkeeper/blob/main/LICENSE.md

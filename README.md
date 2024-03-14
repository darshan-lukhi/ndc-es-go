# NDC example with CSV

The connector is ported from [NDC Reference Connector](https://github.com/hasura/ndc-spec/tree/main/ndc-reference) that read CSV files into memory. It is intended to illustrate the concepts involved, and should be complete, in the sense that every specification feature is covered. It is not optimized and is not intended for production use, but might be useful for testing.

## Getting Started

```bash
go run . serve
```

## Using the reference connector

The reference connector runs on http://localhost:8080:

```sh
curl http://localhost:8080/schema | jq .
```

- Technical Design Document
- Define steps for development 
- Basic skeleton for connector
- Write elasticsearch client to support all query
- Implement ScaScalerler and Object type for Elasticsearch
- Implement the Collection NDA specification
- Implement the Simple queries
- Implement the Complex queries
- Implement Sorting, pagination, filtering, etc.
- Implement the Aggregation
- Add testcases for each functionality


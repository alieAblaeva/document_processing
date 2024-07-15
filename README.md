# Document Processing
VK test assignment
___


This repository contains a document processing service that updates document fields based on specific rules. The service is designed to process incoming document updates and maintain consistent document states.



## Overview

The service processes incoming document updates with the following rules:
- `Text` and `FetchTime` are updated to the values from the document with the greatest `FetchTime`.
- `PubDate` is taken from the document with the smallest `FetchTime`.
- `FirstFetchTime` is set to the smallest `FetchTime`.

Document format:
``` proto
message TDocument {
    string url = 1; 
    uint64 pub_date = 2;
    uint64 fetch_time = 3;
    string text = 4; 
    uint64 first_fetch_time = 5; 
}
```

## Features

- Processes document updates and maintains consistent states.
- Supports MongoDB for document storage.
- Provides unit tests to ensure functionality.

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/alieAblaeva/document_processing.git
    cd document_processing
    ```

2. **Install dependencies**:
    Make sure you have [Go](https://golang.org/dl/) installed. Then, install any Go dependencies:
    ```sh
    go mod tidy
    ```

3. **Set up MongoDB**:
    Ensure you have a running MongoDB instance.:
    ```sh
    docker-compose up -d


## Testing
To run the tests, use the following command:

```sh
go test .
```
This will run the tests and display the results.

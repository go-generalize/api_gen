## Development Guide

### E2E Tests
In `e2e/`, you can write E2E tests to be executed on CI. While running the server with the actual API definitions, the generated clients can be used to run the API and test it.
For example, in `multipart/`, `api/` contains API definitions, `controller/` contains server-side processing, and `tests/` contains client-side processing.
Refer to `multipart/` when writing new tests.
The TypeScript client uses headless Chrome to run in the browser.

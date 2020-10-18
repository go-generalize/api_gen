import { APIClient, MockOption } from "./api/api_client";

// the simplest
(async () => {
    const client = new APIClient();

    const resp = await client.postCreateUser({
        ID: "id",
        Password: "password",
        Gender: 0,
    });

    console.log(resp);
})();

// with options
(async () => {
    const client = new APIClient(
        "pass", // [optional] token for Authorization: Bearer
        {
            "X-Foobar": "foobar", // [optional] custom headers
        },
        "http://localhost:8080",  // [optional] custom endpoint
        {
            cache: "no-cache", // [optional] options for fetch API
        },
    );

    const resp = await client.postCreateUser(
        {
            ID: "id",
            Password: "password",
            Gender: 0,
        },
        {
            "X-Foobar": "hoge", // [optional] custom headers
        },
        {
            mode: "cors" // [optional] options for fetch API 
        },
    );

    console.log(resp);
})();

// mock mode
(async () => {
    const client = new APIClient();
    const mockOption: MockOption = {
        wait_ms: 10,
        target_file: 'error'
    }

    const resp = await client.postCreateUser({
        ID: "id",
        Password: "password",
        Gender: 0,
    }, undefined, {
        'mock_option': mockOption,
    });

    console.log(resp);
})();

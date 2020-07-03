import { PostCreateUserRequest } from "./api/classes/PostCreateUserRequest";
import { APIClient } from "./api/api_client";

// the simplest
(async () => {
    const client = new APIClient();

    const resp = await client.postCreateUser(new PostCreateUserRequest({
        ID: "id",
        Password: "password",
        Gender: 0,
    }));

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
        new PostCreateUserRequest({
            ID: "id",
            Password: "password",
            Gender: 0,
        }),
        {
            "X-Foobar": "hoge", // [optional] custom headers
        },
        {
            mode: "cors" // [optional] options for fetch API 
        },
    );

    console.log(resp);
})();
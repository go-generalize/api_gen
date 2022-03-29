import { APIClient } from "../../clients/ts/api_client";

async function postA(baseURL: string) {
    const client = new APIClient({
        baseURL,
    });

    const file = new Blob(['1'], {type: 'text/plain'});
    const files = [
        new File(['0'], '0.txt', {type: 'text/plain'}),
        new File(['1'], '1.txt', {type: 'text/plain'}),
    ]

    const result = await client.postA({
        File: file,
        Files: files,
        Payload: 'payload',
    })

    if (result.Message !== "OK") {
        throw new Error("Unexpected result: " + result.Message);
    }
}

async function postB(baseURL: string) {
    const client = new APIClient({
        baseURL,
    });

    const file = new Blob(['1'], {type: 'text/plain'});
    const files = [
        new File(['0'], '0.txt', {type: 'text/plain'}),
        new File(['1'], '1.txt', {type: 'text/plain'}),
    ]

    const result = await client._param.postB({
        File: file,
        Files: files,
        payload: 'payload',
        Param: 'param',
    })

    if (result.message !== "OK") {
        throw new Error("Unexpected result: " + result.message);
    }
}

await execute((async () => {
    const baseURL = arguments[0];

    await postA(baseURL);
    await postB(baseURL);
}));

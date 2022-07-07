// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import {
	PostUserRequest as FooBarPostUserRequest,
	PostUserResponse as FooBarPostUserResponse,
} from './classes/foo/bar/types';

const filterUndefinedParam = (param: Object) => {
	return Object.fromEntries(
		Object.entries(param)
		.filter(([_key, value]) => typeof value !== 'undefined')
	);
}

export interface MiddlewareContext {
	httpMethod: string;
	endpoint: string;
	request: unknown;
	response?: unknown;
	responseText?: string;
	baseURL: string;
	headers: {[key: string]: string};
	options: {[key: string]: any};
}

export enum MiddlewareResult {
	CONTINUE = 1,
	MIDDLEWARE_STOP = 2,
	STOP = 4
}

export type ApiClientMiddlewareFunc = (context: MiddlewareContext) => Promise<MiddlewareResult|boolean>;

export interface middlewareSet {
	beforeMiddleware?: ApiClientMiddlewareFunc[];
	afterMiddleware?: ApiClientMiddlewareFunc[];
}

class FooBarClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
	}

	getRequestObject(obj: any, routingPath: string[], isGET: boolean): { [key: string]: any } {
		let res: { [key: string]: any } = {};
		Object.keys(obj).forEach((key) => {
			if (isGET && obj[key] == null) {
				return;
			}
			if (routingPath.indexOf(key) === -1) {
				let val = obj[key];
				if (isGET) {
					val = val.toString();
				}
				res[key] = val;
			}
		});
		return res;
	}

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
	}

	async postUser(
		param?: FooBarPostUserRequest,
		options?: {
			headers?: {[key: string]: string},
			options?: {[key: string]: any}
		}
	): Promise<FooBarPostUserResponse>;
	/** @deprecated */
	async postUser(
		param?: FooBarPostUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<FooBarPostUserResponse>;

	async postUser(
		param?: FooBarPostUserRequest,
		arg1?: any,
		arg2?: any
	): Promise<FooBarPostUserResponse> {
		let headers: {[key: string]: string} | undefined;
		let options: {[key: string]: any} | undefined;
		if (!param) {
			param = {};
		}

		const filteredParam = filterUndefinedParam(param);

		if (
			arg2 !== undefined || arg1 === undefined ||
			Object.values(arg1).filter(v => typeof v !== 'string').length === 0
		) {
			headers = arg1;
			options = arg2;
		} else {
			headers = arg1.headers;
			options = arg1.options;
		}

	    const excludeParams: string[] = [];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
			'Content-Type': 'application/json',
			...this.headers,
			...headers,
			...mockHeaders,
		};
		const reqOption = {
			...this.options,
			...options,
		};
		const context: MiddlewareContext = {
			httpMethod: 'POST',
			endpoint: `${this.baseURL}/foo/bar/user`,
			request: filteredParam,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/foo/bar/user`;
		const resp = await fetch(
			url,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(filteredParam, excludeParams, false)),
				headers: reqHeader,
				...reqOption,
			}
		);

		const responseText = await resp.text();
		context.responseText = responseText;

		if (Math.floor(resp.status / 100) !== 2) {
			await this.callMiddleware(this.afterMiddleware, context);
			throw new ApiError(resp, responseText);
		}
		await resp.text();
		const res = {} as FooBarPostUserResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

class FooClient {
	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];
	public bar: FooBarClient;
	constructor(
		private headers: {[key: string]: string},
		private options: {[key: string]: any},
		private baseURL: string,
		middleware: middlewareSet
	) {
		this.beforeMiddleware = middleware.beforeMiddleware!;
		this.afterMiddleware = middleware.afterMiddleware!;
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};
		this.bar = new FooBarClient(
			headers,
			options,
			baseURL,
			childMiddlewareSet
		);
	}

	getRequestObject(obj: any, routingPath: string[], isGET: boolean): { [key: string]: any } {
		let res: { [key: string]: any } = {};
		Object.keys(obj).forEach((key) => {
			if (isGET && obj[key] == null) {
				return;
			}
			if (routingPath.indexOf(key) === -1) {
				let val = obj[key];
				if (isGET) {
					val = val.toString();
				}
				res[key] = val;
			}
		});
		return res;
	}

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
	}
}

export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];

	public foo: FooClient;

	constructor(opt: {
		token?: string,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions?: {[key: string]: any},
		middleware?: middlewareSet
	});
	constructor(
		token?: string,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions?: {[key: string]: any},
		middleware?: middlewareSet
	);

	constructor(
		token?: any,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions?: {[key: string]: any},
		middleware?: middlewareSet
	) {
		if (token !== null && (typeof token === 'object')) {
			commonHeaders = token.commonHeaders;
			baseURL = token.baseURL;
			commonOptions = token.commonOptions;
			middleware = token.middleware;
			token = token.token;
		}

		const headers: {[key: string]: string} =  {
			...commonHeaders,
		};

		if (token !== undefined) {
			headers['Authorization'] = 'Bearer ' + token;
		}

		this.baseURL =  (baseURL === undefined) ? "" : baseURL;
		this.options = commonOptions ?? {};
		this.headers = headers;

		if (middleware) {
			this.beforeMiddleware = middleware.beforeMiddleware ?? [];
			this.afterMiddleware = middleware.afterMiddleware ?? [];
		}
		const childMiddlewareSet: middlewareSet = {
			beforeMiddleware: this.beforeMiddleware,
			afterMiddleware: this.afterMiddleware
		};

		this.foo = new FooClient(
			headers,
			this.options,
			this.baseURL,
			childMiddlewareSet
		);
	}

	getRequestObject(obj: any, routingPath: string[], isGET: boolean): { [key: string]: any } {
		let res: { [key: string]: any } = {};
		Object.keys(obj).forEach((key) => {
			if (isGET && obj[key] == null) {
				return;
			}
			if (routingPath.indexOf(key) === -1) {
				let val = obj[key];
				if (isGET) {
					val = val.toString();
				}
				res[key] = val;
			}
		});
		return res;
	}

	async callMiddleware(
		middlewares: ApiClientMiddlewareFunc[],
		context: MiddlewareContext
	) {
		for (const m of middlewares) {
			const func: ApiClientMiddlewareFunc = m;
			const mr = await func(context);
			if (typeof mr === 'boolean') {
				if (!mr) {
					break;
				}
			} else {
				if (mr === MiddlewareResult.CONTINUE) {
					continue;
				} else if (mr === MiddlewareResult.MIDDLEWARE_STOP) {
					break;
				} else if (mr === MiddlewareResult.STOP) {
					throw new ApiMiddlewareStop();
				}
			}
		}
	}
}

export class ApiError extends Error {
	private _statusCode: number;
	private _statusText: string;
	private _response: string;

	constructor(response: Response, responseText: string) {
		super();
		this._statusCode = response.status;
		this._statusText = response.statusText;
		this._response = responseText
	}

	get statusCode(): number {
		return this._statusCode;
	}

	get statusText(): string {
		return this._statusText;
	}

	get response(): string {
		return this._response;
	}
}

export class ApiMiddlewareStop extends Error {}

export interface MockOption {
	wait_ms: number;
	target_file: string;
}

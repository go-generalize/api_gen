// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

import {
	PostBRequest as ParamPostBRequest,
	PostBResponse as ParamPostBResponse,
} from './classes/_param/types';

import {
	PostARequest as PostARequest,
	PostAResponse as PostAResponse,
} from './classes/types';

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

class ParamClient {
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

	async postB(
		param: ParamPostBRequest,
		options?: {
			headers?: {[key: string]: string},
			options?: {[key: string]: any}
		}
	): Promise<ParamPostBResponse>;
	/** @deprecated */
	async postB(
		param: ParamPostBRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ParamPostBResponse>;

	async postB(
		param: ParamPostBRequest,
		arg1?: any,
		arg2?: any
	): Promise<ParamPostBResponse> {
		let headers: {[key: string]: string} | undefined;
		let options: {[key: string]: any} | undefined;

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

	    const excludeParams: string[] = ['File', 'Files', 'Param', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
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
			endpoint: `${this.baseURL}/${encodeURI(param.Param.toString())}/b`,
			request: filteredParam,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/${encodeURI(param.Param.toString())}/b`;
		const resp = await fetch(
			url,
			{
				method: "POST",
				body: (() => {
					const body = new FormData();

					body.append(
						'x-multipart-json-binder-request-json',
						new Blob([JSON.stringify(this.getRequestObject(filteredParam, excludeParams, false))], {type: 'application/json'}),
						'x-multipart-json-binder-request-json'
					);
					if (filteredParam.File !== undefined) {
						body.append('file', filteredParam.File);
					}
					if (filteredParam.Files !== undefined) {
						filteredParam.Files.filter((f: unknown) => f !== undefined).forEach((f: File | Blob) => body.append('files', f));
					}
					return body;
				})(),
				headers: reqHeader,
				...reqOption,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			await this.callMiddleware(this.afterMiddleware, context);
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as ParamPostBResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
	}
}

export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	private beforeMiddleware: ApiClientMiddlewareFunc[] = [];
	private afterMiddleware: ApiClientMiddlewareFunc[] = [];

	public _param: ParamClient;

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

		this._param = new ParamClient(
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

	async postA(
		param: PostARequest,
		options?: {
			headers?: {[key: string]: string},
			options?: {[key: string]: any}
		}
	): Promise<PostAResponse>;
	/** @deprecated */
	async postA(
		param: PostARequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostAResponse>;

	async postA(
		param: PostARequest,
		arg1?: any,
		arg2?: any
	): Promise<PostAResponse> {
		let headers: {[key: string]: string} | undefined;
		let options: {[key: string]: any} | undefined;

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

	    const excludeParams: string[] = ['File', 'Files', ];
	    let mockHeaders: {[key: string]: string} = {};
	    if (options && options['mock_option']) {
			mockHeaders['Api-Gen-Option'] = JSON.stringify(options['mock_option']);
			delete options['mock_option'];
		}

		const reqHeader = {
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
			endpoint: `${this.baseURL}/a`,
			request: filteredParam,
			baseURL: this.baseURL,
			headers: reqHeader,
			options: reqOption,
		};
		await this.callMiddleware(this.beforeMiddleware, context);
		const url = `${this.baseURL}/a`;

		const resp = await fetch(
			url,
			{
				method: "POST",
				body: (() => {
					const body = new FormData();

					body.append(
						'x-multipart-json-binder-request-json',
						new Blob([JSON.stringify(this.getRequestObject(filteredParam, excludeParams, false))], {type: 'application/json'}),
						'x-multipart-json-binder-request-json'
					);
					if (filteredParam.File !== undefined) {
						body.append('file', filteredParam.File);
					}
					if (filteredParam.Files !== undefined) {
						filteredParam.Files.filter((f: unknown) => f !== undefined).forEach((f: File | Blob) => body.append('files', f));
					}
					return body;
				})(),
				headers: {
					...this.headers,
					...headers,
					...mockHeaders,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			const responseText = await resp.text();
			throw new ApiError(resp, responseText);
		}
		const res = (await resp.json()) as PostAResponse;
		context.response = res;
		await this.callMiddleware(this.afterMiddleware, context);
		return res;
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

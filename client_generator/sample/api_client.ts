// THIS CODE WAS GENERATED AUTOMATICALLY
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS

import { GetArticleRequest as ServiceGetArticleRequest } from './classes/service/GetArticleRequest';
export { GetArticleRequest as ServiceGetArticleRequest } from './classes/service/GetArticleRequest';
import { GetArticleResponse as ServiceGetArticleResponse } from './classes/service/GetArticleResponse';
export { GetArticleResponse as ServiceGetArticleResponse } from './classes/service/GetArticleResponse';
import { GetUserJobGetRequest as ServiceUserUserIDGetUserJobGetRequest } from './classes/service/user/_UserID/GetUserJobGetRequest';
export { GetUserJobGetRequest as ServiceUserUserIDGetUserJobGetRequest } from './classes/service/user/_UserID/GetUserJobGetRequest';
import { GetUserJobGetResponse as ServiceUserUserIDGetUserJobGetResponse } from './classes/service/user/_UserID/GetUserJobGetResponse';
export { GetUserJobGetResponse as ServiceUserUserIDGetUserJobGetResponse } from './classes/service/user/_UserID/GetUserJobGetResponse';
import { PostCreateTableRequest as PostCreateTableRequest } from './classes//PostCreateTableRequest';
export { PostCreateTableRequest as PostCreateTableRequest } from './classes//PostCreateTableRequest';
import { PostCreateTableResponse as PostCreateTableResponse } from './classes//PostCreateTableResponse';
export { PostCreateTableResponse as PostCreateTableResponse } from './classes//PostCreateTableResponse';
import { PostCreateUserRequest as PostCreateUserRequest } from './classes//PostCreateUserRequest';
export { PostCreateUserRequest as PostCreateUserRequest } from './classes//PostCreateUserRequest';
import { PostCreateUserResponse as PostCreateUserResponse } from './classes//PostCreateUserResponse';
export { PostCreateUserResponse as PostCreateUserResponse } from './classes//PostCreateUserResponse';
import { PostUpdateUserNameRequest as ServiceUserPostUpdateUserNameRequest } from './classes/service/user/PostUpdateUserNameRequest';
export { PostUpdateUserNameRequest as ServiceUserPostUpdateUserNameRequest } from './classes/service/user/PostUpdateUserNameRequest';
import { PostUpdateUserNameRequest as ServiceUser2PostUpdateUserNameRequest } from './classes/service/user2/PostUpdateUserNameRequest';
export { PostUpdateUserNameRequest as ServiceUser2PostUpdateUserNameRequest } from './classes/service/user2/PostUpdateUserNameRequest';
import { PostUpdateUserNameResponse as ServiceUserPostUpdateUserNameResponse } from './classes/service/user/PostUpdateUserNameResponse';
export { PostUpdateUserNameResponse as ServiceUserPostUpdateUserNameResponse } from './classes/service/user/PostUpdateUserNameResponse';
import { PostUpdateUserNameResponse as ServiceUser2PostUpdateUserNameResponse } from './classes/service/user2/PostUpdateUserNameResponse';
export { PostUpdateUserNameResponse as ServiceUser2PostUpdateUserNameResponse } from './classes/service/user2/PostUpdateUserNameResponse';
import { PostUpdateUserPasswordRequest as ServiceUserPostUpdateUserPasswordRequest } from './classes/service/user/PostUpdateUserPasswordRequest';
export { PostUpdateUserPasswordRequest as ServiceUserPostUpdateUserPasswordRequest } from './classes/service/user/PostUpdateUserPasswordRequest';
import { PostUpdateUserPasswordRequest as ServiceUser2PostUpdateUserPasswordRequest } from './classes/service/user2/PostUpdateUserPasswordRequest';
export { PostUpdateUserPasswordRequest as ServiceUser2PostUpdateUserPasswordRequest } from './classes/service/user2/PostUpdateUserPasswordRequest';
import { PostUpdateUserPasswordResponse as ServiceUserPostUpdateUserPasswordResponse } from './classes/service/user/PostUpdateUserPasswordResponse';
export { PostUpdateUserPasswordResponse as ServiceUserPostUpdateUserPasswordResponse } from './classes/service/user/PostUpdateUserPasswordResponse';
import { PostUpdateUserPasswordResponse as ServiceUser2PostUpdateUserPasswordResponse } from './classes/service/user2/PostUpdateUserPasswordResponse';
export { PostUpdateUserPasswordResponse as ServiceUser2PostUpdateUserPasswordResponse } from './classes/service/user2/PostUpdateUserPasswordResponse';
import { PutJobRequest as ServiceUserUserIDJobIDPutJobRequest } from './classes/service/user/_UserID/_JobID/PutJobRequest';
export { PutJobRequest as ServiceUserUserIDJobIDPutJobRequest } from './classes/service/user/_UserID/_JobID/PutJobRequest';
import { PutJobResponse as ServiceUserUserIDJobIDPutJobResponse } from './classes/service/user/_UserID/_JobID/PutJobResponse';
export { PutJobResponse as ServiceUserUserIDJobIDPutJobResponse } from './classes/service/user/_UserID/_JobID/PutJobResponse';



class ServiceClient {

	public static_page: ServiceStaticPageClient;
	public table: ServiceTableClient;
	public user: ServiceUserClient;
	public user2: ServiceUser2Client;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

		this.static_page = new ServiceStaticPageClient(headers, options, baseURL);
		this.table = new ServiceTableClient(headers, options, baseURL);
		this.user = new ServiceUserClient(headers, options, baseURL);
		this.user2 = new ServiceUser2Client(headers, options, baseURL);
	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

	async getArticle(
		param: ServiceGetArticleRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceGetArticleResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/service/article?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString(),
			{
				method: "GET",
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceGetArticleResponse(await resp.json());
	}
}

class ServiceStaticPageClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

	async getStaticPage(
		param: ServiceStaticPageGetStaticPageRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceStaticPageGetStaticPageResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/service/static_page/static_page?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString(),
			{
				method: "GET",
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceStaticPageGetStaticPageResponse(await resp.json());
	}
}

class ServiceTableClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

}

class ServiceUser2Client {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

	async postUpdateUserName(
		param: ServiceUser2PostUpdateUserNameRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2PostUpdateUserNameResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/service/user2/update_user_name`,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUser2PostUpdateUserNameResponse(await resp.json());
	}
	async postUpdateUserPassword(
		param: ServiceUser2PostUpdateUserPasswordRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUser2PostUpdateUserPasswordResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/service/user2/update_user_password`,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUser2PostUpdateUserPasswordResponse(await resp.json());
	}
}

class ServiceUserClient {

	public _UserID: ServiceUserUserIDClient;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

		this._UserID = new ServiceUserUserIDClient(headers, options, baseURL);
	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

	async postUpdateUserName(
		param: ServiceUserPostUpdateUserNameRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserPostUpdateUserNameResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/service/user/update_user_name`,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUserPostUpdateUserNameResponse(await resp.json());
	}
	async postUpdateUserPassword(
		param: ServiceUserPostUpdateUserPasswordRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserPostUpdateUserPasswordResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/service/user/update_user_password`,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUserPostUpdateUserPasswordResponse(await resp.json());
	}
}

class ServiceUserUserIDClient {

	public _JobID: ServiceUserUserIDJobIDClient;
	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

		this._JobID = new ServiceUserUserIDJobIDClient(headers, options, baseURL);
	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

	async getUserJobGet(
		param: ServiceUserUserIDGetUserJobGetRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserUserIDGetUserJobGetResponse> {
	    const excludeParams = ['UserID', ];
		const resp = await fetch(
			`${this.baseURL}/service/user/${encodeURI(param.UserID)}/user_job_get?` + (new URLSearchParams(this.getRequestObject(param, excludeParams))).toString(),
			{
				method: "GET",
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUserUserIDGetUserJobGetResponse(await resp.json());
	}
}

class ServiceUserUserIDJobIDClient {

	constructor(private headers: {[key: string]: string}, private options: {[key: string]: any}, private baseURL: string) {

	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

	async putJob(
		param: ServiceUserUserIDJobIDPutJobRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<ServiceUserUserIDJobIDPutJobResponse> {
	    const excludeParams = ['UserID', 'JobID', ];
		const resp = await fetch(
			`${this.baseURL}/service/user/${encodeURI(param.UserID)}/${encodeURI(param.JobID)}/job`,
			{
				method: "PUT",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new ServiceUserUserIDJobIDPutJobResponse(await resp.json());
	}
}


export class APIClient {
	private headers: {[key: string]: string};
	private options: {[key: string]: any};
	private baseURL: string;

	public service: ServiceClient;

	constructor(
		token?: string,
		commonHeaders?: {[key: string]: string},
		baseURL?: string,
		commonOptions: {[key: string]: any} = {}
	) {
		const headers: {[key: string]: string} =  {
			'Content-Type': 'application/json',
			...commonHeaders,
		};

		if (token !== undefined) {
			headers['Authorization'] = 'Bearer ' + token;
		}

		this.baseURL =  (baseURL === undefined) ? "" : baseURL;
		this.options = commonOptions;
		this.headers = headers;


		this.service = new ServiceClient(headers, this.options, this.baseURL);
	}

	getRequestObject(param: any, routingPath: string[]): any {
		const obj = param.toObject();
		return Object.keys(obj).filter((key) =>{
			return routingPath.indexOf(key) !== -1;
		});
	}

	async postCreateUser(
		param: PostCreateUserRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateUserResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/create_user`,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new PostCreateUserResponse(await resp.json());
	}

	async postCreateTable(
		param: PostCreateTableRequest,
		headers?: {[key: string]: string},
		options?: {[key: string]: any}
	): Promise<PostCreateTableResponse> {
	    const excludeParams = [];
		const resp = await fetch(
			`${this.baseURL}/create_table`,
			{
				method: "POST",
				body: JSON.stringify(this.getRequestObject(param, excludeParams)),
				headers: {
					...this.headers,
					...headers,
				},
				...this.options,
				...options,
			}
		);

		if (Math.floor(resp.status / 100) !== 2) {
			throw new Error(resp.statusText + ": " + await resp.text());
		}

		return new PostCreateTableResponse(await resp.json());
	}

}

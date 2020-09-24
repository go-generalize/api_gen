export type GetUserRequest = {
	id: string;
	search_request: string;
}
export type GetUserResponse = {
	ID: string;
	SearchRequest: string;
}
export type PostUpdateUserNameRequest = {
	Name: string;
}
export type PostUpdateUserNameResponse = {
	Message: string;
	Status: boolean;
}
export type PostUpdateUserPasswordRequest = {
	Password: string;
	PasswordConfirm: string;
}
export type PostUpdateUserPasswordResponse = {
	Message: string;
	Status: boolean;
}

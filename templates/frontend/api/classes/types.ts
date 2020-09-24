export type PostCreateUserRequest = {
	Gender: number;
	ID: string;
	Password: string;
}
export type PostCreateUserResponse = {
	CreatedType: number;
	Message: string;
	Status: boolean;
}

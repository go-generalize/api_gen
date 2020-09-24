export type PostCreateTableRequest = {
	ID: string;
	Text: string;
}
export type PostCreateTableResponse = {
	ID: string;
	Payload: {
		Pos: {
			X: number;
			Y: number;
		};
	};
}
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

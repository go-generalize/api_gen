// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: unknown

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

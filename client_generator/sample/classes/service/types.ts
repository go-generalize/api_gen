export type GetArticleRequest = {
	ID: number;
}
export type GetArticleResponse = {
	Body: string;
	Group: string[] | null;
	ID: number;
}

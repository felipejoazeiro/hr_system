export type LoginRequest = {
    login: string;
    password: string;
}

export type LoginReponse = {
    token: string;
    position: string;
    expires_in: number;
}
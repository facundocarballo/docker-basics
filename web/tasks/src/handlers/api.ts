const BASE_URL = "http://localhost:8080/"

export type ApiTask = {
    id: number,
    description: string,
    createdAt: Uint8Array
}

export class API {
    static BaseUrl(): string {
        return BASE_URL;
    }
}